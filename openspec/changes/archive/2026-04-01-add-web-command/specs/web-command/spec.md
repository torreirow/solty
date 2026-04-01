## ADDED Requirements

### Requirement: Web command opens browser
The system SHALL provide a `web` command that opens the configured Solidtime base URL in the user's default browser.

#### Scenario: User executes web command
- **WHEN** user runs `soltty web`
- **THEN** the system opens the configured Solidtime base URL in the default browser

#### Scenario: Web command with missing configuration
- **WHEN** user runs `soltty web` without a configured API endpoint
- **THEN** the system displays an error message indicating that the API endpoint must be configured first

### Requirement: Browser URL derivation
The system SHALL derive the web URL from the configured API endpoint by using the base URL portion (protocol and host).

#### Scenario: API endpoint to web URL conversion
- **WHEN** the API endpoint is configured as `https://app.solidtime.io/api/v1`
- **THEN** the web command opens `https://app.solidtime.io`

#### Scenario: Custom domain support
- **WHEN** the API endpoint is configured with a custom domain like `https://time.company.com/api/v1`
- **THEN** the web command opens `https://time.company.com`

### Requirement: Cross-platform browser support
The system SHALL open the browser using platform-appropriate mechanisms for Linux, macOS, and Windows.

#### Scenario: Linux browser opening
- **WHEN** the web command is executed on Linux
- **THEN** the system uses `xdg-open` or equivalent to open the browser

#### Scenario: macOS browser opening
- **WHEN** the web command is executed on macOS
- **THEN** the system uses `open` command to open the browser

#### Scenario: Windows browser opening
- **WHEN** the web command is executed on Windows
- **THEN** the system uses appropriate Windows API to open the browser

### Requirement: User feedback
The system SHALL provide feedback to the user when the browser is being opened.

#### Scenario: Successful browser launch
- **WHEN** the web command successfully opens the browser
- **THEN** the system displays a message indicating the URL being opened

#### Scenario: Browser launch failure
- **WHEN** the browser fails to open for any reason
- **THEN** the system displays an error message with details about the failure

### Requirement: Automatic login support
The system MAY support automatic login by including authentication tokens in the URL when available and supported by the Solidtime instance.

#### Scenario: Auto-login when token available
- **WHEN** user has a valid authentication token configured
- **THEN** the system includes the token in the URL to enable automatic login

#### Scenario: Auto-login not available
- **WHEN** no authentication token is available or auto-login is not supported
- **THEN** the system opens the base URL without authentication parameters
