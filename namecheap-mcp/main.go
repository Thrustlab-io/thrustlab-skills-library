package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	sandboxAPIURL    = "https://api.sandbox.namecheap.com/xml.response"
	productionAPIURL = "https://api.namecheap.com/xml.response"
)

type NamecheapClient struct {
	apiUser  string
	apiKey   string
	userName string
	clientIP string
	baseURL  string
}

// XML Response Structures
type ApiResponse struct {
	XMLName         xml.Name        `xml:"ApiResponse"`
	Status          string          `xml:"Status,attr"`
	Errors          Errors          `xml:"Errors"`
	CommandResponse CommandResponse `xml:"CommandResponse"`
}

type Errors struct {
	Error []Error `xml:"Error"`
}

type Error struct {
	Number  string `xml:"Number,attr"`
	Message string `xml:",chardata"`
}

type CommandResponse struct {
	Type              string              `xml:"Type,attr"`
	DomainCheckResult []DomainCheckResult `xml:"DomainCheckResult"`
	DomainCreateResult DomainCreateResult `xml:"DomainCreate"`
	DomainGetListResult DomainGetListResult `xml:"DomainGetListResult"`
}

type DomainCheckResult struct {
	Domain                   string `xml:"Domain,attr"`
	Available                string `xml:"Available,attr"`
	IsPremiumName            string `xml:"IsPremiumName,attr"`
	PremiumRegistrationPrice string `xml:"PremiumRegistrationPrice,attr"`
}

type DomainCreateResult struct {
	Domain      string `xml:"Domain,attr"`
	Registered  string `xml:"Registered,attr"`
	ChargedAmount string `xml:"ChargedAmount,attr"`
	DomainID    string `xml:"DomainID,attr"`
	OrderID     string `xml:"OrderID,attr"`
	TransactionID string `xml:"TransactionID,attr"`
}

type DomainGetListResult struct {
	Domains []DomainInfo `xml:"Domain"`
}

type DomainInfo struct {
	ID         string `xml:"ID,attr"`
	Name       string `xml:"Name,attr"`
	User       string `xml:"User,attr"`
	Created    string `xml:"Created,attr"`
	Expires    string `xml:"Expires,attr"`
	IsExpired  string `xml:"IsExpired,attr"`
	IsLocked   string `xml:"IsLocked,attr"`
	AutoRenew  string `xml:"AutoRenew,attr"`
	WhoisGuard string `xml:"WhoisGuard,attr"`
}

var namecheapClient *NamecheapClient

