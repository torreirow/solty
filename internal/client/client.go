package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is the Solidtime API client
type Client struct {
	baseURL     string
	token       string
	workspaceID string
	httpClient  *http.Client
	memberID    string // Cached member ID
}

// NewClient creates a new Solidtime API client
func NewClient(baseURL, token, workspaceID string) *Client {
	return &Client{
		baseURL:     baseURL,
		token:       token,
		workspaceID: workspaceID,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// doRequest performs an HTTP request with authentication
func (c *Client) doRequest(method, endpoint string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	url := fmt.Sprintf("%s/%s", c.baseURL, endpoint)
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Handle HTTP errors
	if resp.StatusCode >= 400 {
		switch resp.StatusCode {
		case 401:
			return nil, fmt.Errorf("authentication failed. Check your API token in config.json")
		case 403:
			return nil, fmt.Errorf("permission denied")
		case 404:
			return nil, fmt.Errorf("resource not found")
		case 422:
			return nil, fmt.Errorf("validation error: %s", string(respBody))
		default:
			return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, string(respBody))
		}
	}

	return respBody, nil
}

// GetCurrentMemberID fetches and caches the current user's member ID
func (c *Client) GetCurrentMemberID() (string, error) {
	if c.memberID != "" {
		return c.memberID, nil
	}

	endpoint := fmt.Sprintf("organizations/%s/members", c.workspaceID)
	respBody, err := c.doRequest("GET", endpoint, nil)
	if err != nil {
		return "", err
	}

	var response struct {
		Data []struct {
			ID string `json:"id"`
		} `json:"data"`
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return "", fmt.Errorf("failed to parse members response: %w", err)
	}

	if len(response.Data) == 0 {
		return "", fmt.Errorf("no members found in workspace")
	}

	c.memberID = response.Data[0].ID
	return c.memberID, nil
}
