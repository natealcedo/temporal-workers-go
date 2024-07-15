# Define Go related variables to reuse them throughout the Makefile
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=myapp

# Default command when you run just `make`
all: test build

# Build the project
build:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v

# Test the project
test:
	$(GOTEST) -v ./...

# Clean the binary
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Run the project (assumes the binary is already built)
run:
	$(GOCMD) run main.go

# Fetch dependencies
deps:
	$(GOGET) -v ./...