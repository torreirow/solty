package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/torreirow/soltty/internal/client"
)

var (
	listProjectsClientFilter string
)

var listProjectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "List all projects",
	Long:  `Display all active projects with their client names.`,
	Run:   runListProjects,
}

func init() {
	listProjectsCmd.Flags().StringVarP(&listProjectsClientFilter, "client", "c", "", "Filter projects by client name (partial match)")
}

func runListProjects(cmd *cobra.Command, args []string) {
	c, err := getClient()
	if err != nil {
		fmt.Println(formatError(err))
		return
	}

	// Fetch clients for name lookup
	clients, err := c.GetClients()
	if err != nil {
		fmt.Println(formatError(fmt.Errorf("failed to fetch clients: %w", err)))
		return
	}

	// Build client ID -> name map
	clientMap := make(map[string]string)
	for _, cl := range clients {
		clientMap[cl.ID] = cl.Name
	}

	// Fetch projects
	projects, err := c.GetProjects()
	if err != nil {
		fmt.Println(formatError(fmt.Errorf("failed to fetch projects: %w", err)))
		return
	}

	// Filter archived projects
	var activeProjects []client.Project
	for _, p := range projects {
		if !p.IsArchived {
			activeProjects = append(activeProjects, p)
		}
	}

	// Apply client filter if provided
	if listProjectsClientFilter != "" {
		var filteredProjects []client.Project
		filterLower := strings.ToLower(listProjectsClientFilter)

		for _, p := range activeProjects {
			if p.ClientID != nil {
				clientName := clientMap[*p.ClientID]
				if strings.Contains(strings.ToLower(clientName), filterLower) {
					filteredProjects = append(filteredProjects, p)
				}
			}
		}

		activeProjects = filteredProjects

		if len(activeProjects) == 0 {
			fmt.Printf("No projects found for client: %s\n", listProjectsClientFilter)
			return
		}
	}

	// Check if empty
	if len(activeProjects) == 0 {
		fmt.Println("No projects found")
		return
	}

	// Sort by client name, then project name
	sort.Slice(activeProjects, func(i, j int) bool {
		clientNameI := "(no client)"
		clientNameJ := "(no client)"

		if activeProjects[i].ClientID != nil {
			if name, ok := clientMap[*activeProjects[i].ClientID]; ok {
				clientNameI = name
			} else {
				clientNameI = "(unknown client)"
			}
		}

		if activeProjects[j].ClientID != nil {
			if name, ok := clientMap[*activeProjects[j].ClientID]; ok {
				clientNameJ = name
			} else {
				clientNameJ = "(unknown client)"
			}
		}

		if clientNameI != clientNameJ {
			return clientNameI < clientNameJ
		}
		return activeProjects[i].Name < activeProjects[j].Name
	})

	// Display table header
	fmt.Println("Client            | Project")
	fmt.Println(strings.Repeat("-", 18) + "|" + strings.Repeat("-", 30))

	// Display projects
	for _, p := range activeProjects {
		clientName := "(no client)"
		if p.ClientID != nil {
			if name, ok := clientMap[*p.ClientID]; ok {
				clientName = name
			} else {
				clientName = "(unknown client)"
			}
		}

		fmt.Printf("%-17s | %s\n", clientName, p.Name)
	}
}
