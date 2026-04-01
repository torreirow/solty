## 1. Dependencies

- [x] 1.1 Add github.com/pkg/browser dependency to go.mod
- [x] 1.2 Run go mod tidy to update dependencies

## 2. Core Implementation

- [x] 2.1 Create cmd/web.go with basic Cobra command structure
- [x] 2.2 Implement URL derivation function to extract web URL from base_url
- [x] 2.3 Add configuration loading and validation in web command
- [x] 2.4 Implement browser opening using pkg/browser library
- [x] 2.5 Add user feedback messages for successful browser launch
- [x] 2.6 Add error handling and user-friendly error messages

## 3. Command Registration

- [x] 3.1 Register web command in cmd/root.go init() function
- [x] 3.2 Update root command's Long description to include web command example

## 4. Testing

- [x] 4.1 Create unit tests for URL derivation logic
- [x] 4.2 Test with various base_url formats (different domains, ports, paths)
- [x] 4.3 Test error scenarios (missing config, invalid base_url)
- [x] 4.4 Manual testing on Linux
- [x] 4.5 Manual testing on macOS (if available)
- [x] 4.6 Manual testing on Windows (if available)

## 5. Documentation

- [x] 5.1 Add web command to README.md usage examples
- [x] 5.2 Document browser opening behavior and limitations
- [x] 5.3 Add note about auto-login being deferred to future release
