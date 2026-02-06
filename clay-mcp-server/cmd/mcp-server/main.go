package main

import (
	"fmt"
	"os"

	"clay-mcp/pkg/clay"

	"github.com/mark3labs/mcp-go/server"
)

var clayClient *clay.Client

func main() {
	workspaceID := os.Getenv("CLAY_WORKSPACE_ID")
	sessionCookie := os.Getenv("CLAY_SESSION_COOKIE")
	frontendVersion := os.Getenv("CLAY_FRONTEND_VERSION")

	if workspaceID == "" || sessionCookie == "" {
		fmt.Fprintln(os.Stderr, "Required environment variables:")
		fmt.Fprintln(os.Stderr, "  CLAY_WORKSPACE_ID - Your Clay workspace ID")
		fmt.Fprintln(os.Stderr, "  CLAY_SESSION_COOKIE - Session cookie from browser")
		fmt.Fprintln(os.Stderr, "  CLAY_FRONTEND_VERSION (optional) - Frontend version header")
		os.Exit(1)
	}

	clayClient = clay.NewClient(workspaceID, sessionCookie)
	if frontendVersion != "" {
		clayClient.FrontendVersion = frontendVersion
	}

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
