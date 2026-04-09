# Changelog

All notable changes to Soltty will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## NEXT VERSION

### Added
- **Continue command**: New `soltty continue <entry-id>` starts a timer using an existing entry as template
  - Copies description and project from any previous entry
  - Accepts short IDs (6-36 characters) or full UUID
  - Displays 8-character short IDs in list output for easy reference
  - Same auto-stop behavior as start command if timer is running
  - Helpful error messages for: not found, ambiguous matches, invalid format
- **Enhanced list output**: List command now shows entry IDs and project names by default
  - New ID column shows first 8 characters of UUID (e.g., "985d7cb2")
  - New Project column between Duration and Description
  - Full UUID display maintained with `--id` flag
  - Project names fetched once and mapped for efficient display
  - Shows "No project" for entries without assigned projects
- **Short ID matching**: Flexible ID prefix matching system
  - Accept 6-36 character prefixes (minimum 6 for safety)
  - Case-insensitive matching
  - Searches last 1000 entries for performance
  - Clear error messages with suggested actions for ambiguous or missing IDs
  - Statistical analysis shows 8-char IDs are safe (< 0.01% collision under 9,300 entries)

### Changed
- **List output format**: Default list output now includes ID and Project columns (wider output)
  - Previous format: Date | Start | Duration | Description
  - New format: ID | Date | Start | Duration | Project | Description
  - Output width increased from ~60 to ~90 characters
  - `--id` flag still available for full UUID display

## 0.3.0 - 01 Apr 2026

### Added
- **List clients subcommand**: New `soltty list clients` shows all clients alphabetically with project counts
  - Displays clients sorted by name
  - Shows project count per client (e.g., "Acme Corp (14 projects)")
  - Automatically filters out archived clients
- **List projects subcommand**: New `soltty list projects` shows all projects in table format
  - Table view with Client | Project columns
  - Shows projects sorted by client name, then project name
  - Automatically filters out archived projects
  - Optional `-c` flag to filter by client name (partial match, case-insensitive)
  - Projects without a client show as "(no client)"
  - Maintains backwards compatibility: `soltty list` still shows time entries
- **Web command**: New `soltty web` opens the Solidtime web interface in your default browser
  - Automatically derives web URL from configured API endpoint
  - Cross-platform support (Linux, macOS, Windows)
  - Shows URL if browser opening fails for manual copy-paste
  - Seamless transition between CLI and web interface

## 0.2.0 - 01 Apr 2026

### Added
- **Auto-stop on start**: When starting a new timer while one is already running, soltty now prompts to stop the current timer automatically. Eliminates manual `soltty stop` step.
  - Interactive confirmation with [y/N] prompt (defaults to safe 'no')
  - Shows currently running timer details (description and elapsed time)
  - Displays confirmation after stopping and starting timers
  - Addresses #1
- **Configurable API endpoint**: API base URL must now be configured via `base_url` field in config.json
  - ⚠️ **BREAKING**: `base_url` is now a required field (no default)
  - Enables use of self-hosted Solidtime instances
  - Explicit configuration prevents confusion about which instance is being used
  - URL validation with helpful error messages
- **Dedicated config directory**: Config location moved to `~/.config/soltty/` to align with tool name
  - Backward compatible: falls back to `~/.config/solidtime/` if new location doesn't exist
  - Search order: `~/.config/soltty/`, `~/.config/solidtime/`, `~/.solidtime/`, `./config.json`
- **List command improvements**: Running timers now show "running" instead of confusing "0s" duration
  - Makes it easy to identify active timer at a glance
  - Completed entries still show calculated duration

### Changed
- ⚠️ **BREAKING**: `base_url` is now a required field in config.json
  - Existing configs must add: `"base_url": "https://app.example.com/api/v1"`
  - Clear error message guides users if field is missing
  - For self-hosted instances, use your instance URL

## 0.1.2 - 01 Apr 2026

### Fixed
- Fix Nix flake vendorHash for Go module dependencies

## 0.1.1 - 31 Mar 2026

## 0.1.0 - 31 Mar 2026

Initial release of Soltty - Solidtime CLI time tracking tool.

### Added
- **Commands**: start, stop, current, add, list, delete
- **Start command**: Begin time tracking with optional custom start time via --time flag
- **Stop command**: End currently running timer
- **Current command**: Display active timer with elapsed time
- **Add command**: Create completed entries with specific start/end times
- **List command**: Display recent entries with optional --id flag for UUIDs
- **Delete command**: Remove time entries by ID
- **Project support**: Assign time entries to projects by name (case-insensitive lookup)
- **Time formats**: Support for ISO8601 and HH:MM time formats
- **XDG compliance**: Config stored in ~/.config/solidtime/config.json
- **Timezone handling**: UTC for API, local timezone for display
- **API integration**: Full JSON:API support for Solidtime API
- **Member ID resolution**: Automatic member ID lookup
- **Billable field**: Default false for all entries

### Technical
- Built with Go 1.21+ and Cobra CLI framework
- Modular architecture with internal packages (config, client)
- Comprehensive error handling and user-friendly messages
- OpenSpec proposal and documentation included
