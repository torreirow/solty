# Design: solidtime-cli

## Context

The solidtime-exact project currently provides Python scripts for batch export and transformation of time entries for Exact Online integration. However, there is no tool for real-time time tracking operations from the command line.

**Stakeholders:**
- Developers who prefer terminal-based workflows
- Users familiar with toggl-cli or similar tools
- Current users of solidtimexport.py who want real-time tracking

**Constraints:**
- Must reuse existing `config.json` for consistency with Python scripts
- Must work with self-hosted Solidtime instance at `solidtime.tools.technative.cloud`
- API uses JSON:API format (`application/vnd.api+json`)
- No breaking changes to existing Python scripts

## Goals / Non-Goals

**Goals:**
- Provide CLI commands: `start`, `stop`, `add`, `current`, `list`
- Enable quick time tracking without leaving the terminal
- XDG Base Directory compliant configuration
- Follow toggl-cli patterns for familiar UX
- Support project assignment

**Non-Goals (v1):**
- Tags support (API requires pre-existing tag IDs, complex to implement)
- Editing existing time entries
- Deleting time entries
- Reporting or analytics features
- Integration with solidtimexport.py workflow
- OS keychain/credential manager integration
- Interactive prompts (start command uses direct arguments)
- Sync with multiple workspaces

## Decisions

### Technology: Go

**Decision:** Implement in Go (not Python, Rust, or Node.js)

**Rationale:**
- Single binary distribution (no runtime dependencies)
- Fast compilation and execution
- Strong standard library for HTTP/JSON
- Cross-platform builds are straightforward
- cobra CLI framework is mature and well-documented

**Alternatives considered:**
- Python: Would match existing scripts but requires runtime, slower startup, distribution complexity
- Rust: Excellent performance but steeper learning curve, longer compile times
- Node.js: Requires runtime, larger distribution size

### Configuration: XDG Base Directory Compliant

**Decision:** Read from `~/.config/solidtime/config.json` (XDG compliant) with fallbacks

**Rationale:**
- Follows XDG Base Directory specification (Linux/Unix standard)
- Clean separation from Python scripts (which use project directory)
- User-specific configuration in standard location
- Simple JSON parsing in Go

**Search order:**
1. `~/.config/solidtime/config.json` (primary)
2. `~/.solidtime/config.json` (fallback)
3. `./config.json` (fallback for compatibility)

**Alternatives considered:**
- OS keychain: More secure but adds complexity, platform-specific code
- Environment variables only: Less discoverable, harder for new users
- New config format: Would fragment configuration management

### API Client Architecture

**Decision:** Single API client package with methods per resource

**Pattern:**
```go
type Client struct {
    baseURL     string
    token       string
    workspaceID string
    httpClient  *http.Client
}

func (c *Client) StartTimeEntry(description string, projectID *string, tags []string) (*TimeEntry, error)
func (c *Client) StopTimeEntry(entryID string) (*TimeEntry, error)
func (c *Client) GetCurrentTimeEntry() (*TimeEntry, error)
func (c *Client) CreateTimeEntry(entry TimeEntryCreate) (*TimeEntry, error)
func (c *Client) ListTimeEntries(limit int) ([]*TimeEntry, error)
```

**Rationale:**
- Encapsulates HTTP details
- Reusable across commands
- Easy to test with mocks
- Matches Go idioms

### Command Structure

**Decision:** Use cobra for CLI framework with subcommands

**Pattern:**
```
solidtime-cli start "Working on feature X" --project "PSB-Project"
solidtime-cli stop
solidtime-cli current
solidtime-cli add "Meeting" --start "14:00" --end "15:30" --project "TN-Meetings"
solidtime-cli list --limit 10
```

**Rationale:**
- Familiar pattern from toggl-cli, kubectl, git
- Built-in help generation
- Flag parsing included
- Subcommand discoverability

### Time Format Handling

**Decision:** Accept ISO8601 for `--start`/`--end`, default to "today" if only time provided

**Examples:**
- `--start "2026-03-31T14:00:00Z"` (full ISO8601)
- `--start "14:00"` (today at 14:00 in local time)
- `--start "14:00" --end "15:30"` (today from 14:00 to 15:30)

**Rationale:**
- ISO8601 is unambiguous
- Natural time format ("14:00") is user-friendly
- Matches Python script patterns

**Alternatives considered:**
- Natural language parsing ("2 hours ago"): Complex, error-prone
- Unix timestamps only: Not user-friendly

### Project Lookup

**Decision:** Accept project name via `--project` flag, lookup project_id via API

