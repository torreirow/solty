# Exact Prepare Capability

## ADDED Requirements

### Requirement: Automatic Input File Detection
The prepareexact.py script SHALL automatically use time_entries.csv as input when no `-i` argument is provided.

#### Scenario: Default input file exists
- **WHEN** user runs `./prepareexact.py` without `-i` argument
- **AND** time_entries.csv exists in current directory
- **THEN** the script uses time_entries.csv as input automatically

#### Scenario: Default input file missing
- **WHEN** user runs `./prepareexact.py` without `-i` argument
- **AND** time_entries.csv does not exist
- **THEN** the script prints error "Input file 'time_entries.csv' not found"
- **AND** exits with code 1

#### Scenario: Explicit input file specified
- **WHEN** user runs `./prepareexact.py -i custom.csv`
- **THEN** the script uses custom.csv as input
- **AND** ignores time_entries.csv even if it exists

### Requirement: Smart Output Filename Suggestion
The script SHALL analyze the date range in the input CSV and suggest an appropriate output filename based on ISO week number.

#### Scenario: Single week date range
- **WHEN** input CSV contains dates from 2025-01-27 to 2025-02-02 (week 5)
- **AND** no `-o` argument provided
- **THEN** the script suggests filename "week-05.csv"
- **AND** displays "Date range: 2025-01-27 to 2025-02-02 (week 5)"
- **AND** prompts "Save as week-05.csv? (Y/n): "

#### Scenario: Multiple weeks date range
- **WHEN** input CSV contains dates spanning multiple ISO weeks
- **AND** no `-o` argument provided
- **THEN** the script displays "Date range spans multiple weeks"
- **AND** prompts "Enter output filename: "
- **AND** requires manual filename input

#### Scenario: User accepts suggestion
- **WHEN** script suggests "week-05.csv"
- **AND** user enters "y" or "Y" or just presses Enter
- **THEN** the script uses "week-05.csv" as output filename

#### Scenario: User provides custom filename
- **WHEN** script suggests "week-05.csv"
- **AND** user enters "n" or "N"
- **THEN** the script prompts "Enter output filename: "
- **AND** uses the user-provided filename

### Requirement: File Collision Detection
The script SHALL detect when the output file already exists and automatically generate a timestamped alternative filename to prevent data loss.

#### Scenario: Output file does not exist
- **WHEN** suggested output file "week-05.csv" does not exist
- **THEN** the script creates "week-05.csv" normally

#### Scenario: Output file already exists
- **WHEN** suggested output file "week-05.csv" already exists
- **THEN** the script generates timestamp in YYYYMMDD_HHMMSS format
- **AND** creates filename "week-05-20260202_143022.csv"
- **AND** displays "File 'week-05.csv' exists, saving as: week-05-20260202_143022.csv"
- **AND** does NOT overwrite the existing file

#### Scenario: Timestamped file also exists
- **WHEN** both "week-05.csv" and timestamp variant exist
- **THEN** the script generates new timestamp (current time)
- **AND** creates unique filename with new timestamp
- **AND** ensures no file is overwritten

### Requirement: Override Mode
The script SHALL support explicit output filename specification via `-o` argument, bypassing all auto-detection and confirmation prompts.

#### Scenario: Explicit output with -o
- **WHEN** user runs `./prepareexact.py -i input.csv -o output.csv`
- **THEN** the script uses "output.csv" as output filename
- **AND** shows NO confirmation prompts
- **AND** applies file collision detection if output.csv exists

#### Scenario: Override with -o only
- **WHEN** user runs `./prepareexact.py -o output.csv`
- **THEN** the script uses time_entries.csv as input (default)
- **AND** uses "output.csv" as output (explicit)
- **AND** shows NO confirmation prompts

### Requirement: Date Range Validation
The script SHALL validate that all dates in the CSV are parseable and provide meaningful error messages for invalid data.

#### Scenario: Valid date formats
- **WHEN** CSV contains dates in YYYY-MM-DD format
- **THEN** the script parses them correctly for week calculation

#### Scenario: Invalid date format
- **WHEN** CSV contains unparseable dates
- **THEN** the script displays "Warning: Could not parse date: {date_value}"
- **AND** skips invalid dates for week calculation
- **AND** continues processing valid dates

#### Scenario: Empty CSV
- **WHEN** input CSV has no data rows (only header)
- **THEN** the script displays "Error: No time entries found in CSV"
- **AND** exits with code 1

### Requirement: Backward Compatibility
The script SHALL maintain full backward compatibility with existing usage patterns.

#### Scenario: Legacy explicit mode
- **WHEN** user runs `./prepareexact.py -i time_entries.csv -o week-05.csv`
- **THEN** the script behaves exactly as before
- **AND** shows NO new prompts or messages
- **AND** processes silently (except errors)

#### Scenario: Existing scripts continue working
- **WHEN** automated scripts use `-i` and `-o` arguments
- **THEN** they work without modification
- **AND** no user interaction is required
- **AND** behavior is identical to previous version