func main() {
	// Initialize Namecheap client
	apiUser := os.Getenv("NAMECHEAP_API_USER")
	apiKey := os.Getenv("NAMECHEAP_API_KEY")
	userName := os.Getenv("NAMECHEAP_USERNAME")
	clientIP := os.Getenv("NAMECHEAP_CLIENT_IP")
	useSandbox := os.Getenv("NAMECHEAP_SANDBOX") == "true"

	if apiUser == "" || apiKey == "" || userName == "" || clientIP == "" {
		fmt.Fprintln(os.Stderr, "Required environment variables: NAMECHEAP_API_USER, NAMECHEAP_API_KEY, NAMECHEAP_USERNAME, NAMECHEAP_CLIENT_IP")
		os.Exit(1)
	}

	baseURL := productionAPIURL
	if useSandbox {
		baseURL = sandboxAPIURL
	}

	namecheapClient = &NamecheapClient{
		apiUser:  apiUser,
		apiKey:   apiKey,
		userName: userName,
		clientIP: clientIP,
		baseURL:  baseURL,
	}

	// Create MCP server
	s := server.NewMCPServer(
		"Namecheap MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	// Register tools
	registerTools(s)

	// Start stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}

func registerTools(s *server.MCPServer) {
	// Tool: Check Domain Availability
	checkDomainTool := mcp.NewTool("check_domain",
		mcp.WithDescription("Check if one or more domains are available for registration"),
		mcp.WithString("domains",
			mcp.Required(),
			mcp.Description("Comma-separated list of domain names to check (e.g., example.com,mysite.net)"),
		),
	)
	s.AddTool(checkDomainTool, checkDomainHandler)

	// Tool: Create/Register Domain
	createDomainTool := mcp.NewTool("create_domain",
		mcp.WithDescription("Register a new domain name"),
		mcp.WithString("domain",
			mcp.Required(),
			mcp.Description("Domain name to register (e.g., example.com)"),
		),
		mcp.WithNumber("years",
			mcp.Description("Number of years to register the domain (default: 1)"),
		),
		mcp.WithString("registrant_first_name",
			mcp.Required(),
			mcp.Description("Registrant first name"),
		),
		mcp.WithString("registrant_last_name",
			mcp.Required(),
			mcp.Description("Registrant last name"),
		),
		mcp.WithString("registrant_address",
			mcp.Required(),
			mcp.Description("Registrant address"),
		),
		mcp.WithString("registrant_city",
			mcp.Required(),
			mcp.Description("Registrant city"),
		),
		mcp.WithString("registrant_state",
			mcp.Required(),
			mcp.Description("Registrant state/province"),
		),
		mcp.WithString("registrant_postal_code",
			mcp.Required(),
			mcp.Description("Registrant postal code"),
		),
		mcp.WithString("registrant_country",
			mcp.Required(),
			mcp.Description("Registrant country (2-letter code, e.g., US)"),
		),
		mcp.WithString("registrant_phone",
			mcp.Required(),
			mcp.Description("Registrant phone number"),
		),
		mcp.WithString("registrant_email",
			mcp.Required(),
			mcp.Description("Registrant email address"),
		),
	)
	s.AddTool(createDomainTool, createDomainHandler)

	// Tool: List Domains
	listDomainsTool := mcp.NewTool("list_domains",
		mcp.WithDescription("List all domains in your Namecheap account"),
		mcp.WithNumber("page",
			mcp.Description("Page number (default: 1)"),
		),
		mcp.WithNumber("page_size",
			mcp.Description("Number of domains per page (default: 20, max: 100)"),
		),
	)
	s.AddTool(listDomainsTool, listDomainsHandler)
}

func (c *NamecheapClient) makeRequest(command string, params map[string]string) (*ApiResponse, error) {
	// Build URL with parameters
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("ApiUser", c.apiUser)
	q.Set("ApiKey", c.apiKey)
	q.Set("UserName", c.userName)
	q.Set("ClientIp", c.clientIP)
	q.Set("Command", command)

	for key, value := range params {
		q.Set(key, value)
	}

	u.RawQuery = q.Encode()

	// Make HTTP request
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse XML response
	var apiResp ApiResponse
	if err := xml.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %v\nResponse: %s", err, string(body))
	}

	// Check for API errors
	if apiResp.Status != "OK" {
		if len(apiResp.Errors.Error) > 0 {
			errMsgs := make([]string, len(apiResp.Errors.Error))
			for i, e := range apiResp.Errors.Error {
				errMsgs[i] = fmt.Sprintf("Error %s: %s", e.Number, e.Message)
			}
			return nil, fmt.Errorf("API error: %s", strings.Join(errMsgs, "; "))
		}
		return nil, fmt.Errorf("API returned status: %s", apiResp.Status)
	}

	return &apiResp, nil
}

func checkDomainHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	domains := args["domains"].(string)

	params := map[string]string{
		"DomainList": domains,
	}

	apiResp, err := namecheapClient.makeRequest("namecheap.domains.check", params)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to check domains: %v", err)), nil
	}

	// Format results
	var result strings.Builder
	result.WriteString("Domain Availability Check Results:\n\n")

	for _, domain := range apiResp.CommandResponse.DomainCheckResult {
		available := domain.Available == "true"
		isPremium := domain.IsPremiumName == "true"

		result.WriteString(fmt.Sprintf("Domain: %s\n", domain.Domain))
		if available {
			result.WriteString("  Status: ✓ Available\n")
			if isPremium {
				result.WriteString(fmt.Sprintf("  Premium: Yes (Registration: $%s)\n", domain.PremiumRegistrationPrice))
			} else {
				result.WriteString("  Premium: No\n")
			}
		} else {
			result.WriteString("  Status: ✗ Not Available\n")
		}
		result.WriteString("\n")
	}

	return mcp.NewToolResultText(result.String()), nil
}

func createDomainHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	domain := args["domain"].(string)

	years := 1
	if v, ok := args["years"].(float64); ok {
		years = int(v)
	}

	params := map[string]string{
		"DomainName": domain,
		"Years":      fmt.Sprintf("%d", years),

		// Registrant Contact
		"RegistrantFirstName":   args["registrant_first_name"].(string),
		"RegistrantLastName":    args["registrant_last_name"].(string),
		"RegistrantAddress1":    args["registrant_address"].(string),
		"RegistrantCity":        args["registrant_city"].(string),
		"RegistrantStateProvince": args["registrant_state"].(string),
		"RegistrantPostalCode":  args["registrant_postal_code"].(string),
		"RegistrantCountry":     args["registrant_country"].(string),
		"RegistrantPhone":       args["registrant_phone"].(string),
		"RegistrantEmailAddress": args["registrant_email"].(string),

		// Admin Contact (same as registrant)
		"AdminFirstName":        args["registrant_first_name"].(string),
		"AdminLastName":         args["registrant_last_name"].(string),
		"AdminAddress1":         args["registrant_address"].(string),
		"AdminCity":             args["registrant_city"].(string),
		"AdminStateProvince":    args["registrant_state"].(string),
		"AdminPostalCode":       args["registrant_postal_code"].(string),
		"AdminCountry":          args["registrant_country"].(string),
		"AdminPhone":            args["registrant_phone"].(string),
		"AdminEmailAddress":     args["registrant_email"].(string),

		// Tech Contact (same as registrant)
		"TechFirstName":         args["registrant_first_name"].(string),
		"TechLastName":          args["registrant_last_name"].(string),
		"TechAddress1":          args["registrant_address"].(string),
		"TechCity":              args["registrant_city"].(string),
		"TechStateProvince":     args["registrant_state"].(string),
		"TechPostalCode":        args["registrant_postal_code"].(string),
		"TechCountry":           args["registrant_country"].(string),
		"TechPhone":             args["registrant_phone"].(string),
		"TechEmailAddress":      args["registrant_email"].(string),

		// Billing Contact (same as registrant)
		"AuxBillingFirstName":   args["registrant_first_name"].(string),
		"AuxBillingLastName":    args["registrant_last_name"].(string),
		"AuxBillingAddress1":    args["registrant_address"].(string),
		"AuxBillingCity":        args["registrant_city"].(string),
		"AuxBillingStateProvince": args["registrant_state"].(string),
		"AuxBillingPostalCode":  args["registrant_postal_code"].(string),
		"AuxBillingCountry":     args["registrant_country"].(string),
		"AuxBillingPhone":       args["registrant_phone"].(string),
		"AuxBillingEmailAddress": args["registrant_email"].(string),
	}

	apiResp, err := namecheapClient.makeRequest("namecheap.domains.create", params)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to create domain: %v", err)), nil
	}

	result := apiResp.CommandResponse.DomainCreateResult
	return mcp.NewToolResultText(fmt.Sprintf(
		"Domain registered successfully!\n\n"+
			"Domain: %s\n"+
			"Registered: %s\n"+
			"Domain ID: %s\n"+
			"Order ID: %s\n"+
			"Transaction ID: %s\n"+
			"Charged Amount: $%s",
		result.Domain,
		result.Registered,
		result.DomainID,
		result.OrderID,
		result.TransactionID,
		result.ChargedAmount,
	)), nil
}

func listDomainsHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)

	page := 1
	if v, ok := args["page"].(float64); ok {
		page = int(v)
	}

	pageSize := 20
	if v, ok := args["page_size"].(float64); ok {
		pageSize = int(v)
		if pageSize > 100 {
			pageSize = 100
		}
	}

	params := map[string]string{
		"Page":     fmt.Sprintf("%d", page),
		"PageSize": fmt.Sprintf("%d", pageSize),
	}

	apiResp, err := namecheapClient.makeRequest("namecheap.domains.getList", params)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to list domains: %v", err)), nil
	}

	// Format results
	var result strings.Builder
	result.WriteString("Your Namecheap Domains:\n\n")

	domains := apiResp.CommandResponse.DomainGetListResult.Domains
	if len(domains) == 0 {
		result.WriteString("No domains found.\n")
	} else {
		for i, domain := range domains {
			result.WriteString(fmt.Sprintf("%d. %s\n", i+1, domain.Name))
			result.WriteString(fmt.Sprintf("   ID: %s\n", domain.ID))
			result.WriteString(fmt.Sprintf("   Created: %s\n", domain.Created))
			result.WriteString(fmt.Sprintf("   Expires: %s\n", domain.Expires))
			result.WriteString(fmt.Sprintf("   Locked: %s\n", domain.IsLocked))
			result.WriteString(fmt.Sprintf("   Auto-Renew: %s\n", domain.AutoRenew))
			result.WriteString(fmt.Sprintf("   WhoisGuard: %s\n", domain.WhoisGuard))
			result.WriteString("\n")
		}
	}

	return mcp.NewToolResultText(result.String()), nil
}
