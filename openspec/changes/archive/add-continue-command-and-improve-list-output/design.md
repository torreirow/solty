## Context

Currently, `soltty list` displays time entries in a compact format showing Date, Start, Duration, and Description. The `--id` flag shows full UUIDs (36 characters) which is verbose. There's no way to see which project an entry belongs to without additional commands, and no way to quickly restart a timer with the same context.

**Current list output:**
```
Date       | Start | Duration | Description
------------------------------------------------------------
2026-04-09 | 14:30 | 2h 15m   | Working on feature
```

**With --id flag:**
```
ID                                   | Date       | Start | Duration | Description
----------------------------------------------------------------------------------------------------
01234567-89ab-cdef-0123-456789abcdef | 2026-04-09 | 14:30 | 2h 15m   | Working on feature
```

Analysis of 500 real entries shows 8-character UUID prefixes have zero collisions, with <0.01% collision probability at current scale.

## Goals / Non-Goals

**Goals:**
- Make entry IDs visible by default using short 8-character prefixes
- Show project names in list output for quick context
- Enable quick timer restart via `continue` command using short IDs
- Maintain backward compatibility with existing `--id` flag (full UUID display)
- Provide clear error messages when ID matching fails

**Non-Goals:**
- Changing the underlying UUID storage format
- Modifying other commands (start, stop, add, delete) at this time
- Adding interactive selection UI for continue command
- Supporting fuzzy matching beyond prefix matching

## Decisions

### 1. Short ID Length: 8 Characters

**Decision:** Display first 8 characters of UUID in default list output.

**Rationale:**
- Statistical analysis of 500 entries: zero collisions at 8 chars
- Birthday paradox: ~1% collision risk at 9,300 entries, 50% at 65,000 entries
- 8 chars is the sweet spot: short enough to type, long enough to be safe
- Git uses 7-8 chars for commit SHAs with similar reasoning

**Alternatives considered:**
- 6 chars: Already safe (no collisions in dataset), but less margin for growth
- 12 chars: Safer but unnecessarily long to type
- 4 chars: Already has collisions in current dataset

### 2. Column Order: ID | Date | Start | Duration | Project | Description

**Decision:** Place ID first, Project before Description (after Duration).

**Rationale:**
- ID first: Establishes entry identity immediately
- Project before Description: Provides context for the description
- Keeps Date/Start/Duration together (temporal info block)
- Description last: Variable length, natural reading flow

**Alternative considered:**
- Project after ID: Groups identity info, but separates temporal fields awkwardly

### 3. Prefix Matching for Short IDs

**Decision:** Accept 6-36 characters, match by UUID prefix (case-insensitive).

**Rationale:**
- 6 chars minimum: Balances convenience with safety (still safe in current data)
- 36 chars maximum: Full UUID still accepted
- Prefix matching: Simple, predictable, no fuzzy logic surprises
- Case-insensitive: UUIDs are hex, case doesn't matter

**Algorithm:**
```go
func FindEntryByShortID(shortID string) (*TimeEntry, error) {
    // 1. Validate: 6-36 chars, hex + dashes only
    // 2. Fetch recent entries (last 1000)
    // 3. Filter: entry.ID starts with shortID (case-insensitive)
    // 4. Return based on match count (see error handling below)
}
```

**Alternatives considered:**
- Fuzzy matching: Too unpredictable, users prefer explicit behavior
- Database query with LIKE: Requires backend changes, not needed at current scale

### 4. Continue Command Implementation

**Decision:** New `soltty continue <short-id>` command that copies description and project only.

**Rationale:**
- Dedicated command: Clear intent (vs. `start --from-id`)
- Copy description + project: These define the "what" and "where"
- Don't copy start time/duration: These are temporal to the old entry
- Leverage existing start flow: Check for running timer, project resolution, etc.

**Flow:**
```
continue <short-id>
  └─> FindEntryByShortID(shortID)
       ├─> Error if not found/ambiguous
       └─> Call startTimeEntry(entry.Description, entry.ProjectID, nil)
            └─> (existing auto-stop logic applies)
```

