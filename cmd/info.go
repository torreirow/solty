package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/torreirow/soltty/internal/config"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show account information and IDs",
	Long: `Display your workspace ID, member ID, and other account information.
Useful for configuring browser extensions and integrations.`,
	Run: runInfo,
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

func runInfo(cmd *cobra.Command, args []string) {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	c, err := getClient()
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	// Get member ID
	memberID, err := c.GetCurrentMemberID()
	if err != nil {
		fmt.Println(formatError(fmt.Errorf("failed to get member ID: %w", err)))
		return
	}

	// Get member details
	memberInfo, err := getMemberInfo(c, cfg.WorkspaceID, memberID)
	if err != nil {
		fmt.Println(formatError(fmt.Errorf("failed to get member info: %w", err)))
		return
	}

	// Display information
	fmt.Println("Account Information:")
	fmt.Println("═══════════════════")
	fmt.Printf("Username:     %s\n", cfg.Username)
	fmt.Printf("Member ID:    %s\n", memberID)
	fmt.Printf("Workspace ID: %s\n", cfg.WorkspaceID)
	if memberInfo.UserID != "" {
		fmt.Printf("User ID:      %s\n", memberInfo.UserID)
	}
	fmt.Printf("API Endpoint: %s\n", cfg.BaseURL)
}

type memberInfoResponse struct {
	UserID string
}

func getMemberInfo(c interface{}, workspaceID, memberID string) (*memberInfoResponse, error) {
	// Type assert to access doRequest
	type requester interface {
		GetCurrentMemberID() (string, error)
	}

	// For now, just return basic info
	// We can extend this later if we need more details from the API
	return &memberInfoResponse{
		UserID: "", // Will be populated if we fetch from API
	}, nil
}
