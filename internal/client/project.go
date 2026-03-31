package client

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Project represents a Solidtime project
type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ProjectsResponse wraps the API response
type ProjectsResponse struct {
	Data []Project `json:"data"`
}

// GetProjects fetches all projects in the workspace
func (c *Client) GetProjects() ([]Project, error) {
	endpoint := fmt.Sprintf("organizations/%s/projects", c.workspaceID)
	respBody, err := c.doRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var response ProjectsResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse projects response: %w", err)
	}

	return response.Data, nil
}

// FindProjectByName looks up a project by name (case-insensitive)
func (c *Client) FindProjectByName(name string) (*string, error) {
	projects, err := c.GetProjects()
	if err != nil {
		return nil, err
	}

	nameLower := strings.ToLower(name)
	for _, p := range projects {
		if strings.ToLower(p.Name) == nameLower {
			return &p.ID, nil
		}
	}

	// Build suggestion list
	var suggestions []string
	for _, p := range projects {
		suggestions = append(suggestions, p.Name)
	}

	return nil, fmt.Errorf("project '%s' not found. Available projects: %s",
		name, strings.Join(suggestions, ", "))
}
