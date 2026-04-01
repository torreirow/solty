## Context

Currently, the `start` command in `cmd/start.go` creates a new time entry without checking if a timer is already running. The Solidtime API allows multiple running timers, but this leads to data quality issues. Users must manually run `soltty stop` before `soltty start`, which is cumbersome.

Current flow:
1. User runs `soltty start "New task"`
2. Command creates time entry via API
3. No check for existing running timers

The `current` command already has logic to fetch the currently running timer, which can be reused.

## Goals / Non-Goals

**Goals:**
- Detect running timer before creating a new one
- Prompt user for confirmation to stop the current timer
- Automatically stop current timer if user confirms
- Provide clear feedback about what was stopped and started
- Maintain backward compatibility (no breaking changes)

**Non-Goals:**
- Automatically stop without confirmation (too destructive)
- Support multiple concurrent running timers
- Add configuration option to skip confirmation (can add later if requested)
- Handle edge cases like network failures during stop (use existing error handling)

## Decisions

### Decision 1: When to check for running timer

**Choice**: Check for running timer at the beginning of the start command, before any API calls

**Rationale**:
- Fail fast if user declines confirmation
- Avoid creating orphaned data if stop fails
- Clear separation of concerns

**Alternatives considered**:
- Check after parsing args: Less efficient, wastes validation if user declines
- Check during API call: Too late, harder to rollback

### Decision 2: Confirmation mechanism

**Choice**: Use interactive yes/no prompt with default to 'no' (safe choice)

**Rationale**:
- Explicit user consent prevents accidental data loss
- Default 'no' is safer for scripts/automation
- Matches pattern used in other CLI tools (git, rm -i)

**Implementation**:
```go
fmt.Printf("A timer is currently running: \"%s\" (started %s ago)\n",
    currentEntry.Description, elapsed)
fmt.Print("Stop this timer and start a new one? [y/N]: ")
```

**Alternatives considered**:
- Auto-stop without prompt: Too destructive, users might lose data
- Flag `--force` to skip prompt: Adds complexity, can add later if needed
- Configurable default: Over-engineering for v1

### Decision 3: Error handling for stop operation

**Choice**: If stop fails, abort the new start command and show error

**Rationale**:
- Prevents data inconsistency (two running timers)
- Clear error message helps user resolve issue
- Matches principle of least surprise

**Approach**:
```go
if err := stopCurrentTimer(currentEntry.ID); err != nil {
    return fmt.Errorf("failed to stop current timer: %w", err)
}
// Only start new timer if stop succeeded
```

**Alternatives considered**:
- Proceed with new timer anyway: Creates data quality issues
- Retry stop operation: Adds complexity, network issues are rare

### Decision 4: Code reuse

**Choice**: Extract current timer detection logic into a shared function

**Rationale**:
- `current` command already fetches running timer
- DRY principle
- Easier to test and maintain

**Approach**:
Create shared function in `client` package:
```go
func (c *Client) GetCurrentTimer() (*TimeEntry, error)
```

Use in both `current` and `start` commands.

**Alternatives considered**:
- Duplicate logic: Harder to maintain
- Complex shared state: Over-engineered

### Decision 5: Feedback messages

**Choice**: Show both stop and start confirmations

**Rationale**:
- Users need to know both operations succeeded
- Matches existing command output format
- Helps with debugging if something goes wrong

**Format**:
```
✓ Stopped: "Previous task" (duration: 1h 23m)
✓ Started: "New task"
```

## Risks / Trade-offs

**Risk: User accidentally stops long-running timer**
→ Mitigation: Default to 'no', show timer details before prompting, clear confirmation message

**Risk: Network failure between stop and start**
→ Mitigation: Use existing error handling, show clear error message. User can manually run start again.

**Risk: Breaking existing scripts that use `soltty start`**
→ Mitigation: Prompt only appears if timer is running (rare in scripts). Scripts can handle stdin differently. Future: add `--yes` flag if needed.

**Trade-off: Extra API call to check for running timer**
→ Acceptable: Only one extra call, worth it for better UX. Most users don't start multiple timers per minute.

**Trade-off: Interactive prompt not suitable for automation**
→ Acceptable: Can add `--force` or `--yes` flag in future if users request it. Start simple.

## Migration Plan

No migration needed - this is a backward-compatible enhancement.

Deployment:
1. Implement changes in `cmd/start.go` and `internal/client/`
2. Test manually with running timer
3. Test that normal flow (no running timer) still works
4. Release as part of next version

Rollback: No special rollback needed, previous behavior is still available (manual stop then start).

## Open Questions

None - design is straightforward and low risk.
