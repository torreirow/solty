# Change: Add Smart Output Naming to prepareexact.py

## Why

The prepareexact.py script currently requires manual output filename specification via `-o` argument. This creates friction in the workflow as users must:
- Manually determine the appropriate week number
- Risk overwriting existing exports accidentally
- Remember naming conventions (week-XX.csv format)

This increases cognitive load and error potential in the weekly export routine.

## What Changes

- Add automatic input file detection (time_entries.csv as default)
- Add automatic week number calculation from CSV date range
- Add smart output filename suggestion (week-XX.csv format)
- Add interactive confirmation for suggested filename
- Add override capability (-o still works)
- Add file collision detection with automatic timestamp suffix
- Add validation for multi-week date ranges

## Impact

- **Affected specs**: `exact-prepare` (new capability)
- **Affected code**: `prepareexact.py` (enhanced argument parsing and filename logic)
- **User experience**: Simplified workflow, safer (no overwrites), more intuitive
- **Backward compatibility**: Fully maintained - all existing usage patterns still work
- **Dependencies**: No new external dependencies (uses stdlib datetime)
