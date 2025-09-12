#!/bin/bash
# Build script for oura-go

set -e

echo "Building oura-go server..."
go build -o bin/server ./cmd/server

echo "Build complete! Binary located at: bin/server"
echo "Run with: ./bin/server"