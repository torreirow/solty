package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the current timer",
	Long: `Stop the currently running time entry.

Example:
  solty stop`,
	Run: runStop,
}

func runStop(cmd *cobra.Command, args []string) {
	c, err := getClient()
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	// Get current running entry
	current, err := c.GetCurrentTimeEntry()
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	if current == nil {
		fmt.Println("No timer is currently running")
		return
	}

	// Stop the timer
	entry, err := c.StopTimeEntry(current.ID)
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	duration := formatDuration(entry.Duration)
	fmt.Printf("✓ Timer stopped: \"%s\"\n", entry.Description)
	fmt.Printf("  Duration: %s\n", duration)
	fmt.Printf("  Entry ID: %s\n", entry.ID)
}
