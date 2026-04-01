## Why

Users need to quickly view available clients and projects from the CLI to know which project names to use when starting time entries. Currently, users must either remember project names, check the web interface, or rely on error messages that show available projects. This creates friction in the workflow and slows down time tracking.

The Solidtime API already provides clients and projects data with their relationships, but Soltty doesn't expose this information through list commands.

## What Changes

- Extend the `list` command to support listing clients and projects via subcommands
- Add `soltty list clients` to show all clients with project counts
- Add `soltty list projects` to show all projects with their associated clients
- Add `soltty list projects -c <client>` to filter projects by client (partial match, case-insensitive)
- Update data structures to include client information on projects
- Maintain backwards compatibility: `soltty list` continues to show time entries (default behavior)
- Filter out archived clients and projects from listings

## Capabilities

### New Capabilities
- `list-clients`: CLI command to list all active clients with project counts in alphabetical order
- `list-projects`: CLI command to list all active projects with client information, optionally filtered by client name

### Modified Capabilities
<!-- No existing capabilities require requirement changes - this extends the existing list command -->

## Impact

- Extends existing `list` command with new subcommands (backwards compatible)
- Updates `Project` data structure to include `client_id` field
- Adds new `Client` data structure and API methods in `internal/client/`
- New command files: `cmd/list_clients.go` and `cmd/list_projects.go`
- No breaking changes to existing commands or APIs
