## Why

Users need to frequently restart timers for recurring tasks with the same project context. Currently, they must manually re-enter the description and project each time. Additionally, the list output lacks key information (entry IDs and project names) that would make it more useful for quick reference and task continuation.

## What Changes

- Add new `continue` command to start a timer based on an existing entry (copies description and project)
- Modify `list` command default output to always show 8-character short IDs in a new column
- Add Project column to `list` command output between Duration and Description
- Maintain existing `--id` flag behavior (shows full 36-character UUID)
- Implement short ID matching with prefix-based lookup (accepts 6-36 characters)
- Add comprehensive error handling for ID matching (not found, ambiguous matches)

## Capabilities

### New Capabilities
- `continue-timer`: Start a new timer using an existing entry as template (copies description and project)
- `short-id-matching`: Match entries by UUID prefix (8-char default display, 6-36 char input)

### Modified Capabilities
- `list-entries`: Enhanced output format with ID and Project columns always visible

## Impact

**Files to modify:**
- `cmd/list.go` - Update default output format to include ID (8-char) and Project columns
- `cmd/continue.go` - New file for continue command
- `internal/client/timeentry.go` - Add method to find entry by short ID prefix

**User experience:**
- List output becomes wider but more informative
- Users can quickly see IDs and projects without needing the `--id` flag
- Faster workflow for recurring tasks via continue command
