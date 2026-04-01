## ADDED Requirements

### Requirement: List clients command
The system SHALL provide a `list clients` command that displays all active (non-archived) clients in alphabetical order with their project counts.

#### Scenario: User lists all clients
- **WHEN** user runs `soltty list clients`
- **THEN** the system displays all active clients alphabetically with project counts in format: "{ClientName} ({count} projects)" or "{ClientName} ({count} project)" for singular

#### Scenario: Empty client list
- **WHEN** user runs `soltty list clients` and there are no active clients
- **THEN** the system displays "No clients found"

### Requirement: Archived client filtering
The system SHALL exclude archived clients from the list output.

#### Scenario: Archived clients are hidden
- **WHEN** the API returns clients with `is_archived: true`
- **THEN** the system does not display those clients in the output

#### Scenario: Only archived clients exist
- **WHEN** all clients are archived
- **THEN** the system displays "No clients found"

### Requirement: Alphabetical sorting
The system SHALL sort clients alphabetically by name (case-insensitive).

#### Scenario: Clients sorted alphabetically
- **WHEN** clients are named "Zebra", "Apple", "Microsoft"
- **THEN** the output shows them in order: "Apple", "Microsoft", "Zebra"

#### Scenario: Case-insensitive sorting
- **WHEN** clients are named "apple", "Banana", "CHERRY"
- **THEN** the output shows them in alphabetical order regardless of case

### Requirement: Project count display
The system SHALL display the number of active (non-archived) projects for each client.

#### Scenario: Client with multiple projects
- **WHEN** a client has 5 active projects
- **THEN** the output shows "{ClientName} (5 projects)"

#### Scenario: Client with one project
- **WHEN** a client has 1 active project
- **THEN** the output shows "{ClientName} (1 project)" (singular)

#### Scenario: Client with no projects
- **WHEN** a client has 0 active projects
- **THEN** the output shows "{ClientName} (0 projects)"

#### Scenario: Client with archived projects only
- **WHEN** a client has 3 projects but all are archived
- **THEN** the output shows "{ClientName} (0 projects)"

### Requirement: Error handling
The system SHALL display appropriate error messages when the API request fails.

#### Scenario: API connection failure
- **WHEN** the API request to fetch clients fails
- **THEN** the system displays an error message indicating the connection failure

#### Scenario: Authentication failure
- **WHEN** the API returns 401 unauthorized
- **THEN** the system displays an error message indicating authentication failed
