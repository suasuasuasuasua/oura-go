# oura-go

Oura data analysis implemented in Go.

## Project Structure

This project follows Go best practices with a clean package structure:

```
.
├── cmd/
│   └── server/          # Main application entry point
├── internal/
│   └── handlers/        # HTTP handlers and business logic
├── bin/                 # Built binaries (ignored by git)
└── build.sh            # Build script
```

## Building

### Using the build script
```bash
./build.sh
```

### Manual build
```bash
go build -o bin/server ./cmd/server
```

## Running

### From source
```bash
go run ./cmd/server
```

### From binary
```bash
./bin/server
```

The server will start on port 8080. Visit http://localhost:8080 to see the application.

## Testing

Run all tests:
```bash
go test ./...
```

## Development

This project uses Go modules and follows standard Go project layout conventions.
