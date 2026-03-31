# Project Context

## Purpose

This tool exports time entries from Solidtime and converts them to a CSV format that can be imported into Exact Online.

**Workflow:**
1. `solidtimexport.py` - Fetches time entries from Solidtime API and generates `time_entries.csv`
2. `prepareexact.py` - Converts `time_entries.csv` to Exact Online import format

## Tech Stack

- **Python 3.11+** - Primary programming language (tested with 3.11.14)
- **requests** - For Solidtime API calls (external dependency)
- **csv** - For CSV processing (stdlib)
- **json** - For configuration and mapping files (stdlib)
- **argparse** - Command-line argument parsing in solidtimexport.py and prepareexact.py (stdlib)
- **datetime** - Date validation, manipulation, and ISO week calculations (stdlib, includes timedelta)
- **collections** - Data structures (defaultdict for grouping entries) (stdlib)
- **os** - File system operations in prepareexact.py (stdlib)
- **webbrowser** - Opening URLs in default browser in prepareexact.py (stdlib)
- **sys** - System operations (exit codes) in solidtimexport.py (stdlib)

## Project Conventions

### Code Style
- Python scripts with shebang `#!/usr/bin/env python3`
- English variable names, function names, and comments throughout all scripts
- Function-based architecture without OOP
- Scripts are executable (`chmod +x`)
- Docstrings for utility functions in English
- No type hints
- JSON configuration file keys remain in Dutch for backward compatibility (e.g., "medewerker", "Relatie", "Uursoort")

### Architecture Patterns
- Two-step process: export → transformation
- Configuration via JSON files
- Mapping-based transformation (projects.json, employees.json)
- Simple procedural scripts without classes
- Direct error handling with print statements and exit codes

### Error Handling
- Date validation with try/except for datetime parsing
- HTTP status code checking (200, 403, 500)
- File existence validation before processing
- Missing configuration checks (username, project mapping)
- Exit with code 1 on validation errors

### Data Flow
1. **Input**: Solidtime API (time entries, projects, clients)
2. **Intermediate**: `time_entries.csv` (Date;Worked;Titles;Projects) - semicolon delimited
3. **Output**: Exact Online CSV (Medewerker,Artikel,Datum,Relatie,Project,Notities,Aantal) - comma delimited

### File Structure
```
.
├── solidtimexport.py      # Step 1: Export from Solidtime API
├── prepareexact.py        # Step 2: Transform for Exact
├── openexact.sh           # Helper: Open Exact import page
├── config.json            # User credentials & API token (gitignored)
├── config.json.example    # Template for config.json
├── employees.json         # User to employee number mapping
├── projects.json          # Project to Exact data mapping
├── time_entries.csv       # Intermediate export (gitignored)
├── week-*.csv             # Output files for Exact import (gitignored)
├── 2025/                  # Directory for archived exports
└── temp/                  # Optional: Full Solidtime CSV exports (gitignored)
    └── solidtime-export/  # Example export with all Solidtime data
        ├── meta.json
        ├── time_entries.csv
        ├── projects.csv
        ├── clients.csv
        └── members.csv
```

## Domain Context

### Solidtime
- Time tracking application at `solidtime.tools.technative.cloud`
- API authentication via Bearer token
- Workspace-based structure (organizations)
- Time entries contain: date, duration, description, project_id
- Export functionality: Can export all data as CSV files via web interface
  - Format: Multiple CSV files with relational structure
  - Files: time_entries.csv, projects.csv, clients.csv, members.csv, etc.
  - Includes: organization, member, project, client, task, and tag data
  - Export metadata in meta.json (version, export date, organization IDs)

### Exact Online
- Import via ProTime Import functionality
- URL: `https://start.exactonline.nl/docs/ProTimeImport.aspx?_Division_=3314255`
- Requires specific CSV format with comma separator
- Time in decimal hours (not hour:minute format)

### Mapping Files

**config.json** (gitignored - use config.json.example as template)
- `username` - For lookup in employees.json (e.g., "Wouter")
- `api_token` - Solidtime API token (Bearer JWT, long-lived)
- `workspace_id` - Solidtime organization/workspace UUID
- Example: See `config.json.example` for structure

