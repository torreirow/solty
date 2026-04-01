# Changelog

All notable changes to Soltty will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## NEXT VERSION

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
