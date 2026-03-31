## Context

The project currently has a manual release process via `release.sh` that updates VERSION, builds locally, and requires manual GitHub release creation. The codebase is a Go CLI application using Cobra, with version information embedded via ldflags during build. The project supports multiple platforms through Nix flakes and mentions pre-built binaries in the README, but currently lacks automated cross-platform binary generation.

## Goals / Non-Goals

**Goals:**
- Automate release process triggered by Git tags
- Build binaries for common platforms (Linux, macOS, Windows) and architectures (amd64, arm64)
- Automatically create GitHub releases with compiled binaries attached
- Maintain existing version embedding mechanism (ldflags)
- Integrate seamlessly with existing tag-based workflow

**Non-Goals:**
- Replacing Nix flake builds (they remain as alternative installation method)
- Automated version bumping or changelog generation (release.sh can still be used for this)
- Publishing to package managers (Homebrew, apt, etc.) beyond GitHub Releases
- Docker image builds

## Decisions

### Decision 1: Use GoReleaser with GitHub Actions
**Rationale**: GoReleaser is the standard tool for Go binary releases, providing built-in support for multiple platforms, archives, checksums, and GitHub release integration. GitHub Actions provides free CI/CD for public repositories and integrates natively with GitHub Releases.

**Alternatives considered**:
- **Manual GitHub Actions with go build**: More control but requires reinventing release asset management, checksums, and archive creation
- **Goreleaser Pro features**: Not needed for this use case; free version is sufficient

### Decision 2: Trigger on Git tags matching v*.*.*
**Rationale**: Follows semantic versioning convention and allows testing workflow changes without triggering releases. The existing release.sh already creates tags, so this integrates with current workflow.

**Alternatives considered**:
- **Trigger on GitHub release creation**: Creates chicken-egg problem where you need release first, then attach binaries
- **Trigger on push to main**: Would create releases for every commit, not just versioned releases

### Decision 3: Embed version via ldflags matching current approach
**Rationale**: Project already uses `-ldflags "-X github.com/torreirow/soltty/cmd.version=${VERSION}"` pattern. GoReleaser should replicate this to maintain consistency and avoid breaking existing version command.

**Alternatives considered**:
- **Use GoReleaser's built-in version**: Would require code changes to how version is accessed
- **Build time stamping**: Less predictable and harder to correlate with git tags

### Decision 4: Build matrix covering Linux, macOS, Windows for amd64 and arm64
**Rationale**: Covers the vast majority of users. Linux and macOS users often use arm64 (Apple Silicon, server ARM), and Windows amd64 remains dominant.

**Alternatives considered**:
- **Add 386 architecture**: Diminishing returns; very few modern systems
- **Add BSD variants**: Small user base, can be added later if requested

### Decision 5: Use tar.gz for Unix and zip for Windows
**Rationale**: Follows platform conventions for archive formats. GoReleaser handles this automatically with sensible defaults.

## Risks / Trade-offs

**[Risk: Version embedding might break]** → Mitigation: Test build locally with GoReleaser before pushing workflow, verify version command output in release notes

**[Risk: GitHub Actions permissions for creating releases]** → Mitigation: Use `contents: write` permission in workflow, which is standard for release workflows

**[Risk: Binary size increase from cross-compilation]** → Trade-off: Acceptable; users prefer platform-specific binaries over universal but larger builds

**[Risk: Breaking existing Nix flake build]** → Mitigation: GoReleaser only adds automation; Nix builds remain independent and unaffected

**[Trade-off: Automated releases on every tag]** → Consider: Could require manual approval step if desired, but adds friction; prefer fast feedback loop

**[Risk: Build failures block releases]** → Mitigation: GoReleaser runs in parallel for different platforms; one failure doesn't block others. Workflow will show which platforms failed.
