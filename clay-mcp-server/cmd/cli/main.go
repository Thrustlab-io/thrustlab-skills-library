package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"clay-mcp/pkg/clay"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	workspaceID := os.Getenv("CLAY_WORKSPACE_ID")
	sessionCookie := os.Getenv("CLAY_SESSION_COOKIE")
	if workspaceID == "" || sessionCookie == "" {
		fmt.Fprintln(os.Stderr, "Set CLAY_WORKSPACE_ID and CLAY_SESSION_COOKIE environment variables")
		os.Exit(1)
	}

	client := clay.NewClient(workspaceID, sessionCookie)

	if v := os.Getenv("CLAY_FRONTEND_VERSION"); v != "" {
		client.FrontendVersion = v
	}

	switch os.Args[1] {
	case "create-workbook":
		cmdCreateWorkbook(client, os.Args[2:])
	case "search-companies":
		cmdSearchCompanies(client, os.Args[2:])
	case "search-businesses":
		cmdSearchBusinesses(client, os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func cmdCreateWorkbook(client *clay.Client, args []string) {
	fs := flag.NewFlagSet("create-workbook", flag.ExitOnError)
	name := fs.String("name", "", "Workbook name (required)")
	folderID := fs.String("folder-id", "", "Folder ID (optional)")
	fs.Parse(args)

	if *name == "" {
		fmt.Fprintln(os.Stderr, "Error: -name is required")
		fs.Usage()
		os.Exit(1)
	}

	result, err := client.CreateWorkbook(clay.CreateWorkbookParams{
		Name:     *name,
		FolderID: *folderID,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created workbook: %s (ID: %s)\n", result.Name, result.ID)
	fmt.Printf("URL: https://app.clay.com/workspaces/%s/workbooks/%s\n", client.WorkspaceID, result.ID)
}

func cmdSearchCompanies(client *clay.Client, args []string) {
	fs := flag.NewFlagSet("search-companies", flag.ExitOnError)
	workbookID := fs.String("workbook-id", "", "Workbook ID (required)")
	industries := fs.String("industries", "", "Comma-separated LinkedIn industries (e.g., Accounting,Computer Software)")
	keywords := fs.String("keywords", "", "Comma-separated description keywords")
	countries := fs.String("countries", "", "Comma-separated countries")
	sizes := fs.String("sizes", "", "Comma-separated size codes (1,2,10,50,200,500,1000,5000,10000)")
	revenues := fs.String("revenues", "", "Comma-separated revenue ranges (e.g., 1M-5M,10M-25M)")
	minMembers := fs.Int("min-linkedin-members", 0, "Minimum number of LinkedIn members")
	maxMembers := fs.Int("max-linkedin-members", 0, "Maximum number of LinkedIn members")
	fs.Parse(args)

	if *workbookID == "" || (*keywords == "" && *industries == "") {
		fmt.Fprintln(os.Stderr, "Error: -workbook-id is required, and at least one of -industries or -keywords must be provided")
		fs.Usage()
		os.Exit(1)
	}

	params := clay.SearchCompaniesParams{
		WorkbookID: *workbookID,
	}
	if *industries != "" {
		params.Industries = splitAndTrim(*industries)
	}
	if *keywords != "" {
		params.Keywords = splitAndTrim(*keywords)
	}
	if *countries != "" {
		params.Countries = splitAndTrim(*countries)
	}
	if *sizes != "" {
		params.CompanySizes = splitAndTrim(*sizes)
	}
	if *revenues != "" {
		params.AnnualRevenues = splitAndTrim(*revenues)
	}
	if *minMembers > 0 {
		params.MinLinkedInMembers = minMembers
	}
	if *maxMembers > 0 {
		params.MaxLinkedInMembers = maxMembers
	}

	result, err := client.SearchCompanies(params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Table created: %s (ID: %s)\n", result.TableName, result.TableID)
	fmt.Printf("Records found: %.0f\n", result.RecordCount)
	fmt.Printf("URL: https://app.clay.com/workspaces/%s/workbooks/%s/tables/%s\n",
		client.WorkspaceID, *workbookID, result.TableID)
}

func cmdSearchBusinesses(client *clay.Client, args []string) {
	fs := flag.NewFlagSet("search-businesses", flag.ExitOnError)
	workbookID := fs.String("workbook-id", "", "Workbook ID (required)")
	lat := fs.Float64("lat", 0, "Latitude (required)")
	lon := fs.Float64("lon", 0, "Longitude (required)")
	radius := fs.Float64("radius", 0, "Radius in km (required)")
	types := fs.String("types", "", "Comma-separated business types")
	query := fs.String("query", "", "Free text search query")
	numResults := fs.Int("num-results", 100, "Maximum number of results")
	tableName := fs.String("table-name", "", "Custom table name")
	fs.Parse(args)

	if *workbookID == "" || *radius <= 0 {
		fmt.Fprintln(os.Stderr, "Error: -workbook-id and -radius are required")
		fs.Usage()
		os.Exit(1)
	}

	params := clay.SearchBusinessesParams{
		WorkbookID:  *workbookID,
		Latitude:    *lat,
		Longitude:   *lon,
		ProximityKm: *radius,
		Query:       *query,
		NumResults:  *numResults,
		TableName:   *tableName,
	}
	if *types != "" {
		params.BusinessTypes = splitAndTrim(*types)
	}

	result, err := client.SearchBusinessesByGeography(params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Table created: %s (ID: %s)\n", result.TableName, result.TableID)
	fmt.Printf("URL: https://app.clay.com/workspaces/%s/workbooks/%s/tables/%s\n",
		client.WorkspaceID, *workbookID, result.TableID)
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

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage: clay-cli <command> [flags]")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Commands:")
	fmt.Fprintln(os.Stderr, "  create-workbook     Create a new Clay workbook")
	fmt.Fprintln(os.Stderr, "  search-companies    Search companies by industry")
	fmt.Fprintln(os.Stderr, "  search-businesses   Search local businesses by geography")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Environment variables:")
	fmt.Fprintln(os.Stderr, "  CLAY_WORKSPACE_ID      Your Clay workspace ID (required)")
	fmt.Fprintln(os.Stderr, "  CLAY_SESSION_COOKIE    Session cookie from browser (required)")
	fmt.Fprintln(os.Stderr, "  CLAY_FRONTEND_VERSION  Frontend version header (optional)")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Run 'clay-cli <command> -h' for command-specific flags.")
}
