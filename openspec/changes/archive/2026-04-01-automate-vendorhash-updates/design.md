## Context

The Nix flake build configuration includes a `vendorHash` that must match the hash of Go module dependencies. When dependencies change, the hash must be manually updated or builds fail. The release.sh script currently handles version updates and changelog management but does not manage vendorHash synchronization.

Current state:
- release.sh updates VERSION and CHANGELOG.md
- flake.nix has a hardcoded vendorHash value
- When go.mod/go.sum changes, vendorHash becomes stale
- Manual hash calculation required: build fails with correct hash in error message, or use `nix-prefetch`

## Goals / Non-Goals

**Goals:**
- Automatically detect when Go dependencies have changed since the last release
- Calculate the correct vendorHash using Nix tooling
- Update flake.nix programmatically during the release process
- Provide clear feedback when vendorHash is updated
- Allow the process to gracefully skip if Nix tooling is unavailable

**Non-Goals:**
- Managing vendorHash outside of the release process (development builds handle this separately)
- Validating the correctness of go.mod/go.sum (go mod tidy checks already exist)
- Handling non-Go dependency changes (flake.lock updates are separate)

## Decisions

### Decision 1: When to check for vendorHash updates

**Choice**: Check during release.sh execution, after version update but before git commit

**Rationale**:
- Natural integration point in existing workflow
- Ensures hash is current before tagging the release
- Allows including flake.nix in the same release commit if updated

**Alternatives considered**:
- Pre-commit hook: Too late, doesn't help with releases specifically
- CI/CD check: Detects but doesn't fix the problem automatically
- Development-time updates: Unnecessary churn for WIP changes

### Decision 2: Hash calculation method

**Choice**: Use `nix-prefetch` with a temporary build to calculate vendorHash

**Rationale**:
- Standard Nix approach for getting package hashes
- Works offline (no need for network access beyond what go mod already requires)
- Returns correct hash format (sha256-...)

**Implementation approach**:
```bash
# Temporarily set vendorHash to empty string to force calculation
# Build will fail but output the correct hash
nix build .#soltty 2>&1 | grep -oP 'got:\s+\K\S+' ||
  nix-prefetch '{ sha256 }: (import ./. {}).packages.x86_64-linux.soltty.goModules.overrideAttrs (_: { vendorHash = sha256; })'
```

**Alternatives considered**:
- `vendorHash = null`: Works but insecure, defeats reproducibility
- Manual calculation: Already the problem we're solving
- Git-based hashing of go.mod: Doesn't match Nix's actual vendoring hash

### Decision 3: flake.nix update mechanism

**Choice**: Use `sed` to replace the vendorHash line in-place

**Rationale**:
- Simple, doesn't require parsing Nix code
- vendorHash line has a predictable format
- Minimal dependencies (sed is ubiquitous)

**Pattern**:
```bash
sed -i "s|vendorHash = \"sha256-[^\"]*\";|vendorHash = \"${NEW_HASH}\";|" flake.nix
```

**Alternatives considered**:
- Nix language editing tools: Overcomplicated, adds dependencies
- Full file rewrite: Risks formatting changes
- Manual edit prompt: Defeats automation purpose

### Decision 4: Detecting dependency changes

**Choice**: Compare go.mod and go.sum against the last release tag

**Rationale**:
- Release tags mark stable dependency points
- Changes since last release are what matter for the new release
- Works even if development included many intermediate dependency changes

**Implementation**:
```bash
LAST_TAG=$(git describe --tags --abbrev=0 2>/dev/null || echo "")
if git diff ${LAST_TAG}..HEAD -- go.mod go.sum | grep -q .; then
  # Dependencies changed, update vendorHash
fi
```

**Alternatives considered**:
- Check against last commit: Too granular, triggers on every change
- Always recalculate: Wasteful and slow
- User prompt: Reduces automation benefits

### Decision 5: Error handling

**Choice**: Make vendorHash update optional with graceful fallback

**Rationale**:
- Nix might not be installed in all development environments
- Hash calculation can fail for various reasons
- Release process should continue even if this step fails (with warning)

**Approach**:
- Check for Nix availability before attempting updates
- Capture errors during hash calculation
- Warn user but allow release to continue
- User can manually fix if needed before pushing

## Risks / Trade-offs

**Risk: Hash calculation fails or produces wrong hash**
→ Mitigation: Test the updated flake with `nix build` before committing. If build fails, rollback flake.nix changes and warn user.

**Risk: sed replacement matches wrong line or corrupts file**
→ Mitigation: Use specific regex pattern that requires exact vendorHash format. Create backup of flake.nix before modification. Validate file parses after edit.

**Risk: Slows down release process**
→ Trade-off: Hash calculation adds ~10-30 seconds. Acceptable for release process that already includes builds and checks. Only runs when dependencies actually changed.

**Risk: Works on maintainer's system but not others**
→ Mitigation: Document Nix requirement in release.sh. Check for `nix` command availability. Provide clear error messages with manual hash calculation fallback instructions.

**Trade-off: Automated but not validated**
→ The script updates the hash automatically but doesn't verify the build succeeds with the new hash. The existing "Verify Nix flake" step in release.sh (line 241-271) provides this validation, happening after the hash update.

## Migration Plan

1. **Add vendorHash update logic to release.sh** between the "Update VERSION" step (line 236) and "Verify Nix flake" step (line 241)
2. **Test with a dependency change**: Modify go.mod, run release.sh, verify hash updates correctly
3. **Document in README.md**: Note that Nix is required for automated vendorHash updates during releases
4. **Fallback documentation**: Add comment in flake.nix explaining manual hash update process if automation fails

No rollback complexity - if the feature doesn't work, the old manual process still works.

## Open Questions

- Should we validate the new hash by attempting a nix build before committing? (Decision: Yes, existing "Verify Nix flake" step covers this)
- Should we commit flake.nix separately or in the same release commit? (Decision: Same commit, keeps VERSION/CHANGELOG/flake.nix synchronized)
