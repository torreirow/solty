package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config represents the Solidtime CLI configuration
type Config struct {
	Username    string `json:"username"`
	APIToken    string `json:"api_token"`
	WorkspaceID string `json:"workspace_id"`
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
		filepath.Join(configDir, "solidtime", "config.json"),     // Primary: ~/.config/solidtime/config.json
		filepath.Join(homeDir, ".solidtime", "config.json"),      // Fallback 1
		filepath.Join(".", "config.json"),                        // Fallback 2
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
			return cfg, nil
		}
		lastErr = err
	}

	// If we get here, no config was found
	return nil, fmt.Errorf("config.json not found in any of these locations:\n  %s\n  %s\n  %s\n\nPlease create config.json in %s with:\n{\n  \"username\": \"Your Name\",\n  \"api_token\": \"your-token\",\n  \"workspace_id\": \"your-workspace-id\"\n}\n\nLast error: %v",
		searchPaths[0], searchPaths[1], searchPaths[2],
		filepath.Join(configDir, "solidtime"),
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
