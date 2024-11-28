# Define the root directory of the Go project (where go.mod is located)
ROOT_DIR := $(shell pwd)

# Target to run the app using Air
run:
	cd $(ROOT_DIR) && air -c $(ROOT_DIR)/.air.toml

# Target to clean up the binary file (if any)
clean:
	rm -f $(ROOT_DIR)/tmp/main
