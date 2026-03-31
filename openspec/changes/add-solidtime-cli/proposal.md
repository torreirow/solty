# Change: Add solidtime-cli tool

## Why

Currently, time tracking with Solidtime requires using the web interface or desktop app. There is no command-line tool similar to toggl-cli that allows developers to quickly start/stop timers and add time entries from the terminal. This creates friction in the developer workflow, especially when working in terminal-heavy environments.

The existing Python scripts (`solidtimexport.py`, `prepareexact.py`) are focused on batch export and transformation for Exact Online, not real-time time tracking operations.

## What Changes

- **NEW**: Command-line tool `solidtime-cli` written in Go
- **NEW**: Core commands: `start`, `stop`, `add`, `current`, `list`
- **NEW**: Configuration via existing `config.json` (reuses workspace_id and api_token)
- **NEW**: Capability: `solidtime-cli` for interactive time tracking operations

The tool will:
- Start timers with description, optional project, and optional custom start time
- Stop the currently running timer
- Add completed time entries with specific start/end times
- Show the current running timer
- List recent time entries (with optional ID display)
- Delete time entries by ID

## Impact

- Affected specs: **NEW capability** `solidtime-cli`
- Affected code: **NEW Go module** in `cmd/solidtime-cli/`
- Dependencies:
  - Go 1.21+ required
  - Config stored in `~/.config/solidtime/config.json` (XDG compliant)
  - No impact on existing Python scripts (solidtimexport.py, prepareexact.py)
- Users: Developers who want CLI-based time tracking
- Deployment: Binary distribution via `go install` or releases

## Features Added Beyond Initial Proposal

**List Command Enhancements:**
- `--id` flag to display entry UUIDs (needed for deletion)

**Delete Command:**
- NEW command to delete time entries by ID
- Permanent deletion (no undo)
- Required for correcting mistakes

**Start Command Enhancement:**
- `--time` flag for custom start time (if you forgot to start the timer)
- Supports ISO8601 and HH:MM formats
- Default: current time (when --time is not provided)

## Non-Goals (v1)

- No tags support (tags must be pre-created in Solidtime; API requires tag IDs)
- No integration with solidtimexport.py workflow (separate concern)
- No editing existing entries (only create/read operations)
- No reporting or analytics (use web interface)
- No OS keychain integration (uses ~/.config/solidtime/config.json)
- No interactive mode for start command (direct arguments only)
