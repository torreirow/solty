# solidtime-export Specification

## Purpose
This specification defines the date selection functionality for the Solidtime export script (`solidtimexport.py`). The script exports time entries from the Solidtime API to CSV format, supporting both automated (command-line) and interactive modes for flexible date range selection. This capability enables weekly time entry exports, custom date ranges, and ISO week number selections.
## Requirements
### Requirement: Command-Line Date Selection
The solidtimexport.py script SHALL support command-line arguments for non-interactive date selection to enable automation and scripting.

#### Scenario: Explicit start and end dates
- **WHEN** user runs `./solidtimexport.py --start 2025-01-20 --end 2025-01-26`
- **THEN** the script uses those dates without prompting
- **AND** displays "Using dates: 2025-01-20 to 2025-01-26" before processing

#### Scenario: Last week shortcut
- **WHEN** user runs `./solidtimexport.py --last-week` on 2025-02-05 (Wednesday)
- **THEN** the script calculates previous week as 2025-01-27 (Monday) to 2025-02-02 (Sunday)
- **AND** displays the calculated date range before processing

#### Scenario: ISO week number
- **WHEN** user runs `./solidtimexport.py --week 5`
- **THEN** the script calculates week 5 date range (Monday to Sunday of that week)
- **AND** displays the calculated date range before processing

#### Scenario: Invalid date range
- **WHEN** user provides start date after end date (e.g., `--start 2025-01-26 --end 2025-01-20`)
- **THEN** the script prints error "Start date must be before end date"
- **AND** exits with code 1

### Requirement: Interactive Menu Mode
When no command-line arguments are provided, the script SHALL present an interactive menu with multiple date selection options.

#### Scenario: Interactive menu presentation
- **WHEN** user runs `./solidtimexport.py` without arguments
- **THEN** the script displays a menu with three options:
  - [1] Last week (automatic calculation)
  - [2] Enter custom start/end dates
  - [3] Enter specific week number
- **AND** prompts user to select an option

#### Scenario: Menu option 1 - Last week
- **WHEN** user selects option 1 in interactive mode
- **THEN** the script calculates previous Monday-Sunday automatically
- **AND** displays the calculated dates
- **AND** prompts "Continue with these dates? (y/n)"
- **AND** proceeds only if user confirms with 'y'

#### Scenario: Menu option 2 - Custom dates
- **WHEN** user selects option 2 in interactive mode
- **THEN** the script prompts for start date (YYYY-MM-DD)
- **AND** prompts for end date (YYYY-MM-DD)
- **AND** validates that start < end
- **AND** displays the entered dates
- **AND** prompts "Continue with these dates? (y/n)"

#### Scenario: Menu option 3 - Week number
- **WHEN** user selects option 3 in interactive mode
- **THEN** the script prompts for ISO week number
- **AND** calculates Monday-Sunday for that week
- **AND** displays the calculated dates
- **AND** prompts "Continue with these dates? (y/n)"

#### Scenario: Confirmation rejection
- **WHEN** user responds 'n' to date confirmation prompt
- **THEN** the script returns to the main menu
- **AND** allows user to select again

### Requirement: Week Calculation Logic
The script SHALL define weeks as Monday 00:00 to Sunday 23:59 and use ISO 8601 week numbering.

#### Scenario: Last week calculation mid-week
- **WHEN** current date is Wednesday 2025-02-05
- **THEN** "last week" is 2025-01-27 (Monday) to 2025-02-02 (Sunday)

#### Scenario: Last week calculation on Monday
- **WHEN** current date is Monday 2025-02-03
- **THEN** "last week" is 2025-01-27 (Monday) to 2025-02-02 (Sunday)

#### Scenario: ISO week number across year boundary
- **WHEN** user requests week 1 of 2025
- **THEN** the script calculates dates using ISO 8601 standard
- **AND** handles year boundaries correctly (week 1 may include December dates)

### Requirement: Date Format Validation
The script SHALL validate all date inputs follow YYYY-MM-DD format and represent valid calendar dates.

#### Scenario: Valid date format
- **WHEN** user enters "2025-01-20"
- **THEN** the date is accepted and parsed correctly

#### Scenario: Invalid date format
- **WHEN** user enters "01/20/2025" or "2025-1-20"
- **THEN** the script prints error "Invalid date. Use the format YYYY-MM-DD."
- **AND** prompts for input again (interactive) or exits with code 1 (non-interactive)

#### Scenario: Invalid calendar date
- **WHEN** user enters "2025-02-30"
- **THEN** the script prints error "Invalid date. Use the format YYYY-MM-DD."
- **AND** handles the error appropriately

### Requirement: Backward Compatibility
The script SHALL maintain backward compatibility with existing usage patterns when no arguments are provided.

#### Scenario: Legacy interactive mode still accessible
- **WHEN** user runs `./solidtimexport.py` without arguments
- **THEN** the script presents the new interactive menu
- **AND** option 2 (custom dates) provides the same manual entry experience as before
- **AND** all existing workflows continue to function

