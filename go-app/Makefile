# Makefile for PDFminion

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=pdfminion
VERSION=0.3.1

# Directories
OUTPUT_DIR=output
DIST_DIR=dist

# Build information
BUILDTIME=$(shell date -u +'%Y %b %d %H:%M')
BUILD_PLATFORM=$(shell go env GOOS)-$(shell go env GOARCH)

# Build flags
LDFLAGS=-ldflags "-s -w \
   -X 'pdfminion/internal/cli.buildTime=$(BUILDTIME)' \
   -X 'pdfminion/internal/cli.buildPlatform=$(BUILD_PLATFORM)'"

# Platform specific settings
WINDOWS_AMD64=windows-amd64
LINUX_AMD64=linux-amd64
DARWIN_AMD64=darwin-amd64
DARWIN_ARM64=darwin-arm64

# Install directory
INSTALL_DIR=/usr/local/bin

.PHONY: default all clean test compile compile-all run install uninstall release mac \
        compile-windows-amd64 compile-linux-amd64 compile-darwin-amd64 compile-darwin-arm64 \
        package-windows-amd64 package-linux-amd64 package-darwin-amd64 package-darwin-arm64 \
        docker-test-linux


# Directory targets
$(OUTPUT_DIR):
	mkdir -p $(OUTPUT_DIR)

# Directories must exist
$(OUTPUT_DIR)/$(BINARY_NAME)-$(WINDOWS_AMD64):
	mkdir -p $(OUTPUT_DIR)/$(BINARY_NAME)-$(WINDOWS_AMD64)

$(OUTPUT_DIR)/$(BINARY_NAME)-$(LINUX_AMD64):
	mkdir -p $(OUTPUT_DIR)/$(BINARY_NAME)-$(LINUX_AMD64)

$(OUTPUT_DIR)/$(BINARY_NAME)-$(DARWIN_AMD64):
	mkdir -p $(OUTPUT_DIR)/$(BINARY_NAME)-$(DARWIN_AMD64)

$(OUTPUT_DIR)/$(BINARY_NAME)-$(DARWIN_ARM64):
	mkdir -p $(OUTPUT_DIR)/$(BINARY_NAME)-$(DARWIN_ARM64)


$(DIST_DIR):
	mkdir -p $(DIST_DIR)

# Default target for current platform only
default: test compile

# Build for all platforms
all: test compile-all


# Compile for all platforms
compile-all:
	@echo "Compiling for all platforms..."
	$(MAKE) compile-windows-amd64
	$(MAKE) compile-linux-amd64
	$(MAKE) compile-darwin-amd64
	$(MAKE) compile-darwin-arm64
	@echo "All platforms compiled successfully"


# Build for current platform
compile: $(OUTPUT_DIR)
	$(ensure_dir)
	@echo "Building for current platform..."
	@echo "Build Time: $(BUILDTIME)"
	$(GOBUILD) $(LDFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME) -v ./cmd/pdfminion

# Platform specific builds
compile-windows-amd64: $(OUTPUT_DIR)/$(BINARY_NAME)-$(WINDOWS_AMD64)
	@mkdir -p $(OUTPUT_DIR)/$(BINARY_NAME)-$(WINDOWS_AMD64)
	@echo "Building for Windows (amd64)..."
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) \
	-o $(OUTPUT_DIR)/$(BINARY_NAME)-$(WINDOWS_AMD64)/$(BINARY_NAME).exe \
	-v ./cmd/pdfminion

compile-linux-amd64: $(OUTPUT_DIR)/$(BINARY_NAME)-$(LINUX_AMD64)
	@mkdir -p $(OUTPUT_DIR)/$(BINARY_NAME)-$(LINUX_AMD64)
	@echo "Building for Linux (amd64)..."
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) \
	-o $(OUTPUT_DIR)/$(BINARY_NAME)-$(LINUX_AMD64)/$(BINARY_NAME) \
	-v ./cmd/pdfminion

compile-darwin-amd64: $(OUTPUT_DIR)/$(BINARY_NAME)-$(DARWIN_AMD64)
	@mkdir -p $(OUTPUT_DIR)/$(BINARY_NAME)-$(DARWIN_AMD64)
	@echo "Building for macOS (amd64)..."
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) \
   	-o $(OUTPUT_DIR)/$(BINARY_NAME)-$(DARWIN_AMD64)/$(BINARY_NAME) \
   	-v ./cmd/pdfminion

