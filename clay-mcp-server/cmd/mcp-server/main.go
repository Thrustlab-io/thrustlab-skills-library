package main

import (
	"fmt"
	"os"

	"clay-mcp/pkg/clay"

	"github.com/mark3labs/mcp-go/server"
)

var clayClient *clay.Client

func main() {
	var workspaceID, sessionCookie, frontendVersion string
	var initMethod string

	// Try to initialize from environment variables first (backward compatibility)
	workspaceID = os.Getenv("CLAY_WORKSPACE_ID")
	sessionCookie = os.Getenv("CLAY_SESSION_COOKIE")
	frontendVersion = os.Getenv("CLAY_FRONTEND_VERSION")

	if workspaceID != "" && sessionCookie != "" {
		initMethod = "environment variables"
	} else {
		// Try to load from default workspace profile
		pm, err := clay.NewProfileManager()
		if err == nil {
			profile, err := pm.GetDefaultProfile()
			if err == nil {
				workspaceID = profile.WorkspaceID
				sessionCookie = profile.SessionCookie
				if profile.FrontendVersion != "" {
					frontendVersion = profile.FrontendVersion
				}
				initMethod = fmt.Sprintf("default profile ('%s')", pm.GetDefault())
			}
		}
	}

	// If still no credentials, show error
	if workspaceID == "" || sessionCookie == "" {
		fmt.Fprintln(os.Stderr, "Clay MCP Server requires workspace credentials.")
		fmt.Fprintln(os.Stderr, "\nOption 1: Environment Variables")
		fmt.Fprintln(os.Stderr, "  CLAY_WORKSPACE_ID - Your Clay workspace ID")
		fmt.Fprintln(os.Stderr, "  CLAY_SESSION_COOKIE - Session cookie from browser")
		fmt.Fprintln(os.Stderr, "  CLAY_FRONTEND_VERSION (optional) - Frontend version header")
		fmt.Fprintln(os.Stderr, "\nOption 2: Workspace Profiles")
		fmt.Fprintln(os.Stderr, "  Use add_workspace_profile tool to save credentials")
		fmt.Fprintln(os.Stderr, "  Set a default profile for automatic initialization")
		fmt.Fprintln(os.Stderr, "\nNote: You can use either method. Environment variables take precedence.")
		os.Exit(1)
	}

	clayClient = clay.NewClient(workspaceID, sessionCookie)
	if frontendVersion != "" {
		clayClient.FrontendVersion = frontendVersion
	}

	// Log initialization method to stderr (won't interfere with MCP protocol)
	fmt.Fprintf(os.Stderr, "Clay MCP Server initialized from %s\n", initMethod)
	fmt.Fprintf(os.Stderr, "Workspace ID: %s\n", workspaceID)

	s := server.NewMCPServer(
		"Clay MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(true, false),
	)

	registerTools(s)
	registerResources(s)

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
