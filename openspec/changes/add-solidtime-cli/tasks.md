# Implementation Tasks

## 1. Project Setup

- [x] 1.1 Initialize Go module `cmd/solidtime-cli`
- [x] 1.2 Add `go.mod` and `go.sum` with dependencies (cobra for CLI, standard library for HTTP/JSON)
- [x] 1.3 Create basic CLI structure with cobra commands
- [x] 1.4 Add .gitignore entries for Go build artifacts

## 2. Configuration Management

- [x] 2.1 Implement config loader to read existing `config.json`
- [x] 2.2 Parse `api_token` and `workspace_id` from config
- [x] 2.3 Validate required config fields on startup
- [x] 2.4 Add helpful error messages for missing/invalid config

## 3. Solidtime API Client

- [x] 3.1 Create API client package with base URL and authentication
- [x] 3.2 Implement HTTP client with Bearer token authentication
- [x] 3.3 Add Accept header: `application/vnd.api+json`
- [x] 3.4 Implement error handling for HTTP responses (401, 403, 404, 500)

## 4. Start Command

- [x] 4.1 Implement `solidtime-cli start <description>` command
- [x] 4.2 Add flags: `--project <name>` and `--tags <tag1,tag2>`
- [x] 4.3 Create API call: POST `/organizations/{workspace_id}/time-entries`
- [x] 4.4 Handle project name to project_id lookup (fetch projects from API)
- [x] 4.5 Display confirmation with entry ID and start time
- [x] 4.6 Handle error: timer already running

## 5. Stop Command

- [x] 5.1 Implement `solidtime-cli stop` command
- [x] 5.2 Fetch current running timer via GET `/organizations/{workspace_id}/time-entries` with filter
- [x] 5.3 Stop timer via PUT/PATCH with end timestamp
- [x] 5.4 Display stopped entry with duration
- [x] 5.5 Handle error: no timer running

## 6. Current Command

- [x] 6.1 Implement `solidtime-cli current` command
- [x] 6.2 Fetch running timer from API (filter by end=null)
- [x] 6.3 Display description, project, start time, and elapsed duration
- [x] 6.4 Handle case: no timer running (friendly message)

## 7. Add Command

- [x] 7.1 Implement `solidtime-cli add <description>` command
- [x] 7.2 Add flags: `--start <time>`, `--end <time>`, `--project <name>`, `--tags <tags>`
- [x] 7.3 Parse time formats (support ISO8601 and natural formats like "14:30")
- [x] 7.4 Create completed time entry via POST with both start and end times
- [x] 7.5 Display confirmation with calculated duration

## 8. List Command

- [x] 8.1 Implement `solidtime-cli list` command
- [x] 8.2 Add flag: `--limit <n>` (default: 10)
- [x] 8.3 Fetch recent entries via GET with pagination
- [x] 8.4 Display table: start time, duration, description, project
- [x] 8.5 Format durations as human-readable (e.g., "2h 30m")

## 9. Testing & Documentation

- [x] 9.1 Add README.md in `cmd/solidtime-cli/` with installation instructions
- [x] 9.2 Document all commands with examples
- [x] 9.3 Add usage examples for common workflows
- [x] 9.4 Test all commands against live Solidtime API
- [x] 9.5 Add error scenario testing (network errors, auth failures)

## 10. Build & Release

- [x] 10.1 Add Makefile or build script for cross-compilation
- [x] 10.2 Create build targets for Linux, macOS, Windows
- [x] 10.3 Document installation via `go install` command
- [x] 10.4 Update main README.md to mention solidtime-cli tool

## 11. Config Migration to XDG

- [x] 11.1 Update config loader to use ~/.config/solidtime as primary location
- [x] 11.2 Implement fallback search paths
- [x] 11.3 Update error messages with new config location
- [x] 11.4 Create config.json.example template
- [x] 11.5 Update all documentation with new config path

## 12. Tags Removal

- [x] 12.1 Remove --tags flag from start command
- [x] 12.2 Remove --tags flag from add command
- [x] 12.3 Update API client to not send tags
- [x] 12.4 Update documentation to reflect v1 limitation
- [x] 12.5 Add note about tags support planned for v2

## 13. List Command Enhancement

- [x] 13.1 Add --id flag to list command
- [x] 13.2 Implement conditional ID column display
- [x] 13.3 Adjust table formatting for ID column
- [x] 13.4 Update documentation with --id flag usage

## 14. Delete Command

- [x] 14.1 Create delete command file
- [x] 14.2 Implement DeleteTimeEntry in API client
- [x] 14.3 Add DELETE HTTP method support
- [x] 14.4 Display confirmation message
- [x] 14.5 Document delete workflow with list --id

## 15. Start Command Enhancement

- [x] 15.1 Add --start flag to start command
- [x] 15.2 Create shared parseTime utility function
- [x] 15.3 Update StartTimeEntry to accept custom start time
- [x] 15.4 Handle both ISO8601 and HH:MM formats
- [x] 15.5 Display custom start time in confirmation message
- [x] 15.6 Update documentation with custom start time examples
