## Why

Users often forget to stop their current timer before starting a new task, leading to inaccurate time tracking. Currently, running `soltty start` while a timer is already running fails, requiring users to manually run `soltty stop` first. This adds friction to the workflow and interrupts the user's context switch between tasks.

## What Changes

- Enhance the `start` command to detect when a time entry is currently running
- Add interactive confirmation prompt before stopping the current timer
- Automatically stop the current timer upon confirmation
- Start the new timer with the provided description after stopping the old one
- Abort the new start command if user declines to stop the current timer
- Display clear feedback showing which timer was stopped and which was started

## Capabilities

### New Capabilities
<!-- No new capabilities being introduced -->

### Modified Capabilities
- `solidtime-cli`: Add auto-stop behavior to the start command when a timer is already running

## Impact

- **Files Modified**: `cmd/start.go` (start command implementation)
- **User Experience**: Streamlined workflow eliminates manual stop step
- **Backward Compatibility**: No breaking changes; enhancement is opt-in via confirmation prompt
- **API Calls**: Additional API call to stop running timer before starting new one
- **Error Handling**: Enhanced to detect running timers and handle user confirmation

**GitHub Issue**: Addresses #1 - https://github.com/torreirow/soltty/issues/1
