## Context

Soltty uses Cobra for CLI command structure. The current `list` command (`cmd/list.go`) shows time entries. Commands are organized as individual files in the `cmd/` directory.

The Solidtime API provides:
- `GET /organizations/{id}/clients` - returns clients with `id`, `name`, `is_archived` fields
- `GET /organizations/{id}/projects` - returns projects with `id`, `name`, `client_id`, `is_archived` fields

Current workspace has ~15 clients and ~46 projects. The largest client (TechNative) has 14 projects.

## Goals / Non-Goals

**Goals:**
- Add `list clients` and `list projects` subcommands to the existing `list` command
- Show clients alphabetically with project counts
- Show projects in table format with client names
- Filter projects by client name (partial match, case-insensitive)
- Filter out archived clients and projects
- Maintain backwards compatibility (`soltty list` shows time entries)

**Non-Goals:**
- Interactive filtering (fzf-style) - users can pipe to grep if needed
- Pagination - current dataset is small enough, can add later if needed
- Sorting by project count or recent usage
- Showing archived items (even with flags)

## Decisions

### 1. Command Structure: Subcommands

**Decision:** Use Cobra subcommands under the existing `list` parent command.

**Rationale:**
- Cobra has excellent subcommand support
- Explicit and discoverable: `soltty list --help` shows all options
- Backwards compatible: `soltty list` with no args runs default behavior
- Consistent with CLI best practices (like `git log`, `kubectl get`)

**Structure:**
```go
listCmd (parent)
├─ listClientsCmd    // soltty list clients
└─ listProjectsCmd   // soltty list projects [-c client]
```

Default behavior: if no subcommand, run existing time entries list.

**Alternatives considered:**
- Single file with if/else logic → Rejected: harder to maintain, less discoverable
- Flags like --clients, --projects → Rejected: user requested subcommands (option A)

### 2. Data Structure Updates

**Decision:** Extend `Project` struct to include `ClientID` and create new `Client` struct.

**Current:**
```go
type Project struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
```

**New:**
```go
type Project struct {
    ID         string  `json:"id"`
    Name       string  `json:"name"`
    ClientID   *string `json:"client_id"`
    IsArchived bool    `json:"is_archived"`
}

type Client struct {
    ID         string `json:"id"`
    Name       string `json:"name"`
    IsArchived bool   `json:"is_archived"`
}
```

**Rationale:**
- Matches actual API response fields
- `ClientID` is pointer to handle projects without clients (nullable in API)
- `IsArchived` needed to filter archived items
- Minimal - only fields we actually use

### 3. Client Lookup Strategy

**Decision:** Fetch all clients into a map for name lookups when displaying projects.

**Rationale:**
- Projects API returns `client_id` (UUID), not client name
- To show "TechNative" instead of "uuid-123", we need a lookup
- Two API calls: one for clients, one for projects
- Build `map[clientID]clientName` for O(1) lookups

**Flow:**
```
soltty list projects
  1. Fetch all clients → build map[id]name
  2. Fetch all projects
  3. For each project, lookup client name via map
  4. Display in table format
```

**Alternatives considered:**
- Separate API call per project → Rejected: N+1 queries problem, too slow
- Embed client info in project response → Rejected: API doesn't support this

### 4. Filtering Archived Items

**Decision:** Filter archived items in Go code after fetching from API.

**Rationale:**
- API returns all items (including archived)
- No query parameter to exclude archived items in API
- Simple filter: `if !item.IsArchived { ... }`
- Consistent across both clients and projects listings

### 5. Client Name Filtering

**Decision:** Partial match, case-insensitive, in Go code.

**Implementation:**
```go
clientFilter := strings.ToLower(flagClientName)
for _, client := range clients {
    if strings.Contains(strings.ToLower(client.Name), clientFilter) {
        // match!
    }
}
```

**Rationale:**
- User-friendly: `-c tech` matches "TechNative"
- Avoids typos: case-insensitive
- Simple implementation with standard library
- Fast enough for ~15 clients

**Alternatives considered:**
- Exact match → Rejected: less user-friendly
- Regex support → Rejected: overkill for this use case
- Fuzzy matching → Rejected: adds complexity, not needed

### 6. Output Format

**Decision:**
- Clients: Simple list with counts
- Projects: Table with Client | Project columns

**Clients format:**
```
TechNative (14 projects)
BeNext (3 projects)
...
```

**Projects format:**
```
Client            | Project
------------------|---------------------------
TechNative        | TN-General
TechNative        | TMCS-Meetings
...
```

**Rationale:**
- Consistent with existing `list` command table format
- Client column helps users see project context
- Project counts give overview without showing all projects
- Alphabetical sorting makes items easy to find

**Alternatives considered:**
- Grouped output (clients as headers, projects indented) → Rejected: harder to grep/parse
- JSON output → Rejected: could add later with --json flag if needed

### 7. Sorting

**Decision:** Sort clients alphabetically by name, projects by client name then project name.

**Rationale:**
- Predictable and easy to find items
- Standard practice for CLI tools
- Go's `sort.Slice` with custom comparator

**Code pattern:**
```go
sort.Slice(clients, func(i, j int) bool {
    return clients[i].Name < clients[j].Name
})
```

## Risks / Trade-offs

**Risk:** API returns many items, could slow down command
- **Mitigation:** Current dataset is small (~15 clients, ~46 projects). If it grows, can add pagination later with `--limit` flag.

**Risk:** Client filter might match multiple clients unexpectedly
- **Mitigation:** Use partial match (contains), not prefix. User can be more specific if needed.

**Trade-off:** Two API calls for `list projects` (clients + projects)
- **Mitigation:** Necessary to show client names. Calls are fast (~100ms each). Could cache if becomes an issue.

**Risk:** Projects without `client_id` (orphaned projects)
- **Mitigation:** Handle nil `ClientID` gracefully, show "(no client)" in output.

## Migration Plan

**Deployment:**
1. Update `Project` struct in `internal/client/project.go` to include `ClientID` and `IsArchived`
2. Add `Client` struct and `GetClients()` method in `internal/client/`
3. Create `cmd/list_clients.go` with clients listing logic
4. Create `cmd/list_projects.go` with projects listing logic and `-c` flag
5. Update `cmd/list.go` to register subcommands and maintain default behavior
6. Update root command help text to mention new subcommands

**Backwards Compatibility:**
- `soltty list` continues to work as before (shows time entries)
- No changes to existing command flags or output

**Testing:**
- Manual testing with real API data
- Test archived item filtering
- Test client name filtering (partial match, case-insensitive)
- Test projects without client_id
- Verify sorting is alphabetical

## Open Questions

None - all design decisions are resolved for implementation.
