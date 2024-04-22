# Makefile to build a Golang project for multiple platforms

# Project setup
BINARY_NAME=canvafix
SRC=main.go
BIN_DIR=bin

# Build versions for all platforms
all: windows linux mac

# Build for Windows
windows:
	@mkdir -p $(BIN_DIR)
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o '${BIN_DIR}/${BINARY_NAME}-windows-amd64.exe' $(SRC)

# Build for Linux
linux:
	@mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o '${BIN_DIR}/${BINARY_NAME}-linux-amd64' $(SRC)

# Build for macOS
mac:
	@mkdir -p $(BIN_DIR)
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o '${BIN_DIR}/${BINARY_NAME}-darwin-amd64' $(SRC)

# Clean up build artifacts
clean:
	rm -rf $(BIN_DIR)

# Help
help:
	@echo "Usage:"
	@echo "  make windows      Build for Windows"
	@echo "  make linux        Build for Linux"
	@echo "  make mac          Build for macOS"
	@echo "  make clean        Remove all binaries and the bin directory"
	@echo "  make all          Build for all platforms"
