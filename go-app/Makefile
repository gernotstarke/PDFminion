# Makefile for PDFminion

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=pdfminion
VERSION=0.3.1

# Build directories
BUILD_DIR=build
DIST_DIR=dist

# Build information
BUILDTIME=$(shell date -u +'%Y %b %d %H:%M')
COMMIT=$(shell git rev-parse --short HEAD)

# Detect host platform
UNAME_S := $(shell uname -s)
UNAME_M := $(shell uname -m)

# Convert to normalized platform string
ifeq ($(UNAME_S),Darwin)
    HOST_OS := MacOS
else ifeq ($(UNAME_S),Linux)
    HOST_OS := linux
else
    HOST_OS := unknown
endif

ifeq ($(UNAME_M),x86_64)
    HOST_ARCH := amd64
else ifeq ($(UNAME_M),arm64)
    HOST_ARCH := arm64
else
    HOST_ARCH := unknown
endif

HOST_PLATFORM := $(HOST_OS)-$(HOST_ARCH)

# Add to LDFLAGS
LDFLAGS=-ldflags "-s -w \
    -X 'pdfminion/internal/cli.buildTime=$(BUILDTIME)' \
    -X 'pdfminion/internal/cli.hostPlatform=$(HOST_PLATFORM)'"
# Install directory (for Unix-like systems)
INSTALL_DIR=/usr/local/bin

# Platform specific settings
WINDOWS_AMD64=windows-amd64
LINUX_AMD64=linux-amd64
DARWIN_AMD64=darwin-amd64
DARWIN_ARM64=darwin-arm64

.PHONY: all build clean test run install uninstall release \
        build-windows-amd64 build-linux-amd64 build-darwin-amd64 build-darwin-arm64 \
        package-windows-amd64 package-linux-amd64 package-darwin-amd64 package-darwin-arm64 \
        mac

# Directory creation
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

$(DIST_DIR):
	mkdir -p $(DIST_DIR)

# Standard development build (current platform)
build: $(BUILD_DIR)
	@echo "Building for current platform..."
	@echo "Build Time: $(BUILDTIME)"
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) -v ./cmd/pdfminion

# Shortcut for Apple Silicon build
mac: build-darwin-arm64
	@echo "Apple Silicon build available in $(BUILD_DIR)/$(BINARY_NAME)-$(DARWIN_ARM64)/$(BINARY_NAME)"

# Platform specific builds
build-windows-amd64: $(BUILD_DIR)
	@echo "Building for Windows (amd64)..."
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) \
		-o $(BUILD_DIR)/$(BINARY_NAME)-$(WINDOWS_AMD64)/$(BINARY_NAME).exe \
		-v ./cmd/pdfminion

build-linux-amd64: $(BUILD_DIR)
	@echo "Building for Linux (amd64)..."
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) \
		-o $(BUILD_DIR)/$(BINARY_NAME)-$(LINUX_AMD64)/$(BINARY_NAME) \
		-v ./cmd/pdfminion

build-darwin-amd64: $(BUILD_DIR)
	@echo "Building for macOS (amd64)..."
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) \
		-o $(BUILD_DIR)/$(BINARY_NAME)-$(DARWIN_AMD64)/$(BINARY_NAME) \
		-v ./cmd/pdfminion

build-darwin-arm64: $(BUILD_DIR)
	@echo "Building for macOS (Apple Silicon)..."
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) \
		-o $(BUILD_DIR)/$(BINARY_NAME)-$(DARWIN_ARM64)/$(BINARY_NAME) \
		-v ./cmd/pdfminion

# Platform specific packaging
package-windows-amd64: build-windows-amd64 $(DIST_DIR)
	@echo "Packaging Windows (amd64) build..."
	cd $(BUILD_DIR) && \
	zip -r ../$(DIST_DIR)/$(BINARY_NAME)-$(WINDOWS_AMD64)-$(VERSION).zip \
		$(BINARY_NAME)-$(WINDOWS_AMD64)
	@echo "Windows package created in $(DIST_DIR)"

package-linux-amd64: build-linux-amd64 $(DIST_DIR)
	@echo "Packaging Linux (amd64) build..."
	cd $(BUILD_DIR) && \
	tar czf ../$(DIST_DIR)/$(BINARY_NAME)-$(LINUX_AMD64)-$(VERSION).tar.gz \
		$(BINARY_NAME)-$(LINUX_AMD64)
	@echo "Linux package created in $(DIST_DIR)"

package-darwin-amd64: build-darwin-amd64 $(DIST_DIR)
	@echo "Packaging macOS (amd64) build..."
	cd $(BUILD_DIR) && \
	tar czf ../$(DIST_DIR)/$(BINARY_NAME)-$(DARWIN_AMD64)-$(VERSION).tar.gz \
		$(BINARY_NAME)-$(DARWIN_AMD64)
	@echo "macOS package created in $(DIST_DIR)"

package-darwin-arm64: build-darwin-arm64 $(DIST_DIR)
	@echo "Packaging macOS (Apple Silicon) build..."
	cd $(BUILD_DIR) && \
	tar czf ../$(DIST_DIR)/$(BINARY_NAME)-$(DARWIN_ARM64)-$(VERSION).tar.gz \
		$(BINARY_NAME)-$(DARWIN_ARM64)
	@echo "macOS (Apple Silicon) package created in $(DIST_DIR)"

# Build all platforms
release: package-windows-amd64 package-linux-amd64 package-darwin-amd64 package-darwin-arm64
	@echo "All platform builds completed!"
	@ls -l $(DIST_DIR)

# Build darwin universal binary
build-darwin-universal: $(BUILD_DIR)
	@echo "Building Universal macOS binary..."
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) \
		-o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64/$(BINARY_NAME) \
		-v ./cmd/pdfminion
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) \
		-o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64/$(BINARY_NAME) \
		-v ./cmd/pdfminion
	lipo -create \
		$(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64/$(BINARY_NAME) \
		$(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64/$(BINARY_NAME) \
		-output $(BUILD_DIR)/$(BINARY_NAME)-darwin-universal/$(BINARY_NAME)

package-darwin-universal: build-darwin-universal $(DIST_DIR)
	@echo "Packaging Universal macOS build..."
	cd $(BUILD_DIR) && \
	tar czf ../$(DIST_DIR)/$(BINARY_NAME)-darwin-universal-$(VERSION).tar.gz \
		$(BINARY_NAME)-darwin-universal
	@echo "macOS Universal package created in $(DIST_DIR)"

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -rf $(DIST_DIR)

# Run tests
test:
	$(GOTEST) -v ./...

# Run development build
run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

# Install (Unix-like systems only)
install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)"
	@sudo mv $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Installation complete. You can now run '$(BINARY_NAME)' from anywhere."

# Uninstall
uninstall:
	@echo "Uninstalling $(BINARY_NAME) from $(INSTALL_DIR)"
	@sudo rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Uninstallation complete."