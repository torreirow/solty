# Implementation Tasks

## 1. Input File Auto-Detection
- [x] 1.1 Make `-i` argument optional (default: time_entries.csv)
- [x] 1.2 Check if time_entries.csv exists when no -i specified
- [x] 1.3 Show clear error if default file missing

## 2. Date Range Analysis
- [x] 2.1 Create function to parse dates from CSV file
- [x] 2.2 Calculate ISO week numbers for all dates
- [x] 2.3 Detect if dates span single week or multiple weeks
- [x] 2.4 Calculate week number from date range

## 3. Smart Filename Suggestion
- [x] 3.1 Generate filename format: week-{WW}.csv for single-week ranges
- [x] 3.2 Show suggestion to user with date range info
- [x] 3.3 Prompt for confirmation: "Save as week-05.csv? (Y/n)"
- [x] 3.4 Allow user to provide custom filename on 'n'
- [x] 3.5 Handle multi-week ranges: show warning, require manual filename

## 4. File Collision Handling
- [x] 4.1 Check if suggested/specified output file already exists
- [x] 4.2 Generate timestamp suffix: YYYYMMDD_HHMMSS format
- [x] 4.3 Create new filename: {basename}-{timestamp}.csv
- [x] 4.4 Show message: "File exists, saving as: week-05-20260202_143022.csv"

## 5. Argument Handling
- [x] 5.1 Keep `-o` argument for explicit output filename (override mode)
- [x] 5.2 Make `-o` optional when smart naming is used
- [x] 5.3 When `-o` specified, skip auto-detection and confirmation
- [x] 5.4 Update argument parser help text

## 6. User Interaction Flow
- [x] 6.1 Silent mode: `-i input.csv -o output.csv` (no prompts, like before)
- [x] 6.2 Auto mode: no args → detect input, suggest output, prompt confirm
- [x] 6.3 Semi-auto: `-i input.csv` → suggest output, prompt confirm
- [x] 6.4 Display date range and week number in suggestions

## 7. Error Handling
- [x] 7.1 Handle missing input file gracefully
- [x] 7.2 Handle CSV parse errors
- [x] 7.3 Handle invalid date formats in CSV
- [x] 7.4 Handle empty CSV files
- [x] 7.5 Validate user input (y/n responses)

## 8. Testing & Validation
- [x] 8.1 Test with single-week date range
- [x] 8.2 Test with multi-week date range
- [x] 8.3 Test with existing output file (collision)
- [x] 8.4 Test with -o override (no prompts)
- [x] 8.5 Test with missing time_entries.csv
- [x] 8.6 Test with various date formats
- [x] 8.7 Test confirmation prompts (y/n/custom)
- [x] 8.8 Verify backward compatibility
