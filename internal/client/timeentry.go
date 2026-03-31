package client

import (
	"encoding/json"
	"fmt"
	"time"
)

// TimeEntry represents a Solidtime time entry
type TimeEntry struct {
	ID          string     `json:"id"`
	Description string     `json:"description"`
	Start       time.Time  `json:"start"`
	End         *time.Time `json:"end"`
	ProjectID   *string    `json:"project_id"`
	Duration    int        `json:"duration"` // seconds
}

// TimeEntryResponse wraps the API response
type TimeEntryResponse struct {
	Data TimeEntry `json:"data"`
}

// TimeEntriesResponse wraps the API list response
type TimeEntriesResponse struct {
	Data []TimeEntry `json:"data"`
}

// StartTimeEntry creates a new running time entry
func (c *Client) StartTimeEntry(description string, projectID *string, customStart *time.Time) (*TimeEntry, error) {
	memberID, err := c.GetCurrentMemberID()
	if err != nil {
		return nil, fmt.Errorf("failed to get member ID: %w", err)
	}

	startTime := time.Now().UTC()
	if customStart != nil {
		startTime = customStart.UTC()
	}

	payload := map[string]interface{}{
		"description":     description,
		"member_id":       memberID,
		"organization_id": c.workspaceID,
		"start":           startTime.Format(time.RFC3339),
		"billable":        false,
	}

	if projectID != nil {
		payload["project_id"] = *projectID
	}

	endpoint := fmt.Sprintf("organizations/%s/time-entries", c.workspaceID)
	respBody, err := c.doRequest("POST", endpoint, payload)
	if err != nil {
		return nil, err
	}

	var response TimeEntryResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &response.Data, nil
}

// StopTimeEntry stops a running time entry
func (c *Client) StopTimeEntry(entryID string) (*TimeEntry, error) {
	now := time.Now().UTC()
	payload := map[string]interface{}{
		"end": now.Format(time.RFC3339),
	}

	endpoint := fmt.Sprintf("organizations/%s/time-entries/%s", c.workspaceID, entryID)
	respBody, err := c.doRequest("PUT", endpoint, payload)
	if err != nil {
		return nil, err
	}

	var response TimeEntryResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &response.Data, nil
}

// GetCurrentTimeEntry fetches the currently running time entry
func (c *Client) GetCurrentTimeEntry() (*TimeEntry, error) {
	memberID, err := c.GetCurrentMemberID()
	if err != nil {
		return nil, fmt.Errorf("failed to get member ID: %w", err)
	}

	endpoint := fmt.Sprintf("organizations/%s/time-entries?member_id=%s&active=true", c.workspaceID, memberID)
	respBody, err := c.doRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var response TimeEntriesResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if len(response.Data) == 0 {
		return nil, nil
	}

	return &response.Data[0], nil
}

// CreateTimeEntry creates a completed time entry with start and end times
func (c *Client) CreateTimeEntry(description string, start, end time.Time, projectID *string) (*TimeEntry, error) {
	memberID, err := c.GetCurrentMemberID()
	if err != nil {
		return nil, fmt.Errorf("failed to get member ID: %w", err)
	}

	payload := map[string]interface{}{
		"description":     description,
		"member_id":       memberID,
		"organization_id": c.workspaceID,
		"start":           start.UTC().Format(time.RFC3339),
		"end":             end.UTC().Format(time.RFC3339),
		"billable":        false,
	}

	if projectID != nil {
		payload["project_id"] = *projectID
	}

	endpoint := fmt.Sprintf("organizations/%s/time-entries", c.workspaceID)
	respBody, err := c.doRequest("POST", endpoint, payload)
	if err != nil {
		return nil, err
	}

	var response TimeEntryResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &response.Data, nil
}

// ListTimeEntries fetches recent time entries
func (c *Client) ListTimeEntries(limit int) ([]*TimeEntry, error) {
	memberID, err := c.GetCurrentMemberID()
	if err != nil {
		return nil, fmt.Errorf("failed to get member ID: %w", err)
	}

	endpoint := fmt.Sprintf("organizations/%s/time-entries?member_id=%s&limit=%d", c.workspaceID, memberID, limit)
	respBody, err := c.doRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var response TimeEntriesResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	entries := make([]*TimeEntry, len(response.Data))
	for i := range response.Data {
		entries[i] = &response.Data[i]
	}

	return entries, nil
}

// DeleteTimeEntry permanently deletes a time entry
func (c *Client) DeleteTimeEntry(entryID string) error {
	endpoint := fmt.Sprintf("organizations/%s/time-entries/%s", c.workspaceID, entryID)
	_, err := c.doRequest("DELETE", endpoint, nil)
	return err
}
