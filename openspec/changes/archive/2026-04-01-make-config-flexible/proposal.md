## Why

Currently, the Solidtime API endpoint (`https://solidtime.tools.technative.cloud/api/v1`) is hardcoded in the Go source code, making it impossible to:
- Use different Solidtime instances (e.g., self-hosted, different environments)
- Test against local development servers
- Support multiple deployment environments without rebuilding

Additionally, soltty currently uses `~/.config/solidtime/` as its configuration directory, which creates confusion:
- Users might expect soltty to have its own dedicated config directory
- The naming inconsistency between the CLI tool name (soltty) and config directory (solidtime) is confusing
- It makes it harder to have separate configurations for different tools that might interact with Solidtime

The `list` command also has a UX issue:
- Currently running timers show "0s" as duration, which is confusing
- Users cannot distinguish running vs completed entries at a glance
- The duration should indicate that the timer is actively running

## What Changes

**Configuration improvements:**
- Move the API base URL from hardcoded constant to a required field in `config.json`
- Add `base_url` as required field in config.json (no default - must be explicitly configured)
- Show clear error message if `base_url` is missing from config
- Change configuration directory from `~/.config/solidtime/` to `~/.config/soltty/`
- Update configuration loading logic to search in the new location
- Update all documentation to reflect the new config path
- Maintain backward compatibility by checking old location as fallback

**List command improvements:**
- Detect currently running timer in `list` output
- Show "running" instead of "0s" for active timers
- Make it visually clear which entry is currently active

## Capabilities

### New Capabilities
<!-- No new capabilities being introduced -->

### Modified Capabilities
- `config-loading`: Add support for configurable API endpoint and new config directory path
- `list-entries`: Show "running" instead of "0s" for currently active timers

## Impact

- **Files Modified**:
  - `internal/client/client.go` (remove hardcoded baseURL constant, accept it as parameter)
  - `internal/config/config.go` (add base_url field, update config paths)
  - `cmd/list.go` (detect running timer and show "running" instead of "0s")
  - `README.md` (update configuration documentation)
  - Any files that instantiate the client
- **User Experience**:
  - More flexible configuration options
  - Users can point to different Solidtime instances
  - Config directory aligns with tool name
  - Clearer list output - active timers show "running" instead of confusing "0s"
- **Backward Compatibility**:
  - Check `~/.config/solidtime/config.json` as fallback if `~/.config/soltty/config.json` doesn't exist
  - **Breaking change**: `base_url` must be explicitly added to existing configs
  - Clear error message guides users to add the required field
- **Migration**:
  - Users must add `base_url` field to their existing config.json
  - Users can optionally move config.json to new location `~/.config/soltty/`
  - Example migration: add `"base_url": "https://solidtime.tools.technative.cloud/api/v1"` to config
