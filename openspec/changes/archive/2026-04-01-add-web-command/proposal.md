## Why

Users need a quick way to access the Solidtime web interface from the CLI. Currently, they must manually navigate to the base URL in their browser, which is inefficient for workflows that frequently switch between CLI and web UI. This change enables seamless transitions between command-line time tracking and the web-based interface.

## What Changes

- Add a new `web` command to the Soltty CLI that opens the configured Solidtime base URL in the default browser
- Support automatic login functionality when possible, reducing authentication friction
- Read the base URL from the existing Soltty configuration (API endpoint)
- Provide user feedback when the browser is launched

## Capabilities

### New Capabilities
- `web-command`: CLI command that opens the Solidtime web interface in the default browser with optional auto-login

### Modified Capabilities
<!-- No existing capabilities require requirement changes -->

## Impact

- New CLI command added to the main command structure
- Requires cross-platform browser opening functionality (Linux, macOS, Windows)
- May need to integrate with existing authentication/configuration system for auto-login
- No breaking changes to existing commands or APIs
