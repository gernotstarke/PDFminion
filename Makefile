# Makefile for PDFminion

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=pdfminion
BINARY_UNIX=$(BINARY_NAME)_unix

# Build information
BUILDTIME=$(shell date -u +'%Y %b %d %H:%M')

# Build flags
LDFLAGS=-ldflags "-s -w -X 'pdfminion/internal/config.BuildTime=$(BUILDTIME)'"

# Install directory
INSTALL_DIR=/usr/local/bin

.PHONY: all build clean test run install uninstall

all: test build

build:
	@echo "Build Time: $(BUILDTIME)"
	@echo "LDFLAGS: $(LDFLAGS)"
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) -v ./cmd/pdfminion

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) -v ./cmd/pdfminion
	./$(BINARY_NAME)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BINARY_UNIX) -v ./cmd/pdfminion

# Install the binary
install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)"
	@sudo mv $(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Installation complete. You can now run '$(BINARY_NAME)' from anywhere."

# Uninstall the binary
uninstall:
	@echo "Uninstalling $(BINARY_NAME) from $(INSTALL_DIR)"
	@sudo rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Uninstallation complete. $(BINARY_NAME) has been removed."