## MODIFIED Requirements

### Requirement: List Recent Entries

The CLI tool SHALL provide a `list` command to display recent time entries with entry IDs, project names, and clear indication of running timers.

#### Scenario: List with default limit and format
- **WHEN** user runs `solidtime-cli list`
- **THEN** the CLI SHALL fetch the 10 most recent time entries
- **AND** display them in a table format with columns: ID (8 chars), Date, Start, Duration, Project, Description
- **AND** format durations as human-readable (e.g., "2h 30m", "45m")
- **AND** display first 8 characters of entry UUID in ID column
- **AND** display project name in Project column (or "No project" if unassigned)

#### Scenario: List with custom limit
- **WHEN** user runs `solidtime-cli list --limit 5`
- **THEN** the CLI SHALL fetch and display only the 5 most recent entries
- **AND** maintain the enhanced format with ID and Project columns

#### Scenario: List with no entries
- **WHEN** user runs `solidtime-cli list`
- **AND** the workspace has no time entries
- **THEN** the CLI SHALL display "No time entries found"
- **AND** exit with code 0

#### Scenario: List with full UUIDs displayed
- **WHEN** user runs `solidtime-cli list --id`
- **THEN** the CLI SHALL display the full 36-character UUID in the ID column
- **AND** maintain all other columns (Date, Start, Duration, Project, Description)
- **AND** expand column width to accommodate full UUID

#### Scenario: List entries with running timer
- **WHEN** user runs `solidtime-cli list`
- **AND** one of the entries has end time set to null (currently running)
- **THEN** the CLI SHALL display "running" in the duration column for that entry
- **AND** SHALL display calculated duration for all completed entries

#### Scenario: List entries with all completed
- **WHEN** user runs `solidtime-cli list`
- **AND** all entries have end times (no running timer)
- **THEN** the CLI SHALL display calculated duration for all entries
- **AND** SHALL NOT display "running" for any entry

#### Scenario: List entries with projects
- **WHEN** user runs `solidtime-cli list`
- **AND** entries have projects assigned
- **THEN** the CLI SHALL display project names in the Project column
- **AND** fetch project data once and map project IDs to names

#### Scenario: List entries without projects
- **WHEN** user runs `solidtime-cli list`
- **AND** some entries have no project assigned (project_id is null)
- **THEN** the CLI SHALL display "No project" in the Project column for those entries
