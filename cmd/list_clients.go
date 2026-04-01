package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"github.com/torreirow/soltty/internal/client"
)

var listClientsCmd = &cobra.Command{
	Use:   "clients",
	Short: "List all clients",
	Long:  `Display all active clients with their project counts.`,
	Run:   runListClients,
}

func runListClients(cmd *cobra.Command, args []string) {
	c, err := getClient()
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	// Fetch clients
	clients, err := c.GetClients()
	if err != nil {
		fmt.Println(formatError(fmt.Errorf("failed to fetch clients: %w", err)))
		return
	}

	// Filter archived clients
	var activeClients []client.SolidtimeClient
	for _, cl := range clients {
		if !cl.IsArchived {
			activeClients = append(activeClients, cl)
		}
	}

	// Check if empty
	if len(activeClients) == 0 {
		fmt.Println("No clients found")
		return
	}

	// Fetch projects for counts
	projects, err := c.GetProjects()
	if err != nil {
		fmt.Println(formatError(fmt.Errorf("failed to fetch projects: %w", err)))
		return
	}

	// Count active projects per client
	projectCounts := make(map[string]int)
	for _, p := range projects {
		if !p.IsArchived && p.ClientID != nil {
			projectCounts[*p.ClientID]++
		}
	}

	// Sort clients alphabetically
	sort.Slice(activeClients, func(i, j int) bool {
		return activeClients[i].Name < activeClients[j].Name
	})

	// Display clients with counts
	for _, cl := range activeClients {
		count := projectCounts[cl.ID]
		if count == 1 {
			fmt.Printf("%s (1 project)\n", cl.Name)
		} else {
			fmt.Printf("%s (%d projects)\n", cl.Name, count)
		}
	}
}
