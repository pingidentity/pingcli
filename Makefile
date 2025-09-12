# Use .ONESHELL to treat each recipe as a single shell script.
# This simplifies complex commands and makes the file more readable.
.ONESHELL:
SHELL := $(shell which bash || which sh)

# ====================================================================================
# VARIABLES
# ====================================================================================

# Go variables
GOCMD := go
GOTIDY := $(GOCMD) mod tidy
GOINSTALL := $(GOCMD) install .
GOFMT := $(GOCMD) fmt ./...
GOVET := $(GOCMD) vet ./...
GOTEST := $(GOCMD) test -v -count=1 -timeout=15m
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

.PHONY: help default install fmt vet test importfmtlint golangcilint devcheck devchecknotest
.PHONY: starttestcontainer removetestcontainer spincontainer openlocalwebapi openapp protogen
.PHONY: _check_env _check_ping_env _check_docker _run_pf_container _wait_for_pf _stop_pf_container

# ====================================================================================
# USER-FACING COMMANDS
# ====================================================================================

# Set the default goal to 'help'.
default: help

help: ## Display this help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install the application binaries
	@echo "  > Tidy: Ensuring dependencies are up to date..."
	@$(GOTIDY)
	@echo "  > Install: Building and installing application..."
	@$(GOINSTALL)

fmt: ## Format Go source code
	@echo "  > Fmt: Formatting Go code..."
	@$(GOFMT)

vet: ## Run go vet to catch suspicious constructs
	@echo "  > Vet: Analyzing source code for potential issues..."
	@$(GOVET)

importfmtlint: ## Format Go import ordering
	@echo "  > ImportFmt: Formatting import statements..."
	@$(IMPI) --skip internal/proto/pingcli_command --local . --scheme stdThirdPartyLocal ./...

golangcilint: ## Run golangci-lint for comprehensive code analysis
	@echo "  > Lint: Running golangci-lint..."
	@$(GOLANGCI_LINT) cache clean
	@$(GOLANGCI_LINT) run --timeout 5m ./...

protogen: ## Generate Go code from .proto files
	@echo "  > Protogen: Generating gRPC code from proto files..."
	@protoc --proto_path=./internal/proto --go_out=./internal --go-grpc_out=./internal ./internal/proto/*.proto

test: _check_ping_env ## Run all tests
	@echo "  > Test: Running all Go tests..."
	set -e
	for dir in $(TEST_DIRS); do
		echo "    -> $$dir"
		$(GOTEST) $$dir
	done
	@echo "  > Test: All tests passed."

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
	@$(OPEN_CMD) "https://localhost:9999/pf-admin-api/api-docs/#/"

openapp: ## Open the PingFederate Admin Console in a browser
	@echo "  > Browser: Opening PingFederate Admin Console..."
	@$(OPEN_CMD) "https://localhost:9999/pingfederate/app"

# ====================================================================================
# INTERNAL HELPER TARGETS (Not intended for direct use)
# ====================================================================================

_check_env:
	@echo "  > Env: Checking Docker container variables..."
	@$(call check_env,TEST_PING_IDENTITY_DEVOPS_USER,See https://devops.pingidentity.com/how-to/devopsRegistration/)
	@$(call check_env,TEST_PING_IDENTITY_DEVOPS_KEY,See https://devops.pingidentity.com/how-to/devopsRegistration/)
	@$(call check_env,TEST_PING_IDENTITY_ACCEPT_EULA,Set to 'YES' to accept the EULA.)

_check_ping_env:
	@echo "  > Env: Checking PingOne test variables..."
	@$(call check_env,TEST_PINGONE_ENVIRONMENT_ID,Specify an unconfigured PingOne environment.)
	@$(call check_env,TEST_PINGONE_REGION_CODE,Specify the region for the PingOne environment.)
	@$(call check_env,TEST_PINGONE_WORKER_CLIENT_ID,Specify a worker app client ID with admin roles.)
	@$(call check_env,TEST_PINGONE_WORKER_CLIENT_SECRET,Specify the secret for the worker app.)

_check_docker:
	@echo "  > Docker: Checking if the Docker daemon is running..."
	@$(DOCKER) info > /dev/null 2>&1

_run_pf_container:
	@echo "  > Docker: Starting the PingFederate container..."
	# Removed '> /dev/null' to ensure docker errors are visible for easier debugging.
	@$(DOCKER) run --name $(CONTAINER_NAME) \
		-d -p 9031:9031 -p 9999:9999 \
		--env PING_IDENTITY_DEVOPS_USER="$(TEST_PING_IDENTITY_DEVOPS_USER)" \
		--env PING_IDENTITY_DEVOPS_KEY="$(TEST_PING_IDENTITY_DEVOPS_KEY)" \
		--env PING_IDENTITY_ACCEPT_EULA="$(TEST_PING_IDENTITY_ACCEPT_EULA)" \
		--env CREATE_INITIAL_ADMIN_USER="true" \
		-v $$(pwd)/internal/testing/pingfederate_container_files/deploy:/opt/in/instance/server/default/deploy \
		pingidentity/pingfederate:latest

_wait_for_pf:
	@echo "  > Docker: Waiting for container to become healthy (up to 4 minutes)..."
	set -e
	timeout=240
	while test $$timeout -gt 0; do
		status=$$(docker inspect --format='{{json .State.Health.Status}}' $(CONTAINER_NAME) 2>/dev/null || echo "")
		if test "$$status" = '"healthy"'; then
			echo "  > Docker: Container is healthy."
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
	@$(DOCKER) rm -f $(CONTAINER_NAME) 2>/dev/null || true