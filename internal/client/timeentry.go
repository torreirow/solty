package client

import (
	"encoding/json"
	"fmt"
	"strings"
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

// FindEntryByShortID finds a time entry by UUID prefix (short ID)
// Accepts 6-36 characters, returns error for invalid format, not found, or ambiguous matches
func (c *Client) FindEntryByShortID(shortID string) (*TimeEntry, error) {
	// Validate length (minimum 6 characters)
	if len(shortID) < 6 {
		return nil, fmt.Errorf("ID too short: '%s'\nPlease provide at least 6 characters\nUse 'soltty list' to see entry IDs", shortID)
	}

	if len(shortID) > 36 {
		return nil, fmt.Errorf("ID too long: '%s' (maximum 36 characters)", shortID)
	}

	// Validate format (hex + dashes only)
	for _, ch := range shortID {
		if !((ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F') || ch == '-') {
			return nil, fmt.Errorf("Invalid ID format '%s'\nIDs must be 6-36 characters (hex digits and dashes only)\nExample: 985d7cb2", shortID)
		}
	}

	// Fetch last 500 entries for search scope (API limit)
	entries, err := c.ListTimeEntries(500)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch entries: %w", err)
	}

	// Find matches using case-insensitive prefix matching
	var matches []*TimeEntry
	shortIDLower := strings.ToLower(shortID)

	for _, entry := range entries {
		entryIDLower := strings.ToLower(entry.ID)
		if strings.HasPrefix(entryIDLower, shortIDLower) {
			matches = append(matches, entry)
		}
	}

	// Handle different match counts
	switch len(matches) {
	case 0:
		return nil, fmt.Errorf("No entry found with ID '%s'\nUse 'soltty list' to see available entries (with 8-char IDs)\nUse 'soltty list --id' to see full UUIDs", shortID)

	case 1:
		return matches[0], nil

	default:
		// Build ambiguous error message with matching entries
		errMsg := fmt.Sprintf("Ambiguous ID '%s' matches multiple entries:\n", shortID)

		// Show up to 5 matches
		maxShow := 5
		if len(matches) < maxShow {
			maxShow = len(matches)
		}

		for i := 0; i < maxShow; i++ {
			entry := matches[i]
			localStart := entry.Start.Local()
			date := localStart.Format("2006-01-02")
			startTime := localStart.Format("15:04")
			errMsg += fmt.Sprintf("  %s - %s %s: %s\n", entry.ID[:8], date, startTime, entry.Description)
		}

		if len(matches) > maxShow {
			errMsg += fmt.Sprintf("  ... and %d more\n", len(matches)-maxShow)
		}

		errMsg += fmt.Sprintf("\nPlease use more characters (e.g., '%s')\n", matches[0].ID[:12])
		errMsg += "Use 'soltty list --id' to see full UUIDs if needed"

		return nil, fmt.Errorf(errMsg)
	}
}
