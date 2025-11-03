# Use .ONESHELL to treat each recipe as a single shell script.
# This simplifies complex commands and makes the file more readable.
.ONESHELL:
SHELL := $(shell which bash || which sh)
.SHELLFLAGS := -ec

# ====================================================================================
# VARIABLES
# ====================================================================================

# Go variables
GOCMD := go
GOTIDY := $(GOCMD) mod tidy
GOINSTALL := $(GOCMD) install .
GOFMT := $(GOCMD) fmt ./...
GOVET := $(GOCMD) vet ./...
GOTEST := $(GOCMD) test -count=1
GOLANGCI_LINT := $(GOCMD) tool golangci-lint
IMPI := $(GOCMD) tool impi

# Find all directories containing Go test files.
TEST_DIRS := $(shell find . -type f -name '*_test.go' -print0 | xargs -0 -n1 dirname | sort -u)

# Docker variables
DOCKER := docker
CONTAINER_NAME := pingcli_test_pingfederate_container

# Cross-platform 'open' command.
OPEN_CMD := xdg-open
ifeq ($(shell uname), Darwin)
	OPEN_CMD := open
endif

# Helper to check for required environment variables.
# If not set, it stops execution with an error.
define check_env
	$(if $(value $1),,$(error ERROR: Environment variable '$1' is not set. $(2)))
endef

# ====================================================================================
# PHONY TARGETS
# ====================================================================================

.PHONY: help default install fmt vet test test-auth importfmtlint golangcilint devcheck devchecknotest
.PHONY: starttestcontainer removetestcontainer spincontainer openlocalwebapi openapp protogen
.PHONY: _check_env _check_ping_env _check_docker _run_pf_container _wait_for_pf _stop_pf_container
.PHONY: generate-options-docs generate-command-docs generate-all-docs

# ====================================================================================
# USER-FACING COMMANDS
# ====================================================================================

# Set the default goal to 'help'.
default: help

help: ## Display this help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install the application binaries
	@echo "  > Tidy: Ensuring dependencies are up to date..."
	$(GOTIDY)
	echo "✅ Dependencies are up to date."
	echo "  > Install: Building and installing application..."
	$(GOINSTALL)
	echo "✅ Application installed."

fmt: ## Format Go source code
	@echo "  > Fmt: Formatting Go code..."
	$(GOFMT)
	echo "✅ Go code formatted."

vet: ## Run go vet to catch suspicious constructs
	@echo "  > Vet: Analyzing source code for potential issues..."
	$(GOVET)
	echo "✅ No issues found."

importfmtlint: ## Format Go import ordering
	@echo "  > ImportFmt: Formatting import statements..."
	$(IMPI) --skip internal/proto/pingcli_command --local . --scheme stdThirdPartyLocal ./...
	echo "✅ Import statements formatted."

golangcilint: ## Run golangci-lint for comprehensive code analysis
	@echo "  > Lint: Running golangci-lint..."
	$(GOLANGCI_LINT) cache clean
	$(GOLANGCI_LINT) run --timeout 5m ./...
	echo "✅ No linting issues found."

generate-options-docs: ## Generate configuration options documentation then validate via golden tests
	@echo "  > Docs: Generating options documentation..."
	@if [ -z "$(OUTPUT)" ]; then \
		mkdir -p ./docs/dev-ux-portal-docs/general; \
		$(GOCMD) run ./tools/generate-options-docs -asciidoc -o ./docs/dev-ux-portal-docs/general/cli-configuration-settings-reference.adoc; \
		echo "✅ Documentation generated at docs/dev-ux-portal-docs/general/cli-configuration-settings-reference.adoc"; \
	else \
		$(GOCMD) run ./tools/generate-options-docs $(OUTPUT); \
		echo "✅ Documentation generated with custom OUTPUT $(OUTPUT)"; \
	fi
	@echo "  > Docs: Running golden tests for options docs..."
	@$(GOCMD) test ./tools/generate-options-docs/docgen -run TestOptionsDocGeneration >/dev/null && echo "✅ Options documentation golden test passed."

generate-command-docs: ## Generate per-command AsciiDoc pages then validate via golden tests
	@echo "  > Docs: Generating command documentation..."
	mkdir -p ./docs/dev-ux-portal-docs
	$(GOCMD) run ./tools/generate-command-docs -o ./docs/dev-ux-portal-docs $(COMMAND_DOCS_ARGS)
	echo "✅ Command docs generated in docs/dev-ux-portal-docs"
	@echo "  > Docs: Running golden tests for command docs..."
	@$(GOCMD) test ./tools/generate-command-docs -run TestCommandDocGeneration >/dev/null && echo "✅ Command documentation golden test passed."

generate-all-docs: ## Rebuild ALL docs then run golden tests for both sets
	@echo "  > Docs: Rebuilding all documentation (clean + regenerate)..."
	mkdir -p ./docs/dev-ux-portal-docs/general
	$(MAKE) generate-options-docs OUTPUT='-o docs/dev-ux-portal-docs/general/cli-configuration-settings-reference.adoc'
	$(MAKE) generate-command-docs
	@echo "✅ All documentation rebuilt and validated via golden tests."

