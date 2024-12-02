# Detect the OS
ifeq ($(OS),Windows_NT)
    # Windows-specific settings
    ROOT_DIR := $(CURDIR)
    RM := del /f
    PATH_SEP := \\
    GO_BUILD := go build
    GO_RUN := go run
else
    # Linux/macOS-specific settings
    ROOT_DIR := $(shell pwd)
    RM := rm -f
    PATH_SEP := /
    GO_BUILD := go build
    GO_RUN := go run
endif

# Define the root directory of the Go project (where go.mod is located)
# ROOT_DIR is already set conditionally above.

# Target to run the app using Air (ensure Air is available on Windows)
.PHONY: run
run:
ifeq ($(OS),Windows_NT)
		cd "$(ROOT_DIR)" && air -c "$(ROOT_DIR)/.air.toml"
else
		cd "$(ROOT_DIR)" && air -c "$(ROOT_DIR)/.air.toml"
endif

# Target to clean up the binary file (if any)
.PHONY: clean
clean:
ifeq ($(OS),Windows_NT)
		$(RM) "$(ROOT_DIR)$(PATH_SEP)tmp$(PATH_SEP)main.exe"
else
		$(RM) "$(ROOT_DIR)$(PATH_SEP)tmp$(PATH_SEP)main"
endif

# Target to build the application binary
.PHONY: build
build:
ifeq ($(OS),Windows_NT)
		$(GO_BUILD) -o "$(ROOT_DIR)$(PATH_SEP)tmp$(PATH_SEP)main.exe" "$(ROOT_DIR)$(PATH_SEP)cmd"
else
		$(GO_BUILD) -o "$(ROOT_DIR)$(PATH_SEP)tmp$(PATH_SEP)main" "$(ROOT_DIR)$(PATH_SEP)cmd"
endif

# Run database migrations
.PHONY: migrate
migrate:
		$(GO_RUN) "$(ROOT_DIR)$(PATH_SEP)main.go" migrate

# Target to run database seeding
.PHONY: seed
seed:
		$(GO_RUN) "$(ROOT_DIR)$(PATH_SEP)main.go" seed
