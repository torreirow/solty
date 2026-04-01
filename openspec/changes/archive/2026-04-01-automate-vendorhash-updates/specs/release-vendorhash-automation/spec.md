## ADDED Requirements

### Requirement: Detect Go dependency changes

The release.sh script SHALL detect when Go module dependencies (go.mod or go.sum) have changed since the last release tag.

#### Scenario: Dependencies changed since last release
- **WHEN** release.sh runs and go.mod or go.sum differs from the last release tag
- **THEN** the script SHALL identify that vendorHash needs updating

#### Scenario: No dependency changes
- **WHEN** release.sh runs and go.mod and go.sum are unchanged from the last release tag
- **THEN** the script SHALL skip vendorHash update steps

#### Scenario: No previous release tag exists
- **WHEN** release.sh runs and there is no previous release tag
- **THEN** the script SHALL assume dependencies have changed and proceed with vendorHash update

### Requirement: Calculate correct vendorHash

The release.sh script SHALL automatically calculate the correct vendorHash value for the current Go module dependencies using Nix tooling.

#### Scenario: Successful hash calculation
- **WHEN** vendorHash needs updating and Nix is available
- **THEN** the script SHALL execute Nix build or nix-prefetch to obtain the correct hash

#### Scenario: Hash calculation with correct format
- **WHEN** hash is calculated successfully
- **THEN** the result SHALL be in the format "sha256-<base64-hash>"

#### Scenario: Nix not available
- **WHEN** vendorHash needs updating but Nix command is not found
- **THEN** the script SHALL warn the user and skip automatic hash update

#### Scenario: Hash calculation fails
- **WHEN** Nix command fails during hash calculation
- **THEN** the script SHALL warn the user with error details and skip automatic hash update

### Requirement: Update flake.nix programmatically

The release.sh script SHALL automatically update the vendorHash value in flake.nix when a new hash is calculated.

#### Scenario: Replace vendorHash in flake.nix
- **WHEN** a new vendorHash is calculated
- **THEN** the script SHALL replace the existing vendorHash line in flake.nix with the new value

#### Scenario: Preserve flake.nix formatting
- **WHEN** flake.nix is updated
- **THEN** the file SHALL maintain its original formatting except for the vendorHash line

#### Scenario: Validate flake.nix after update
- **WHEN** flake.nix is modified
- **THEN** the script SHALL verify the file is valid before proceeding

#### Scenario: Include flake.nix in release commit
- **WHEN** flake.nix is updated during release process
- **THEN** the modified flake.nix SHALL be included in the release commit

### Requirement: Provide clear feedback

The release.sh script SHALL provide clear feedback about vendorHash update operations to the user.

#### Scenario: Log when checking dependencies
- **WHEN** the script checks for Go dependency changes
- **THEN** it SHALL display a message indicating the check is in progress

#### Scenario: Log when hash is being calculated
- **WHEN** vendorHash calculation starts
- **THEN** it SHALL display a message indicating hash calculation is in progress

#### Scenario: Log successful hash update
- **WHEN** vendorHash is successfully updated in flake.nix
- **THEN** it SHALL display a success message with the new hash value

#### Scenario: Log when skipping hash update
- **WHEN** vendorHash update is skipped (no changes, Nix unavailable, or failure)
- **THEN** it SHALL display a clear message explaining why the update was skipped

#### Scenario: Warn on error with recovery instructions
- **WHEN** hash calculation or file update fails
- **THEN** it SHALL display a warning with instructions for manual hash update

### Requirement: Maintain backward compatibility

The release.sh script SHALL continue to function correctly in environments where Nix is not installed or vendorHash automation fails.

#### Scenario: Release succeeds without Nix
- **WHEN** release.sh runs in an environment without Nix installed
- **THEN** the release process SHALL complete successfully with a warning about skipped vendorHash update

#### Scenario: Release succeeds on hash calculation failure
- **WHEN** automatic hash calculation fails for any reason
- **THEN** the release process SHALL continue with a warning, allowing manual hash update later

#### Scenario: Manual hash update still possible
- **WHEN** automatic vendorHash update fails or is skipped
- **THEN** the user SHALL be able to manually update the hash using traditional methods
