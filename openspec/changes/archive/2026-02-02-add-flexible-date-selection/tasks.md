# Implementation Tasks

## 1. Date Utility Functions
- [x] 1.1 Create function to calculate "last week" (previous Monday-Sunday)
- [x] 1.2 Create function to convert ISO week number to date range (Monday-Sunday)
- [x] 1.3 Create date validation function (start < end)
- [x] 1.4 Create date confirmation prompt function for interactive mode

## 2. Command-Line Argument Parsing
- [x] 2.1 Import argparse module (already imported)
- [x] 2.2 Add `--start` and `--end` arguments for manual date range
- [x] 2.3 Add `--last-week` flag for automatic previous week calculation
- [x] 2.4 Add `--week` argument for ISO week number input
- [x] 2.5 Add argument validation (mutually exclusive groups)

## 3. Interactive Menu Mode
- [x] 3.1 Create interactive menu function with three options
- [x] 3.2 Implement option 1: "Last week" with automatic calculation
- [x] 3.3 Implement option 2: "Custom start/end dates" with manual entry
- [x] 3.4 Implement option 3: "Specific week number" with ISO week conversion
- [x] 3.5 Add date confirmation after selection in all modes

## 4. Main Script Refactoring
- [x] 4.1 Refactor date input logic to check for command-line arguments first
- [x] 4.2 Fall back to interactive mode when no arguments provided
- [x] 4.3 Display selected date range before API call
- [x] 4.4 Ensure all date inputs result in YYYY-MM-DD format

## 5. Testing & Validation
- [x] 5.1 Test non-interactive mode with `--start` and `--end`
- [x] 5.2 Test `--last-week` flag with various current dates
- [x] 5.3 Test `--week` with various ISO week numbers (including year boundaries)
- [x] 5.4 Test interactive menu with all three options
- [x] 5.5 Test date validation (reject invalid dates, start > end)
- [x] 5.6 Test backward compatibility (no args = old behavior + menu)
- [x] 5.7 Verify output CSV format unchanged
