package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <entry-id>",
	Short: "Delete a time entry",
	Long: `Permanently delete a time entry by ID.

Warning: This action cannot be undone.

Example:
  solty list --id                                    # Get entry IDs
  solty delete 01234567-89ab-cdef-0123-456789abcdef  # Delete by ID`,
	Args: cobra.ExactArgs(1),
	Run:  runDelete,
}

func runDelete(cmd *cobra.Command, args []string) {
	entryID := args[0]

	c, err := getClient()
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	// Delete the entry
	err = c.DeleteTimeEntry(entryID)
	if err != nil {
		fmt.Println(formatError(err))
		fmt.Println("\nTip: Use 'solty list --id' to see valid entry IDs")
		return
	}

	fmt.Printf("✓ Time entry deleted: %s\n", entryID)
}
