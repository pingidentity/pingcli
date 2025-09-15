# Ping CLI

[![Code Analysis and Tests](https://github.com/pingidentity/pingcli/actions/workflows/code-analysis-lint-test.yaml/badge.svg)](https://github.com/pingidentity/pingcli/actions/workflows/code-analysis-lint-test.yaml)
[![CodeQL](https://github.com/pingidentity/pingcli/actions/workflows/codeql.yaml/badge.svg)](https://github.com/pingidentity/pingcli/actions/workflows/codeql.yaml)
[![Docker Pulls](https://img.shields.io/docker/pulls/pingidentity/pingcli.svg)](https://hub.docker.com/r/pingidentity/pingcli)
[![GitHub release](https://img.shields.io/github/v/release/pingidentity/pingcli?include_prereleases&sort=semver)](https://github.com/pingidentity/pingcli/releases)

The Ping CLI is a unified command line interface for configuring and managing Ping Identity Services.

## Table of Contents

- [Install](#install)
	- [Docker](#docker)
	- [macOS](#macos)
	- [Linux](#linux)
	- [Windows](#windows)
- [Verify](#verify)
	- [Checksums](#checksums)
	- [GPG Signatures](#gpg-signatures)
- [Configure Ping CLI](#configure-ping-cli)
- [Commands](#commands)
	- [Platform Export](#platform-export)
	- [Custom Request](#custom-request)
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

##### Homebrew

Use PingIdentity's Homebrew tap to install Ping CLI

``` shell
brew tap pingidentity/tap
brew install pingcli
```

##### Manual Installation

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

##### Manual Installation

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

##### Manual Installation

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

##### Add our public GPG Key via OpenPGP Public Key Server

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

## Configure Ping CLI

Before using the Ping CLI, you need to configure your Ping Identity Service profile(s). The following steps show the quickest path to configuration.

Start by running the command to create a new profile and answering the prompts.

```shell
$ pingcli config add-profile
Ping CLI configuration file '$HOME/.pingcli/config.yaml' does not exist. - No Action (Warning)
Creating new Ping CLI configuration file at: $HOME/.pingcli/config.yaml
New profile name: : dev
New profile description: : configuration for development environment
Set new profile as active: : y
Adding new profile 'dev'...
Profile created. Update additional profile attributes via 'pingcli config set' or directly within the config file at '$HOME/.pingcli/config.yaml' - Success
Profile 'dev' set as active. - Success
```

The newly created profile can now be configured via the `pingcli config set` command. General Ping Identity service connection settings are found under the `service` key, and settings relevant to individual commands are found under their command names e.g. `export` and `request`.

See [Configuration Key Documentation](./docs/tool-configuration/configuration-key.md) for more information on configuration keys
and their purposes.

See [Autocompletion Documentation](./docs/autocompletion/autocompletion.md) for information on loading autocompletion for select command flags.

### Generating Configuration Options Documentation

To regenerate the markdown table of configuration options used in documentation, use the provided Makefile target instead of the previous ad-hoc test:

```shell
make generate-options-docs
```

This prints the markdown to stdout. To write it directly to a file (for example updating docs/options.md):

```shell
make generate-options-docs OUTPUT='-o docs/options.md'
```

The generator code lives in `cmd/generate-options-docs` and can be invoked manually as well:

```shell
go run ./cmd/generate-options-docs -o /path/to/file.md
```

The logic was migrated from the skipped test `internal/configuration/options/options_test.go` (Test_outputOptionsMDInfo) to make doc generation an explicit developer action.

## Commands

Ping CLI commands have the following structure:

```shell
pingcli <command> <subcommand> [options and parameters]
```

To get the version of Ping CLI:

```shell
pingcli --version
```

### Platform Export

The `pingcli platform export` command uses your configured settings to connect to the requested services and generate [Terraform import blocks](https://developer.hashicorp.com/terraform/language/import) for every supported resource available.

An example command to export a PingOne environment for HCL generation looks like:

```shell
pingcli platform export --services "pingone-platform,pingone-sso"
```

The generated import blocks are organized into one folder with a file per resource type found. These import blocks can be used to [generate terraform configuration](https://developer.hashicorp.com/terraform/language/import/generating-configuration).

### Custom Request

The `pingcli request` command uses your configured settings to authenticate to the desired ping service before executing your API request.

An example command to view PingOne Environments looks like:

```shell
pingcli request --http-method GET --service pingone environments
```

## Getting Help

The best way to interact with our team is through Github. You can [open an issue](https://github.com/pingidentity/pingcli/issues/new) for guidance, bug reports, or feature requests.

Please check for similar open issues before opening a new one.
