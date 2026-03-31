# code-translation Specification

## Purpose
TBD - created by archiving change improve-export-output. Update Purpose after archive.
## Requirements
### Requirement: English Variable Names
All variable names in Python scripts SHALL use English words following Python naming conventions (snake_case).

#### Scenario: Dutch variables translated
- **WHEN** reviewing code in prepareexact.py
- **THEN** variables like `medewerker` are renamed to `employee_number`
- **AND** variables like `bestand` are renamed to `file`
- **AND** variables like `datum` are renamed to `date`
- **AND** variables like `uursoort` are renamed to `hour_type`
- **AND** variables like `relatie` are renamed to `relation`
- **AND** variables like `artikel` are renamed to `article`
- **AND** variables like `notities` are renamed to `notes`
- **AND** variables like `aantal` are renamed to `quantity`

#### Scenario: Compound variable names
- **WHEN** Dutch compound words are encountered (e.g., `bestandsnaam`)
- **THEN** they are translated to English equivalents (e.g., `filename`)
- **AND** proper snake_case is used for multi-word names (e.g., `output_filename`)

#### Scenario: Loop variables and temporaries
- **WHEN** translating short-lived variables in loops
- **THEN** use clear English names (e.g., `row`, `entry`, `value`)
- **AND** maintain readability without overly verbose names

### Requirement: English Function Names
All function names SHALL use English words with descriptive names following Python conventions.

#### Scenario: Dutch function names translated
- **WHEN** functions have Dutch names (e.g., `suggereer_output_bestandsnaam`)
- **THEN** they are renamed to English (e.g., `suggest_output_filename`)
- **AND** the function behavior remains identical
- **AND** all call sites are updated accordingly

#### Scenario: Utility function names
- **WHEN** renaming utility functions
- **THEN** use clear English verbs (analyze, validate, convert, format)
- **AND** maintain consistency across the codebase

### Requirement: English Comments
All comments and inline documentation SHALL be written in English.

#### Scenario: Dutch comments translated
- **WHEN** code contains Dutch comments (e.g., "# Utility functies voor datum")
- **THEN** they are translated to English (e.g., "# Utility functions for date processing")
- **AND** the meaning is preserved accurately

#### Scenario: Inline explanations
- **WHEN** complex logic has inline comments
- **THEN** translate them to clear English explanations
- **AND** improve clarity where possible without changing meaning

### Requirement: English Docstrings
All function docstrings SHALL be written in English following Python documentation conventions.

#### Scenario: Dutch docstrings translated
- **WHEN** functions have Dutch docstrings
- **THEN** translate them to English
- **AND** maintain the same level of detail
- **AND** follow Python docstring conventions (triple quotes, clear descriptions)

#### Scenario: Docstring completeness
- **WHEN** translating docstrings
- **THEN** include function purpose, parameters, and return values where appropriate
- **AND** ensure consistency in style across all functions

### Requirement: Preserved External Interfaces
External interfaces (CLI arguments, file formats, JSON keys) SHALL remain unchanged for backward compatibility.

#### Scenario: JSON configuration keys unchanged
- **WHEN** code accesses JSON configuration files (employees.json, projects.json)
- **THEN** keys remain Dutch (e.g., "medewerker", "Relatie", "Project", "Uursoort")
- **AND** only the internal variable names are translated
- **AND** existing configuration files continue to work without modification

#### Scenario: CSV column headers unchanged
- **WHEN** reading or writing CSV files
- **THEN** column headers remain unchanged (e.g., "Medewerker", "Artikel", "Datum", "Relatie", "Project", "Notities", "Aantal")
- **AND** internal variable names use English translations

#### Scenario: Command-line arguments unchanged
- **WHEN** scripts use command-line arguments
- **THEN** argument names remain in English (already the case)
- **AND** help messages remain in English
- **AND** script behavior is identical

### Requirement: Functional Equivalence
Translated code SHALL maintain identical functionality and behavior.

#### Scenario: Logic preservation
- **WHEN** translating variable and function names
- **THEN** all algorithms and logic remain unchanged
- **AND** control flow is identical
- **AND** calculations produce the same results

#### Scenario: Error handling unchanged
- **WHEN** translating error messages and handling
- **THEN** error conditions are caught the same way
- **AND** error messages remain clear (in English)
- **AND** exit codes are unchanged

#### Scenario: Integration testing
- **WHEN** running the translated scripts with test data
- **THEN** output files are identical to pre-translation output
- **AND** all edge cases behave the same way
- **AND** performance is unchanged

### Requirement: Code Readability
Translated code SHALL improve readability and follow Python community conventions.

#### Scenario: Consistent naming
- **WHEN** reviewing translated code
- **THEN** variable names are consistent across functions
- **AND** similar concepts use similar naming patterns
- **AND** names are clear and self-documenting

#### Scenario: Python conventions
- **WHEN** variable names are chosen
- **THEN** they follow PEP 8 naming guidelines
- **AND** use snake_case for functions and variables
- **AND** avoid abbreviations unless widely recognized

