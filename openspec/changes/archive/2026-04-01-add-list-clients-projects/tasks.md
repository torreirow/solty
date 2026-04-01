## 1. Update Data Structures

- [x] 1.1 Update Project struct in internal/client/project.go to add ClientID and IsArchived fields
- [x] 1.2 Add Client struct in internal/client/project.go with ID, Name, and IsArchived fields
- [x] 1.3 Update ProjectsResponse struct if needed for new fields

## 2. API Client Methods

- [x] 2.1 Add GetClients() method to fetch all clients from API
- [x] 2.2 Update GetProjects() if needed to ensure it returns all project fields
- [x] 2.3 Add ClientsResponse struct for API response parsing

## 3. List Clients Command

- [x] 3.1 Create cmd/list_clients.go with listClientsCmd Cobra command
- [x] 3.2 Implement client fetching logic
- [x] 3.3 Implement archived client filtering
- [x] 3.4 Fetch all projects to count projects per client
- [x] 3.5 Filter archived projects from counts
- [x] 3.6 Implement alphabetical sorting by client name
- [x] 3.7 Implement output formatting with project counts (singular/plural)
- [x] 3.8 Add error handling for API failures
- [x] 3.9 Add "No clients found" message for empty results

## 4. List Projects Command

- [x] 4.1 Create cmd/list_projects.go with listProjectsCmd Cobra command
- [x] 4.2 Add --client/-c flag for client filtering
- [x] 4.3 Implement client fetching and map building (id → name)
- [x] 4.4 Implement project fetching logic
- [x] 4.5 Implement archived project filtering
- [x] 4.6 Implement client name resolution (handle null/unknown client_id)
- [x] 4.7 Implement client filter logic (partial match, case-insensitive)
- [x] 4.8 Implement sorting (by client name, then project name)
- [x] 4.9 Implement table output formatting
- [x] 4.10 Add error handling for API failures
- [x] 4.11 Add "No projects found" message for empty results

## 5. Command Registration

- [x] 5.1 Update cmd/list.go to register listClientsCmd as subcommand
- [x] 5.2 Update cmd/list.go to register listProjectsCmd as subcommand
- [x] 5.3 Ensure backwards compatibility: soltty list runs existing time entries list
- [x] 5.4 Update list command description to mention subcommands

## 6. Root Command Updates

- [x] 6.1 Update root.go Long description to include list clients and list projects examples
- [x] 6.2 Verify help text shows new subcommands

## 7. Testing

- [x] 7.1 Test soltty list clients with real API data
- [x] 7.2 Test soltty list projects with real API data
- [x] 7.3 Test soltty list projects -c with partial match
- [x] 7.4 Test soltty list projects -c with case-insensitive match
- [x] 7.5 Test archived client filtering
- [x] 7.6 Test archived project filtering
- [x] 7.7 Test project count accuracy (excluding archived)
- [x] 7.8 Test sorting is alphabetical
- [x] 7.9 Test projects without client_id show "(no client)"
- [x] 7.10 Test error handling (network failure, auth failure)
- [x] 7.11 Test backwards compatibility: soltty list still shows time entries
- [x] 7.12 Test empty results messages

## 8. Documentation

- [x] 8.1 Update README.md with list clients and list projects examples
- [x] 8.2 Document the -c/--client flag for filtering
- [x] 8.3 Add note about archived items being hidden
