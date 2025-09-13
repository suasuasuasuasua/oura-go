# Oura-Go Development Instructions

**ALWAYS** follow these instructions first and only fallback to additional search and context gathering if the information in these instructions is incomplete or found to be in error.

Oura-go is a Go 1.25.0 web server application that provides free Oura ring data analysis as an alternative to Oura's paid subscription service. The project consists of a single main.go file that implements a basic HTTP server.

## Working Effectively

### Bootstrap and Build
- Ensure Go 1.25.0 is installed: `go version` should show `go version go1.25.0 linux/amd64`
- Navigate to repository root: `cd /path/to/oura-go`
- Build the application: `go build -o oura-go main.go` -- takes 11-12 seconds (cold build). NEVER CANCEL. Set timeout to 30+ minutes for safety.
- Verify build: `ls -la oura-go` should show executable binary (~8MB)

### Testing
- Run tests: `go test -v ./...` -- takes <1 second. Currently no test files exist.
- Run Go module tidy: `go mod tidy` -- instant
- Format code: `go fmt ./...` -- instant
- Vet code: `go vet ./...` -- instant

### Running the Application
- **CRITICAL BUG**: Flag parsing has a bug where custom ports and addresses are ignored due to init() vs main() order issue
- Default execution: `./oura-go` -- starts server on http://127.0.0.1:8080
- Verbose mode: `./oura-go -v` -- shows startup messages
- Custom flags (DO NOT WORK due to bug): `./oura-go -port 9000 -address 0.0.0.0` -- still uses defaults
- **WORKAROUND**: Always use default port 8080 and localhost address

### Development Environment (Optional)
This project includes devenv.nix configuration but Go works fine without it:
- If using Nix/devenv: `devenv shell` (if available)
- Git hooks are configured for: gofmt, golangci-lint, golines, govet, staticcheck
- Without devenv: Use standard Go toolchain as documented above

## Validation

### Manual Testing Requirements
**ALWAYS** run through these validation steps after making any changes:

1. **Build Validation**: 
   - `go build -o oura-go main.go`
   - Verify binary exists and is executable

2. **Server Functionality**: 
   - Start server: `./oura-go -v`
   - In another terminal: `curl -s http://127.0.0.1:8080`
   - Expect response: `Hello`
   - Stop server with Ctrl+C

3. **Help System**:
   - Test help: `./oura-go -help`
   - Verify flag descriptions are shown

4. **Code Quality**:
   - `go fmt ./...` (should make no changes if code is properly formatted)
   - `go vet ./...` (should show no issues)
   - `go mod tidy` (should not modify go.mod/go.sum)

5. **Flag Bug Verification**:
   - `./oura-go -v -port 9999` (should still show "Starting server on port 8080")
   - This confirms the known flag parsing bug is still present

### Timing Expectations
- **NEVER CANCEL** any build operations
- Build time: 11-12 seconds cold, <1 second cached (set timeout 30+ minutes)
- Test time: <1 second (set timeout 10+ minutes)
- Formatting/vetting: instant
- Server startup: instant

## Common Issues and Workarounds

### Flag Parsing Bug
The application has a known bug where command-line flags for port and address are ignored:
- **Issue**: `netip.AddrPort` is constructed in `init()` before `flag.Parse()` in `main()`
- **Symptom**: `./oura-go -port 9000` still runs on port 8080
- **Workaround**: Always assume server runs on http://127.0.0.1:8080
- **Fix Location**: If fixing, move flag parsing logic from `init()` to `main()` after `flag.Parse()`

### No Test Infrastructure
- Currently no tests exist (`go test` reports "[no test files]")
- If adding tests, create `*_test.go` files following Go conventions
- Test the HTTP handler: verify `/` endpoint returns "Hello"

## Project Structure

### Repository Root Contents
```
.
├── .envrc              # direnv configuration
├── .git/               # git repository
├── .gitignore          # git ignore rules
├── LICENSE             # MIT license
├── README.md           # project documentation
├── devenv.lock         # devenv lockfile
├── devenv.nix          # Nix development environment
├── devenv.yaml         # devenv configuration
├── go.mod              # Go module definition
└── main.go             # main application code (HTTP server)
```

### Key Files
- **main.go**: Single file containing the entire application
  - HTTP server listening on configurable address/port (defaults to 127.0.0.1:8080)
  - Simple handler returning "Hello" for all requests
  - Command-line flags for verbose mode, address, and port (flags have parsing bug)
- **go.mod**: Defines module as `github.com/suasuasuasuasua/oura-go` with Go 1.25.0
- **devenv.nix**: Nix configuration for development environment with Go 1.25.0

## Quick Reference Commands

### Daily Development
```bash
# Build and test cycle
go build -o oura-go main.go    # 11-12 seconds cold, <1 second cached
go test -v ./...               # <1 second  
go fmt ./...                   # instant
go vet ./...                   # instant

# Run server (always uses port 8080 due to bug)
./oura-go -v

# Test server response
curl -s http://127.0.0.1:8080  # Should return "Hello"
```

### Troubleshooting
```bash
# Check Go version
go version  # Must be go1.25.0

# Clean build
rm -f oura-go && go build -o oura-go main.go

# Verify flags (will show bug - custom values ignored)
./oura-go -help
./oura-go -v -port 9999  # Still uses port 8080
```

## Architecture Notes

The application is currently a minimal HTTP server proof-of-concept. Future roadmap includes:
- CSV file upload and processing for Oura data
- Web-based health statistics dashboard
- Both guest mode (in-session) and account-based persistence
- Potential iOS companion app integration

The server uses Go's `net/http` package and `net/netip` for address handling. All request handling currently returns a simple "Hello" response regardless of the endpoint.