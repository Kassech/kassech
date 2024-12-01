# Define the root directory of the Go project (where go.mod is located)
ROOT_DIR := $(shell pwd)
GO_RUN := go run

# Target to run the app using Air
run:
	cd $(ROOT_DIR) && air -c $(ROOT_DIR)/.air.toml

# Target to clean up the binary file (if any)
clean:
	rm -f $(ROOT_DIR)/tmp/main

# Target to build the application binary
build:
	$(GO_BUILD) -o $(ROOT_DIR)/tmp/main $(ROOT_DIR)/cmd

# Run database migrations
migrate: 
.PHONY: migrate
migrate:
	$(GO_RUN) $(ROOT_DIR)/main.go migrate

# Target to run database seeding
.PHONY: seed
seed:
	$(GO_RUN) $(ROOT_DIR)/main.go seed
