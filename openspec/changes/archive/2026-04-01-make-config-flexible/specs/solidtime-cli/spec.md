## MODIFIED Requirements

### Requirement: Configuration Loading

The CLI tool SHALL load configuration from `config.json` located in XDG-compliant directory with fallback locations.

#### Scenario: Config found in new primary location
- **WHEN** config.json exists in ~/.config/soltty/config.json
- **AND** contains valid api_token and workspace_id fields
- **THEN** the CLI SHALL use these credentials for API calls
- **AND** SHALL use base_url if provided, or default endpoint if omitted

#### Scenario: Config found in XDG location (legacy fallback)
- **WHEN** config.json does not exist in ~/.config/soltty/config.json
- **AND** exists in ~/.config/solidtime/config.json
- **AND** contains valid api_token and workspace_id fields
- **THEN** the CLI SHALL use these credentials for API calls
- **AND** SHALL use base_url if provided, or default endpoint if omitted

#### Scenario: Config found in other fallback location
- **WHEN** config.json does not exist in primary or legacy XDG location
- **AND** exists in fallback location (~/.solidtime/config.json or ./config.json)
- **THEN** the CLI SHALL use the fallback config

#### Scenario: Config with valid base_url
- **WHEN** config.json contains a base_url field
- **THEN** the CLI SHALL use the specified base_url for all API calls
- **AND** SHALL validate that base_url is a valid URL format

#### Scenario: Config missing base_url (required field)
- **WHEN** config.json does not contain a base_url field
- **THEN** the CLI SHALL display an error message
- **AND** show that base_url is a required field
- **AND** provide example value for TechNative Cloud: https://solidtime.tools.technative.cloud/api/v1
- **AND** exit with code 1

#### Scenario: Config with invalid base_url format
- **WHEN** config.json contains a base_url field with invalid URL format
- **THEN** the CLI SHALL display an error message showing the invalid URL
- **AND** show expected format: https://your-instance.com/api/v1
- **AND** exit with code 1

#### Scenario: Config missing required fields
- **WHEN** config.json is found but missing api_token or workspace_id
- **THEN** the CLI SHALL display an error message listing missing fields
- **AND** exit with code 1

#### Scenario: Config file not found
- **WHEN** config.json does not exist in any search location
- **THEN** the CLI SHALL display an error with all searched paths
- **AND** searched paths SHALL include: ~/.config/soltty/, ~/.config/solidtime/, ~/.solidtime/, ./
- **AND** suggest creating config.json in ~/.config/soltty/
- **AND** show setup instructions
- **AND** exit with code 1

## MODIFIED Requirements

### Requirement: List Recent Entries

The CLI tool SHALL provide a `list` command to display recent time entries with clear indication of running timers.

#### Scenario: List entries with running timer
- **WHEN** user runs `soltty list`
- **AND** one of the entries has end time set to null (currently running)
- **THEN** the CLI SHALL display "running" in the duration column for that entry
- **AND** SHALL display calculated duration for all completed entries

#### Scenario: List entries with all completed
- **WHEN** user runs `soltty list`
- **AND** all entries have end times (no running timer)
- **THEN** the CLI SHALL display calculated duration for all entries
- **AND** SHALL NOT display "running" for any entry
