# Solty - Solidtime CLI

Command-line interface for Solidtime time tracking.

## Features

- **Start/Stop timers** - Quick time tracking from the terminal
- **Add entries** - Create completed time entries with specific times
- **Current timer** - See what's running
- **List entries** - View recent time entries
- **Delete entries** - Remove mistakes
- **Project support** - Assign time to projects
- **Custom start times** - Backdate timers if you forgot to start

## Installation

```bash
# Clone the repository
git clone https://github.com/torreirow/solty.git
cd solty

# Build
go build -o solty

# Optional: Install to PATH
go install
```

## Configuration

Create `~/.config/solidtime/config.json`:

```json
{
  "username": "Your Name",
  "api_token": "your-solidtime-api-token",
  "workspace_id": "your-workspace-uuid"
}
```

**Getting your credentials:**
- **api_token**: Generate in Solidtime → Settings → API Tokens
- **workspace_id**: Found in Solidtime URL or organization settings

## Usage

### Start a timer

```bash
# Simple start
solty start "Working on feature X"

# With project
solty start "Bug fix" --project "Customer-Project"

# With custom start time (if you forgot to start)
solty start "Morning work" --time "09:00"
solty start "Task" --time "2026-03-31T08:00:00Z"
```

### Stop the timer

```bash
solty stop
```

### Show current timer

```bash
solty current
```

### Add completed entry

```bash
# Add entry with specific times
solty add "Meeting" --start "14:00" --end "15:30"

# With project
solty add "Sprint planning" --start "10:00" --end "12:00" --project "TN-Meetings"

# Full ISO8601 timestamps
solty add "Client call" --start "2026-03-31T14:00:00Z" --end "2026-03-31T15:30:00Z"
```

### List recent entries

```bash
# Show last 10 entries
solty list

# Show last 5 entries
solty list --limit 5

# Show with IDs (for deletion)
solty list --id
```

### Delete entry

```bash
# Get entry ID first
solty list --id

# Delete by ID
solty delete 01234567-89ab-cdef-0123-456789abcdef
```

## Configuration Locations

Solty searches for `config.json` in this order:

1. `~/.config/solidtime/config.json` (recommended)
2. `~/.solidtime/config.json`
3. `./config.json`

## Time Formats

Solty supports two time formats:

- **ISO8601**: `2026-03-31T14:00:00Z` (full timestamp)
- **Time only**: `14:00` (assumes today in local timezone)

## Project Names

Use project names (not IDs) with the `--project` flag. Solty will:
- Look up the project ID automatically
- Match names case-insensitively
- Suggest available projects if not found

## Exit Codes

- `0` - Success
- `1` - User error (invalid input, validation failed)
- `2` - System error (network issues, API errors)

## Examples

**Typical workflow:**

```bash
# Morning: Start working
solty start "Daily standup" --project "TN-General"
solty stop

solty start "Feature development" --project "Customer-Project"
# ... work for a few hours ...
solty stop

# Check what you tracked today
solty list
```

**Forgot to start timer:**

```bash
# Oops, been working since 9am
solty start "Morning coding" --time "09:00" --project "Customer-Project"
solty stop
```

**Add past entry:**

```bash
# Add yesterday's meeting you forgot to track
solty add "Client meeting" --start "2026-03-30T14:00:00Z" --end "2026-03-30T15:30:00Z" --project "Customer-Project"
```

**Fix mistakes:**

```bash
# List with IDs
solty list --id

# Delete wrong entry
solty delete <entry-id>
```

## Development

Built with:
- Go 1.21+
- Cobra CLI framework
- Solidtime API (JSON:API format)

## License

Internal tool - not for external distribution.
