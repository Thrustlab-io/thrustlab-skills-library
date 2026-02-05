package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	clayAPIBase = "https://api.clay.com/v3"

	// Action Package IDs from HAR analysis
	mixrankActionPackageID = "e251a70e-46d7-4f3a-b3ef-a211ad3d8bd2"
	mixrankActionKey       = "find-lists-of-companies-with-mixrank-source"

	googleMapsActionPackageID = "3282a1c7-6bb0-497e-a34b-32268e104e55"
	googleMapsActionKey       = "google-review-source-v3"
)

var (
	workspaceID     string
	sessionCookie   string
	frontendVersion string
	httpClient      *http.Client
	currentWorkbookID string // Store the current workbook ID
)

func main() {
	// Load configuration from environment
	workspaceID = os.Getenv("CLAY_WORKSPACE_ID")
	sessionCookie = os.Getenv("CLAY_SESSION_COOKIE")
	frontendVersion = os.Getenv("CLAY_FRONTEND_VERSION")

	if workspaceID == "" || sessionCookie == "" {
		fmt.Fprintln(os.Stderr, "Required environment variables:")
		fmt.Fprintln(os.Stderr, "  CLAY_WORKSPACE_ID - Your Clay workspace ID")
		fmt.Fprintln(os.Stderr, "  CLAY_SESSION_COOKIE - Session cookie from browser")
		fmt.Fprintln(os.Stderr, "  CLAY_FRONTEND_VERSION (optional) - Frontend version header")
		os.Exit(1)
	}

	if frontendVersion == "" {
		frontendVersion = "v20260205_151537Z_19e7945c5e"
	}

	httpClient = &http.Client{}

	// Create MCP server
	s := server.NewMCPServer(
		"Clay MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(true, false),
	)

	// Register tools and resources
	registerTools(s)
	registerResources(s)

	// Start stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}

func registerTools(s *server.MCPServer) {
	// Tool: Set Workspace ID
	setWorkspaceTool := mcp.NewTool("set_workspace_id",
		mcp.WithDescription("Set your Clay workspace ID for API access"),
		mcp.WithString("workspace_id",
			mcp.Required(),
			mcp.Description("Your Clay workspace ID (e.g., 757984)"),
		),
	)
	s.AddTool(setWorkspaceTool, setWorkspaceIDHandler)

	// Tool: Set Session Cookie
	setSessionCookieTool := mcp.NewTool("set_session_cookie",
		mcp.WithDescription("Set your Clay session cookie for authentication"),
		mcp.WithString("session_cookie",
			mcp.Required(),
			mcp.Description("Session cookie value (starts with s%3A...)"),
		),
	)
	s.AddTool(setSessionCookieTool, setSessionCookieHandler)

	// Tool: Create Workbook
	createWorkbookTool := mcp.NewTool("create_workbook",
		mcp.WithDescription("Create a new Clay workbook and set it as the default for subsequent operations"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name for the new workbook"),
		),
		mcp.WithString("folder_id",
			mcp.Description("Optional folder ID to create the workbook in"),
		),
	)
	s.AddTool(createWorkbookTool, createWorkbookHandler)

	// Tool: Search Companies by Industry
	industrySearchTool := mcp.NewTool("search_companies_by_industry",
		mcp.WithDescription("Search for companies by industry using Clay's Mixrank/LinkedIn data source"),
		mcp.WithString("workbook_id",
			mcp.Description("The workbook ID to create the table in (uses current workbook if not specified)"),
		),
		mcp.WithString("industries",
			mcp.Required(),
			mcp.Description("Comma-separated list of industries (e.g., 'Accounting,Consulting')"),
		),
		mcp.WithString("countries",
			mcp.Description("Comma-separated list of country names (e.g., 'Belgium,Netherlands')"),
		),
		mcp.WithString("company_sizes",
			mcp.Description("Comma-separated company sizes: 1,2-10,11-50,50,200,500,1000,5000,10000"),
		),
		mcp.WithString("keywords",
			mcp.Description("Comma-separated description keywords to filter by"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of results (default: 25000)"),
		),
	)
	s.AddTool(industrySearchTool, searchCompaniesByIndustryHandler)

	// Tool: Search Businesses by Geography (Google Maps)
	geographySearchTool := mcp.NewTool("search_businesses_by_geography",
		mcp.WithDescription("Search for local businesses by geography using Google Maps. Supports two search modes: business types (e.g., book_store, restaurant) or free text query (e.g., 'slager', 'pizza'). Provide either business_types OR query, not both."),
		mcp.WithString("workbook_id",
			mcp.Description("The workbook ID to create the table in (uses current workbook if not specified)"),
		),
		mcp.WithNumber("latitude",
			mcp.Required(),
			mcp.Description("Latitude coordinate (e.g., 51.049)"),
		),
		mcp.WithNumber("longitude",
			mcp.Required(),
			mcp.Description("Longitude coordinate (e.g., 3.725)"),
		),
		mcp.WithNumber("proximity_km",
			mcp.Required(),
			mcp.Description("Search radius in kilometers (e.g., 13)"),
		),
		mcp.WithString("business_types",
			mcp.Description("Comma-separated Google Maps business types. Valid types include: accounting, art_gallery, bakery, bank, bar, beauty_salon, book_store, cafe, car_dealer, car_repair, clothing_store, dentist, doctor, electrician, florist, furniture_store, gym, hair_care, hardware_store, hospital, jewelry_store, lawyer, library, liquor_store, lodging, museum, night_club, pharmacy, plumber, real_estate_agency, restaurant, school, shoe_store, spa, store, supermarket, veterinary_care, etc. Required if query is not provided."),
		),
		mcp.WithString("query",
			mcp.Description("Free text search query (e.g., 'slager', 'pizza', 'kapper'). Use this for searches that don't match a standard Google Maps business type. Required if business_types is not provided."),
		),
		mcp.WithNumber("num_results",
			mcp.Description("Maximum number of results to return (default: 100)"),
		),
		mcp.WithString("table_name",
			mcp.Description("Custom table name (default: '⚡️ Find local businesses using Google Maps Table')"),
		),
		mcp.WithString("table_emoji",
			mcp.Description("Table emoji icon (default: '🌙')"),
		),
	)
	s.AddTool(geographySearchTool, searchBusinessesByGeographyHandler)
}

func registerResources(s *server.MCPServer) {
	// Resource: List all available commands
	commandsResource := mcp.NewResource(
		"clay://commands",
		"List of all available Clay MCP commands",
		mcp.WithResourceDescription("Complete list of available commands and their descriptions"),
		mcp.WithMIMEType("text/plain"),
	)
	s.AddResource(commandsResource, listCommandsHandler)

	// Resource: List all Google Maps business types
	businessTypesResource := mcp.NewResource(
		"clay://business-types",
		"List of all Google Maps business types",
		mcp.WithResourceDescription("Complete list of valid business types for geography searches"),
		mcp.WithMIMEType("text/plain"),
	)
	s.AddResource(businessTypesResource, listBusinessTypesHandler)
}

func listCommandsHandler(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	commands := `Clay MCP Server - Available Commands
==========================================

Configuration Commands:
-----------------------
1. set_workspace_id
   Set your Clay workspace ID for API access
   Parameters:
   - workspace_id (required): Your Clay workspace ID (e.g., 757984)

2. set_session_cookie
   Set your Clay session cookie for authentication
   Parameters:
   - session_cookie (required): Session cookie value (starts with s%3A...)

Workbook Commands:
------------------
3. create_workbook
   Create a new Clay workbook and set it as default
   Parameters:
   - name (required): Name for the new workbook
   - folder_id (optional): Folder ID to create the workbook in

Search Commands:
----------------
4. search_companies_by_industry
   Search for companies by industry using Clay's Mixrank/LinkedIn data
   Parameters:
   - workbook_id (optional): Uses current workbook if not specified
   - industries (required): Comma-separated list (e.g., 'Accounting,Consulting')
   - countries (optional): Comma-separated country names
   - company_sizes (optional): Comma-separated sizes (1,2-10,11-50,50,200,500,1000,5000,10000)
   - keywords (optional): Comma-separated description keywords
   - limit (optional): Maximum results (default: 25000)

5. search_businesses_by_geography
   Search for local businesses by geography using Google Maps
   Two search modes: business types or free text query (provide one or the other)
   Parameters:
   - workbook_id (optional): Uses current workbook if not specified
   - latitude (required): Latitude coordinate (e.g., 51.049)
   - longitude (required): Longitude coordinate (e.g., 3.725)
   - proximity_km (required): Search radius in kilometers
   - business_types (optional): Comma-separated types (see clay://business-types)
   - query (optional): Free text search (e.g., 'slager', 'pizza', 'kapper')
   - num_results (optional): Maximum number of results (default: 100)
   - table_name (optional): Custom table name
   - table_emoji (optional): Table emoji icon

Resources:
----------
- clay://commands - This list of commands
- clay://business-types - List of all valid Google Maps business types
`

	content := mcp.TextResourceContents{
		URI:      req.Params.URI,
		MIMEType: "text/plain",
		Text:     commands,
	}

	return []mcp.ResourceContents{&content}, nil
}

func listBusinessTypesHandler(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	businessTypes := `Google Maps Business Types
===========================

Common Business Types:
----------------------
accounting, airport, amusement_park, aquarium, art_gallery, atm, bakery, bank, bar,
beauty_salon, bicycle_store, book_store, bowling_alley, bus_station, cafe, campground,
car_dealer, car_rental, car_repair, car_wash, casino, cemetery, church, city_hall,
clothing_store, convenience_store, courthouse, dentist, department_store, doctor,
drugstore, electrician, electronics_store, embassy, fire_station, florist, funeral_home,
furniture_store, gas_station, gym, hair_care, hardware_store, hindu_temple, home_goods_store,
hospital, insurance_agency, jewelry_store, laundry, lawyer, library, light_rail_station,
liquor_store, local_government_office, locksmith, lodging, meal_delivery, meal_takeaway,
mosque, movie_rental, movie_theater, moving_company, museum, night_club, painter, park,
parking, pet_store, pharmacy, physiotherapist, plumber, police, post_office, primary_school,
real_estate_agency, restaurant, roofing_contractor, rv_park, school, secondary_school,
shoe_store, shopping_mall, spa, stadium, storage, store, subway_station, supermarket,
synagogue, taxi_stand, tourist_attraction, train_station, transit_station, travel_agency,
university, veterinary_care, zoo

Usage Examples:
---------------
Single type:
  book_store

Multiple types (comma-separated):
  book_store,library,university

Restaurant types:
  restaurant,cafe,bar,meal_delivery

Retail:
  clothing_store,shoe_store,jewelry_store,department_store

Services:
  hair_care,beauty_salon,spa,gym

Note: Use underscores (_) not spaces in business type names.
`

	content := mcp.TextResourceContents{
		URI:      req.Params.URI,
		MIMEType: "text/plain",
		Text:     businessTypes,
	}

	return []mcp.ResourceContents{&content}, nil
}

func createWorkbookHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	name := args["name"].(string)

	var folderID *string
	if v, ok := args["folder_id"].(string); ok && v != "" {
		folderID = &v
	}

	// Build request payload
	payload := map[string]any{
		"name":        name,
		"workspaceId": workspaceID,
	}

	if folderID != nil {
		payload["folderId"] = *folderID
	}

	url := fmt.Sprintf("%s/workbooks", clayAPIBase)
	resp, err := makeClayRequest("POST", url, payload)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to create workbook: %v", err)), nil
	}

	// Parse response
	var result map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to parse response: %v", err)), nil
	}

	// Store the workbook ID for future use
	currentWorkbookID = result["id"].(string)

	return mcp.NewToolResultText(fmt.Sprintf(
		"✅ Workbook created successfully!\n\n"+
			"Workbook ID: %s\n"+
			"Workbook Name: %s\n\n"+
			"This workbook is now set as the default for subsequent operations.\n"+
			"You no longer need to specify workbook_id in other commands.\n\n"+
			"View in Clay: https://app.clay.com/workspaces/%s/workbooks/%s",
		currentWorkbookID,
		result["name"],
		workspaceID,
		currentWorkbookID,
	)), nil
}

func setWorkspaceIDHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	newWorkspaceID := args["workspace_id"].(string)

	// Update the global workspace ID (for immediate use)
	workspaceID = newWorkspaceID

	// Update Claude Desktop config file
	if err := updateClaudeConfig(func(config map[string]any) error {
		// Navigate to mcpServers.clay.env
		servers, ok := config["mcpServers"].(map[string]any)
		if !ok {
			servers = make(map[string]any)
			config["mcpServers"] = servers
		}

		clay, ok := servers["clay"].(map[string]any)
		if !ok {
			clay = make(map[string]any)
			servers["clay"] = clay
		}

		env, ok := clay["env"].(map[string]any)
		if !ok {
			env = make(map[string]any)
			clay["env"] = env
		}

		env["CLAY_WORKSPACE_ID"] = newWorkspaceID
		return nil
	}); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to update config: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf(
		"✅ Clay workspace ID updated successfully!\n\n"+
			"Workspace ID: %s\n\n"+
			"Updated in: ~/Library/Application Support/Claude/claude_desktop_config.json\n\n"+
			"⚠️ Please restart Claude Desktop for changes to take effect.",
		workspaceID,
	)), nil
}

func setSessionCookieHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	newCookie := args["session_cookie"].(string)

	// Update the global session cookie (for immediate use)
	sessionCookie = newCookie

	// Mask the cookie for display (show first and last 10 chars)
	maskedCookie := newCookie
	if len(newCookie) > 20 {
		maskedCookie = newCookie[:10] + "..." + newCookie[len(newCookie)-10:]
	}

	// Update Claude Desktop config file
	if err := updateClaudeConfig(func(config map[string]any) error {
		// Navigate to mcpServers.clay.env
		servers, ok := config["mcpServers"].(map[string]any)
		if !ok {
			servers = make(map[string]any)
			config["mcpServers"] = servers
		}

		clay, ok := servers["clay"].(map[string]any)
		if !ok {
			clay = make(map[string]any)
			servers["clay"] = clay
		}

		env, ok := clay["env"].(map[string]any)
		if !ok {
			env = make(map[string]any)
			clay["env"] = env
		}

		env["CLAY_SESSION_COOKIE"] = newCookie
		return nil
	}); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to update config: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf(
		"✅ Clay session cookie updated successfully!\n\n"+
			"Cookie: %s\n\n"+
			"Updated in: ~/Library/Application Support/Claude/claude_desktop_config.json\n\n"+
			"⚠️ Please restart Claude Desktop for changes to take effect.",
		maskedCookie,
	)), nil
}

func searchCompaniesByIndustryHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)

	// Use provided workbook_id or fall back to current workbook
	workbookID := currentWorkbookID
	if v, ok := args["workbook_id"].(string); ok && v != "" {
		workbookID = v
	}

	if workbookID == "" {
		return mcp.NewToolResultError("No workbook specified. Either provide workbook_id or create a workbook first using create_workbook."), nil
	}

	industries := strings.Split(args["industries"].(string), ",")

	// Optional parameters
	var countries []string
	if v, ok := args["countries"].(string); ok && v != "" {
		countries = strings.Split(v, ",")
	}

	var sizes []string
	if v, ok := args["company_sizes"].(string); ok && v != "" {
		sizes = strings.Split(v, ",")
	}

	var keywords []string
	if v, ok := args["keywords"].(string); ok && v != "" {
		keywords = strings.Split(v, ",")
	}

	limit := 25000.0
	if v, ok := args["limit"].(float64); ok {
		limit = v
	}

	// Build request payload matching HAR structure
	payload := map[string]any{
		"workbookId":       workbookID,
		"wizardId":         "find-companies",
		"wizardStepId":     "companies-search",
		"sessionId":        generateSessionID(),
		"currentStepIndex": 0,
		"outputs":          []any{},
		"firstUseCase":     nil,
		"parentFolderId":   nil,
		"formInputs": map[string]any{
			"clientSettings": map[string]any{
				"tableType": "company",
			},
			"requiredDataPoint": nil,
			"type":              "companies",
			"typeSettings": map[string]any{
				"name":              "Find companies",
				"iconType":          "Buildings",
				"actionKey":         mixrankActionKey,
				"actionPackageId":   mixrankActionPackageID,
				"previewTextPath":   "name",
				"defaultPreviewText": "Profile",
				"recordsPath":       "companies",
				"idPath":            "linkedin_company_id",
				"scheduleConfig": map[string]any{
					"runSettings": "once",
				},
				"inputs": map[string]any{
					"industries":                         industries,
					"country_names":                      countries,
					"sizes":                              sizes,
					"description_keywords":               keywords,
					"description_keywords_exclude":       []string{},
					"limit":                              limit,
					"types":                              []string{},
					"locations":                          []string{},
					"locations_exclude":                  []string{},
					"funding_amounts":                    []string{},
					"annual_revenues":                    []string{},
					"industries_exclude":                 []string{},
					"minimum_follower_count":             nil,
					"minimum_member_count":               nil,
					"maximum_member_count":               nil,
					"semantic_description":               "",
					"company_identifier":                 []string{},
					"startFromCompanyType":               "company_identifier",
					"exclude_company_identifiers_mixed":  []string{},
					"exclude_entities_configuration":     []string{},
					"exclude_entities_bitmap":            nil,
					"previous_entities_bitmap":           nil,
					"derived_industries":                 []string{},
					"derived_subindustries":              []string{},
					"derived_subindustries_exclude":      []string{},
					"derived_revenue_streams":            []string{},
					"derived_business_types":             []string{},
					"tableId":                            nil,
					"domainFieldId":                      nil,
					"useRadialKnn":                       false,
					"radialKnnMinScore":                  nil,
					"has_resolved_domain":                nil,
					"resolved_domain_is_live":            nil,
					"resolved_domain_redirects":          nil,
					"name":                               "",
				},
				"hasEvaluatedInputs": true,
			},
		},
	}

	url := fmt.Sprintf("%s/workspaces/%s/wizard/evaluate-step", clayAPIBase, workspaceID)
	resp, err := makeClayRequest("POST", url, payload)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to create company search: %v", err)), nil
	}

	// Parse response
	var result map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to parse response: %v", err)), nil
	}

	output := result["output"].(map[string]any)
	table := output["table"].(map[string]any)

	return mcp.NewToolResultText(fmt.Sprintf(
		"✅ Company search table created successfully!\n\n"+
			"Table ID: %s\n"+
			"Table Name: %s\n"+
			"Records Found: %.0f\n"+
			"Industries: %s\n"+
			"Countries: %s\n\n"+
			"View in Clay: https://app.clay.com/workspaces/%s/workbooks/%s/tables/%s",
		table["tableId"],
		table["tableName"],
		output["recordCount"],
		strings.Join(industries, ", "),
		strings.Join(countries, ", "),
		workspaceID,
		workbookID,
		table["tableId"],
	)), nil
}

func searchBusinessesByGeographyHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)

	// Use provided workbook_id or fall back to current workbook
	workbookID := currentWorkbookID
	if v, ok := args["workbook_id"].(string); ok && v != "" {
		workbookID = v
	}

	if workbookID == "" {
		return mcp.NewToolResultError("No workbook specified. Either provide workbook_id or create a workbook first using create_workbook."), nil
	}

	latitude := args["latitude"].(float64)
	longitude := args["longitude"].(float64)
	proximityKm := args["proximity_km"].(float64)

	// Determine search mode: business types or free text
	var businessTypes []string
	var query string
	isFreeText := false

	if v, ok := args["query"].(string); ok && v != "" {
		query = v
		isFreeText = true
		// Free text mode uses "store" as default business type
		businessTypes = []string{"store"}
	}

	if v, ok := args["business_types"].(string); ok && v != "" {
		raw := strings.Split(v, ",")
		for _, bt := range raw {
			cleaned := strings.TrimSpace(bt)
			if cleaned != "" {
				businessTypes = append(businessTypes, cleaned)
			}
		}
		if !isFreeText {
			// Only business types mode
		}
	}

	if !isFreeText && len(businessTypes) == 0 {
		return mcp.NewToolResultError("Either business_types or query must be provided."), nil
	}

	// Number of results (default: 100)
	numResults := 100
	if v, ok := args["num_results"].(float64); ok && v > 0 {
		numResults = int(v)
	}

	tableName := "⚡️ Find local businesses using Google Maps Table"
	if v, ok := args["table_name"].(string); ok && v != "" {
		tableName = v
	}

	tableEmoji := "🌙"
	if v, ok := args["table_emoji"].(string); ok && v != "" {
		tableEmoji = v
	}

	// Build Google Maps search configuration (matching exact HAR format)
	mapConfig := fmt.Sprintf(`{"latitude":%f,"longitude":%f,"proximity":%d}`,
		latitude, longitude, int(proximityKm))

	// Build business types JSON array
	businessTypesJSON, _ := json.Marshal(businessTypes)

	// Build inputs based on search mode
	inputs := map[string]any{
		"usePreferredGoogleApi": "true",
		"map":                   mapConfig,
		"numResults":            fmt.Sprintf("%d", numResults),
		"dynamicFields|businessTypes": string(businessTypesJSON),
	}

	if isFreeText {
		inputs["dynamicFields|searchType"] = `"freeText"`
		inputs["dynamicFields|searchType_displayName"] = `"Free text"`
		inputs["dynamicFields|query"] = fmt.Sprintf(`"%s"`, query)
	} else {
		inputs["dynamicFields|searchType"] = `"businessTypes"`
		inputs["dynamicFields|searchType_displayName"] = `"Business types"`
	}

	// Build request payload matching HAR structure exactly
	payload := map[string]any{
		"icon": map[string]string{
			"emoji": tableEmoji,
		},
		"workspaceId": workspaceID,
		"type":        "company",
		"template":    "empty",
		"name":        tableName,
		"workbookId":  workbookID,
		"callerName":  "source creator modal",
		"sourceSettings": map[string]any{
			"addSource": map[string]any{
				"name": "Find local businesses using Google Maps",
				"source": map[string]any{
					"name":        "⚡️ Find local businesses using Google Maps",
					"workspaceId": workspaceID,
					"type":        "v3-action",
					"typeSettings": map[string]any{
						"name":                   "Find local businesses using Google Maps",
						"description":            "Pull local businesses from a specific location on Google Maps",
						"stages":                 []string{"Inputs"},
						"categories":             []string{"FIND"},
						"customSignalSettings": map[string]any{
							"categories":  []string{"SOURCING"},
							"rank":        3,
							"title":       "Monitor local businesses using Google Maps",
							"description": "Monitor local businesses from a specific location on Google Maps",
						},
						"iconType":               "GoogleMapsSource",
						"actionKey":              googleMapsActionKey,
						"actionPackageId":        googleMapsActionPackageID,
						"defaultPreviewText":     "Business Found",
						"recordsPath":            "results",
						"idPath":                 "id",
						"isAdmin":                false,
						"dedupeOnUniqueIds":      true,
						"isPaginationAvailable":  true,
						"tableType":              "company",
						"costEstimate":           1,
						"sourceTableOutputs": map[string]any{
							"spreadsheet": map[string]any{
								"newFieldsNameToPaths": map[string]any{
									"Name":            map[string]any{"path": []string{"name"}, "type": "text"},
									"Google Maps URL": map[string]any{"path": []string{"googleMapsPlaceLink"}, "type": "url"},
									"Description":     map[string]any{"path": []string{"description"}, "type": "text"},
									"Website":         map[string]any{"path": []string{"website"}, "type": "url"},
									"Phone":           map[string]any{"path": []string{"phone"}, "type": "text"},
									"Address":         map[string]any{"path": []string{"address"}, "type": "text"},
									"Rating":          map[string]any{"path": []string{"rating"}, "type": "number"},
									"Reviews Count":   map[string]any{"path": []string{"reviews.count"}, "type": "number"},
								},
							},
						},
						"scheduleConfig": map[string]any{
							"runSettings": "once",
						},
						"inputs": inputs,
					},
				},
				"isPinned": true,
			},
			"addBasicFields": []map[string]any{
				{"name": "Name", "dataType": "text", "formulaText": "{{source}}.name"},
				{"name": "Google Maps URL", "dataType": "url", "formulaText": "{{source}}.googleMapsPlaceLink"},
				{"name": "Description", "dataType": "text", "formulaText": "{{source}}.description"},
				{"name": "Website", "dataType": "url", "formulaText": "{{source}}.website"},
				{"name": "Phone", "dataType": "text", "formulaText": "{{source}}.phone"},
				{"name": "Address", "dataType": "text", "formulaText": "{{source}}.address"},
				{"name": "Rating", "dataType": "number", "formulaText": "{{source}}.rating"},
				{"name": "Reviews Count", "dataType": "number", "formulaText": "{{source}}.reviews.count"},
			},
		},
	}

	url := fmt.Sprintf("%s/tables", clayAPIBase)
	resp, err := makeClayRequest("POST", url, payload)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to create geography search: %v", err)), nil
	}

	// Parse response
	var result map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to parse response: %v", err)), nil
	}

	// Extract table info from response
	tableInfo, ok := result["table"].(map[string]any)
	var tableID, tableNameResult string
	if ok {
		if id, ok := tableInfo["id"].(string); ok {
			tableID = id
		}
		if name, ok := tableInfo["name"].(string); ok {
			tableNameResult = name
		}
	}
	// Fallback to root level if not in nested structure
	if tableID == "" {
		if id, ok := result["id"].(string); ok {
			tableID = id
		}
	}
	if tableNameResult == "" {
		if name, ok := result["name"].(string); ok {
			tableNameResult = name
		}
	}

	searchMode := fmt.Sprintf("Business Types: %s", strings.Join(businessTypes, ", "))
	if isFreeText {
		searchMode = fmt.Sprintf("Free Text Query: %s", query)
	}

	return mcp.NewToolResultText(fmt.Sprintf(
		"✅ Geography search table created successfully!\n\n"+
			"Table ID: %s\n"+
			"Table Name: %s\n"+
			"Location: %.4f, %.4f (radius: %.0f km)\n"+
			"Search Mode: %s\n"+
			"Max Results: %d\n\n"+
			"View in Clay: https://app.clay.com/workspaces/%s/workbooks/%s/tables/%s",
		tableID,
		tableNameResult,
		latitude, longitude, proximityKm,
		searchMode,
		numResults,
		workspaceID,
		workbookID,
		tableID,
	)), nil
}

