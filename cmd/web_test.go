package cmd

import (
	"testing"
)

func TestDeriveWebURL(t *testing.T) {
	tests := []struct {
		name        string
		baseURL     string
		expected    string
		expectError bool
	}{
		{
			name:        "Standard API URL",
			baseURL:     "https://app.example.com/api/v1",
			expected:    "https://app.example.com",
			expectError: false,
		},
		{
			name:        "Simple API URL",
			baseURL:     "https://app.solidtime.io/api/v1",
			expected:    "https://app.solidtime.io",
			expectError: false,
		},
		{
			name:        "Custom domain",
			baseURL:     "https://time.company.com/api/v1",
			expected:    "https://time.company.com",
			expectError: false,
		},
		{
			name:        "URL with port",
			baseURL:     "https://localhost:8080/api/v1",
			expected:    "https://localhost:8080",
			expectError: false,
		},
		{
			name:        "HTTP instead of HTTPS",
			baseURL:     "http://localhost:3000/api/v1",
			expected:    "http://localhost:3000",
			expectError: false,
		},
		{
			name:        "URL without path",
			baseURL:     "https://solidtime.io",
			expected:    "https://solidtime.io",
			expectError: false,
		},
		{
			name:        "URL with trailing slash",
			baseURL:     "https://solidtime.io/api/v1/",
			expected:    "https://solidtime.io",
			expectError: false,
		},
		{
			name:        "Invalid URL - no scheme",
			baseURL:     "solidtime.io/api/v1",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Invalid URL - empty string",
			baseURL:     "",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Invalid URL - malformed",
			baseURL:     "ht!tp://invalid url",
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := deriveWebURL(tt.baseURL)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("Expected %q but got %q", tt.expected, result)
				}
			}
		})
	}
}
