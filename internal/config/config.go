package config

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// Config represents the Solidtime CLI configuration
type Config struct {
	Username    string `json:"username"`
	APIToken    string `json:"api_token"`
	WorkspaceID string `json:"workspace_id"`
	BaseURL     string `json:"base_url"` // Required: API endpoint URL
}

// Load reads the config from XDG-compliant location with fallbacks
func Load() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	// XDG Base Directory compliant path
	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		configDir = filepath.Join(homeDir, ".config")
	}

	// Search paths in priority order
	searchPaths := []string{
		filepath.Join(configDir, "soltty", "config.json"),        // Primary: ~/.config/soltty/config.json
		filepath.Join(configDir, "solidtime", "config.json"),     // Fallback 1: ~/.config/solidtime/config.json (legacy)
		filepath.Join(homeDir, ".solidtime", "config.json"),      // Fallback 2
		filepath.Join(".", "config.json"),                        // Fallback 3
	}

	var lastErr error
	for _, path := range searchPaths {
		cfg, err := loadFromPath(path)
		if err == nil {
			// Validate required fields
			if cfg.APIToken == "" {
				return nil, fmt.Errorf("missing required field: api_token in %s", path)
			}
			if cfg.WorkspaceID == "" {
				return nil, fmt.Errorf("missing required field: workspace_id in %s", path)
			}
			if cfg.BaseURL == "" {
				return nil, fmt.Errorf("missing required field: base_url in %s\n\nPlease add \"base_url\" to your config.json:\n{\n  ...\n  \"base_url\": \"https://app.example.com/api/v1\"\n}\n\nUse your Solidtime instance URL (e.g., https://app.example.com/api/v1)", path)
			}

			// Validate base_url format
			if err := validateBaseURL(cfg.BaseURL); err != nil {
				return nil, fmt.Errorf("invalid base_url in %s: %w", path, err)
			}

			return cfg, nil
		}
		lastErr = err
	}

	// If we get here, no config was found
	return nil, fmt.Errorf("config.json not found in any of these locations:\n  %s\n  %s\n  %s\n  %s\n\nPlease create config.json in %s with:\n{\n  \"username\": \"Your Name\",\n  \"api_token\": \"your-token\",\n  \"workspace_id\": \"your-workspace-id\",\n  \"base_url\": \"https://app.example.com/api/v1\"\n}\n\nNote: base_url is required. Use your Solidtime instance URL.\n\nLast error: %v",
		searchPaths[0], searchPaths[1], searchPaths[2], searchPaths[3],
		filepath.Join(configDir, "soltty"),
		lastErr)
}

func loadFromPath(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("invalid JSON in %s: %w", path, err)
	}

	return &cfg, nil
}

// validateBaseURL checks if the provided base URL is valid
func validateBaseURL(baseURL string) error {
	// Parse as URL
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("invalid URL format: %w\nExpected format: https://your-solidtime-instance.com/api/v1", err)
	}

	// Check for http/https scheme
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return fmt.Errorf("URL must use http:// or https:// scheme\nExpected format: https://your-solidtime-instance.com/api/v1")
	}

	// Check that host is not empty
	if parsedURL.Host == "" {
		return fmt.Errorf("URL must include a host\nExpected format: https://your-solidtime-instance.com/api/v1")
	}

	// Warn if URL has trailing slash (we'll handle it, but it's a common mistake)
	if strings.HasSuffix(baseURL, "/") {
		return fmt.Errorf("base_url should not have a trailing slash\nGot: %s\nExpected: %s", baseURL, strings.TrimSuffix(baseURL, "/"))
	}

	return nil
}
