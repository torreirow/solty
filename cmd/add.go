package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	addStart   string
	addEnd     string
	addProject string
)

var addCmd = &cobra.Command{
	Use:   "add <description>",
	Short: "Add a completed time entry",
	Long: `Create a time entry with specific start and end times.

Examples:
  solty add "Meeting" --start "2026-03-31T14:00:00Z" --end "2026-03-31T15:30:00Z"
  solty add "Code review" --start "14:00" --end "15:30"
  solty add "Sprint planning" --start "10:00" --end "12:00" --project "TN-Meetings"`,
	Args: cobra.ExactArgs(1),
	Run:  runAdd,
}

func init() {
	addCmd.Flags().StringVar(&addStart, "start", "", "Start time (ISO8601 or HH:MM) [required]")
	addCmd.Flags().StringVar(&addEnd, "end", "", "End time (ISO8601 or HH:MM) [required]")
	addCmd.Flags().StringVarP(&addProject, "project", "p", "", "Project name")
	addCmd.MarkFlagRequired("start")
	addCmd.MarkFlagRequired("end")
}

func runAdd(cmd *cobra.Command, args []string) {
	description := args[0]

	// Parse times
	startTime, err := parseTime(addStart)
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	endTime, err := parseTime(addEnd)
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	// Validate: end must be after start
	if endTime.Before(startTime) || endTime.Equal(startTime) {
		fmt.Println("Error: End time must be after start time")
		return
	}

	c, err := getClient()
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	// Resolve project if specified
	var projectID *string
	if addProject != "" {
		pid, err := c.FindProjectByName(addProject)
		if err != nil {
			fmt.Println(formatError(err))
			return
		}
		projectID = pid
	}

	// Create the entry
	entry, err := c.CreateTimeEntry(description, startTime, endTime, projectID)
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	duration := formatDuration(entry.Duration)
	fmt.Printf("✓ Time entry added: \"%s\"\n", entry.Description)
	fmt.Printf("  Start: %s\n", entry.Start.Format("15:04"))
	fmt.Printf("  End: %s\n", entry.End.Format("15:04"))
	fmt.Printf("  Duration: %s\n", duration)
	if addProject != "" {
		fmt.Printf("  Project: %s\n", addProject)
	}
	fmt.Printf("  Entry ID: %s\n", entry.ID)
}
