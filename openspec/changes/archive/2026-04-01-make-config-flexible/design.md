## Context

Currently, the Solidtime API endpoint is hardcoded in `internal/client/client.go` as:
```go
const baseURL = "https://solidtime.tools.technative.cloud/api/v1"
```

This prevents users from:
- Using self-hosted Solidtime instances
- Testing against local development servers
- Using different environments without rebuilding

Additionally, soltty uses `~/.config/solidtime/` for configuration, which creates naming inconsistency with the CLI tool name (`soltty`).

Current config structure in `internal/config/config.go`:
- Searches: `~/.config/solidtime/`, `~/.solidtime/`, `./config.json`
- Contains: `username`, `api_token`, `workspace_id`

The `list` command also has a UX issue in `cmd/list.go`:
- Currently running timers (end time = null) show duration as "0s"
- Users cannot easily identify which timer is actively running
- Should show "running" to indicate active state

## Goals / Non-Goals

**Goals:**
- Make API endpoint configurable via config.json as required field
- Require explicit configuration (no hidden defaults)
- Move config directory to `~/.config/soltty/` (align with tool name)
- Maintain backward compatibility for config location
- Show "running" instead of "0s" for active timers in list command
- Update all documentation

**Non-Goals:**
- Support multiple endpoints simultaneously
- Auto-discover endpoint from workspace
- Migrate configs automatically (fallback is sufficient)
- Support environment variables for endpoint (can add later)

## Decisions

### Decision 1: Config field name and format

**Choice**: Add `base_url` field to config.json (required, no default)

**Rationale**:
- Clear and descriptive name
- Explicit configuration prevents confusion about which instance is being used
- No hidden magic - users know exactly where their data goes
- Makes self-hosted vs cloud usage obvious

**Example config.json**:
```json
{
  "username": "Wouter",
  "api_token": "eyJ0eXAi...",
  "workspace_id": "13516df0-...",
  "base_url": "https://solidtime.tools.technative.cloud/api/v1"
}
```

If `base_url` is omitted, show clear error:
```
Error: missing required field: base_url in ~/.config/soltty/config.json

Please add "base_url" to your config.json:
{
  ...
  "base_url": "https://solidtime.tools.technative.cloud/api/v1"
}

For TechNative Cloud, use: https://solidtime.tools.technative.cloud/api/v1
For self-hosted, use your instance URL.
```

**Alternatives considered**:
- `api_url` or `endpoint`: Less clear about what it includes
- Optional with default: Hides which instance is being used, not explicit enough
- Separate host/port fields: Over-engineered

### Decision 2: Config directory migration path

**Choice**: Primary location is `~/.config/soltty/`, fallback to old locations

**Rationale**:
- Aligns tool name with config directory
- Backward compatible via fallback
- Users can migrate when convenient
- No forced migration reduces friction

**Search order**:
1. `~/.config/soltty/config.json` (new primary)
2. `~/.config/solidtime/config.json` (old primary, now fallback)
3. `~/.solidtime/config.json` (old fallback)
4. `./config.json` (current directory, last resort)

**Alternatives considered**:
- Auto-migrate on first run: Too aggressive, might confuse users
- Keep old location: Doesn't fix naming inconsistency
- Deprecation warning: Can add later if needed

### Decision 3: Client initialization changes

**Choice**: Pass baseURL to `NewClient()` instead of hardcoding

**Rationale**:
- Separation of concerns (config vs client)
- Easier to test with different endpoints
- Client doesn't need to know about defaults

**Before**:
```go
func NewClient(token, workspaceID string) *Client {
    return &Client{
        baseURL:     baseURL, // const
        token:       token,
        workspaceID: workspaceID,
        ...
    }
}
```

**After**:
```go
func NewClient(baseURL, token, workspaceID string) *Client {
    return &Client{
        baseURL:     baseURL,
        token:       token,
        workspaceID: workspaceID,
        ...
    }
}
```

Config package ensures base_url is required:
```go
// In config.Load(), base_url is validated as required field
// No default - error if missing
client := client.NewClient(cfg.BaseURL, cfg.APIToken, cfg.WorkspaceID)
```

**Alternatives considered**:
- Keep NewClient signature, add SetBaseURL: Less clear initialization flow
- Default in client package: Tight coupling between config and client
- Optional with default in config: Hides which instance is active

