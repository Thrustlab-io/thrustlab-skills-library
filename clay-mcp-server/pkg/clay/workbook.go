package clay

import (
	"encoding/json"
	"fmt"
)

func (c *Client) CreateWorkbook(params CreateWorkbookParams) (*CreateWorkbookResult, error) {
	payload := map[string]any{
		"name":        params.Name,
		"workspaceId": c.WorkspaceID,
	}
	if params.FolderID != "" {
		payload["folderId"] = params.FolderID
	}

	url := fmt.Sprintf("%s/workbooks", c.APIBase)
	resp, err := c.makeRequest("POST", url, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to create workbook: %w\nResponse: %s", err, string(resp))
	}

	var result map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	id, _ := result["id"].(string)
	name, _ := result["name"].(string)

	c.CurrentWorkbookID = id

	return &CreateWorkbookResult{ID: id, Name: name}, nil
}