**Implementation:**
- Cache project list on first use (in-memory for command duration)
- Match project names case-insensitively
- Error if project not found with suggestion of available projects

**Rationale:**
- Users remember names, not UUIDs
- Matches existing Python script behavior (projects.json)
- Single API call per invocation is acceptable

**Alternatives considered:**
- Require project ID: Poor UX
- Local project cache file: Adds stale data complexity
- Fuzzy matching: Can cause unexpected behavior

## Data Structures

### TimeEntry

```go
type TimeEntry struct {
    ID            string    `json:"id"`
    OrganizationID string   `json:"organization_id"`
    UserID        string    `json:"user_id"`
    ProjectID     *string   `json:"project_id"`
    TaskID        *string   `json:"task_id"`
    Description   string    `json:"description"`
    Start         time.Time `json:"start"`
    End           *time.Time `json:"end"` // nil if running
    Duration      int       `json:"duration"` // seconds
    Billable      bool      `json:"billable"`
    Tags          []string  `json:"tags"`
}
```

### Config

```go
type Config struct {
    Username    string `json:"username"`
    APIToken    string `json:"api_token"`
    WorkspaceID string `json:"workspace_id"`
}
```

## Error Handling

**Strategy:**
- Return errors up to command level
- Display user-friendly messages (not raw HTTP errors)
- Exit codes: 0 = success, 1 = user error, 2 = system error

**Examples:**
- API token invalid: "Authentication failed. Check your API token in config.json"
- No timer running: "No timer is currently running"
- Project not found: "Project 'XYZ' not found. Available projects: ABC, DEF"

## Risks / Trade-offs

### Risk: API Changes

**Mitigation:**
- Version API calls if Solidtime adds versioning
- Document tested Solidtime version in README
- Monitor Solidtime releases for breaking changes

### Risk: Config.json Location

**Issue:** Tool must find config.json in standard location

**Mitigation:**
- Primary location: `~/.config/solidtime/config.json` (XDG compliant)
- Fallback locations: `~/.solidtime/config.json`, `./config.json`
- Clear error message showing search paths with setup instructions
- Config example file included: `cmd/solidtime-cli/config.json.example`
- Future: Add `--config` flag for explicit path

### Trade-off: No credential manager

**Benefit:** Simpler implementation, consistent with Python scripts
**Cost:** Less secure than OS keychain (token in plaintext file)
**Justification:** Internal tool, existing pattern, users accept risk

### Trade-off: In-memory project cache

**Benefit:** No file I/O, no stale data
**Cost:** API call per invocation
**Justification:** Acceptable latency (~200ms), projects change infrequently

## Migration Plan

**Phase 1: Initial Release (v0.1.0)**
- Core commands: start, stop, current, add, list
- Uses config.json from current directory
- Manual installation via `go install`

**Phase 2: Future Enhancements** (not in this proposal)
- Edit commands (update existing entries)
- Tags support (requires tag name → ID lookup)
- Better project caching (persistent cache)
- Integration with solidtimexport.py
- OS keychain support
- Interactive mode for start command
- Delete confirmation prompts

**Rollback:** N/A (new tool, no migration needed)

## Decisions Made During Implementation

### Tags Removed from v1

**Decision:** Remove tags support from initial release

**Rationale:**
- Solidtime API requires tag IDs, not tag names
- Tags must be pre-created in Solidtime web interface
- Implementing tag name → ID lookup adds complexity
- Users can add tags via web interface after creation

**For v2:** Implement `/organizations/{id}/tags` endpoint integration

### Delete Command Added

**Decision:** Add delete command for correcting mistakes

**Rationale:**
- Users need ability to remove incorrectly created entries
- List --id flag provides the UUIDs needed for deletion
- Simple API: DELETE `/organizations/{id}/time-entries/{entry_id}`
- No confirmation prompt in v1 (keep it simple, warn in docs)

### Custom Start Time Added

**Decision:** Add --start flag to start command

**Rationale:**
- Common use case: forgot to start timer when work began
- Supports both ISO8601 and HH:MM formats (consistent with add command)
- More user-friendly than deleting and using add command
- API accepts any start time in the past

**Pattern:** `solidtime-cli start "Task" --start "09:00"`

## Open Questions

None - all clarifications received from user:
- ✓ Commands: start, stop, add, current, list, delete
- ✓ Configuration: ~/.config/solidtime (XDG compliant)
- ✓ Tags: Removed from v1 (planned for v2)
- ✓ Custom start time: Added to start command