### 5. Error Handling Strategy

**Decision:** Three error states with helpful messages and recovery suggestions.

**Error states:**

1. **Not found:**
   ```
   Error: No entry found with ID '98765432'
   Use 'soltty list' to see available entries (with 8-char IDs)
   Use 'soltty list --id' to see full UUIDs
   ```

2. **Ambiguous (multiple matches):**
   ```
   Error: Ambiguous ID '985d' matches multiple entries:
     985d7cb2 - 2026-04-09 14:30: TNIIT-105 am-iris-prod-03
     985d1234 - 2026-04-08 10:00: Other task

   Please use more characters (e.g., '985d7cb2')
   Use 'soltty list --id' to see full UUIDs if needed
   ```

3. **Invalid format:**
   ```
   Error: Invalid ID format 'xyz123'
   IDs must be 6-36 characters (hex digits and dashes only)
   Example: 985d7cb2
   ```

**Rationale:**
- Always suggest `soltty list` as the discovery mechanism
- Mention `--id` flag for cases where 8 chars aren't enough
- Show matching entries for ambiguous cases (helps user choose)
- Practical examples in error messages

### 6. Project Name Resolution in List

**Decision:** Fetch projects once, build a map, display names (or "No project" for null).

**Rationale:**
- Projects are relatively static, fetching once is sufficient
- Map lookup is O(1) for each entry
- Graceful degradation: "No project" for unassigned entries or lookup failures
- Project names are more useful than UUIDs for humans

**Implementation:**
```go
// In runList:
projects := c.ListProjects()  // Fetch once
projectMap := make(map[string]string)  // UUID -> Name
for _, p := range projects {
    projectMap[p.ID] = p.Name
}

// Then for each entry:
projectName := "No project"
if entry.ProjectID != nil {
    if name, ok := projectMap[*entry.ProjectID]; ok {
        projectName = name
    }
}
```

## Risks / Trade-offs

### Risk: Terminal width limitations
**Issue:** New columns make output wider, may wrap on narrow terminals (< 120 chars).

**Mitigation:**
- Most modern terminals default to 120+ columns
- Users can still use old behavior if needed (future: add `--compact` flag)
- Description truncation already happens naturally at terminal edge

### Risk: Collision as dataset grows
**Issue:** With 8-char prefixes, collisions become likely after ~9,300 entries.

**Mitigation:**
- Continue command accepts longer prefixes (up to 36 chars)
- Error message guides users to use more characters or `--id` flag
- Can be monitored: if collisions become common, revisit minimum length

### Risk: Performance of prefix matching with large entry lists
**Issue:** Fetching 1000 entries to search could be slow.

**Mitigation:**
- Current implementation already fetches entries for list
- In-memory prefix matching is fast (O(n) scan)
- If needed later: backend can add a prefix search endpoint

### Risk: Breaking change for scripts parsing list output
**Issue:** Scripts scraping `soltty list` output will break.

**Mitigation:**
- Add `--format` flag in future for machine-readable output (JSON/CSV)
- Document the change in CHANGELOG as output format update
- Most users likely use list interactively, not in scripts

## Migration Plan

**Deployment:**
1. Update `cmd/list.go` to show new columns by default
2. Add `cmd/continue.go` with short ID matching
3. Add `FindEntryByShortID()` to `internal/client/timeentry.go`
4. Update README with new list output examples and continue command docs
5. Update CHANGELOG with output format change note

**Rollback:**
- No data migration required (UUIDs unchanged)
- Rollback is simple: revert code changes
- No backward compatibility issues (new command, enhanced display only)

**Testing:**
- Manual testing with real data (already done: 500 entries analyzed)
- Test continue command with various short ID lengths (6, 8, 12, 36 chars)
- Test error cases: not found, ambiguous, invalid format
- Verify list output formatting on different terminal widths

## Open Questions

None - all design decisions resolved during exploration phase.
