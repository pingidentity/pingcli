SHELL := /bin/bash

.PHONY: install fmt vet test devchecknotest devcheck importfmtlint

default: install

install:
	go mod tidy
	go install .

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test -parallel=4 -count=1 ./...

devchecknotest: install importfmtlint fmt vet golangcilint

devcheck: devchecknotest spincontainer test removetestcontainer

importfmtlint:
	impi --local . --scheme stdThirdPartyLocal ./...

golangcilint:
	golangci-lint run --timeout 5m ./...

starttestcontainer:
	docker run --name pingfederate_terraform_provider_container \
		-d -p 9031:9031 \
		-p 9999:9999 \
		--env-file "${HOME}/.pingidentity/config" \
		-v $$(pwd)/server-profiles/shared-profile:/opt/in \
		-v $$(pwd)/server-profiles/12.2/data.json:/opt/in/instance/bulk-config/data.json \
		pingidentity/pingfederate:12.2.0-latest
# Wait for the instance to become ready
	sleep 1
	duration=0
	while (( duration < 240 )) && ! docker logs pingfederate_terraform_provider_container 2>&1 | grep -q "Removing Imported Bulk File\|CONTAINER FAILURE"; \
	do \
	    duration=$$((duration+1)); \
		sleep 1; \
	done
# Fail if the container didn't become ready in time
	docker logs pingfederate_terraform_provider_container 2>&1 | grep -q "Removing Imported Bulk File" || \
		{ echo "PingFederate container did not become ready in time or contains errors. Logs:"; docker logs pingfederate_terraform_provider_container; exit 1; }

removetestcontainer:
	docker rm -f pingfederate_terraform_provider_container

spincontainer: removetestcontainer starttestcontainer

openlocalwebapi:
	open "https://localhost:9999/pf-admin-api/api-docs/#/"

openapp:
	open "https://localhost:9999/pingfederate/app"
