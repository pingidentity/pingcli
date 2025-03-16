SHELL := /bin/bash

.PHONY: install fmt vet test devchecknotest devcheck importfmtlint golangcilint starttestcontainer removetestcontainer spincontainer openlocalwebapi openapp

default: install

install:
	@echo -n "Running 'go mod tidy' to ensure all dependencies are up to date..."
	@if go mod tidy; then \
        echo " SUCCESS"; \
    else \
        echo " FAILED"; \
        exit 1; \
    fi

	@echo -n "Running 'go install' to install pingcli..."
	@if go install .; then \
		echo " SUCCESS"; \
	else \
		echo " FAILED"; \
		exit 1; \
	fi

fmt:
	@echo -n "Running 'go fmt' to format the code..."
	@if go fmt ./...; then \
		echo " SUCCESS"; \
	else \
		echo " FAILED"; \
		exit 1; \
	fi

vet:
	@echo -n "Running 'go vet' to check for potential issues..."
	@if go vet ./...; then \
		echo " SUCCESS"; \
	else \
		echo " FAILED"; \
		exit 1; \
	fi

test:
	@echo "Running 'go test' to execute all pingcli tests..."
	@if go test -count=1 ./...; then \
		echo "'go test' - SUCCESS"; \
	else \
		echo "'go test' - FAILED"; \
		exit 1; \
	fi

devchecknotest: install importfmtlint fmt vet golangcilint

devcheck: devchecknotest spincontainer test removetestcontainer

importfmtlint:
	@echo -n "Running 'impi' to format import ordering..."
	@if impi --local . --scheme stdThirdPartyLocal ./...; then \
		echo " SUCCESS"; \
	else \
		echo " FAILED"; \
		exit 1; \
	fi

golangcilint:
	@echo -n "Running 'golangci-lint' to check for code quality issues..."
	@if golangci-lint run --timeout 5m ./...; then \
		echo " SUCCESS"; \
	else \
		echo " FAILED"; \
		exit 1; \
	fi

starttestcontainer: --checkneededpfenvvars --checkdocker --dockerrunpf --waitforpfhealthy

--checkneededpfenvvars:
	@echo -n "Checking for required environment variables to run PingFederate container..."
	@test -n "$$PING_IDENTITY_DEVOPS_USER" || { echo " FAILED"; echo "PING_IDENTITY_DEVOPS_USER environment variable is not set.\n\nNot Registered? Register for the DevOps Program at https://devops.pingidentity.com/how-to/devopsRegistration/."; exit 1; }
	@test -n "$$PING_IDENTITY_DEVOPS_KEY" || { echo " FAILED"; echo "PING_IDENTITY_DEVOPS_KEY environment variable is not set.\n\nNot Registered? Register for the DevOps Program at https://devops.pingidentity.com/how-to/devopsRegistration/."; exit 1; }
	@test "YES" = "$$PING_IDENTITY_ACCEPT_EULA" || { echo " FAILED"; echo "You must accept the EULA to use the PingFederate container. Set PING_IDENTITY_ACCEPT_EULA=YES to continue."; exit 1; }
	@echo " SUCCESS"

--checkdocker:
	@echo -n "Checking if Docker is running..."
	@docker info > /dev/null 2>&1 || { echo " FAILED"; echo "Docker is not running. Please start Docker and try again."; exit 1; }
	@echo " SUCCESS"

--dockerrunpf:
	@echo -n "Starting the PingFederate container..."
	@docker run --name pingcli_test_pingfederate_container \
		-d -p 9031:9031 \
		-p 9999:9999 \
		--env PING_IDENTITY_DEVOPS_USER="$${PING_IDENTITY_DEVOPS_USER}" \
		--env PING_IDENTITY_DEVOPS_KEY="$${PING_IDENTITY_DEVOPS_KEY}" \
		--env PING_IDENTITY_ACCEPT_EULA="$${PING_IDENTITY_ACCEPT_EULA}" \
		--env CREATE_INITIAL_ADMIN_USER="true" \
		-v $$(pwd)/internal/testing/pingfederate_deploy:/opt/in/instance/server/default/deploy \
		pingidentity/pingfederate:latest > /dev/null 2>&1 || { echo " FAILED"; echo "Failed to start the PingFederate container. Please check your Docker setup."; exit 1; }
	@echo " SUCCESS"

--waitforpfhealthy:
	@echo -n "Waiting for the PingFederate container to become healthy..."
	@timeout=240; \
	while test $$timeout -gt 0; do \
		status=$$(docker inspect --format='{{json .State.Health.Status}}' pingcli_test_pingfederate_container 2>/dev/null || echo ""); \
		if test "$$status" = '"healthy"'; then \
			echo " SUCCESS"; \
			exit 0; \
		fi; \
		sleep 1; \
		timeout=$$((timeout - 1)); \
	done; \
	echo " FAILED"; \
	echo "PingFederate container did not become healthy within the timeout period."; \
	echo "Current status: $$status"; \
	docker logs pingcli_test_pingfederate_container || echo "No logs available."; \
	exit 1

removetestcontainer: --checkdocker
	@echo -n "Stopping and removing the PingFederate container..."
	@if docker rm -f pingcli_test_pingfederate_container > /dev/null 2>&1; then \
		echo " SUCCESS"; \
	else \
		echo " FAILED"; \
		exit 1; \
	fi

spincontainer: removetestcontainer starttestcontainer

openlocalwebapi:
	@echo -n "Opening the PingFederate Admin API documentation in the default web browser..."
	@if open "https://localhost:9999/pf-admin-api/api-docs/#/"; then \
		echo " SUCCESS"; \
	else \
		echo " FAILED"; \
		exit 1; \
	fi

openapp:
	@echo -n "Opening the PingFederate Admin Console in the default web browser..."
	@if open "https://localhost:9999/pingfederate/app"; then \
		echo " SUCCESS"; \
	else \
		echo " FAILED"; \
		exit 1; \
	fi
