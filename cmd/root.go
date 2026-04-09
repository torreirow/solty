package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/torreirow/soltty/internal/client"
	"github.com/torreirow/soltty/internal/config"
)

// version is set via ldflags during build
var version = "dev"

var rootCmd = &cobra.Command{
	Use:   "soltty",
	Short: "Solidtime CLI - Command-line time tracking",
	Long: `Soltty is a command-line interface for Solidtime time tracking.

Track time directly from your terminal with commands like:
  soltty start "Working on feature"
  soltty stop
  soltty current
  soltty continue <entry-id>
  soltty list
  soltty list clients
  soltty list projects
  soltty web
  soltty info

Configuration is read from ~/.config/soltty/config.json`,
	Version: version,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(currentCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(continueCmd)
	rootCmd.AddCommand(webCmd)
}

// getClient creates an API client from config
func getClient() (*client.Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// BaseURL is validated as required field in config.Load()
	return client.NewClient(cfg.BaseURL, cfg.APIToken, cfg.WorkspaceID), nil
}

// formatError returns a user-friendly error message
func formatError(err error) error {
	return fmt.Errorf("Error: %v", err)
}
