package cmd

import (
	"fmt"
	"net/url"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"github.com/torreirow/soltty/internal/config"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Open the Solidtime web interface in your browser",
	Long: `Open the Solidtime web interface in your default browser.

The web URL is automatically derived from your configured API endpoint.

Example:
  soltty web`,
	Run: runWeb,
}

func runWeb(cmd *cobra.Command, args []string) {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	// Derive web URL from API base URL
	webURL, err := deriveWebURL(cfg.BaseURL)
	if err != nil {
		fmt.Println(formatError(fmt.Errorf("failed to derive web URL: %v", err)))
		return
	}

	// Open browser
	fmt.Printf("Opening %s in your browser...\n", webURL)
	err = browser.OpenURL(webURL)
	if err != nil {
		fmt.Println(formatError(fmt.Errorf("failed to open browser: %v", err)))
		fmt.Printf("\nYou can manually open: %s\n", webURL)
		return
	}
}

// deriveWebURL extracts the base URL (scheme + host) from the API endpoint
func deriveWebURL(baseURL string) (string, error) {
	parsed, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %v", err)
	}

	if parsed.Scheme == "" || parsed.Host == "" {
		return "", fmt.Errorf("base URL must include scheme and host (e.g., https://example.com)")
	}

	// Return just scheme + host (e.g., https://app.example.com)
	return fmt.Sprintf("%s://%s", parsed.Scheme, parsed.Host), nil
}
