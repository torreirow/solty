## Why

When Go module dependencies change (go.mod/go.sum), the Nix flake's vendorHash becomes outdated, causing build failures for Nix users. Currently, this requires manual hash calculation and update, which is error-prone and was missed in the recent release, breaking Nix builds.

## What Changes

- Automatic detection of go.mod/go.sum changes since the last release tag
- Automatic calculation of the correct vendorHash using Nix tooling
- Automatic update of flake.nix with the new hash during the release process
- Optional skip mechanism for cases where automatic hash calculation fails
- Enhanced logging to show when and why vendorHash is being updated

## Capabilities

### New Capabilities
- `release-vendorhash-automation`: Automatic detection and update of Nix vendorHash in release.sh when Go dependencies change

### Modified Capabilities
<!-- No existing capabilities are being modified, only enhanced -->

## Impact

- **Files Modified**: release.sh, potentially flake.nix (during releases)
- **Dependencies**: Requires Nix tooling (nix-prefetch or similar) to be available during release process
- **Build Process**: Release workflow gains automatic vendorHash management
- **User Impact**: Nix flake users will no longer encounter hash mismatch errors after releases
- **Backward Compatibility**: No breaking changes; enhancement is opt-out if Nix tooling unavailable
