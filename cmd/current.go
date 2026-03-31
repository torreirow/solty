package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show the current running timer",
	Long: `Display information about the currently running time entry.

Example:
  solty current`,
	Run: runCurrent,
}

func runCurrent(cmd *cobra.Command, args []string) {
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

	elapsed := formatElapsedTime(current.Start)
	fmt.Printf("Timer running: \"%s\"\n", current.Description)
	fmt.Printf("  Started: %s\n", current.Start.Local().Format("15:04"))
	fmt.Printf("  Elapsed: %s\n", elapsed)
	if current.ProjectID != nil {
		fmt.Printf("  Project ID: %s\n", *current.ProjectID)
	}
	fmt.Printf("  Entry ID: %s\n", current.ID)
}
