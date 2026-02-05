package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const baseURL = "https://api.premiuminboxes.com"

var (
	httpClient *http.Client
	apiToken   string
)

func main() {
	// Initialize API token
	apiToken = os.Getenv("PREMIUMINBOXES_API_TOKEN")
	if apiToken == "" {
		fmt.Fprintln(os.Stderr, "PREMIUMINBOXES_API_TOKEN environment variable required")
		os.Exit(1)
	}

	// Initialize HTTP client
	httpClient = &http.Client{}

	// Create MCP server
	s := server.NewMCPServer(
		"Premium Inboxes MCP Server",
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
	// Tool: Get All Subscriptions
	getAllSubscriptionsTool := mcp.NewTool("get_all_subscriptions",
		mcp.WithDescription("Get all subscriptions with orders, tags, billing dates, items, and pricing"),
	)
	s.AddTool(getAllSubscriptionsTool, getAllSubscriptionsHandler)

	// Tool: Get Single Subscription
	getSubscriptionTool := mcp.NewTool("get_subscription",
		mcp.WithDescription("Get a single subscription by ID with complete details"),
		mcp.WithString("subscription_id",
			mcp.Required(),
			mcp.Description("The subscription ID"),
		),
	)
	s.AddTool(getSubscriptionTool, getSubscriptionHandler)

	// Tool: Cancel Subscription
	cancelSubscriptionTool := mcp.NewTool("cancel_subscription",
		mcp.WithDescription("Cancel a subscription"),
		mcp.WithString("subscription_id",
			mcp.Required(),
			mcp.Description("The subscription ID to cancel"),
		),
		mcp.WithString("reason",
			mcp.Description("Reason for cancellation (optional)"),
		),
		mcp.WithBoolean("remove_immediately",
			mcp.Description("Whether to remove immediately (optional)"),
		),
	)
	s.AddTool(cancelSubscriptionTool, cancelSubscriptionHandler)

	// Tool: Get All Orders
	getAllOrdersTool := mcp.NewTool("get_all_orders",
		mcp.WithDescription("Get all orders with status, domains, inboxes, emails, and tags"),
	)
	s.AddTool(getAllOrdersTool, getAllOrdersHandler)

	// Tool: Get Single Order
	getOrderTool := mcp.NewTool("get_order",
		mcp.WithDescription("Get a single order by ID with complete details"),
		mcp.WithString("order_id",
			mcp.Required(),
			mcp.Description("The order ID"),
		),
	)
	s.AddTool(getOrderTool, getOrderHandler)

	// Tool: Create Purchase
	createPurchaseTool := mcp.NewTool("create_purchase",
		mcp.WithDescription("Create a new purchase order"),
		mcp.WithString("purchase_data",
			mcp.Required(),
			mcp.Description("JSON string containing hosting platform, sequencer details, domain info, email configuration, personas, and coupons"),
		),
	)
	s.AddTool(createPurchaseTool, createPurchaseHandler)
}

// HTTP helper functions
func makeRequest(method, path string, body interface{}) ([]byte, error) {
	url := baseURL + path

	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("x-api-token", apiToken)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// Handler functions
func getAllSubscriptionsHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	respBody, err := makeRequest("GET", "/client/subscription", nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get subscriptions: %v", err)), nil
	}

	// Pretty print JSON
	var result interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to parse response: %v", err)), nil
	}

	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return mcp.NewToolResultText(string(respBody)), nil
	}

	return mcp.NewToolResultText(string(prettyJSON)), nil
}

func getSubscriptionHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	subscriptionID := args["subscription_id"].(string)

	respBody, err := makeRequest("GET", "/client/subscription/"+subscriptionID, nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get subscription: %v", err)), nil
	}

	// Pretty print JSON
	var result interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to parse response: %v", err)), nil
	}

	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return mcp.NewToolResultText(string(respBody)), nil
	}

	return mcp.NewToolResultText(string(prettyJSON)), nil
}

func cancelSubscriptionHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	subscriptionID := args["subscription_id"].(string)

	cancelData := make(map[string]interface{})
	if reason, ok := args["reason"].(string); ok {
		cancelData["reason"] = reason
	}
	if removeImmediately, ok := args["remove_immediately"].(bool); ok {
		cancelData["removeImmediately"] = removeImmediately
	}

	respBody, err := makeRequest("PUT", "/client/subscription/cancel/"+subscriptionID, cancelData)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to cancel subscription: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Subscription %s cancelled successfully!\nResponse: %s", subscriptionID, string(respBody))), nil
}

func getAllOrdersHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	respBody, err := makeRequest("GET", "/client/order", nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get orders: %v", err)), nil
	}

	// Pretty print JSON
	var result interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to parse response: %v", err)), nil
	}

	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return mcp.NewToolResultText(string(respBody)), nil
	}

	return mcp.NewToolResultText(string(prettyJSON)), nil
}

func getOrderHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	orderID := args["order_id"].(string)

	respBody, err := makeRequest("GET", "/client/order/"+orderID, nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get order: %v", err)), nil
	}

	// Pretty print JSON
	var result interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to parse response: %v", err)), nil
	}

	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return mcp.NewToolResultText(string(respBody)), nil
	}

	return mcp.NewToolResultText(string(prettyJSON)), nil
}

func createPurchaseHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	purchaseDataStr := args["purchase_data"].(string)

	// Parse the JSON string into a map
	var purchaseData map[string]interface{}
	if err := json.Unmarshal([]byte(purchaseDataStr), &purchaseData); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to parse purchase data: %v", err)), nil
	}

	respBody, err := makeRequest("POST", "/client/purchase", purchaseData)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to create purchase: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Purchase created successfully!\nOrder ID: %s", string(respBody))), nil
}
