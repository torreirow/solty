## ADDED Requirements

### Requirement: List projects command
The system SHALL provide a `list projects` command that displays all active (non-archived) projects with their client names in a table format.

#### Scenario: User lists all projects
- **WHEN** user runs `soltty list projects`
- **THEN** the system displays all active projects in a table with "Client | Project" columns

#### Scenario: Empty project list
- **WHEN** user runs `soltty list projects` and there are no active projects
- **THEN** the system displays "No projects found"

### Requirement: Archived project filtering
The system SHALL exclude archived projects from the list output.

#### Scenario: Archived projects are hidden
- **WHEN** the API returns projects with `is_archived: true`
- **THEN** the system does not display those projects in the output

#### Scenario: Only archived projects exist
- **WHEN** all projects are archived
- **THEN** the system displays "No projects found"

### Requirement: Client name display
The system SHALL display the client name for each project by resolving the `client_id` field.

#### Scenario: Project with client
- **WHEN** a project has a `client_id` that matches a client
- **THEN** the system displays the client name in the Client column

#### Scenario: Project without client
- **WHEN** a project has `null` or missing `client_id`
- **THEN** the system displays "(no client)" in the Client column

#### Scenario: Client not found
- **WHEN** a project has a `client_id` that doesn't match any client
- **THEN** the system displays "(unknown client)" in the Client column

### Requirement: Sorting
The system SHALL sort projects first by client name alphabetically, then by project name alphabetically (both case-insensitive).

#### Scenario: Projects sorted by client then project name
- **WHEN** projects exist for clients "Zebra" and "Apple"
- **THEN** all "Apple" client projects are shown first, then "Zebra" projects

#### Scenario: Multiple projects per client sorted
- **WHEN** client "TechNative" has projects "Meetings", "General", "Admin"
- **THEN** they are displayed in alphabetical order: "Admin", "General", "Meetings"

### Requirement: Client filter option
The system SHALL support a `-c` or `--client` flag to filter projects by client name using partial, case-insensitive matching.

#### Scenario: Filter by exact client name
- **WHEN** user runs `soltty list projects -c "TechNative"`
- **THEN** the system displays only projects where the client name is "TechNative"

#### Scenario: Filter by partial client name
- **WHEN** user runs `soltty list projects -c "Tech"`
- **THEN** the system displays projects for all clients containing "Tech" (e.g., "TechNative", "FinTech")

#### Scenario: Filter is case-insensitive
- **WHEN** user runs `soltty list projects -c "technative"`
- **THEN** the system displays projects for "TechNative" (case-insensitive match)

#### Scenario: Filter matches no clients
- **WHEN** user runs `soltty list projects -c "NonExistent"`
- **THEN** the system displays "No projects found for client: NonExistent"

#### Scenario: Filter matches multiple clients
- **WHEN** user runs `soltty list projects -c "Tech"` and both "TechNative" and "BioTech" exist
- **THEN** the system displays projects from both clients

### Requirement: Table format
The system SHALL display projects in a table format with header row and column separator.

#### Scenario: Table header displayed
- **WHEN** projects are displayed
- **THEN** the first line shows "Client            | Project"

#### Scenario: Table separator displayed
- **WHEN** projects are displayed
- **THEN** the second line shows a separator like "------------------|---------------------------"

#### Scenario: Table columns aligned
- **WHEN** projects are displayed
- **THEN** client names are left-aligned in a fixed-width column and project names follow the separator

### Requirement: Error handling
The system SHALL display appropriate error messages when API requests fail.

#### Scenario: Clients API failure
- **WHEN** the API request to fetch clients fails
- **THEN** the system displays an error message indicating failure to fetch clients

#### Scenario: Projects API failure
- **WHEN** the API request to fetch projects fails
- **THEN** the system displays an error message indicating failure to fetch projects

#### Scenario: Authentication failure
- **WHEN** the API returns 401 unauthorized
- **THEN** the system displays an error message indicating authentication failed
