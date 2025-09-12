# Development Environment Setup

## Requirements

- [Go](https://golang.org/doc/install) 1.25.1+ (to build and test the CLI)
- [Docker](https://docs.docker.com/get-docker/) (to run PingFederate integration tests)
- [Git](https://git-scm.com/downloads) (for version control)
- Access to a PingOne environment (for integration testing)

## Quick Start

If you wish to work on the CLI, you'll first need [Go](http://www.golang.org) installed on your machine (please check the [requirements](#requirements) before proceeding).

*Note:* This project uses [Go Modules](https://blog.golang.org/using-go-modules) making it safe to work with it outside of your existing [GOPATH](http://golang.org/doc/code.html#GOPATH). The instructions that follow assume a directory in your home directory outside of the standard GOPATH (i.e `$HOME/development/pingidentity/`).

Clone repository to: `$HOME/development/pingidentity/`

```sh
mkdir -p $HOME/development/pingidentity/; cd $HOME/development/pingidentity/
git clone git@github.com:pingidentity/pingcli.git
...
```

To install the CLI for local development, run `make install`. This will build the CLI and install it in your Go binary directory.

```sh
make install
...
pingcli --version
...
```

You can also build without installing:

```sh
go build -o pingcli .
./pingcli --version
...
```

Or run directly from the project root:

```sh
go run ./ --version
...
```

## Testing the CLI

### Unit Tests

To run unit tests locally with no external dependencies, you can run `make test`.

```sh
make test
```

### Integration Tests

To run the full suite of integration tests against live Ping Identity services, you need to set up environment variables and run the complete test suite.

*Note:* Integration tests interact with real Ping Identity services. Please ensure you have appropriate access to PingOne and PingFederate environments.

#### Required Environment Variables

For PingOne integration tests:
```sh
export TEST_PINGONE_ENVIRONMENT_ID="your-environment-id"
export TEST_PINGONE_REGION_CODE="your-region-code"  # e.g., "NA", "EU", "AP"
export TEST_PINGONE_WORKER_CLIENT_ID="your-worker-client-id"
export TEST_PINGONE_WORKER_CLIENT_SECRET="your-worker-client-secret"
```

For PingFederate container tests:
```sh
export TEST_PING_IDENTITY_DEVOPS_USER="your-devops-username"
export TEST_PING_IDENTITY_DEVOPS_KEY="your-devops-key"
export TEST_PING_IDENTITY_ACCEPT_EULA="YES"
```

#### Running Integration Tests

Run all tests including container setup:
```sh
make devcheck
```

Or run tests with an existing container:
```sh
make spincontainer  # Start PingFederate container
make test           # Run tests
make removetestcontainer  # Clean up
```

## Using the CLI

After installing with `make install`, the `pingcli` command will be available in your terminal. 

### Configuration

Before using the CLI, you need to configure it with your Ping Identity service credentials:

```sh
pingcli config add-profile
New profile name: : dev
New profile description: : Development environment configuration
Set new profile as active: : y
```

Configure your profile with service connection details:

```sh
# Set PingOne configuration
pingcli config set service.pingone.region "NA"  # or "EU", "AP"
pingcli config set service.pingone.environment_id "your-environment-id"
pingcli config set service.pingone.client_id "your-client-id"
pingcli config set service.pingone.client_secret "your-client-secret"

# Set PingFederate configuration (if needed)
pingcli config set service.pingfederate.host "https://your-pf-host:9999"
pingcli config set service.pingfederate.username "administrator"
pingcli config set service.pingfederate.password "your-password"
```

### Testing Your Setup

Test your configuration by running a simple command:

```sh
pingcli request --service pingone --http-method GET environments
```

## Local SDK Changes

Occasionally, development may include changes to the [PingOne GO SDK](https://github.com/patrickcping/pingone-go-sdk-v2) or [PingFederate GO Client](https://github.com/pingidentity/pingfederate-go-client). If you'd like to develop the CLI locally using local, modified versions of these SDKs, this can be achieved by adding `replace` directives in the `go.mod` file.

For example, to use local versions of both SDKs:

```go
module github.com/pingidentity/pingcli

go 1.25.1

replace github.com/patrickcping/pingone-go-sdk-v2/management => ../pingone-go-sdk-v2/management
replace github.com/patrickcping/pingone-go-sdk-v2/mfa => ../pingone-go-sdk-v2/mfa
replace github.com/pingidentity/pingfederate-go-client/v1220 => ../pingfederate-go-client/v1220

require (
	github.com/patrickcping/pingone-go-sdk-v2/management v0.60.0
	github.com/patrickcping/pingone-go-sdk-v2/mfa v0.23.1
	github.com/pingidentity/pingfederate-go-client/v1220 v1220.0.0
  
  ...
)

...
```

Once updated, run the following to install the CLI with your local changes:

```shell
make install
```

## Development Tools

The project includes several development tools to maintain code quality:

### Code Formatting
```sh
make fmt          # Format Go code
make importfmtlint # Format import statements
```

### Code Analysis
```sh
make vet          # Run go vet
make golangcilint # Run comprehensive linting
```

### Development Check
```sh
make devchecknotest  # Run all checks except tests
make devcheck        # Run all checks including tests
```
