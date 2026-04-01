# Soltty - Solidtime CLI

Command-line interface for Solidtime time tracking.

<img src="./images/soltty-logo.png">

## Features

- **Start/Stop timers** - Quick time tracking from the terminal
- **Auto-stop on start** - Automatically prompts to stop running timer when starting a new one
- **Add entries** - Create completed time entries with specific times
- **Current timer** - See what's running
- **List entries** - View recent time entries
- **List clients** - View all clients with project counts
- **List projects** - View all projects with client information
- **Delete entries** - Remove mistakes
- **Web interface** - Open Solidtime in your browser with one command
- **Project support** - Assign time to projects
- **Custom start times** - Backdate timers if you forgot to start

## Installation

<details>
<summary><b>Option 1: Nix Flake (Recommended)</b></summary>

```bash
# Run directly without installing
nix run github:torreirow/soltty -- start "My task"

# Install to user profile
nix profile install github:torreirow/soltty

# Or add to your NixOS/home-manager configuration:
# flake.nix
{
  inputs.soltty.url = "github:torreirow/soltty";
  # ...
  outputs = { self, nixpkgs, soltty, ... }: {
    # NixOS configuration
    nixosConfigurations.hostname = nixpkgs.lib.nixosSystem {
      modules = [
        {
          environment.systemPackages = [ soltty.packages.x86_64-linux.soltty ];
        }
      ];
    };

    # Or home-manager configuration
    homeConfigurations.username = home-manager.lib.homeManagerConfiguration {
      modules = [
        {
          home.packages = [ soltty.packages.x86_64-linux.soltty ];
        }
      ];
    };
  };
}
```

</details>

<details>
<summary><b>Option 2: Go Build</b></summary>

```bash
# Clone the repository
git clone https://github.com/torreirow/soltty.git
cd soltty

# Build with version
VERSION=$(cat VERSION)
go build -ldflags "-X github.com/torreirow/soltty/cmd.version=${VERSION}" -o soltty

# Optional: Install to PATH
go install
```

</details>

<details>
<summary><b>Option 3: Pre-built Binaries</b></summary>

