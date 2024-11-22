.DEFAULT_GOAL := build
.PHONY: fmt vet clean build dev release release-snapshot

APP_NAME := tester
BUILD_DIR := bin
MAIN_FILE := ./cmd/$(APP_NAME)

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

fmt: # Format the code
	@go fmt ./...

vet: # Vet the code
	@go vet ./...

clean: # Clean the code
	@rm -rf $(BUILD_DIR)

build: clean fmt vet # Build the code
	@go build -ldflags "-w -s" -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

dev:
	@go run $(MAIN_FILE)

release-snapshot:
	goreleaser release --snapshot --clean
release: build
	@if [ -z "$(GITHUB_TOKEN)" ]; then \
			echo "GITHUB_TOKEN must be set in the environment or .env file"; \
			exit 1; \
		fi
	@if [ -n "$$(git status --porcelain)" ]; then \
			echo "Git working directory is dirty. Commit or stash your changes before releasing."; \
			exit 1; \
		fi
	@if [ -z "$(version)" ] || [ -z "$(message)" ]; then \
		echo "Usage: make release version=<version> message=<message>"; \
		exit 1; \
	fi
	@version=$$(cat version.txt); \
	@message="Release v$$version"; \
	git tag -a $$version -m "$$message"
	git push origin $(version)
	GITHUB_TOKEN=$(GITHUB_TOKEN) goreleaser release --clean
