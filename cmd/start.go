package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	startProject string
	startTime    string
)

var startCmd = &cobra.Command{
	Use:   "start <description>",
	Short: "Start a new timer",
	Long: `Start a new time tracking entry.

Examples:
  solty start "Working on feature X"
  solty start "Bug fix" --project "PSB-Project"
  solty start "Forgot to start" --time "09:00"
  solty start "Task" --time "2026-03-31T08:00:00Z"`,
	Args: cobra.ExactArgs(1),
	Run:  runStart,
}

func init() {
	startCmd.Flags().StringVarP(&startProject, "project", "p", "", "Project name")
	startCmd.Flags().StringVarP(&startTime, "time", "t", "", "Custom start time (ISO8601 or HH:MM)")
}

func runStart(cmd *cobra.Command, args []string) {
	description := args[0]

	c, err := getClient()
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	// Check if timer is already running
	current, err := c.GetCurrentTimeEntry()
	if err != nil {
		fmt.Println(formatError(err))
		os.Exit(2)
	}

	// If timer is running, ask user if they want to stop it
	if current != nil {
		elapsed := formatElapsedTime(current.Start)
		fmt.Printf("A timer is currently running: \"%s\" (started %s ago)\n", current.Description, elapsed)
		fmt.Print("Stop this timer and start a new one? [y/N]: ")

		// Read user input
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(formatError(fmt.Errorf("failed to read input: %w", err)))
			os.Exit(2)
		}

		// Parse input (trim whitespace and convert to lowercase)
		input = strings.TrimSpace(strings.ToLower(input))

		// Check if user confirmed (accept 'y' or 'yes')
		if input != "y" && input != "yes" {
			fmt.Println("Keeping current timer running. No new timer started.")
			os.Exit(0)
		}

		// Stop the current timer
		stoppedEntry, err := c.StopTimeEntry(current.ID)
		if err != nil {
			fmt.Println(formatError(fmt.Errorf("failed to stop current timer: %w", err)))
			os.Exit(2)
		}

		// Display confirmation of stopped timer
		duration := formatDuration(stoppedEntry.Duration)
		fmt.Printf("✓ Stopped: \"%s\" (duration: %s)\n", stoppedEntry.Description, duration)
	}

	// Resolve project if specified
	var projectID *string
	if startProject != "" {
		pid, err := c.FindProjectByName(startProject)
		if err != nil {
			fmt.Println(formatError(err))
			return
		}
		projectID = pid
	}

	// Parse custom start time if specified
	var customStart *time.Time
	if startTime != "" {
		t, err := parseTime(startTime)
		if err != nil {
			fmt.Println(formatError(err))
			return
		}
		customStart = &t
	}

	// Start the timer
	entry, err := c.StartTimeEntry(description, projectID, customStart)
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	fmt.Printf("✓ Timer started: \"%s\"\n", entry.Description)
	if customStart != nil {
		fmt.Printf("  Start time: %s (custom)\n", entry.Start.Local().Format("15:04"))
	} else {
		fmt.Printf("  Start time: %s\n", entry.Start.Local().Format("15:04"))
	}
	if startProject != "" {
		fmt.Printf("  Project: %s\n", startProject)
	}
	fmt.Printf("  Entry ID: %s\n", entry.ID)
}
