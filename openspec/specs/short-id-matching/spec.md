## ADDED Requirements

### Requirement: Match entries by UUID prefix

The system SHALL provide functionality to match time entries by UUID prefix (short ID).

#### Scenario: Match by 8-character prefix
- **WHEN** system searches for entry with short ID "985d7cb2"
- **AND** an entry exists with UUID "985d7cb2-cb20-40a4-ad9a-627ffa5cdc77"
- **THEN** the system SHALL return that entry as a match

#### Scenario: Match by 6-character prefix (minimum)
- **WHEN** system searches for entry with short ID "985d7c"
- **AND** only one entry UUID starts with "985d7c"
- **THEN** the system SHALL return that entry as a match

#### Scenario: Match by full UUID
- **WHEN** system searches for entry with full UUID "985d7cb2-cb20-40a4-ad9a-627ffa5cdc77"
- **AND** an entry exists with that exact UUID
- **THEN** the system SHALL return that entry as a match

#### Scenario: Case-insensitive matching
- **WHEN** system searches for entry with short ID "985D7CB2"
- **AND** an entry exists with UUID "985d7cb2-cb20-40a4-ad9a-627ffa5cdc77"
- **THEN** the system SHALL return that entry as a match (case-insensitive comparison)

#### Scenario: No matching entry
- **WHEN** system searches for entry with short ID "ffffffff"
- **AND** no entry UUID starts with "ffffffff"
- **THEN** the system SHALL return zero matches

#### Scenario: Multiple matching entries
- **WHEN** system searches for entry with short ID "985d"
- **AND** two entries exist with UUIDs starting with "985d"
- **THEN** the system SHALL return both entries as matches

#### Scenario: Search scope limited to recent entries
- **WHEN** system searches for entry by short ID
- **THEN** the system SHALL search within the last 1000 entries
- **AND** SHALL NOT search the entire database (performance optimization)

#### Scenario: Prefix matching only (not substring)
- **WHEN** system searches for entry with short ID "7cb2"
- **AND** an entry exists with UUID "985d7cb2-cb20-40a4-ad9a-627ffa5cdc77"
- **THEN** the system SHALL NOT match (must match from the beginning)

#### Scenario: Validate ID format before searching
- **WHEN** system receives short ID "xyz123"
- **AND** the ID contains non-hex characters (excluding dashes)
- **THEN** the system SHALL reject the ID as invalid format
- **AND** SHALL NOT perform a search

#### Scenario: Validate ID length before searching
- **WHEN** system receives short ID with less than 6 characters
- **THEN** the system SHALL reject the ID as too short
- **AND** SHALL NOT perform a search

#### Scenario: Accept dashes in ID
- **WHEN** system searches for entry with ID "985d7cb2-cb20"
- **AND** an entry exists with UUID starting with "985d7cb2-cb20"
- **THEN** the system SHALL return that entry as a match (dashes are part of UUID format)
