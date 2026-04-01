## 1. Add vendorHash detection logic

- [x] 1.1 Add function to get the last release tag using `git describe --tags --abbrev=0`
- [x] 1.2 Add function to detect if go.mod or go.sum changed since last tag using `git diff`
- [x] 1.3 Add logging to inform user when checking for Go dependency changes
- [x] 1.4 Handle case when no previous release tag exists (first release)

## 2. Implement hash calculation

- [x] 2.1 Add function to check if `nix` command is available
- [x] 2.2 Implement hash calculation using nix build with temporary empty hash
- [x] 2.3 Parse the correct hash from nix build output (grep for "got: sha256-...")
- [x] 2.4 Add fallback to nix-prefetch if nix build approach fails
- [x] 2.5 Validate calculated hash format matches "sha256-[A-Za-z0-9+/=]+"
- [x] 2.6 Add error handling for hash calculation failures with user-friendly messages

## 3. Update flake.nix programmatically

- [x] 3.1 Create backup of flake.nix before modification
- [x] 3.2 Implement sed command to replace vendorHash line with new hash
- [x] 3.3 Validate flake.nix syntax after update (check file can be parsed)
- [x] 3.4 Add rollback mechanism if flake.nix update fails
- [x] 3.5 Stage flake.nix for git commit if successfully updated

## 4. Integrate into release.sh workflow

- [x] 4.1 Insert vendorHash update logic after VERSION update (after line 238)
- [x] 4.2 Insert vendorHash update logic before Nix flake verification (before line 241)
- [x] 4.3 Add flake.nix to files shown in release summary if updated
- [x] 4.4 Include flake.nix in the git add command if it was modified (around line 369-374)
- [x] 4.5 Update release commit message to mention flake.nix if included

## 5. Add user feedback and logging

- [x] 5.1 Add print_info message when checking Go dependencies
- [x] 5.2 Add print_info message when calculating vendorHash
- [x] 5.3 Add print_success message when vendorHash is updated successfully
- [x] 5.4 Add print_warning message when Nix is not available
- [x] 5.5 Add print_warning message when hash calculation fails with error details
- [x] 5.6 Add print_success message when no dependency changes detected
- [x] 5.7 Include instructions for manual hash update in warning messages

## 6. Documentation updates

- [x] 6.1 Add comment in flake.nix explaining manual hash update process
- [x] 6.2 Update README.md to document Nix requirement for automated releases
- [x] 6.3 Add comment in release.sh explaining the vendorHash automation step
- [x] 6.4 Document the hash calculation approach in release.sh header comments

## 7. Testing and validation

- [x] 7.1 Test with Go dependency changes (modify go.mod, run release.sh)
- [x] 7.2 Test with no Go dependency changes (verify skip behavior)
- [x] 7.3 Test in environment without Nix installed (verify graceful degradation)
- [x] 7.4 Test hash calculation failure scenarios
- [x] 7.5 Verify flake.nix builds successfully after automatic hash update
- [x] 7.6 Test first release scenario (no previous tag exists)
