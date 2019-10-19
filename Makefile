# Shell to use with Make
SHELL := /bin/bash

# Build Environment
PACKAGE = rossby
PBPKG = $(CURDIR)/pb
BUILD = $(CURDIR)/_build

# Commands
GOCMD = go
GODOC = godoc
GORUN = $(GOCMD) run
GOGET = $(GOCMD) get
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
PROTOC = protoc

# Export targets not associated with files.
.PHONY: all install build raft test citest clean doc protobuf

# Ensure dependencies are installed, run tests and compile
all: protobuf build test

# Install the commands and create configurations and data directories
install: build
	@ cp $(BUILD)/rossby /usr/local/bin/

# Build the various binaries and sources
build: rossby

# Build the rossby command and store in the build directory
rossby:
	@ $(GOBUILD) -o $(BUILD)/rossby ./cmd/rossby

# Target for simple testing on the command line
test:
	@ $(GOTEST) ./...

# Run Godoc server and open browser to the documentation
doc:
	$(info running go documentation server at http://localhost:6060)
	$(info type CTRL+C to exit the server)
	@ open http://localhost:6060/pkg/github.com/kansaslabs/rossby/
	@ $(GODOC) --http=:6060

# Clean build files
clean:
	@ $(GOCLEAN)
	@ find . -name "*.coverprofile" -print0 | xargs -0 rm -rf
	@ rm -rf $(BUILD)

# Compile protocol buffers
protobuf:
	$(info compiling protocol buffers …)
	@ $(PROTOC) -I $(PBPKG) $(PBPKG)/*.proto --go_out=plugins=grpc:$(PBPKG)