compile-darwin-arm64: $(OUTPUT_DIR)/$(BINARY_NAME)-$(DARWIN_ARM64)
	@mkdir -p $(OUTPUT_DIR)/$(BINARY_NAME)-$(DARWIN_ARM64)
	@echo "Building for macOS (Apple Silicon)..."
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) \
	-o $(OUTPUT_DIR)/$(BINARY_NAME)-$(DARWIN_ARM64)/$(BINARY_NAME) \
	-v ./cmd/pdfminion

# Shortcut for Apple Silicon build
mac: $(OUTPUT_DIR)/$(BINARY_NAME)-$(DARWIN_ARM64)
	$(ensure_dir)
	@echo "Building for macOS (Apple Silicon)..."
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) \
   	-o $(OUTPUT_DIR)/$(BINARY_NAME)-$(DARWIN_ARM64)/$(BINARY_NAME) \
   	-v ./cmd/pdfminion
	@echo "Apple Silicon build available in $(OUTPUT_DIR)/$(BINARY_NAME)-$(DARWIN_ARM64)/$(BINARY_NAME)"

# Platform specific packaging
package-windows-amd64: compile-windows-amd64 $(DIST_DIR)
	@echo "Packaging Windows (amd64) build..."
	cd $(OUTPUT_DIR) && \
	zip -r ../$(DIST_DIR)/$(BINARY_NAME)-$(WINDOWS_AMD64)-$(VERSION).zip \
   	$(BINARY_NAME)-$(WINDOWS_AMD64)
	@echo "Windows package created in $(DIST_DIR)"

package-linux-amd64: compile-linux-amd64 $(DIST_DIR)
	@echo "Packaging Linux (amd64) build..."
	cd $(OUTPUT_DIR) && \
	tar czf ../$(DIST_DIR)/$(BINARY_NAME)-$(LINUX_AMD64)-$(VERSION).tar.gz \
   	$(BINARY_NAME)-$(LINUX_AMD64)
	@echo "Linux package created in $(DIST_DIR)"

package-darwin-amd64: compile-darwin-amd64 $(DIST_DIR)
	@echo "Packaging macOS (amd64) build..."
	cd $(OUTPUT_DIR) && \
   tar czf ../$(DIST_DIR)/$(BINARY_NAME)-$(DARWIN_AMD64)-$(VERSION).tar.gz \
	$(BINARY_NAME)-$(DARWIN_AMD64)
	@echo "macOS package created in $(DIST_DIR)"

package-darwin-arm64: compile-darwin-arm64 $(DIST_DIR)
	@echo "Packaging macOS (Apple Silicon) build..."
	cd $(OUTPUT_DIR) && \
	tar czf ../$(DIST_DIR)/$(BINARY_NAME)-$(DARWIN_ARM64)-$(VERSION).tar.gz \
		$(BINARY_NAME)-$(DARWIN_ARM64)
	@echo "macOS (Apple Silicon) package created in $(DIST_DIR)"

# Build all platforms
release: package-windows-amd64 package-linux-amd64 package-darwin-amd64 package-darwin-arm64
	@echo "All platform builds completed!"
	@ls -l $(DIST_DIR)

test:
	$(GOTEST) -v ./...

# Test binaries in containers
docker-test-linux: compile-linux-amd64
	docker build -f Dockerfile.linux-test -t pdfminion-linux-test .
	@echo "Testing Linux build..."
	docker run --rm pdfminion-linux-test ./pdfminion --version
	docker run --rm pdfminion-linux-test ./pdfminion --help


clean:
	$(GOCLEAN)
	rm -rf $(OUTPUT_DIR)
	rm -rf $(DIST_DIR)

run: compile
	./$(OUTPUT_DIR)/$(BINARY_NAME)

# Install (Unix-like systems only)
install: compile
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)"
	@sudo mv $(OUTPUT_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Installation complete. You can now run '$(BINARY_NAME)' from anywhere."

# Uninstall
uninstall:
	@echo "Uninstalling $(BINARY_NAME) from $(INSTALL_DIR)"
	@sudo rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Uninstallation complete."