**employees.json**
- Array of `{"User": "name", "medewerker": number}`
- Maps Solidtime user to Exact employee number

**projects.json**
- Project name mapping to Exact data
- Keys: Solidtime project names (uppercase normalized)
- Values: `{"Relatie": int, "Project": string, "Uursoort": string}`
- Hour types (Uursoort) used:
  - `CONS-EXT-PROJ` - External consulting project work
  - `CONS-EXT-MEET` - External consulting meetings
  - `CONS-INT-PROJ` - Internal consulting project work
  - `CONS-INT-MEET` - Internal consulting meetings
  - `VAKANTIE` - Vacation/PTO
  - `VERZUIM` - Sick leave/absence
  - `TRAINING` - Training hours
  - `NON BILLABLE` - Non-billable internal work

### Project Naming Conventions
Format: `{CLIENT}-{TYPE}`
- Clients: TN (TechNative), MUS, PSB, INF, IIT, ACF, SMC, WDB, MOO, NBD, AMS, SPL, BNX, DR, TMCS
- Types: Project, General, Meetings, Incidents, MSA, PTO, Learning, Training, ISO, FinOps, PreSales, Managed_Services, Sick, etc.
- Examples from projects.json:
  - `TN-PTO` → Vacation hours (VAKANTIE)
  - `TN-Sick` → Sick leave (VERZUIM)
  - `PSB-Meetings` → External client meeting hours (CONS-EXT-MEET)
  - `TN-Meetings` → Internal meeting hours (CONS-INT-MEET)
  - `IIT-Project` → External project hours (CONS-EXT-PROJ)

## Important Constraints

- CSV delimiter switching: Solidtime export uses `;`, Exact import requires `,`
- Time format conversion: `h:mm` → decimal hours (rounded to 4 decimals)
- Entries without description are skipped in the export (solidtimexport.py:80)
- Entries without project_id are skipped (solidtimexport.py:79)
- Project names must match exactly (case-insensitive, normalized to uppercase) in projects.json
- Username in config.json must exist in employees.json (raises ValueError if not found)
- Time entries are grouped by date and description, summing durations
- `config.json` is gitignored for security (contains API tokens)

## External Dependencies

### Solidtime API
- Base URL: `https://solidtime.tools.technative.cloud/api/v1`
- Endpoints:
  - `/organizations/{workspace_id}/time-entries` (with start/end params)
  - `/organizations/{workspace_id}/projects`
  - `/organizations/{workspace_id}/clients`
- Authentication: `Bearer {api_token}`
- Accept: `application/vnd.api+json`

### Exact Online
- Manual import via web interface
- Division: 3314255
- ProTime Import functionality

## Setup & Installation

### Prerequisites
- Python 3.11 or higher
- `requests` library (install via pip)
- Linux/Unix environment (uses `xdg-open` in openexact.sh)

### Initial Setup
1. Clone the repository
2. Install dependencies:
   ```bash
   pip install requests
   ```
3. Create `config.json` from template:
   ```bash
   cp config.json.example config.json
   ```
4. Edit `config.json` with your credentials:
   - Add your Solidtime API token
   - Add your workspace ID
   - Add your username (must match employees.json)
5. Ensure your username exists in `employees.json`
6. Verify projects in `projects.json` match your Solidtime project names

### Configuration Files
- **config.json** - Personal credentials (gitignored, create from example)
- **employees.json** - Shared team mapping (version controlled)
- **projects.json** - Shared project mapping (version controlled)

## Scripts Usage

### solidtimexport.py

The export script supports both interactive and non-interactive modes with flexible date selection options.

**Non-interactive mode (automation-friendly):**

```bash
# Export specific date range
./solidtimexport.py --start 2025-01-20 --end 2025-01-26

# Export previous week automatically
./solidtimexport.py --last-week

# Export specific ISO week number (current year)
./solidtimexport.py --week 5

# Export specific ISO week number (specific year)
./solidtimexport.py --week 5 --year 2024

# Show help and all options
./solidtimexport.py --help
```

