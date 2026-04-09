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

	// Fetch projects for name lookup
	projects, err := c.GetProjects()
	if err != nil {
		fmt.Println(formatError(err))
		return
	}
	projectMap := make(map[string]string)
	for _, p := range projects {
		projectMap[p.ID] = p.Name
	}

	// Print header
	if listShowID {
		fmt.Println("ID                                   | Date       | Start | Duration | Project        | Description")
		fmt.Println(strings.Repeat("-", 120))
	} else {
		fmt.Println("ID       | Date       | Start | Duration | Project        | Description")
		fmt.Println(strings.Repeat("-", 90))
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

		// Get project name
		projectName := "No project"
		if entry.ProjectID != nil {
			if name, ok := projectMap[*entry.ProjectID]; ok {
				projectName = name
			}
		}

		if listShowID {
			fmt.Printf("%-36s | %-10s | %-5s | %-8s | %-14s | %s\n",
				entry.ID, date, startTime, duration, projectName, entry.Description)
		} else {
			fmt.Printf("%-8s | %-10s | %-5s | %-8s | %-14s | %s\n",
				entry.ID[:8], date, startTime, duration, projectName, entry.Description)
		}
	}
}
