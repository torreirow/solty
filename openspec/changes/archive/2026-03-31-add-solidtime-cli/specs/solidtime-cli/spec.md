# Capability: solidtime-cli

Command-line interface for real-time time tracking operations with Solidtime.

## ADDED Requirements

### Requirement: Configuration Loading

The CLI tool SHALL load configuration from `config.json` located in XDG-compliant directory with fallback locations.

#### Scenario: Config found in XDG location
- **WHEN** config.json exists in ~/.config/solidtime/config.json
- **AND** contains valid api_token and workspace_id fields
- **THEN** the CLI SHALL use these credentials for API calls

#### Scenario: Config found in fallback location
- **WHEN** config.json does not exist in primary location
- **AND** exists in fallback location (~/.solidtime/config.json or ./config.json)
- **THEN** the CLI SHALL use the fallback config

#### Scenario: Config missing required fields
- **WHEN** config.json is found but missing api_token or workspace_id
- **THEN** the CLI SHALL display an error message listing missing fields
- **AND** exit with code 1

#### Scenario: Config file not found
- **WHEN** config.json does not exist in any search location
- **THEN** the CLI SHALL display an error with all searched paths
- **AND** suggest creating config.json in ~/.config/solidtime/
- **AND** show setup instructions
- **AND** exit with code 1

### Requirement: Start Time Entry

The CLI tool SHALL provide a `start` command to begin tracking time.

#### Scenario: Start with description only
- **WHEN** user runs `solidtime-cli start "Working on feature X"`
- **THEN** the CLI SHALL create a new time entry with start time set to now
- **AND** display confirmation showing entry ID and start time
- **AND** the entry SHALL have end time set to null (running state)

#### Scenario: Start with project
- **WHEN** user runs `solidtime-cli start "Bug fix" --project "PSB-Project"`
- **THEN** the CLI SHALL lookup the project ID from the project name
- **AND** create a time entry with the resolved project_id
- **AND** display confirmation with project name

#### Scenario: Start with custom start time
- **WHEN** user runs `solidtime-cli start "Forgot to start" --time "09:00"`
- **THEN** the CLI SHALL parse the time as today at 09:00
- **AND** create a time entry with start time set to 09:00
- **AND** display confirmation showing custom start time

#### Scenario: Start with ISO8601 start time
- **WHEN** user runs `solidtime-cli start "Task" --time "2026-03-31T08:00:00Z"`
- **THEN** the CLI SHALL parse the full ISO8601 timestamp
- **AND** create a time entry with the specified start time
- **AND** display confirmation

#### Scenario: Start with unknown project
- **WHEN** user runs `solidtime-cli start "Task" --project "NonExistent"`
- **AND** the project name does not match any project in the workspace
- **THEN** the CLI SHALL display an error message
- **AND** list available project names as suggestions
- **AND** exit with code 1

#### Scenario: Start when timer already running
- **WHEN** user runs `solidtime-cli start "New task"`
- **AND** there is already a running time entry
- **THEN** the CLI SHALL display an error message indicating a timer is already active
- **AND** show the description of the currently running entry
- **AND** exit with code 1

### Requirement: Stop Time Entry

The CLI tool SHALL provide a `stop` command to end the currently running timer.

#### Scenario: Stop running timer
- **WHEN** user runs `solidtime-cli stop`
- **AND** there is a running time entry
- **THEN** the CLI SHALL set the end time to now
- **AND** update the time entry via API
- **AND** display the stopped entry with description and total duration

#### Scenario: Stop with no running timer
- **WHEN** user runs `solidtime-cli stop`
- **AND** there is no running time entry
- **THEN** the CLI SHALL display a message "No timer is currently running"
- **AND** exit with code 0 (not an error condition)

### Requirement: Show Current Entry

The CLI tool SHALL provide a `current` command to display the active time entry.

#### Scenario: Display running timer
- **WHEN** user runs `solidtime-cli current`
- **AND** there is a running time entry
- **THEN** the CLI SHALL display the entry description
- **AND** show the project name (if assigned)
- **AND** show the start time
- **AND** show the elapsed duration updated in real-time format (e.g., "2h 15m")

#### Scenario: No timer running
- **WHEN** user runs `solidtime-cli current`
- **AND** there is no running time entry
- **THEN** the CLI SHALL display "No timer is currently running"
- **AND** exit with code 0

### Requirement: Add Completed Entry

The CLI tool SHALL provide an `add` command to create time entries with specific start and end times.

#### Scenario: Add entry with explicit times
- **WHEN** user runs `solidtime-cli add "Meeting" --start "2026-03-31T14:00:00Z" --end "2026-03-31T15:30:00Z"`
- **THEN** the CLI SHALL create a completed time entry
- **AND** calculate duration as 90 minutes (5400 seconds)
- **AND** display confirmation with calculated duration

#### Scenario: Add entry with time-only format
- **WHEN** user runs `solidtime-cli add "Code review" --start "14:00" --end "15:30"`
- **THEN** the CLI SHALL interpret times as today in local timezone
- **AND** create entry with full ISO8601 timestamps
- **AND** display confirmation

#### Scenario: Add entry with project
- **WHEN** user runs `solidtime-cli add "Sprint planning" --start "10:00" --end "12:00" --project "TN-Meetings"`
- **THEN** the CLI SHALL resolve project name to ID
- **AND** create entry with project attached
- **AND** display confirmation with all metadata

#### Scenario: Add entry with end before start
- **WHEN** user runs `solidtime-cli add "Task" --start "15:00" --end "14:00"`
- **THEN** the CLI SHALL display an error "End time must be after start time"
- **AND** exit with code 1

