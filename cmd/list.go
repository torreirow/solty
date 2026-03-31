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
	Short: "List recent time entries",
	Long: `Display recent time entries in a table format.

Examples:
  solty list
  solty list --limit 5
  solty list --id           # Show entry IDs for deletion`,
	Run: runList,
}

func init() {
	listCmd.Flags().IntVarP(&listLimit, "limit", "l", 10, "Number of entries to show")
	listCmd.Flags().BoolVar(&listShowID, "id", false, "Show entry IDs")
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
		date := entry.Start.Format("2006-01-02")
		startTime := entry.Start.Format("15:04")
		duration := formatDuration(entry.Duration)

		if listShowID {
			fmt.Printf("%-36s | %-10s | %-5s | %-8s | %s\n",
				entry.ID, date, startTime, duration, entry.Description)
		} else {
			fmt.Printf("%-10s | %-5s | %-8s | %s\n",
				date, startTime, duration, entry.Description)
		}
	}
}
