## 1. Extract shared timer detection logic

- [x] 1.1 Create `GetCurrentTimer()` function in `internal/client/time_entry.go`
- [x] 1.2 Move current timer fetching logic from `cmd/current.go` to new function
- [x] 1.3 Update `cmd/current.go` to use the new shared function
- [x] 1.4 Test that `current` command still works correctly

## 2. Implement running timer detection in start command

- [x] 2.1 Call `GetCurrentTimer()` at the beginning of `cmd/start.go` Run function
- [x] 2.2 Handle case when no timer is running (proceed as normal)
- [x] 2.3 Handle case when timer is running (proceed to confirmation)
- [x] 2.4 Calculate and format elapsed time for currently running timer

## 3. Add user confirmation prompt

- [x] 3.1 Display currently running timer details (description and elapsed time)
- [x] 3.2 Prompt user with "Stop this timer and start a new one? [y/N]: "
- [x] 3.3 Read user input from stdin
- [x] 3.4 Parse input (accept 'y', 'yes', 'Y', 'YES' as confirmation)
- [x] 3.5 Default to 'no' if Enter pressed without input
- [x] 3.6 Handle user decline (display message and exit with code 0)

## 4. Implement automatic stop functionality

- [x] 4.1 Extract stop logic from `cmd/stop.go` into reusable function if needed
- [x] 4.2 Call stop timer API when user confirms
- [x] 4.3 Handle API errors during stop (display error and exit with code 2)
- [x] 4.4 Display confirmation of stopped timer with total duration

## 5. Update start command flow

- [x] 5.1 Proceed with normal start logic after successful stop
- [x] 5.2 Display confirmation of newly started timer
- [x] 5.3 Ensure proper error handling throughout the flow
- [x] 5.4 Verify exit codes match specification (0 for success, 2 for errors)

## 6. Testing

- [x] 6.1 Test start command when no timer is running (existing behavior)
- [x] 6.2 Test start command when timer is running and user confirms stop
- [x] 6.3 Test start command when timer is running and user declines stop
- [x] 6.4 Test start command when stop API call fails
- [x] 6.5 Test with various user inputs ('y', 'yes', 'n', 'no', Enter)
- [x] 6.6 Test that elapsed time is displayed correctly
- [x] 6.7 Verify all confirmation messages are clear and accurate

## 7. Documentation

- [x] 7.1 Update README.md with new start command behavior
- [x] 7.2 Add example of auto-stop workflow to README
- [x] 7.3 Update CHANGELOG.md with new feature description
- [x] 7.4 Add note in GitHub issue #1 when implemented