#### Scenario: Add entry with missing time flags
- **WHEN** user runs `solidtime-cli add "Task" --start "14:00"`
- **AND** no --end flag is provided
- **THEN** the CLI SHALL display an error "Both --start and --end are required for add command"
- **AND** exit with code 1

### Requirement: List Recent Entries

The CLI tool SHALL provide a `list` command to display recent time entries.

#### Scenario: List with default limit
- **WHEN** user runs `solidtime-cli list`
- **THEN** the CLI SHALL fetch the 10 most recent time entries
- **AND** display them in a table format with columns: start time, duration, description, project
- **AND** format durations as human-readable (e.g., "2h 30m", "45m")

#### Scenario: List with custom limit
- **WHEN** user runs `solidtime-cli list --limit 5`
- **THEN** the CLI SHALL fetch and display only the 5 most recent entries

#### Scenario: List with no entries
- **WHEN** user runs `solidtime-cli list`
- **AND** the workspace has no time entries
- **THEN** the CLI SHALL display "No time entries found"
- **AND** exit with code 0

#### Scenario: List with IDs displayed
- **WHEN** user runs `solidtime-cli list --id`
- **THEN** the CLI SHALL display an additional ID column
- **AND** show the full UUID for each entry
- **AND** maintain all other columns (start time, duration, description, project)

### Requirement: Delete Time Entry

The CLI tool SHALL provide a `delete` command to permanently remove a time entry by ID.

#### Scenario: Delete existing entry
- **WHEN** user runs `solidtime-cli delete <entry-id>`
- **AND** the entry exists in the workspace
- **THEN** the CLI SHALL send DELETE request to the API
- **AND** display confirmation message with the deleted entry ID
- **AND** exit with code 0

#### Scenario: Delete non-existent entry
- **WHEN** user runs `solidtime-cli delete <entry-id>`
- **AND** the entry does not exist
- **THEN** the CLI SHALL display an error message
- **AND** suggest using `list --id` to find valid entry IDs
- **AND** exit with code 1

#### Scenario: Delete with invalid UUID format
- **WHEN** user runs `solidtime-cli delete <invalid-id>`
- **AND** the ID is not a valid UUID format
- **THEN** the CLI SHALL display an error message about invalid ID format
- **AND** exit with code 1

#### Scenario: Delete without entry ID
- **WHEN** user runs `solidtime-cli delete` with no entry ID argument
- **THEN** the CLI SHALL display usage information
- **AND** show that entry ID is required
- **AND** exit with code 1

### Requirement: API Authentication

The CLI tool SHALL authenticate all API requests using the Bearer token from configuration.

#### Scenario: Successful authentication
- **WHEN** the CLI makes any API request
- **THEN** it SHALL include the Authorization header with value "Bearer {api_token}"
- **AND** include the Accept header with value "application/vnd.api+json"
- **AND** construct the URL using the workspace_id from config

#### Scenario: Authentication failure
- **WHEN** the API returns 401 Unauthorized
- **THEN** the CLI SHALL display "Authentication failed. Check your API token in config.json"
- **AND** exit with code 1

#### Scenario: Network error
- **WHEN** the API request fails due to network issues
- **THEN** the CLI SHALL display a user-friendly error message
- **AND** suggest checking network connectivity
- **AND** exit with code 2

### Requirement: Help and Usage

The CLI tool SHALL provide built-in help for all commands.

#### Scenario: Display global help
- **WHEN** user runs `solidtime-cli --help` or `solidtime-cli`
- **THEN** the CLI SHALL display a list of available commands
- **AND** show brief description for each command
- **AND** show global flags

#### Scenario: Display command-specific help
- **WHEN** user runs `solidtime-cli start --help`
- **THEN** the CLI SHALL display usage syntax for start command
- **AND** list all available flags with descriptions
- **AND** provide usage examples

### Requirement: Version Information

The CLI tool SHALL provide version information.

#### Scenario: Display version
- **WHEN** user runs `solidtime-cli --version` or `solidtime-cli version`
- **THEN** the CLI SHALL display the version number
- **AND** exit with code 0

#### Scenario: Version embedded at build time
- **WHEN** the CLI is built with ldflags
- **THEN** the version SHALL be embedded from the VERSION file
- **AND** the version SHALL match the git tag

### Requirement: Error Messages

The CLI tool SHALL provide clear, actionable error messages.

#### Scenario: User-friendly error formatting
- **WHEN** any error occurs
- **THEN** the CLI SHALL display the error in a consistent format
- **AND** avoid exposing raw HTTP responses or stack traces
- **AND** suggest corrective actions when applicable

#### Scenario: Validation error
- **WHEN** user provides invalid input (e.g., malformed time format)
- **THEN** the CLI SHALL display what was invalid
- **AND** show the expected format
- **AND** exit with code 1

### Requirement: Timezone Handling

The CLI tool SHALL handle timezone conversions between user's local time and API UTC requirements.

#### Scenario: Send times to API in UTC
- **WHEN** the CLI creates or updates a time entry
- **THEN** it SHALL convert all timestamps to UTC before sending to API
- **AND** format times as ISO8601/RFC3339
- **AND** include billable field (default: false)

#### Scenario: Display times in local timezone
- **WHEN** the CLI displays time information to the user
- **THEN** it SHALL convert UTC times from API to user's local timezone
- **AND** format times in user-friendly format (HH:MM or YYYY-MM-DD HH:MM)

#### Scenario: User inputs local time
- **WHEN** user provides time in HH:MM format (e.g., "14:00")
- **THEN** the CLI SHALL interpret it as local timezone
- **AND** convert to UTC for API submission

#### Scenario: User inputs ISO8601 time
- **WHEN** user provides time in ISO8601 format with timezone
- **THEN** the CLI SHALL preserve the timezone information
- **AND** convert to UTC for API submission
