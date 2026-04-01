# Changelog

All notable changes to Soltty will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
  - Existing configs must add: `"base_url": "https://solidtime.tools.technative.cloud/api/v1"`
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
