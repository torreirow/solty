package cmd

import (
	"fmt"
	"time"
)

// parseTime parses a time string in ISO8601 or HH:MM format
func parseTime(timeStr string) (time.Time, error) {
	// Try ISO8601 format first (e.g., "2026-03-31T14:00:00Z")
	t, err := time.Parse(time.RFC3339, timeStr)
	if err == nil {
		return t, nil
	}

	// Try time-only format (e.g., "14:00"), assume today
	t, err = time.Parse("15:04", timeStr)
	if err == nil {
		now := time.Now()
		return time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), 0, 0, time.Local), nil
	}

	return time.Time{}, fmt.Errorf("invalid time format '%s'. Use ISO8601 (2026-03-31T14:00:00Z) or time (14:00)", timeStr)
}

// formatDuration formats duration in seconds to human-readable format
func formatDuration(seconds int) string {
	if seconds < 60 {
		return fmt.Sprintf("%ds", seconds)
	}

	minutes := seconds / 60
	if minutes < 60 {
		return fmt.Sprintf("%dm", minutes)
	}

	hours := minutes / 60
	mins := minutes % 60
	if mins == 0 {
		return fmt.Sprintf("%dh", hours)
	}
	return fmt.Sprintf("%dh %dm", hours, mins)
}

// formatElapsedTime formats elapsed time from start to now
func formatElapsedTime(start time.Time) string {
	elapsed := int(time.Since(start).Seconds())
	return formatDuration(elapsed)
}
