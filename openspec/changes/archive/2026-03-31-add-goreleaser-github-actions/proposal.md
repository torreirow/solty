## Why

The project currently uses a manual release script (`release.sh`) that requires local builds and manual GitHub release creation. This creates friction in the release process, makes it difficult to provide cross-platform binaries, and lacks automation. Implementing GoReleaser with GitHub Actions will automate binary building for multiple platforms and automatically attach them to GitHub releases whenever a new tag is pushed.

## What Changes

- Add GoReleaser configuration (`.goreleaser.yaml`) to define build targets and release settings
- Create GitHub Actions workflow to trigger GoReleaser on new tags
- Automate binary compilation for multiple platforms (Linux, macOS, Windows) and architectures (amd64, arm64)
- Automatically create GitHub releases with compiled binaries attached
- Maintain version embedding during GoReleaser builds
- Update documentation to reflect new automated release process

## Capabilities

### New Capabilities
- `goreleaser-config`: GoReleaser configuration defining build targets, binary naming, archives, and release settings
- `github-actions-release`: GitHub Actions workflow that triggers on tag creation and runs GoReleaser to build and publish releases

### Modified Capabilities
<!-- No existing capability requirements are changing -->

## Impact

- `.goreleaser.yaml` will be added to project root
- `.github/workflows/release.yml` will be created (new workflows directory)
- `release.sh` script may become optional or deprecated in favor of just creating tags
- GitHub Actions will need GITHUB_TOKEN permissions to create releases
- Users will get pre-built binaries for multiple platforms from GitHub Releases page
- README.md installation section already references GitHub Releases, aligning with this change
