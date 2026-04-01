## Context

Soltty is a CLI tool for Solidtime time tracking, built with Go and Cobra. Users configure the tool via `config.json` which includes an API `base_url` field (e.g., `https://solidtime.tools.technative.cloud/api/v1`). The system needs to add a `web` command that opens the Solidtime web interface in the user's default browser.

Currently, commands are structured as individual files in the `cmd/` directory, each implementing a Cobra command. The `root.go` file aggregates these commands and provides shared utilities like `getClient()` for API access and configuration loading.

## Goals / Non-Goals

**Goals:**
- Add a `web` command that opens the Solidtime web interface in the default browser
- Derive the web URL from the existing `base_url` configuration
- Support cross-platform browser opening (Linux, macOS, Windows)
- Provide clear user feedback on success/failure
- Optionally support automatic login when feasible

**Non-Goals:**
- Modifying the existing configuration structure or adding new config fields
- Creating a custom browser or web view
- Supporting URL customization beyond what's in the configuration
- Implementing OAuth or complex authentication flows for auto-login

## Decisions

### 1. URL Derivation Strategy

**Decision:** Extract the web URL from `base_url` by parsing the URL and using only the scheme and host portions.

**Rationale:**
- The `base_url` already contains the full API endpoint (e.g., `https://solidtime.tools.technative.cloud/api/v1`)
- The web interface is typically at the root of the same domain (e.g., `https://solidtime.tools.technative.cloud`)
- Using Go's `net/url` package provides robust parsing and handles edge cases
- No additional configuration needed from users

**Alternatives considered:**
- Adding a separate `web_url` config field → Rejected: redundant, more configuration burden
- String manipulation (substring/replace) → Rejected: fragile, error-prone with various URL formats

### 2. Browser Opening Implementation

**Decision:** Use the `github.com/pkg/browser` library for cross-platform browser opening.

**Rationale:**
- Well-maintained, proven library used by other Go CLI tools
- Handles platform detection automatically (Linux/macOS/Windows)
- Gracefully falls back across multiple methods per platform
- Simple API: `browser.OpenURL(url)`

**Alternatives considered:**
- Manual implementation with `exec.Command` → Rejected: reinventing the wheel, harder to maintain
- `github.com/skratchdot/open-golang` → Rejected: less maintained, similar functionality

### 3. Command Structure

**Decision:** Follow the existing pattern with a new `cmd/web.go` file containing a Cobra command.

**Rationale:**
- Consistent with existing command architecture (start, stop, current, etc.)
- Easy to discover and maintain
- Natural extension of the CLI structure

### 4. Auto-Login Approach

**Decision:** Defer auto-login implementation for initial release; open the base URL without authentication.

**Rationale:**
- Solidtime API authentication mechanism may not support URL-based token passing
- Security risk if tokens are exposed in browser history or server logs
- Better to ship working basic functionality first, add auto-login if feasible later
- Users can manually log in once and rely on browser session persistence

**Alternatives considered:**
- Include API token as query parameter → Rejected: security risk, may not be supported by Solidtime
- Deep linking with custom protocol → Rejected: complexity, requires Solidtime support

### 5. Error Handling

**Decision:** Validate configuration before attempting browser open, provide specific error messages.

**Rationale:**
- Existing `getClient()` pattern already validates config and provides good error messages
- Browser library errors should be wrapped with context
- Users need to know if the issue is configuration vs. browser opening

## Risks / Trade-offs

**Risk:** Browser library may fail on unusual systems or containerized environments
- **Mitigation:** Provide clear error messages with the URL so users can manually open it. Document the limitation.

**Risk:** Derived web URL may not match actual web interface location for custom deployments
- **Mitigation:** The URL derivation assumes standard Solidtime deployment. If needed, future enhancement could add optional `web_url` config override.

**Trade-off:** No auto-login in initial release reduces friction but requires manual login
- **Mitigation:** Most users will have persistent browser sessions. Can be added later if there's demand and Solidtime supports it.

**Risk:** Browser opening may be blocked by security software or enterprise policies
- **Mitigation:** Clear error messages. Users in restricted environments can copy the URL from error output.

## Migration Plan

**Deployment:**
1. Add dependency: `go get github.com/pkg/browser`
2. Implement `cmd/web.go` with URL derivation and browser opening
3. Register command in `cmd/root.go`
4. Test on Linux, macOS, and Windows
5. Update documentation with usage examples

**Rollback:**
- Low risk: new command doesn't affect existing functionality
- If issues arise, can simply remove the command registration in root.go

**Testing:**
- Unit tests for URL derivation logic
- Manual testing on all three platforms
- Test with various base_url formats (with/without trailing slash, different domains)
- Test error scenarios (missing config, invalid URL)

## Open Questions

None - all design decisions are resolved for initial implementation.
