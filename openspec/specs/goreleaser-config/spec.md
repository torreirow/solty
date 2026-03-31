# goreleaser-config Specification

## Purpose
TBD - created by archiving change add-goreleaser-github-actions. Update Purpose after archive.
## Requirements
### Requirement: GoReleaser configuration file exists
The project SHALL include a `.goreleaser.yaml` configuration file at the repository root that defines all build and release settings.

#### Scenario: Configuration file is present
- **WHEN** the repository is cloned
- **THEN** the file `.goreleaser.yaml` exists at the root level

### Requirement: Binary builds for multiple platforms
The configuration SHALL define builds for Linux, macOS, and Windows operating systems on amd64 and arm64 architectures.

#### Scenario: Linux amd64 build is configured
- **WHEN** GoReleaser processes the configuration
- **THEN** a build target for linux/amd64 is included

#### Scenario: Linux arm64 build is configured
- **WHEN** GoReleaser processes the configuration
- **THEN** a build target for linux/arm64 is included

#### Scenario: macOS amd64 build is configured
- **WHEN** GoReleaser processes the configuration
- **THEN** a build target for darwin/amd64 is included

#### Scenario: macOS arm64 build is configured
- **WHEN** GoReleaser processes the configuration
- **THEN** a build target for darwin/arm64 is included

#### Scenario: Windows amd64 build is configured
- **WHEN** GoReleaser processes the configuration
- **THEN** a build target for windows/amd64 is included

#### Scenario: Windows arm64 build is configured
- **WHEN** GoReleaser processes the configuration
- **THEN** a build target for windows/arm64 is included

### Requirement: Version information is embedded in binaries
The configuration SHALL embed version information using ldflags matching the pattern `-X github.com/torreirow/soltty/cmd.version={{.Version}}` to maintain compatibility with existing version display.

#### Scenario: Version is embedded during build
- **WHEN** GoReleaser builds a binary from a git tag
- **THEN** the binary includes version information accessible via `cmd.version` variable

#### Scenario: Version command shows correct version
- **WHEN** user runs `solty --version` with a GoReleaser-built binary
- **THEN** the output displays the git tag version (e.g., "v0.2.0")

### Requirement: Binary naming follows conventions
The configuration SHALL name binaries using the pattern `solty-{{.Os}}-{{.Arch}}` for clarity in release assets.

#### Scenario: Linux amd64 binary is named correctly
- **WHEN** GoReleaser creates a Linux amd64 binary
- **THEN** the binary is named `solty-linux-amd64`

#### Scenario: Windows binary includes .exe extension
- **WHEN** GoReleaser creates a Windows binary
- **THEN** the binary is named `solty-windows-amd64.exe`

### Requirement: Archives are created in platform-appropriate formats
The configuration SHALL create tar.gz archives for Unix-like systems (Linux, macOS) and zip archives for Windows.

#### Scenario: Unix archive is tar.gz format
- **WHEN** GoReleaser packages Linux or macOS binaries
- **THEN** the archive uses .tar.gz format

#### Scenario: Windows archive is zip format
- **WHEN** GoReleaser packages Windows binaries
- **THEN** the archive uses .zip format

### Requirement: Checksums are generated
The configuration SHALL generate checksums for all release artifacts to enable verification of download integrity.

#### Scenario: Checksum file is created
- **WHEN** GoReleaser creates release artifacts
- **THEN** a checksums file is generated containing SHA256 hashes of all artifacts

### Requirement: GitHub release is automatically created
The configuration SHALL create a GitHub release with all compiled binaries and checksums as release assets.

#### Scenario: Release is created on tag
- **WHEN** GoReleaser runs with a git tag
- **THEN** a GitHub release is created for that tag with all binary archives attached

