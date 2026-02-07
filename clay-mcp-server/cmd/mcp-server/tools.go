package main

import (
	"context"
	"fmt"
	"strings"

	"clay-mcp/pkg/clay"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerTools(s *server.MCPServer) {
	setWorkspaceTool := mcp.NewTool("set_workspace_id",
		mcp.WithDescription("Set your Clay workspace ID for API access"),
		mcp.WithString("workspace_id",
			mcp.Required(),
			mcp.Description("Your Clay workspace ID (e.g., 757984)"),
		),
	)
	s.AddTool(setWorkspaceTool, setWorkspaceIDHandler)

	setSessionCookieTool := mcp.NewTool("set_session_cookie",
		mcp.WithDescription("Set your Clay session cookie for authentication"),
		mcp.WithString("session_cookie",
			mcp.Required(),
			mcp.Description("Session cookie value (starts with s%3A...)"),
		),
	)
	s.AddTool(setSessionCookieTool, setSessionCookieHandler)

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

	companySearchTool := mcp.NewTool("search_companies_by_industry",
		mcp.WithDescription("Search for companies by industry using Clay's Mixrank/LinkedIn data source"),
		mcp.WithString("workbook_id",
			mcp.Description("The workbook ID to create the table in (uses current workbook if not specified)"),
		),
		mcp.WithString("keywords",
			mcp.Required(),
			mcp.Description("Comma-separated description keywords to filter by (e.g., 'accounting,boekhouding,SaaS')"),
		),
		mcp.WithString("countries",
			mcp.Description("Comma-separated list of country names (e.g., 'Belgium,Netherlands')"),
		),
		mcp.WithString("company_sizes",
			mcp.Description("Comma-separated company size codes: 1, 2, 10, 50, 200, 500, 1000, 5000, 10000"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of results (default: 25000)"),
		),
		mcp.WithString("annual_revenues",
			mcp.Description("Comma-separated annual revenue ranges. Valid values: 0-500K, 500K-1M, 1M-5M, 5M-10M, 10M-25M, 25M-75M, 75M-200M, 200M-500M, 500M-1B, 1B-10B, 10B-100B, 100B-1T"),
		),
		mcp.WithNumber("min_linkedin_members",
			mcp.Description("Minimum number of LinkedIn members (e.g., 100)"),
		),
		mcp.WithNumber("max_linkedin_members",
			mcp.Description("Maximum number of LinkedIn members (e.g., 200)"),
		),
	)
	s.AddTool(companySearchTool, searchCompaniesHandler)

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
			mcp.Description("Custom table name (default: '\u26a1\ufe0f Find local businesses using Google Maps Table')"),
		),
		mcp.WithString("table_emoji",
			mcp.Description("Table emoji icon (default: '\U0001f319')"),
		),
	)
	s.AddTool(geographySearchTool, searchBusinessesByGeographyHandler)
}

func setWorkspaceIDHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	newID := args["workspace_id"].(string)

	clayClient.SetWorkspaceID(newID)

	if err := updateClaudeConfig(func(config map[string]any) error {
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
		env["CLAY_WORKSPACE_ID"] = newID
		return nil
	}); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to update config: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf(
		"Clay workspace ID updated successfully!\n\n"+
			"Workspace ID: %s\n\n"+
			"Updated in: ~/Library/Application Support/Claude/claude_desktop_config.json\n\n"+
			"Please restart Claude Desktop for changes to take effect.",
		clayClient.WorkspaceID,
	)), nil
}

func setSessionCookieHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	newCookie := args["session_cookie"].(string)

	clayClient.SetSessionCookie(newCookie)

	maskedCookie := newCookie
	if len(newCookie) > 20 {
		maskedCookie = newCookie[:10] + "..." + newCookie[len(newCookie)-10:]
	}

	if err := updateClaudeConfig(func(config map[string]any) error {
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
		"Clay session cookie updated successfully!\n\n"+
			"Cookie: %s\n\n"+
			"Updated in: ~/Library/Application Support/Claude/claude_desktop_config.json\n\n"+
			"Please restart Claude Desktop for changes to take effect.",
		maskedCookie,
	)), nil
}

func createWorkbookHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	name := args["name"].(string)

	var folderID string
	if v, ok := args["folder_id"].(string); ok {
		folderID = v
	}

	result, err := clayClient.CreateWorkbook(clay.CreateWorkbookParams{
		Name:     name,
		FolderID: folderID,
	})
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to create workbook: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf(
		"Workbook created successfully!\n\n"+
			"Workbook ID: %s\n"+
			"Workbook Name: %s\n\n"+
			"This workbook is now set as the default for subsequent operations.\n"+
			"You no longer need to specify workbook_id in other commands.\n\n"+
			"View in Clay: https://app.clay.com/workspaces/%s/workbooks/%s",
		result.ID,
		result.Name,
		clayClient.WorkspaceID,
		result.ID,
	)), nil
}

func searchCompaniesHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)

	workbookID := clayClient.CurrentWorkbookID
	if v, ok := args["workbook_id"].(string); ok && v != "" {
		workbookID = v
	}
	if workbookID == "" {
		return mcp.NewToolResultError("No workbook specified. Either provide workbook_id or create a workbook first using create_workbook."), nil
	}

	params := clay.SearchCompaniesParams{
		WorkbookID: workbookID,
		Keywords:   splitAndTrim(args["keywords"].(string)),
	}

	if v, ok := args["countries"].(string); ok && v != "" {
		params.Countries = splitAndTrim(v)
	}
	if v, ok := args["company_sizes"].(string); ok && v != "" {
		params.CompanySizes = splitAndTrim(v)
	}
	if v, ok := args["annual_revenues"].(string); ok && v != "" {
		params.AnnualRevenues = splitAndTrim(v)
	}
	if v, ok := args["min_linkedin_members"].(float64); ok {
		n := int(v)
		params.MinLinkedInMembers = &n
	}
	if v, ok := args["max_linkedin_members"].(float64); ok {
		n := int(v)
		params.MaxLinkedInMembers = &n
	}

	result, err := clayClient.SearchCompanies(params)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to create company search: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf(
		"Company search table created successfully!\n\n"+
			"Table ID: %s\n"+
			"Table Name: %s\n"+
			"Records Found: %.0f\n"+
			"Keywords: %s\n"+
			"Countries: %s\n\n"+
			"View in Clay: https://app.clay.com/workspaces/%s/workbooks/%s/tables/%s",
		result.TableID,
		result.TableName,
		result.RecordCount,
		strings.Join(params.Keywords, ", "),
		strings.Join(params.Countries, ", "),
		clayClient.WorkspaceID,
		workbookID,
		result.TableID,
	)), nil
}

func searchBusinessesByGeographyHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)

	workbookID := clayClient.CurrentWorkbookID
	if v, ok := args["workbook_id"].(string); ok && v != "" {
		workbookID = v
	}
	if workbookID == "" {
		return mcp.NewToolResultError("No workbook specified. Either provide workbook_id or create a workbook first using create_workbook."), nil
	}

	params := clay.SearchBusinessesParams{
		WorkbookID:  workbookID,
		Latitude:    args["latitude"].(float64),
		Longitude:   args["longitude"].(float64),
		ProximityKm: args["proximity_km"].(float64),
	}

	if v, ok := args["query"].(string); ok && v != "" {
		params.Query = v
	}
	if v, ok := args["business_types"].(string); ok && v != "" {
		params.BusinessTypes = splitAndTrim(v)
	}
	if v, ok := args["num_results"].(float64); ok && v > 0 {
		params.NumResults = int(v)
	}
	if v, ok := args["table_name"].(string); ok && v != "" {
		params.TableName = v
	}
	if v, ok := args["table_emoji"].(string); ok && v != "" {
		params.TableEmoji = v
	}

	result, err := clayClient.SearchBusinessesByGeography(params)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to create geography search: %v", err)), nil
	}

	searchMode := fmt.Sprintf("Business Types: %s", strings.Join(params.BusinessTypes, ", "))
	if params.Query != "" {
		searchMode = fmt.Sprintf("Free Text Query: %s", params.Query)
	}

	return mcp.NewToolResultText(fmt.Sprintf(
		"Geography search table created successfully!\n\n"+
			"Table ID: %s\n"+
			"Table Name: %s\n"+
			"Location: %.4f, %.4f (radius: %.0f km)\n"+
			"Search Mode: %s\n"+
			"Max Results: %d\n\n"+
			"View in Clay: https://app.clay.com/workspaces/%s/workbooks/%s/tables/%s",
		result.TableID,
		result.TableName,
		params.Latitude, params.Longitude, params.ProximityKm,
		searchMode,
		params.NumResults,
		clayClient.WorkspaceID,
		workbookID,
		result.TableID,
	)), nil
}

func splitAndTrim(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if t := strings.TrimSpace(p); t != "" {
			out = append(out, t)
		}
	}
	return out
}
