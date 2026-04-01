## 1. Update config structure

- [x] 1.1 Add `BaseURL` field to Config struct in `internal/config/config.go`
- [x] 1.2 Update config search paths to prioritize `~/.config/soltty/config.json`
- [x] 1.3 Add fallback to old paths: `~/.config/solidtime/`, `~/.solidtime/`, `./`
- [x] 1.4 Add validation that base_url is REQUIRED (error if missing, no default)
- [x] 1.5 Add URL validation for BaseURL field if provided
- [x] 1.6 Update error message to show base_url is required with example

## 2. Update client initialization

- [x] 2.1 Remove hardcoded `baseURL` constant from `internal/client/client.go`
- [x] 2.2 Add `baseURL` parameter to `NewClient()` function signature
- [x] 2.3 Update `NewClient()` to use the baseURL parameter instead of constant

## 3. Update all client instantiation call sites

- [x] 3.1 Find all locations where `client.NewClient()` is called
- [x] 3.2 Update call site to pass `config.BaseURL` (no default fallback)
- [x] 3.3 Remove default endpoint logic from getClient() function

## 4. Update list command to show "running" for active timers

- [x] 4.1 Read `cmd/list.go` to understand current duration calculation
- [x] 4.2 Add check for null end time in time entry
- [x] 4.3 Display "running" instead of duration when end time is null
- [x] 4.4 Ensure completed entries still show calculated duration

## 5. Update documentation

- [x] 5.1 Update README.md Configuration section with new primary path `~/.config/soltty/`
- [x] 5.2 Update config example to show base_url as REQUIRED field (not optional)
- [x] 5.3 Document config file search order
- [x] 5.4 Add migration note with breaking change warning
- [x] 5.5 Update CHANGELOG.md with breaking change and new features

## 6. Testing

- [x] 6.1 Test with config in new location (`~/.config/soltty/config.json`)
- [x] 6.2 Test with config in old location (`~/.config/solidtime/config.json`) as fallback
- [x] 6.3 Test with base_url field in config (should work)
- [x] 6.4 Test WITHOUT base_url field (should error with clear message)
- [x] 6.5 Test with invalid base_url format (should show error)
- [x] 6.6 Test list command with running timer (should show "running")
- [x] 6.7 Test list command with all completed entries (should show durations)
- [x] 6.8 Verify all commands still work with new config structure
