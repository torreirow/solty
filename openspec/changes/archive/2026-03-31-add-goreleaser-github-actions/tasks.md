## 1. Create GoReleaser Configuration

- [x] 1.1 Create `.goreleaser.yaml` file at repository root
- [x] 1.2 Configure build targets for Linux (amd64, arm64)
- [x] 1.3 Configure build targets for macOS (amd64, arm64)
- [x] 1.4 Configure build targets for Windows (amd64, arm64)
- [x] 1.5 Add ldflags to embed version: `-X github.com/torreirow/soltty/cmd.version={{.Version}}`
- [x] 1.6 Configure binary naming pattern: `solty-{{.Os}}-{{.Arch}}`
- [x] 1.7 Configure archives with tar.gz for Unix and zip for Windows
- [x] 1.8 Configure checksum generation for all artifacts
- [x] 1.9 Configure GitHub release creation with artifact uploads

## 2. Create GitHub Actions Workflow

- [x] 2.1 Create `.github/workflows/` directory
- [x] 2.2 Create `release.yml` workflow file
- [x] 2.3 Configure workflow to trigger on tags matching `v*.*.*`
- [x] 2.4 Add checkout step with `fetch-depth: 0` for full git history
- [x] 2.5 Add Go setup step with version 1.21 or higher
- [x] 2.6 Add GoReleaser action step
- [x] 2.7 Configure `contents: write` permission for the workflow
- [x] 2.8 Configure GITHUB_TOKEN for GoReleaser authentication
- [x] 2.9 Set workflow to run on ubuntu-latest runner

## 3. Testing and Validation

- [x] 3.1 Test GoReleaser configuration locally with `goreleaser build --snapshot --clean`
- [x] 3.2 Verify binary naming matches expected pattern
- [x] 3.3 Verify version embedding works correctly by checking binary output
- [x] 3.4 Test archives are created in correct formats (tar.gz for Unix, zip for Windows)
- [x] 3.5 Verify checksums are generated

## 4. Documentation

- [x] 4.1 Update README.md to document new automated release process
- [x] 4.2 Add note about release.sh being optional (can still be used for version/changelog management)
- [x] 4.3 Document how to trigger releases (push tag matching v*.*.*)
