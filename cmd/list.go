package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	listLimit  int
	listShowID bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List time entries, clients, or projects",
	Long: `Display time entries, clients, or projects.

Subcommands:
  soltty list              # List recent time entries (default)
  soltty list clients      # List all clients with project counts
  soltty list projects     # List all projects with client names
  soltty list projects -c <client>  # Filter projects by client

Examples:
  soltty list
  soltty list --limit 5
  soltty list --id           # Show entry IDs for deletion
  soltty list clients
  soltty list projects
  soltty list projects -c Acme`,
	Run: runList,
}

func init() {
	listCmd.Flags().IntVarP(&listLimit, "limit", "l", 10, "Number of entries to show")
	listCmd.Flags().BoolVar(&listShowID, "id", false, "Show entry IDs")

	// Register subcommands
	listCmd.AddCommand(listClientsCmd)
	listCmd.AddCommand(listProjectsCmd)
}

func runList(cmd *cobra.Command, args []string) {
	c, err := getClient()
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	entries, err := c.ListTimeEntries(listLimit)
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	if len(entries) == 0 {
		fmt.Println("No time entries found")
		return
	}

	// Print header
	if listShowID {
		fmt.Println("ID                                   | Date       | Start | Duration | Description")
		fmt.Println(strings.Repeat("-", 100))
	} else {
		fmt.Println("Date       | Start | Duration | Description")
		fmt.Println(strings.Repeat("-", 60))
	}

	// Print entries
	for _, entry := range entries {
		localStart := entry.Start.Local()
		date := localStart.Format("2006-01-02")
		startTime := localStart.Format("15:04")

		// Check if timer is currently running (no end time)
		var duration string
		if entry.End == nil {
			duration = "running"
		} else {
			duration = formatDuration(entry.Duration)
		}

		if listShowID {
			fmt.Printf("%-36s | %-10s | %-5s | %-8s | %s\n",
				entry.ID, date, startTime, duration, entry.Description)
		} else {
			fmt.Printf("%-10s | %-5s | %-8s | %s\n",
				date, startTime, duration, entry.Description)
		}
	}
}
