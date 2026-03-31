package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/torreirow/solty/internal/client"
	"github.com/torreirow/solty/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "solty",
	Short: "Solidtime CLI - Command-line time tracking",
	Long: `Solty is a command-line interface for Solidtime time tracking.

Track time directly from your terminal with commands like:
  solty start "Working on feature"
  solty stop
  solty current
  solty list

Configuration is read from ~/.config/solidtime/config.json`,
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
}

// getClient creates an API client from config
func getClient() (*client.Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	return client.NewClient(cfg.APIToken, cfg.WorkspaceID), nil
}

// formatError returns a user-friendly error message
func formatError(err error) error {
	return fmt.Errorf("Error: %v", err)
}
