# Change: Add Flexible Date Selection to solidtimexport.py

## Why

The solidtimexport.py script currently only supports interactive date entry, requiring manual input of start and end dates for each export. This creates friction in weekly workflows and prevents automation. Users performing regular weekly exports must manually calculate and enter dates each time, which is error-prone and time-consuming.

## What Changes

- Add command-line options for non-interactive date selection (`--start`, `--end`, `--last-week`, `--week`)
- Add interactive menu mode with three options: last week, custom dates, or week number
- Implement automatic date calculation for "last week" (previous Monday-Sunday)
- Implement ISO week number to date range conversion
- Add date validation (start date must be before end date)
- Add date confirmation prompt in interactive mode
- Display selected date range in both interactive and non-interactive modes
- Maintain backward compatibility (no arguments = interactive mode)

## Impact

- **Affected specs**: `solidtime-export` (new capability)
- **Affected code**: `solidtimexport.py` (major refactoring of date input logic)
- **User experience**: More flexible, automation-friendly, reduces manual errors
- **Backward compatibility**: Fully maintained - existing usage patterns still work
- **Dependencies**: No new external dependencies (uses stdlib datetime)
