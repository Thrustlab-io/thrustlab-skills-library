package clay

import "net/http"

type Client struct {
	APIBase           string
	WorkspaceID       string
	SessionCookie     string
	FrontendVersion   string
	HTTPClient        *http.Client
	CurrentWorkbookID string
}

func NewClient(workspaceID, sessionCookie string) *Client {
	return &Client{
		APIBase:         DefaultAPIBase,
		WorkspaceID:     workspaceID,
		SessionCookie:   sessionCookie,
		FrontendVersion: DefaultFrontendVersion,
		HTTPClient:      &http.Client{},
	}
}

func (c *Client) SetWorkspaceID(id string) {
	c.WorkspaceID = id
}

func (c *Client) SetSessionCookie(cookie string) {
	c.SessionCookie = cookie
}

// SwitchWorkspace switches to a different workspace profile at runtime
func (c *Client) SwitchWorkspace(profile *WorkspaceProfile) {
	c.WorkspaceID = profile.WorkspaceID
	c.SessionCookie = profile.SessionCookie
	if profile.FrontendVersion != "" {
		c.FrontendVersion = profile.FrontendVersion
	}
	// Clear current workbook when switching workspaces
	c.CurrentWorkbookID = ""
}

// GetWorkspaceInfo returns current workspace information
func (c *Client) GetWorkspaceInfo() map[string]string {
	return map[string]string{
		"workspace_id":     c.WorkspaceID,
		"frontend_version": c.FrontendVersion,
		"workbook_id":      c.CurrentWorkbookID,
	}
}