**Interactive mode (when no arguments provided):**

```bash
./solidtimexport.py
# Presents menu with three options:
#   [1] Last week (automatic calculation)
#   [2] Enter custom start/end dates
#   [3] Enter specific week number
# Always shows calculated dates and asks for confirmation
# Output: time_entries.csv
```

**Date selection features:**
- Week definition: Monday 00:00 to Sunday 23:59
- ISO 8601 week numbering
- Date format validation (YYYY-MM-DD required)
- Date range validation (start must be ≤ end)
- Confirmation prompt in interactive mode
- Display selected dates in all modes

### prepareexact.py
```bash
./prepareexact.py -i time_entries.csv -o exact_import.csv
# Or: open Exact import page
./prepareexact.py --open-url https://start.exactonline.nl/docs/ProTimeImport.aspx?_Division_=3314255
```

### openexact.sh
```bash
./openexact.sh
# Opens Exact Online ProTime Import page in browser
```

## Typical Workflow

### Quick Weekly Export (Recommended)

Most common use case - exporting last week's time entries:

```bash
# 1. Export last week's time entries (one command!)
./solidtimexport.py --last-week
# Output: Using dates: 2025-01-27 to 2025-02-02
# Creates: time_entries.csv

# 2. Convert to Exact format
./prepareexact.py -i time_entries.csv -o week-05.csv
# Creates: week-05.csv

# 3. Import into Exact
./openexact.sh
# Opens browser to ProTime Import page
# Upload the week-05.csv file manually
```

### Interactive Workflow

For custom date ranges or when you prefer guided prompts:

```bash
# 1. Export time entries (interactive menu)
./solidtimexport.py
# Select option:
#   [1] Last week → automatic calculation
#   [2] Custom dates → enter 2025-01-20 and 2025-01-26
#   [3] Week number → enter 3 for week 3
# Confirm dates when prompted
# Creates: time_entries.csv

# 2. Convert to Exact format
./prepareexact.py -i time_entries.csv -o week-03.csv

# 3. Import into Exact
./openexact.sh
```

### Automated Workflow (Cron/Scripts)

For automation and batch processing:

```bash
# Weekly cron job example (runs every Monday)
./solidtimexport.py --last-week
./prepareexact.py -i time_entries.csv -o "week-$(date +%V).csv"

# Specific date range for reporting
./solidtimexport.py --start 2025-01-01 --end 2025-01-31
./prepareexact.py -i time_entries.csv -o january-2025.csv
```

## Known Limitations & Issues

- No automated testing (no test suite)
- No logging system (only console prints)
- No validation that all projects in Solidtime have mappings in projects.json
- Projects without mapping will have None values in output (may cause import errors)
- No retry logic for API failures
- API token expiration not handled gracefully
- Time entries without project_id are silently skipped
- Weekly CSV files (week-*.csv) accumulate in root directory
- prepareexact.py must be run separately (no single-command workflow)

## Security Considerations

- **API tokens**: config.json is gitignored and contains sensitive Bearer tokens
- **Token format**: JWT tokens with long expiration (example expires in 2027)
- **No encryption**: Credentials stored in plain text JSON (file system security only)
- **Workspace isolation**: All API calls scoped to specific workspace_id
- **No authentication validation**: Scripts don't verify token validity before use
- **Division ID**: Exact division (3314255) is hardcoded in URLs

## Development Notes

- No virtual environment recommended in documentation
- No requirements.txt file (manual pip install)
- Scripts assume they run from repository root directory
- CSV encoding is UTF-8 with newline handling
- Date format strictly enforced as YYYY-MM-DD (ISO 8601)
- Timezone handling: API uses UTC (Z suffix), but dates are local
- HTTP errors print to console but don't provide detailed debugging info
- ISO week calculations use ISO 8601 standard (week 1 contains first Thursday of year)
- Week definitions: Monday = start (weekday 0), Sunday = end (weekday 6)
- Command-line arguments use argparse with mutually exclusive groups
- Interactive mode uses while loops with input validation and retry logic
- Exit codes: 0 = success, 1 = validation/error
