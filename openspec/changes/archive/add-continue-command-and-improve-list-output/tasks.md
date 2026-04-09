## 1. Update List Command Output

- [x] 1.1 Modify `cmd/list.go` to add ID column (8 chars) as first column in default output
- [x] 1.2 Add Project column between Duration and Description in default output
- [x] 1.3 Fetch all projects once and create ID-to-name map for efficient lookup
- [x] 1.4 Update default table header to show: ID | Date | Start | Duration | Project | Description
- [x] 1.5 Update row format to display 8-char short ID using `entry.ID[:8]`
- [x] 1.6 Update row format to show project name or "No project" for null project_id
- [x] 1.7 Adjust column spacing and separator line length for new format
- [x] 1.8 Maintain existing `--id` flag behavior (show full 36-char UUID with wider ID column)
- [x] 1.9 Update header and separator for `--id` mode to accommodate full UUID width

## 2. Implement Short ID Matching

- [x] 2.1 Add `FindEntryByShortID(shortID string) (*TimeEntry, error)` method to `internal/client/timeentry.go`
- [x] 2.2 Implement ID format validation (6-36 chars, hex + dashes only, case-insensitive)
- [x] 2.3 Implement length validation (minimum 6 characters)
- [x] 2.4 Fetch last 1000 entries for search scope
- [x] 2.5 Implement prefix matching algorithm (case-insensitive comparison)
- [x] 2.6 Return appropriate errors for: not found (0 matches), ambiguous (2+ matches), invalid format
- [x] 2.7 For ambiguous matches, return list of matching entries with details

## 3. Create Continue Command

- [x] 3.1 Create new file `cmd/continue.go` with continue command structure
- [x] 3.2 Define command with usage: `continue <entry-id>`
- [x] 3.3 Validate that entry ID argument is provided (show usage if missing)
- [x] 3.4 Call `FindEntryByShortID()` to look up the referenced entry
- [x] 3.5 Handle error cases: not found, ambiguous, invalid format with helpful messages
- [x] 3.6 For "not found", suggest using `soltty list` and `soltty list --id`
- [x] 3.7 For "ambiguous", show matching entries (ID, date, time, description) and suggest using more characters
- [x] 3.8 Extract description and project_id from found entry
- [x] 3.9 Check for running timer (reuse existing logic from start command)
- [x] 3.10 If timer running, prompt user to stop it (same flow as start command)
- [x] 3.11 Call existing `StartTimeEntry()` with copied description and project_id
- [x] 3.12 Display confirmation showing copied description and project name
- [x] 3.13 Register continue command in `cmd/root.go`

## 4. Update Documentation

- [x] 4.1 Update README.md with new list output format example (showing ID and Project columns)
- [x] 4.2 Add continue command documentation to README usage section
- [x] 4.3 Document short ID usage (6-36 chars accepted, 8 chars shown in list)
- [x] 4.4 Add examples of continue command with error handling scenarios
- [x] 4.5 Update list command documentation to mention `--id` flag shows full UUID
- [x] 4.6 Update CHANGELOG.md with new features under "## NEXT VERSION"

## 5. Testing

- [x] 5.1 Test list command shows 8-char IDs and project names correctly
- [x] 5.2 Test list --id shows full UUIDs with proper column width
- [x] 5.3 Test continue with 6-char, 8-char, 12-char, and 36-char IDs
- [x] 5.4 Test continue with non-existent ID shows proper error
- [x] 5.5 Test continue with ambiguous ID shows matching entries
- [x] 5.6 Test continue with invalid format shows validation error
- [x] 5.7 Test continue with ID too short (< 6 chars) shows error
- [x] 5.8 Test continue copies description and project correctly
- [x] 5.9 Test continue with entry having no project works correctly
- [x] 5.10 Test continue with running timer prompts for stop confirmation
- [x] 5.11 Verify all error messages reference `soltty list --id` for full UUIDs
