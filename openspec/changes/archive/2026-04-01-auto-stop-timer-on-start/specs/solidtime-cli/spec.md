## MODIFIED Requirements

### Requirement: Start Time Entry

The CLI tool SHALL provide a `start` command to begin tracking time.

#### Scenario: Start with description only
- **WHEN** user runs `solidtime-cli start "Working on feature X"`
- **AND** no timer is currently running
- **THEN** the CLI SHALL create a new time entry with start time set to now
- **AND** display confirmation showing entry ID and start time
- **AND** the entry SHALL have end time set to null (running state)

#### Scenario: Start with project
- **WHEN** user runs `solidtime-cli start "Bug fix" --project "PSB-Project"`
- **AND** no timer is currently running
- **THEN** the CLI SHALL lookup the project ID from the project name
- **AND** create a time entry with the resolved project_id
- **AND** display confirmation with project name

#### Scenario: Start with custom start time
- **WHEN** user runs `solidtime-cli start "Forgot to start" --time "09:00"`
- **AND** no timer is currently running
- **THEN** the CLI SHALL parse the time as today at 09:00
- **AND** create a time entry with start time set to 09:00
- **AND** display confirmation showing custom start time

#### Scenario: Start with ISO8601 start time
- **WHEN** user runs `solidtime-cli start "Task" --time "2026-03-31T08:00:00Z"`
- **AND** no timer is currently running
- **THEN** the CLI SHALL parse the full ISO8601 timestamp
- **AND** create a time entry with the specified start time
- **AND** display confirmation

#### Scenario: Start with unknown project
- **WHEN** user runs `solidtime-cli start "Task" --project "NonExistent"`
- **AND** the project name does not match any project in the workspace
- **THEN** the CLI SHALL display an error message
- **AND** list available project names as suggestions
- **AND** exit with code 1

#### Scenario: Start when timer already running - user confirms stop
- **WHEN** user runs `solidtime-cli start "New task"`
- **AND** there is already a running time entry
- **THEN** the CLI SHALL display the currently running entry details (description and elapsed time)
- **AND** prompt the user to confirm stopping the current timer
- **WHEN** user confirms (enters 'y' or 'yes')
- **THEN** the CLI SHALL stop the current timer (set end time to now)
- **AND** display confirmation of stopped timer with duration
- **AND** create a new time entry for "New task" with start time set to now
- **AND** display confirmation of started timer
- **AND** exit with code 0

#### Scenario: Start when timer already running - user declines stop
- **WHEN** user runs `solidtime-cli start "New task"`
- **AND** there is already a running time entry
- **THEN** the CLI SHALL display the currently running entry details
- **AND** prompt the user to confirm stopping the current timer
- **WHEN** user declines (enters 'n', 'no', or presses Enter for default)
- **THEN** the CLI SHALL abort the start command
- **AND** display message indicating the current timer is still running
- **AND** NOT create a new time entry
- **AND** exit with code 0

#### Scenario: Start when timer already running - stop fails
- **WHEN** user runs `solidtime-cli start "New task"`
- **AND** there is already a running time entry
- **AND** user confirms stopping the current timer
- **AND** the API call to stop the current timer fails
- **THEN** the CLI SHALL display an error message with details
- **AND** NOT create a new time entry
- **AND** exit with code 2