### Decision 4: Validation

**Choice**: Validate base_url is present and has correct format

**Rationale**:
- Explicit configuration prevents mistakes
- Early error detection
- Prevent confusing network errors
- Guide users to correct format

**Validation**:
- Must be present (required field)
- Must be valid URL
- Must start with http:// or https://
- Should not have trailing slash (we add it in client)
- Should end with `/api/v1` (common mistake to include or omit)

**Error message if invalid**:
```
Invalid base_url in config.json: "<url>"
Expected format: https://your-solidtime-instance.com/api/v1
```

**Alternatives considered**:
- No validation: Poor UX, cryptic errors
- Strict validation (exact format): Too rigid, might break valid cases
- Auto-fix trailing slashes: Could hide user mistakes

### Decision 5: Documentation updates

**Choice**: Update README.md with new config path and base_url as required field

**Rationale**:
- Users need to know about new config location
- Examples must show base_url as required field
- Clear migration path for existing users
- Breaking change must be clearly documented

**Documentation sections to update**:
1. Configuration section: Show new path as primary
2. Config example: Show base_url as required field with example value
3. Add migration note: Users must add base_url to existing configs
4. Add breaking change warning in CHANGELOG

### Decision 6: List command display for running timers

**Choice**: Show "running" instead of "0s" for active timers in list output

**Rationale**:
- "0s" is confusing and technically incorrect
- Users need to quickly identify active timer
- Better UX alignment with `current` command

**Implementation**:
- Check if entry has null end time
- If null: display "running" in duration column
- If not null: calculate and display actual duration

**Before**:
```
Date       | Start | Duration | Description
------------------------------------------------------------
2026-04-01 | 14:58 | 0s       | MSA-reportingxr
```

**After**:
```
Date       | Start | Duration | Description
------------------------------------------------------------
2026-04-01 | 14:58 | running  | MSA-reportingxr
```

**Alternatives considered**:
- Show elapsed time like `current` command: Too dynamic, list is static snapshot
- Add separate "Status" column: Adds visual clutter for rare case
- Special formatting (color, symbol): Not all terminals support it

## Risks / Trade-offs

**Risk: Breaking change - existing configs will fail**
- Mitigation: Clear error message with exact fix needed. Document in CHANGELOG and release notes. Users can easily add the field.

**Risk: Users confused about which config location is active**
- Mitigation: Document search order clearly. Consider adding `--show-config` command in future.

**Risk: Typo in base_url causes all commands to fail**
- Mitigation: Validate URL format, show clear error message with expected format.

**Risk: Users might use http:// instead of https://**
- Mitigation: Allow both, but document https:// as recommended. API server should redirect anyway.

**Trade-off: More complex config loading logic**
- Acceptable: Code is still straightforward, backward compatibility for location worth it.

**Trade-off: Breaking change for programmatic client instantiation**
- Acceptable: Internal package, no external API guarantees. All call sites are in our codebase.

**Trade-off: Requiring base_url vs providing default**
- Acceptable: Explicit is better than implicit. Users should know which instance they're using.

## Migration Plan

**For users**:
1. **Required**: Add `base_url` field to existing config.json
   - For TechNative Cloud: `"base_url": "https://solidtime.tools.technative.cloud/api/v1"`
   - For self-hosted: use your instance URL
2. Optional: Move `~/.config/solidtime/config.json` to `~/.config/soltty/config.json`
3. Old config location continues to work (fallback)
4. List command will now show "running" for active timers (cosmetic, no action needed)

**For developers**:
1. Update `internal/config/config.go`:
   - Add `BaseURL` field to Config struct
   - Update search paths (new primary, old as fallback)
   - Add validation that base_url is required (error if missing)
   - Add URL format validation
2. Update `internal/client/client.go`:
   - Remove hardcoded `baseURL` const
   - Add `baseURL` parameter to `NewClient()`
3. Update `cmd/list.go`:
   - Check for null end time
   - Display "running" instead of "0s" for active timers
4. Update all call sites that create Client
5. Update README.md and CHANGELOG.md with breaking change notice
6. Test with both config locations
7. Test missing base_url (should error)
8. Test list command with running timer

**Rollback**: Breaking change - users must add base_url. If issues arise, can revert and use optional base_url with default.

## Open Questions

None - design is straightforward and low risk.