protogen: ## Generate Go code from .proto files
	@echo "  > Protogen: Generating gRPC code from proto files..."
	protoc --proto_path=./internal/proto --go_out=./internal --go-grpc_out=./internal ./internal/proto/*.proto
	echo "✅ gRPC code generated."

test: test-auth ## Run all tests
	@echo "  > Test: Running all Go tests..."
	@for dir in $(TEST_DIRS); do \
		$(GOTEST) $$dir; \
	done
	@echo "✅ All tests passed."

devcheck: install importfmtlint fmt vet golangcilint spincontainer test removetestcontainer ## Run the full suite of development checks and tests
	@echo "✅ All development checks passed successfully."

devchecknotest: install importfmtlint fmt vet golangcilint ## Run all development checks except tests
	@echo "✅ All development checks (no tests) passed successfully."

# ====================================================================================
# DOCKER & CONTAINER COMMANDS
# ====================================================================================

starttestcontainer: _check_docker _check_env _run_pf_container _wait_for_pf ## Start the PingFederate test container
removetestcontainer: _check_docker _stop_pf_container ## Stop and remove the PingFederate test container
spincontainer: removetestcontainer starttestcontainer ## Re-spin the test container (remove and start)

openlocalwebapi: ## Open the PingFederate Admin API docs in a browser
	@echo "  > Browser: Opening PingFederate Admin API docs..."
	$(OPEN_CMD) "https://localhost:9999/pf-admin-api/api-docs/#/"
	echo "✅ Opened PingFederate Admin API docs."

openapp: ## Open the PingFederate Admin Console in a browser
	@echo "  > Browser: Opening PingFederate Admin Console..."
	$(OPEN_CMD) "https://localhost:9999/pingfederate/app"
	echo "✅ Opened PingFederate Admin Console."

# ====================================================================================
# INTERNAL HELPER TARGETS (Not intended for direct use)
# ====================================================================================

_check_env:
	@echo "  > Env: Checking Docker container variables..."
	$(call check_env,TEST_PING_IDENTITY_DEVOPS_USER,See https://devops.pingidentity.com/how-to/devopsRegistration/)
	$(call check_env,TEST_PING_IDENTITY_DEVOPS_KEY,See https://devops.pingidentity.com/how-to/devopsRegistration/)
	$(call check_env,TEST_PING_IDENTITY_ACCEPT_EULA,Set to 'YES' to accept the EULA.)
	echo "✅ Required Docker variables are set."

_check_ping_env:
	@echo "  > Env: Checking PingOne test variables..."
	$(call check_env,TEST_PINGONE_ENVIRONMENT_ID,Specify an unconfigured PingOne environment.)
	$(call check_env,TEST_PINGONE_REGION_CODE,Specify the region for the PingOne environment.)
	$(call check_env,TEST_PINGONE_WORKER_CLIENT_ID,Specify a worker app client ID with admin roles.)
	$(call check_env,TEST_PINGONE_WORKER_CLIENT_SECRET,Specify the secret for the worker app.)
	echo "✅ Required PingOne test variables are set."

_check_docker:
	@echo "  > Docker: Checking if the Docker daemon is running..."
	$(DOCKER) info > /dev/null
	echo "✅ Docker daemon is running."

_run_pf_container:
	@echo "  > Docker: Starting the PingFederate container..."
	$(DOCKER) run --name $(CONTAINER_NAME) \
		-d -p 9031:9031 -p 9999:9999 \
		--env PING_IDENTITY_DEVOPS_USER="$(TEST_PING_IDENTITY_DEVOPS_USER)" \
		--env PING_IDENTITY_DEVOPS_KEY="$(TEST_PING_IDENTITY_DEVOPS_KEY)" \
		--env PING_IDENTITY_ACCEPT_EULA="$(TEST_PING_IDENTITY_ACCEPT_EULA)" \
		--env CREATE_INITIAL_ADMIN_USER="true" \
		-v $$(pwd)/internal/testing/pingfederate_container_files/deploy:/opt/in/instance/server/default/deploy \
		pingidentity/pingfederate:latest
	echo "✅ PingFederate container started."

_wait_for_pf:
	@echo "  > Docker: Waiting for container to become healthy (up to 4 minutes)..."
	timeout=240
	while test $$timeout -gt 0; do
		status=$$(docker inspect --format='{{json .State.Health.Status}}' $(CONTAINER_NAME) 2>/dev/null || echo "")
		if test "$$status" = '"healthy"'; then
			echo "✅ Docker: Container is healthy."
			exit 0
		fi
		sleep 1
		timeout=$$((timeout - 1))
	done
	echo "Error: Container did not become healthy within the timeout period."
	$(DOCKER) logs $(CONTAINER_NAME) || echo "No logs available."
	exit 1

_stop_pf_container:
	@echo "  > Docker: Stopping and removing previous container..."
	# Using '|| true' correctly prevents an error if the container doesn't exist.
	$(DOCKER) rm -f $(CONTAINER_NAME) 2>/dev/null || true
	echo "✅ Previous container removed."