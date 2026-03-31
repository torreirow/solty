# export-reporting Specification

## Purpose
TBD - created by archiving change improve-export-output. Update Purpose after archive.
## Requirements
### Requirement: Export Summary Report
The script SHALL display a formatted summary report showing key statistics about the exported time entries.

#### Scenario: Successful export with entries
- **WHEN** the script successfully exports time entries for a date range
- **THEN** it displays a summary section with:
  - Section header "=== Solidtime Export ==="
  - Date range with ISO week number
  - Total number of time entries found
  - Total hours (sum of all durations in decimal format)
  - Number of unique days with entries
  - Number of unique projects
  - Output file location with entry count

#### Scenario: Export with no entries
- **WHEN** the script finds no time entries for the selected date range
- **THEN** it displays "Found 0 time entries"
- **AND** shows "Total hours: 0.0"
- **AND** indicates the date range that was searched

#### Scenario: Export with single entry
- **WHEN** the script finds exactly one time entry
- **THEN** it displays "Found 1 time entry" (singular form)
- **AND** shows the correct hour calculation for that single entry

### Requirement: Progress Indicators
The script SHALL provide clear progress feedback during the export process.

#### Scenario: API fetch in progress
- **WHEN** the script is fetching data from the Solidtime API
- **THEN** it displays "Fetching time entries..." before the API call
- **AND** uses a checkmark (✓) symbol when data is successfully retrieved

#### Scenario: Multi-step process visibility
- **WHEN** the export involves multiple API calls (time entries, projects, clients)
- **THEN** the user sees clear indication of each step
- **AND** can follow the progress through the workflow

### Requirement: Statistics Accuracy
The displayed statistics SHALL accurately reflect the exported data.

#### Scenario: Total hours calculation
- **WHEN** the script calculates total hours
- **THEN** it sums all entry durations converted to decimal hours
- **AND** displays the result rounded to 1 decimal place

#### Scenario: Unique counts
- **WHEN** the script counts unique days with entries
- **THEN** it counts distinct dates (YYYY-MM-DD) from all entries
- **WHEN** the script counts unique projects
- **THEN** it counts distinct project IDs (excluding None/null)

#### Scenario: Entry count matches CSV
- **WHEN** the script displays "X entries"
- **THEN** the CSV file contains exactly X data rows (excluding header)
- **AND** the count excludes any entries skipped due to missing project_id or description

### Requirement: Clean Output Formatting
The console output SHALL be readable, professional, and suitable for both interactive and automated use.

#### Scenario: Section headers and spacing
- **WHEN** displaying the report
- **THEN** use clear section headers with "===" formatting
- **AND** appropriate blank lines between sections for readability

#### Scenario: Symbol usage
- **WHEN** indicating success or completion
- **THEN** use ✓ (checkmark) for successful operations
- **AND** use → (arrow) for output file indication

#### Scenario: Alignment and indentation
- **WHEN** displaying statistics
- **THEN** use consistent indentation (2 spaces) for sub-items
- **AND** align related information clearly

### Requirement: Backward Compatibility
The enhanced output SHALL maintain compatibility with existing automation scripts.

#### Scenario: Exit codes unchanged
- **WHEN** the export succeeds
- **THEN** the script exits with code 0 (as before)
- **WHEN** the export fails
- **THEN** the script exits with code 1 (as before)

#### Scenario: Error messages preserved
- **WHEN** errors occur (403, 500, validation errors)
- **THEN** error messages are still displayed clearly
- **AND** the error handling behavior is unchanged

#### Scenario: Scripting compatibility
- **WHEN** the script is used in automated workflows
- **THEN** the enhanced output doesn't break parsing or monitoring
- **AND** all critical information remains accessible

