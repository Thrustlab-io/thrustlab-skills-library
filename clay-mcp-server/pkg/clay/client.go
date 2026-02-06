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
