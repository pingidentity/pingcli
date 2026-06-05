# Ping CLI

[![Code Analysis and Tests](https://github.com/pingidentity/pingcli/actions/workflows/code-analysis-lint-test.yaml/badge.svg)](https://github.com/pingidentity/pingcli/actions/workflows/code-analysis-lint-test.yaml)
[![Docker Pulls](https://img.shields.io/docker/pulls/pingidentity/pingcli.svg)](https://hub.docker.com/r/pingidentity/pingcli)
[![GitHub release](https://img.shields.io/github/v/release/pingidentity/pingcli?include_prereleases&sort=semver)](https://github.com/pingidentity/pingcli/releases)

The Ping CLI is a unified command line interface for configuring and managing Ping Identity Services.

## Table of Contents

- [Ping CLI](#ping-cli)
  - [Table of Contents](#table-of-contents)
  - [Install](#install)
    - [Docker](#docker)
    - [macOS](#macos)
      - [Homebrew](#homebrew)
      - [Manual Installation](#manual-installation)
    - [Linux](#linux)
        - [Homebrew](#homebrew-1)
        - [Alpine (.apk)](#alpine-apk)
        - [Debian/Ubuntu (.deb)](#debianubuntu-deb)
        - [CentOS/Fedora/RHEL (.rpm)](#centosfedorarhel-rpm)
      - [Manual Installation](#manual-installation-1)
    - [Windows](#windows)
      - [Manual Installation](#manual-installation-2)
  - [Verify](#verify)
    - [Checksums](#checksums)
    - [GPG Signatures](#gpg-signatures)
      - [Add our public GPG Key via OpenPGP Public Key Server](#add-our-public-gpg-key-via-openpgp-public-key-server)
        - [Add our public GPG Key via MIT PGP Public Key Server](#add-our-public-gpg-key-via-mit-pgp-public-key-server)
        - [Verify Artifact via Signature File](#verify-artifact-via-signature-file)
  - [Getting Started](#getting-started)
  - [Configure Ping CLI](#configure-ping-cli)
    - [Documentation Generation](#documentation-generation)
  - [Commands](#commands)
    - [Platform Commands](#platform-commands)
    - [Product Commands](#product-commands)
    - [Export Configuration](#export-configuration)
    - [Custom API Requests](#custom-api-requests)
  - [Telemetry](#telemetry)
    - [Enable Telemetry](#enable-telemetry)
    - [TLS Configuration](#tls-configuration)
  - [Logging](#logging)
    - [Log Levels](#log-levels)
    - [Examples](#examples)
  - [Experimental Features](#experimental-features)
    - [Enable via Environment Variable](#enable-via-environment-variable)
    - [Enable via Configuration](#enable-via-configuration)
  - [AI Agent Skills](#ai-agent-skills)
  - [Getting Help](#getting-help)

## Install

### Docker

Use the [Ping CLI Docker image](https://hub.docker.com/r/pingidentity/pingcli)

Pull Image:

```shell
docker pull pingidentity/pingcli:latest
```

Example Commands:

```shell
docker run --rm pingidentity/pingcli:latest <sub commands>

docker run --rm pingidentity/pingcli:latest --version
```

### macOS

#### Homebrew

Use PingIdentity's Homebrew tap to install Ping CLI

``` shell
brew tap pingidentity/tap
brew install pingcli
```

#### Manual Installation

See [the latest GitHub release](https://github.com/pingidentity/pingcli/releases/latest) for artifact downloads, artifact signatures, and the checksum file. To verify package downloads, see the [Verify Section](#verify).

OR

Use the following single-line command to install Ping CLI into '/usr/local/bin' directly.

```shell
RELEASE_VERSION=$(basename $(curl -Ls -o /dev/null -w %{url_effective} https://github.com/pingidentity/pingcli/releases/latest)); \
OS_NAME=$(uname -s); \
HARDWARE_PLATFORM=$(uname -m | sed s/aarch64/arm64/ | sed s/x86_64/amd64/); \
URL="https://github.com/pingidentity/pingcli/releases/download/${RELEASE_VERSION}/pingcli_${RELEASE_VERSION#v}_${OS_NAME}_${HARDWARE_PLATFORM}"; \
curl -Ls -o pingcli "${URL}"; \
chmod +x pingcli; \
sudo mv pingcli /usr/local/bin/pingcli;
```

### Linux

##### Homebrew

Use PingIdentity's Homebrew tap to install Ping CLI

``` shell
brew tap pingidentity/tap
brew install pingcli
```

##### Alpine (.apk)

See [the latest GitHub release](https://github.com/pingidentity/pingcli/releases/latest) for Alpine (.apk) package downloads. To verify package downloads, see the [Verify Section](#verify).

> **_NOTE:_** The following commands may require `sudo` if not run as the root user.

```shell
apk add --allow-untrusted ./pingcli_<version>_linux_amd64.apk
apk add --allow-untrusted ./pingcli_<version>_linux_arm64.apk
```

##### Debian/Ubuntu (.deb)

See [the latest GitHub release](https://github.com/pingidentity/pingcli/releases/latest) for Debian (.deb) package downloads. To verify package downloads, see the [Verify Section](#verify).

> **_NOTE:_** The following commands may require `sudo` if not run as the root user.

```shell
apt install ./pingcli_<version>_linux_amd64.deb
apt install ./pingcli_<version>_linux_arm64.deb
```

##### CentOS/Fedora/RHEL (.rpm)

See [the latest GitHub release](https://github.com/pingidentity/pingcli/releases/latest) for RPM (.rpm) package downloads. To verify package downloads, see the [Verify Section](#verify).

> **_NOTE:_** The following commands may require `sudo` if not run as the root user.

```shell
yum install ./pingcli_<version>_linux_amd64.rpm
yum install ./pingcli_<version>_linux_arm64.rpm
```

OR

> **_NOTE:_** The following commands may require `sudo` if not run as the root user.

```shell
dnf install ./pingcli_<version>_linux_amd64.rpm
dnf install ./pingcli_<version>_linux_arm64.rpm
```

> **_NOTE:_**

> - Use `yum` for CentOS/RHEL 7 and earlier, and for older Fedora systems.
> - Use `dnf` for Fedora 22+ and CentOS/RHEL 8+.
> Both commands achieve the same result; use the one appropriate for your distribution.

#### Manual Installation

See [the latest GitHub release](https://github.com/pingidentity/pingcli/releases/latest) for artifact downloads, artifact signatures, and the checksum file. To verify package downloads, see the [Verify Section](#verify).

OR

Use the following single-line command to install Ping CLI into '/usr/local/bin' directly.

```shell
RELEASE_VERSION=$(basename $(curl -Ls -o /dev/null -w %{url_effective} https://github.com/pingidentity/pingcli/releases/latest)); \
OS_NAME=$(uname -s); \
HARDWARE_PLATFORM=$(uname -m | sed s/aarch64/arm64/ | sed s/x86_64/amd64/); \
URL="https://github.com/pingidentity/pingcli/releases/download/${RELEASE_VERSION}/pingcli_${RELEASE_VERSION#v}_${OS_NAME}_${HARDWARE_PLATFORM}"; \
curl -Ls -o pingcli "${URL}"; \
chmod +x pingcli; \
sudo mv pingcli /usr/local/bin/pingcli;
```

### Windows

#### Manual Installation

See [the latest GitHub release](https://github.com/pingidentity/pingcli/releases/latest) for artifact downloads, artifact signatures, and the checksum file. To verify package downloads, see the [Verify Section](#verify).

OR

Use the following single-line PowerShell 7.4 command to install Ping CLI into '%LOCALAPPDATA%\Programs' directly.
>**_NOTE:_** After installation, ensure that `%LOCALAPPDATA%\Programs` is included in your PATH environment variable. If it is not already present, add it so you can call `pingcli` directly in your terminal.

```powershell
$latestReleaseUrl = Invoke-WebRequest -Uri "https://github.com/pingidentity/pingcli/releases/latest" -MaximumRedirection 0 -ErrorAction Ignore -UseBasicParsing -SkipHttpErrorCheck; `
$RELEASE_VERSION = [System.IO.Path]::GetFileName($latestReleaseUrl.Headers.Location); `
$RELEASE_VERSION_NO_PREFIX = $RELEASE_VERSION -replace "^v", ""; `
$HARDWARE_PLATFORM = $env:PROCESSOR_ARCHITECTURE -replace "ARM64", "arm64" -replace "x86", "386" -replace "AMD64", "amd64" -replace "EM64T", "amd64"; `
$URL = "https://github.com/pingidentity/pingcli/releases/download/${RELEASE_VERSION}/pingcli_${RELEASE_VERSION_NO_PREFIX}_windows_${HARDWARE_PLATFORM}.exe"
Invoke-WebRequest -Uri $URL -OutFile "pingcli.exe"; `
Move-Item -Path pingcli.exe -Destination "${env:LOCALAPPDATA}\Programs"
```

## Verify

### Checksums

See [the latest GitHub release](https://github.com/pingidentity/pingcli/releases/latest) for the checksums.txt file. The checksums are in the format of SHA256.

### GPG Signatures

See [the latest GitHub release](https://github.com/pingidentity/pingcli/releases/latest) for the artifact downloads and signature files.

#### Add our public GPG Key via OpenPGP Public Key Server

```shell
gpg --keyserver keys.openpgp.org --recv-key 0x6703FFB15B36A7AC
```

OR

##### Add our public GPG Key via MIT PGP Public Key Server

```shell
gpg --keyserver pgp.mit.edu --recv-key 0x6703FFB15B36A7AC
```

##### Verify Artifact via Signature File

```shell
gpg --verify <artifact-name>.sig <artifact-name>
```

## Getting Started

1. **Install** Ping CLI for your platform — see [Install](#install).

2. **Verify** your download (optional but recommended) — see [Verify](#verify).

3. **Run the setup wizard** to configure a profile and connect your Ping Identity services:

   ```shell
   pingcli init
   ```

   This walks you through creating a configuration profile, connecting one or more Ping Identity services, and setting feature preferences. When complete, you are ready to use product commands such as `pingcli pingone export` or `pingcli pingfederate api`.

   Individual product `init` commands are also available if you want to connect a single service independently:

   ```shell
   pingcli pingone init
   pingcli pingfederate init
   ```

## Configure Ping CLI

Before using the Ping CLI, you need to configure your Ping Identity Service profile(s). The following steps show the quickest path to configuration.

Start by running the command to create a new profile and answering the prompts.

```shell
pingcli config profiles create
```

### Configuration Model Migration

Ping CLI stores a configuration model version in `configModelVersion` and can
automatically migrate legacy config files as models evolve.

- New config files are stamped with `configModelVersion: 2`.
- Legacy files without `configModelVersion` are treated as model v1.
- Before migration, Ping CLI creates a backup file in the same directory.

Startup migration behavior:

- Interactive terminals: prompts for approval by default.
- Non-interactive terminals: does not mutate config unless explicitly approved.

Startup policy override (values: `ask`, `yes`, `no`):

```shell
export PINGCLI_CONFIG_MIGRATION_POLICY=yes
```

Or create a profile non-interactively and set it as active in one step:

```shell
pingcli config profiles create --name dev --description "My development environment" --set-active
```

The newly created profile can now be configured via the `pingcli config set` command. General Ping Identity service connection settings are found under the `service` key, and settings relevant to individual commands are found under their command names e.g. `export`.

See [Configuration Key Documentation](./docs/tool-configuration/configuration-key.md) for more information on configuration keys
and their purposes.

See [Autocompletion Documentation](./docs/autocompletion/autocompletion.md) for information on loading autocompletion for select command flags.

### Documentation Generation

Documentation generation instructions (configuration options reference, per-command pages, navigation, rebuild workflow, and golden test usage) have moved to a dedicated guide:

See: `tools/README_DocumentGeneration.md`

## Commands

Ping CLI commands have the following structure:

```shell
pingcli <command> <subcommand> [options and parameters]
```

To get the version of Ping CLI:

```shell
pingcli --version
```

### Platform Commands

These commands operate across all connected products and services configured in a profile.

| Command | Description |
|---|---|
| `auth` | Authenticate (or refresh authentication) for all connected products and services at once. |
| `export` | Export configuration for all connected products and services as a bundle to the local filesystem. |
| `init` | Initialize Ping CLI with a guided setup wizard. |

### Product Commands

Each product has its own top-level command grouping the capabilities Ping CLI supports for that service.

| Command | Description |
|---|---|
| `pingone` | Administration tools for the PingOne platform. |
| `davinci` | Administration tools for the PingOne DaVinci universal service. |
| `pingfederate` | Administration tools for PingFederate deployed as software. |

Each product command exposes a consistent set of subcommands where supported:

| Subcommand | Description |
|---|---|
| `init` | Interactive setup wizard for the product. |
| `auth` | Authenticate Ping CLI to the product's management APIs. |
| `export` | Export product configuration as Terraform import blocks to the local filesystem. |
| `api` | Send a custom REST API request to the product's management API. |

### Export Configuration

Ping CLI generates [Terraform import blocks](https://developer.hashicorp.com/terraform/language/import) for every supported resource. Export is available per product or as a multi-service bundle.

Export all connected products at once:

```shell
pingcli export
```

Export a specific product:

```shell
pingcli pingone export
pingcli pingfederate export
pingcli davinci export
```

The generated import blocks are organized into one folder with a file per resource type, and can be used to [generate Terraform configuration](https://developer.hashicorp.com/terraform/language/import/generating-configuration).

### Custom API Requests

Each product exposes an `api` subcommand for sending arbitrary REST requests — a cURL-like experience where authentication and connection details are filled automatically from the active profile.

Send a GET request to list PingOne environments:

```shell
pingcli pingone api environments
```

Send a POST request with a JSON body:

```shell
pingcli pingone api --http-method POST --data ./my-environment.json environments
```

The same pattern applies to other products:

```shell
pingcli pingfederate api --http-method GET pf-admin-api/v1/serverSettings
```

### Resource Management

Supported resources expose `create`, `replace` (alias `update`), and `delete` subcommands. Body-bearing operations (`create`, `replace`) require a `--from-file` (`-f`) flag pointing to a JSON file whose shape matches the SDK request type directly.

Create a PingOne environment from a JSON file:

```shell
pingcli pingone environments create --from-file env.json
```

Read the JSON body from stdin (pipe from `jq` or redirect a file):

```shell
pingcli pingone environments create --from-file - < env.json
cat env.json | jq '.name = "new-env"' | pingcli pingone environments create --from-file -
```

The same `--from-file` flag works on `replace` alongside required path-parameter flags:

```shell
pingcli pingone environments replace --environment-id <env-uuid> --from-file env.json
pingcli pingone populations create --environment-id <env-uuid> --from-file population.json
pingcli pingone groups create --environment-id <env-uuid> --from-file group.json
pingcli davinci variables create --environment-id <env-uuid> --from-file variable.json
pingcli davinci variables replace --environment-id <env-uuid> --variable-id <var-uuid> --from-file - < variable.json
```

Path-parameter flags such as `--environment-id` are separate from the body and must be provided as CLI flags.

## Telemetry

Ping CLI includes optional [OpenTelemetry](https://opentelemetry.io/) tracing support. Telemetry is **disabled by default** and no data is collected or exported unless you explicitly enable it and configure an OTLP backend.

When enabled, Ping CLI emits trace spans for command executions and API calls. All trace data is sent exclusively to the OTLP endpoint you configure — no data is sent to Ping Identity.

### Enable Telemetry

Set `telemetry.enabled` in your profile configuration:

```shell
pingcli config set telemetry.enabled true
pingcli config set telemetry.otlp.endpoint http://localhost:4318
```

Or use environment variables:

```shell
export PINGCLI_TELEMETRY_ENABLED=true
export OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4318
```

The OTLP protocol defaults to `http`. To use gRPC, set:

```shell
pingcli config set telemetry.otlp.protocol grpc
# or
export OTEL_EXPORTER_OTLP_PROTOCOL=grpc
```

### TLS Configuration

For OTLP backends that require TLS, use the standard OpenTelemetry environment variables:

| Variable | Description |
|---|---|
| `OTEL_EXPORTER_OTLP_TLS_ENABLED` | Enable TLS for the OTLP connection |
| `OTEL_EXPORTER_OTLP_CERTIFICATE` | Path to the client certificate file (mTLS) |
| `OTEL_EXPORTER_OTLP_CLIENT_KEY` | Path to the client private key file (mTLS) |
| `OTEL_EXPORTER_OTLP_CA_CERTIFICATE` | Path to the CA certificate file |
| `OTEL_EXPORTER_OTLP_INSECURE` | Skip server certificate verification (not recommended for production) |

These can also be set via their corresponding `telemetry.tls.*` configuration keys.

## Logging

Ping CLI logging is **disabled by default**. It is configured exclusively via environment variables and takes effect for the lifetime of the process.

| Variable | Description |
|---|---|
| `PINGCLI_LOG_LEVEL` | Sets the minimum log level. When unset, logging is disabled. |
| `PINGCLI_LOG_PATH` | Path to a log file. When unset, logs are written to stdout. |

### Log Levels

Accepted values for `PINGCLI_LOG_LEVEL`, from most to least verbose:

| Level | Description |
|---|---|
| `TRACE` | Finest-grained tracing |
| `DEBUG` | Detailed diagnostic output |
| `INFO` | General informational messages |
| `WARN` | Warnings that do not prevent execution |
| `ERROR` | Errors encountered during execution |
| `FATAL` | Fatal errors that terminate the process |
| `PANIC` | Panics |
| `NOLEVEL` | Log all messages regardless of level |

### Examples

Enable debug logging to a file:

```shell
PINGCLI_LOG_LEVEL=DEBUG PINGCLI_LOG_PATH=./pingcli.log pingcli pingone api environments
```

Enable debug logging to stdout:

```shell
PINGCLI_LOG_LEVEL=DEBUG pingcli config profiles list
```

## Experimental Features

Ping CLI includes features that are still in active development and subject to change. These are **disabled by default** and must be explicitly opted in to.

When enabled, experimental commands and options become available in the CLI. Because experimental status is resolved before the command tree is built, the flag must be set via the environment variable or the config file — not a command-line flag.

### Enable via Environment Variable

```shell
PINGCLI_EXPERIMENTAL=true pingcli <command>
```

### Enable via Configuration

```shell
pingcli config set experimental.enabled true
```

> **Note:** Experimental features may change or be removed without notice. Do not rely on them in production workflows.

## AI Agent Skills

Ping CLI ships agent skills in `.agents/skills/` — structured reference files that AI coding assistants (GitHub Copilot, Claude, etc.) automatically load when relevant tasks are detected. No installation is needed; the skills are picked up from the repository automatically by compatible IDEs.

| Skill | Path | Description |
|---|---|---|
| `pingcli-usage` | `.agents/skills/pingcli-usage/` | Complete command reference, flags, and usage examples for every Ping CLI command. Auto-generated from the command tree — refresh with `make generate-skill-docs`. |
| `pingcli-plugin-builder` | `.agents/skills/pingcli-plugin-builder/` | Step-by-step guide for building Ping CLI plugins using gRPC and the HashiCorp go-plugin library. |
| `pingcli-developer` | `.agents/skills/pingcli-developer/` | Architectural patterns for contributing to Ping CLI — connectors, commands, configuration options, and authentication flows. |

## Getting Help

The best way to interact with our team is through Github. You can [open an issue](https://github.com/pingidentity/pingcli/issues/new) for guidance, bug reports, or feature requests.

Please check for similar open issues before opening a new one.