// Helper function to update Claude Desktop config file
func updateClaudeConfig(updateFn func(config map[string]any) error) error {
	// Get config file path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	configPath := filepath.Join(homeDir, "Library", "Application Support", "Claude", "claude_desktop_config.json")

	// Read existing config
	var config map[string]any
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Create new config if it doesn't exist
			config = make(map[string]any)
		} else {
			return fmt.Errorf("failed to read config: %w", err)
		}
	} else {
		if err := json.Unmarshal(data, &config); err != nil {
			return fmt.Errorf("failed to parse config: %w", err)
		}
	}

	// Apply the update
	if err := updateFn(config); err != nil {
		return err
	}

	// Write back to file
	updatedData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, updatedData, 0644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}

// Helper function to make authenticated requests to Clay API
func makeClayRequest(method, url string, payload any) ([]byte, error) {
	var body io.Reader
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request: %w", err)
		}
		body = bytes.NewReader(jsonData)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers based on HAR analysis
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:147.0) Gecko/20100101 Firefox/147.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Clay-Frontend-Version", frontendVersion)
	req.Header.Set("Origin", "https://app.clay.com")
	req.Header.Set("Referer", "https://app.clay.com/")
	req.Header.Set("Cookie", fmt.Sprintf("claysession=%s", sessionCookie))

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// Generate a session ID (simplified version)
func generateSessionID() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		randomBytes(4), randomBytes(2), randomBytes(2), randomBytes(2), randomBytes(6))
}

func randomBytes(n int) []byte {
	b := make([]byte, n)
	// Simple pseudo-random for session ID
	for i := range b {
		b[i] = byte(i * 17 % 256)
	}
	return b
}