Download from [GitHub Releases](https://github.com/torreirow/soltty/releases)

</details>

## Configuration

<details>
<summary><b>Setup config.json</b></summary>

Create `~/.config/soltty/config.json`:

```json
{
  "username": "Your Name",
  "api_token": "your-solidtime-api-token",
  "workspace_id": "your-workspace-uuid",
  "base_url": "https://app.example.com/api/v1"
}
```

**Getting your credentials:**
- **api_token**: Generate in Solidtime → Settings → API Tokens
- **workspace_id**: Found in Solidtime URL or organization settings
- **base_url** (required): API endpoint URL
  - Use your Solidtime instance URL (e.g., `https://app.example.com/api/v1`)
  - For self-hosted: use your instance URL

</details>

## Usage

### Start a timer

```bash
# Simple start
soltty start "Working on feature X"

# With project
soltty start "Bug fix" --project "Customer-Project"

# With custom start time (if you forgot to start)
soltty start "Morning work" --time "09:00"
soltty start "Task" --time "2026-03-31T08:00:00Z"
```

**Auto-stop feature**: If a timer is already running, `soltty start` will prompt you to stop it first:
```bash
$ soltty start "New task"
A timer is currently running: "Old task" (started 1h 23m ago)
Stop this timer and start a new one? [y/N]: y
✓ Stopped: "Old task" (duration: 1h 23m)
✓ Timer started: "New task"
```

This eliminates the need to manually run `soltty stop` before starting a new timer.

### Stop the timer

```bash
soltty stop
```

### Show current timer

```bash
soltty current
```

### Check version

```bash
soltty --version
# or
soltty version
```

### Add completed entry

```bash
# Add entry with specific times
soltty add "Meeting" --start "14:00" --end "15:30"

# With project
soltty add "Sprint planning" --start "10:00" --end "12:00" --project "Meetings"

# Full ISO8601 timestamps
soltty add "Client call" --start "2026-03-31T14:00:00Z" --end "2026-03-31T15:30:00Z"
```

### List recent entries

```bash
# Show last 10 entries
soltty list

# Show last 5 entries
soltty list --limit 5

# Show with IDs (for deletion)
soltty list --id
```

### List clients

```bash
# Show all clients with project counts
soltty list clients
```

### List projects

```bash
# Show all projects with client names
soltty list projects

# Filter projects by client (partial match, case-insensitive)
soltty list projects -c Acme
soltty list projects -c acme
soltty list projects --client "Customer Name"
```

**Note**: Archived clients and projects are automatically hidden from listings.

### Delete entry

```bash
# Get entry ID first
soltty list --id

# Delete by ID
soltty delete 01234567-89ab-cdef-0123-456789abcdef
```

### Open web interface

```bash
# Open Solidtime web interface in your browser
soltty web
```

The web URL is automatically derived from your configured API endpoint. For example, if your `base_url` is `https://app.example.com/api/v1`, the web command will open `https://app.example.com` in your default browser.

**Notes:**
- Works cross-platform (Linux, macOS, Windows)
- Uses your system's default browser
- If browser opening fails, the URL is displayed so you can manually copy it
- Requires existing browser session for authentication (no automatic login yet)

<details>
<summary><b>Configuration Locations (Advanced)</b></summary>

soltty searches for `config.json` in this order:

1. `~/.config/soltty/config.json` (recommended)
2. `~/.config/solidtime/config.json` (legacy - for backward compatibility)
3. `~/.solidtime/config.json`
4. `./config.json`

**Migration notes:**

⚠️ **Breaking change**: `base_url` is now a required field. If you have an existing config, you must add:
```json
"base_url": "https://app.example.com/api/v1"
```

Config location: If you have an existing config at `~/.config/solidtime/config.json`, it will continue to work. You can optionally move it to the new location `~/.config/soltty/config.json` to align with the tool name.

</details>

## Time Formats

soltty supports two time formats:

- **ISO8601**: `2026-03-31T14:00:00Z` (full timestamp)
- **Time only**: `14:00` (assumes today in local timezone)

## Project Names

Use project names (not IDs) with the `--project` flag. soltty will:
- Look up the project ID automatically
- Match names case-insensitively
- Suggest available projects if not found

**Finding available projects:**
```bash
# List all projects to see available names
soltty list projects

# Filter by client to find specific projects
soltty list projects -c "Client Name"
```

## Exit Codes

- `0` - Success
- `1` - User error (invalid input, validation failed)
- `2` - System error (network issues, API errors)

## Examples

**Typical workflow:**

```bash
# Morning: Start working
soltty start "Daily standup" --project "General"
soltty stop

soltty start "Feature development" --project "Customer-Project"
# ... work for a few hours ...
soltty stop

# Check what you tracked today
soltty list
```

**Forgot to start timer:**

```bash
# Oops, been working since 9am
soltty start "Morning coding" --time "09:00" --project "Customer-Project"
soltty stop
```

**Add past entry:**

```bash
# Add yesterday's meeting you forgot to track
soltty add "Client meeting" --start "2026-03-30T14:00:00Z" --end "2026-03-30T15:30:00Z" --project "Customer-Project"
```

**Fix mistakes:**

```bash
# List with IDs
soltty list --id

# Delete wrong entry
soltty delete <entry-id>
```

## Development

### Tech Stack

- **Language**: Go 1.21+
- **CLI Framework**: Cobra
- **API**: Solidtime JSON:API
- **Build**: Nix flakes
- **Documentation**: OpenSpec

### Development Workflow

This project uses [OpenSpec](https://openspec.dev) for specification-driven development:

```bash
# View all specifications
openspec list

# View specific capability spec
openspec show solidtime-cli --type spec

# Create new feature proposal
openspec proposal "Add new feature"

# Apply approved changes
openspec apply <change-id>

# Archive completed changes
openspec archive <change-id>
```

**Documentation structure:**
- `openspec/specs/` - Current capability specifications
- `openspec/changes/` - Active change proposals
- `openspec/changes/archive/` - Completed and archived changes

### Building from Source

```bash
# Clone repository
git clone https://github.com/torreirow/soltty.git
cd soltty

# Build with version
VERSION=$(cat VERSION)
go build -ldflags "-X github.com/torreirow/soltty/cmd.version=${VERSION}" -o soltty

# Or use Nix
nix build
```

### Making Changes

1. **Create proposal**: Use OpenSpec to document the change
2. **Update specs**: Modify or create specifications in `openspec/changes/`
3. **Implement**: Write code following the spec
4. **Test**: Verify all scenarios from the spec
5. **Update docs**: Keep README and CHANGELOG in sync
6. **Mark complete**: Mark the change as complete in OpenSpec

### Releasing

Releases are automated via GitHub Actions and GoReleaser. To create a new release:

**Prerequisites**: Nix must be installed for automated vendorHash updates when Go dependencies change.

1. **Update version and changelog** (use `release.sh`):
   ```bash
   ./release.sh
   ```
   This will:
   - Update VERSION and CHANGELOG.md
   - Automatically update Nix flake vendorHash if Go dependencies changed
   - Create a commit and tag

2. **Push the tag** to trigger the release:
   ```bash
   git push origin v<version>
   ```
   Example: `git push origin v0.3.0`

3. **Automated build**: GitHub Actions will automatically:
   - Build binaries for Linux, macOS, and Windows (amd64 and arm64)
   - Create archives (tar.gz for Unix, zip for Windows)
   - Generate checksums
   - Create a GitHub release with all artifacts attached

The `release.sh` script is still useful for version management and changelog updates, but the binary building and GitHub release creation are now fully automated.

## Contributing

Contributions are welcome! Please:
1. Check existing OpenSpec proposals for planned features
2. Create a new proposal for significant changes
3. Follow the OpenSpec workflow
4. Ensure tests pass and version embeds correctly
5. Update documentation

## License

MIT License - see [LICENSE](LICENSE) file for details.

Copyright (c) 2026 Wouter van der Toorren
