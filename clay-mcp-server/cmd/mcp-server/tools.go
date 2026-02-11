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
		mcp.WithString("industries",
			mcp.Description("Comma-separated LinkedIn industries to filter by (e.g., 'Accounting,Computer Software'). See clay://industries for valid values."),
		),
		mcp.WithString("keywords",
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

	// Workspace profile management tools
	addWorkspaceProfileTool := mcp.NewTool("add_workspace_profile",
		mcp.WithDescription("Add or update a named workspace profile for easy switching between workspaces"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Profile name (e.g., 'thrustlab', 'client-acme')"),
		),
		mcp.WithString("workspace_id",
			mcp.Required(),
			mcp.Description("Clay workspace ID"),
		),
		mcp.WithString("session_cookie",
			mcp.Required(),
			mcp.Description("Session cookie value"),
		),
		mcp.WithString("description",
			mcp.Description("Optional description for this workspace"),
		),
		mcp.WithString("frontend_version",
			mcp.Description("Optional frontend version header"),
		),
		mcp.WithBoolean("set_as_default",
			mcp.Description("Set this profile as the default workspace (default: false)"),
		),
	)
	s.AddTool(addWorkspaceProfileTool, addWorkspaceProfileHandler)

	listWorkspaceProfilesTool := mcp.NewTool("list_workspace_profiles",
		mcp.WithDescription("List all saved workspace profiles"),
	)
	s.AddTool(listWorkspaceProfilesTool, listWorkspaceProfilesHandler)

	switchWorkspaceTool := mcp.NewTool("switch_workspace",
		mcp.WithDescription("Switch to a different workspace profile at runtime (no restart required)"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Profile name to switch to"),
		),
	)
	s.AddTool(switchWorkspaceTool, switchWorkspaceHandler)

	removeWorkspaceProfileTool := mcp.NewTool("remove_workspace_profile",
		mcp.WithDescription("Remove a saved workspace profile"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Profile name to remove"),
		),
	)
	s.AddTool(removeWorkspaceProfileTool, removeWorkspaceProfileHandler)

	getCurrentWorkspaceTool := mcp.NewTool("get_current_workspace",
		mcp.WithDescription("Get information about the currently active workspace"),
	)
	s.AddTool(getCurrentWorkspaceTool, getCurrentWorkspaceHandler)
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
	}

	if v, ok := args["industries"].(string); ok && v != "" {
		params.Industries = splitAndTrim(v)
	}
	if v, ok := args["keywords"].(string); ok && v != "" {
		params.Keywords = splitAndTrim(v)
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

// Workspace profile management handlers
func addWorkspaceProfileHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	name := args["name"].(string)
	workspaceID := args["workspace_id"].(string)
	sessionCookie := args["session_cookie"].(string)

	profile := &clay.WorkspaceProfile{
		WorkspaceID:   workspaceID,
		SessionCookie: sessionCookie,
	}

	if v, ok := args["description"].(string); ok && v != "" {
		profile.Description = v
	}
	if v, ok := args["frontend_version"].(string); ok && v != "" {
		profile.FrontendVersion = v
	}

	pm, err := clay.NewProfileManager()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to initialize profile manager: %v", err)), nil
	}

	if err := pm.AddProfile(name, profile); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to add profile: %v", err)), nil
	}

	setAsDefault := false
	if v, ok := args["set_as_default"].(bool); ok {
		setAsDefault = v
	}

	if setAsDefault {
		if err := pm.SetDefault(name); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to set default: %v", err)), nil
		}
	}

	maskedCookie := sessionCookie
	if len(sessionCookie) > 20 {
		maskedCookie = sessionCookie[:10] + "..." + sessionCookie[len(sessionCookie)-10:]
	}

	defaultMsg := ""
	if setAsDefault {
		defaultMsg = "\n\nThis profile has been set as the default workspace."
	}

	return mcp.NewToolResultText(fmt.Sprintf(
		"Workspace profile '%s' saved successfully!\n\n"+
			"Workspace ID: %s\n"+
			"Session Cookie: %s\n"+
			"Description: %s\n\n"+
			"Saved in: ~/.clay/workspaces.json\n"+
			"Use 'switch_workspace' to activate this profile.%s",
		name,
		workspaceID,
		maskedCookie,
		profile.Description,
		defaultMsg,
	)), nil
}

func listWorkspaceProfilesHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	pm, err := clay.NewProfileManager()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to initialize profile manager: %v", err)), nil
	}

	profiles := pm.ListProfiles()
	defaultProfile := pm.GetDefault()
	currentInfo := clayClient.GetWorkspaceInfo()

	if len(profiles) == 0 {
		return mcp.NewToolResultText(
			"No workspace profiles saved.\n\n" +
				"Use 'add_workspace_profile' to save a workspace for easy switching.",
		), nil
	}

	var result strings.Builder
	result.WriteString("Saved Workspace Profiles\n")
	result.WriteString("=========================\n\n")

	for name, profile := range profiles {
		isDefault := name == defaultProfile
		isCurrent := profile.WorkspaceID == currentInfo["workspace_id"]

		result.WriteString(fmt.Sprintf("Profile: %s", name))
		if isDefault {
			result.WriteString(" [DEFAULT]")
		}
		if isCurrent {
			result.WriteString(" [ACTIVE]")
		}
		result.WriteString("\n")

		result.WriteString(fmt.Sprintf("  Workspace ID: %s\n", profile.WorkspaceID))
		if profile.Description != "" {
			result.WriteString(fmt.Sprintf("  Description: %s\n", profile.Description))
		}
		result.WriteString("\n")
	}

	result.WriteString("\nCurrent Active Workspace:\n")
	result.WriteString(fmt.Sprintf("  Workspace ID: %s\n", currentInfo["workspace_id"]))
	if currentInfo["workbook_id"] != "" {
		result.WriteString(fmt.Sprintf("  Current Workbook: %s\n", currentInfo["workbook_id"]))
	}

	return mcp.NewToolResultText(result.String()), nil
}

func switchWorkspaceHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	name := args["name"].(string)

	pm, err := clay.NewProfileManager()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to initialize profile manager: %v", err)), nil
	}

	profile, err := pm.GetProfile(name)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get profile: %v", err)), nil
	}

	clayClient.SwitchWorkspace(profile)

	return mcp.NewToolResultText(fmt.Sprintf(
		"Switched to workspace '%s'!\n\n"+
			"Workspace ID: %s\n"+
			"Description: %s\n\n"+
			"All subsequent operations will use this workspace.\n"+
			"No restart required - the switch is effective immediately.",
		name,
		profile.WorkspaceID,
		profile.Description,
	)), nil
}

func removeWorkspaceProfileHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments.(map[string]any)
	name := args["name"].(string)

	pm, err := clay.NewProfileManager()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to initialize profile manager: %v", err)), nil
	}

	if err := pm.RemoveProfile(name); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to remove profile: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf(
		"Workspace profile '%s' removed successfully!\n\n"+
			"The profile has been deleted from ~/.clay/workspaces.json",
		name,
	)), nil
}

func getCurrentWorkspaceHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	info := clayClient.GetWorkspaceInfo()

	result := fmt.Sprintf(
		"Current Active Workspace\n"+
			"========================\n\n"+
			"Workspace ID: %s\n"+
			"Frontend Version: %s\n",
		info["workspace_id"],
		info["frontend_version"],
	)

	if info["workbook_id"] != "" {
		result += fmt.Sprintf("Current Workbook: %s\n", info["workbook_id"])
		result += fmt.Sprintf("\nView in Clay: https://app.clay.com/workspaces/%s/workbooks/%s",
			info["workspace_id"],
			info["workbook_id"])
	} else {
		result += "\nNo workbook currently selected.\n"
		result += "Create a workbook with 'create_workbook' to get started."
	}

	return mcp.NewToolResultText(result), nil
}
