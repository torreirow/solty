## ADDED Requirements

### Requirement: Continue existing entry

The CLI tool SHALL provide a `continue` command to start a new timer based on an existing entry's description and project.

#### Scenario: Continue entry by short ID
- **WHEN** user runs `solidtime-cli continue 985d7cb2`
- **AND** an entry with UUID starting with "985d7cb2" exists
- **THEN** the CLI SHALL look up the entry using short ID matching
- **AND** create a new time entry with the same description and project as the found entry
- **AND** set start time to now
- **AND** display confirmation showing the copied description and project

#### Scenario: Continue entry by full UUID
- **WHEN** user runs `solidtime-cli continue 985d7cb2-cb20-40a4-ad9a-627ffa5cdc77`
- **AND** an entry with that UUID exists
- **THEN** the CLI SHALL look up the entry
- **AND** create a new time entry with the same description and project
- **AND** display confirmation

#### Scenario: Continue when timer already running - user confirms
- **WHEN** user runs `solidtime-cli continue <short-id>`
- **AND** there is already a running time entry
- **THEN** the CLI SHALL display the currently running entry details
- **AND** prompt the user to confirm stopping the current timer (same as start command)
- **WHEN** user confirms
- **THEN** the CLI SHALL stop the current timer
- **AND** create a new time entry with description and project from the referenced entry
- **AND** display confirmation

#### Scenario: Continue when timer already running - user declines
- **WHEN** user runs `solidtime-cli continue <short-id>`
- **AND** there is already a running time entry
- **AND** user declines stopping the timer
- **THEN** the CLI SHALL abort the continue command
- **AND** NOT create a new time entry
- **AND** exit with code 0

#### Scenario: Continue entry with no project
- **WHEN** user runs `solidtime-cli continue <short-id>`
- **AND** the referenced entry has no project assigned (project_id is null)
- **THEN** the CLI SHALL create a new time entry with the description only
- **AND** NOT assign a project to the new entry
- **AND** display confirmation showing "No project"

#### Scenario: Continue with short ID not found
- **WHEN** user runs `solidtime-cli continue abc12345`
- **AND** no entry UUID starts with "abc12345"
- **THEN** the CLI SHALL display error: "No entry found with ID 'abc12345'"
- **AND** suggest: "Use 'soltty list' to see available entries (with 8-char IDs)"
- **AND** suggest: "Use 'soltty list --id' to see full UUIDs"
- **AND** exit with code 1

#### Scenario: Continue with ambiguous short ID
- **WHEN** user runs `solidtime-cli continue 985d`
- **AND** multiple entry UUIDs start with "985d"
- **THEN** the CLI SHALL display error: "Ambiguous ID '985d' matches multiple entries:"
- **AND** list matching entries showing: short ID, date, time, and description (up to first 3 matches)
- **AND** suggest: "Please use more characters (e.g., '985d7cb2')"
- **AND** suggest: "Use 'soltty list --id' to see full UUIDs if needed"
- **AND** exit with code 1

#### Scenario: Continue with invalid ID format
- **WHEN** user runs `solidtime-cli continue xyz123`
- **AND** the ID contains non-hex characters (excluding dashes)
- **THEN** the CLI SHALL display error: "Invalid ID format 'xyz123'"
- **AND** show: "IDs must be 6-36 characters (hex digits and dashes only)"
- **AND** show example: "Example: 985d7cb2"
- **AND** exit with code 1

#### Scenario: Continue with ID too short
- **WHEN** user runs `solidtime-cli continue 985d`
- **AND** the ID is less than 6 characters
- **THEN** the CLI SHALL display error: "ID too short: '985d'"
- **AND** show: "Please provide at least 6 characters"
- **AND** suggest: "Use 'soltty list' to see entry IDs"
- **AND** exit with code 1

#### Scenario: Continue without ID argument
- **WHEN** user runs `solidtime-cli continue` with no ID argument
- **THEN** the CLI SHALL display usage information
- **AND** show that entry ID is required
- **AND** show: "Usage: soltty continue <entry-id>"
- **AND** show example: "Example: soltty continue 985d7cb2"
- **AND** exit with code 1
