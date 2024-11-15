# Ping CLI

The Ping CLI is a unified command line interface for configuring and managing Ping Identity Services.

## Install

#### macOS/Linux - Homebrew

Use PingIdentity's Homebrew tap to install Ping CLI

```text
brew install pingidentity/tap/pingcli

or

brew tap pingidentity/tap
brew install pingcli
```

#### Manual Installation - macOS/Linux

See [the latest GitHub release](https://github.com/pingidentity/pingcli/releases/latest) for binary downloads and SHA256 checksum files.

OR

Use the following single-line command to install Ping CLI directly.

```text
RELEASE_VERSION=$(basename $(curl -Ls -o /dev/null -w %{url_effective} https://github.com/pingidentity/pingcli/releases/latest)); \
OS_NAME=$(uname -s); \
HARDWARE_PLATFORM=$(uname -m | sed s/aarch64/arm64/ | sed s/x86_64/amd64/); \
TAR_FILENAME="pingcli.tar.gz"; \
curl --location --silent --output "${TAR_FILENAME}" --request GET \
"https://github.com/pingidentity/pingcli/releases/download/${RELEASE_VERSION}/pingcli_${RELEASE_VERSION#v}_${OS_NAME}_${HARDWARE_PLATFORM}.tar.gz"; \
tar -zxf "${TAR_FILENAME}" pingcli; \
rm -f "${TAR_FILENAME}"
```

#### Manual Installation - Windows

See [the latest GitHub release](https://github.com/pingidentity/pingcli/releases/latest) for binary downloads and SHA256 checksum files.

OR

Use the following single-line powershell command to install Ping CLI directly.
```TEXT
$latestReleaseUrl = Invoke-WebRequest -Uri "https://github.com/pingidentity/pingcli/releases/latest" -MaximumRedirection 0 -ErrorAction Ignore -UseBasicParsing; `
$RELEASE_VERSION = [System.IO.Path]::GetFileName($latestReleaseUrl.Headers.Location); `
$RELEASE_VERSION_NO_PREFIX = $RELEASE_VERSION -replace "^v", ""; `
Invoke-WebRequest -Uri "https://github.com/pingidentity/pingcli/releases/download/${RELEASE_VERSION}/pingcli_${RELEASE_VERSION_NO_PREFIX}_windows_amd64.tar.gz" -OutFile pingcli.tar.gz; `
tar -zxf pingcli.tar.gz pingcli.exe; `
Remove-Item pingcli.tar.gz
```

## Configure Ping CLI

Before using the Ping CLI, you need to configure your Ping Identity Service profile(s). The following steps show the quickest path to configuration.

Start by running the command to create a new profile and answering the prompts.

```text
$ pingcli config add-profile
Pingcli configuration file '/Users/<me>/.pingcli/config.yaml' does not exist. - No Action (Warning)
Creating new Ping CLI configuration file at: /Users/<me>/.pingcli/config.yaml
New profile name: : dev
New profile description: : configuration for development environment
Set new profile as active: : y
Adding new profile 'dev'...
Profile created. Update additional profile attributes via 'pingcli config set' or directly within the config file at '/Users/<me>/.pingcli/config.yaml' - Success
Profile 'dev' set as active. - Success
```

The newly create profile can now be configured via the `pingcli config set` command. General Ping Identity service connection settings are found under the `service` key, and settings relevant to individual commands are found under their command names e.g. `export` and `request`.

See [Configuration Key Documentation](./docs/tool-configuration/configuration-key.md) for more information on configuration keys
and their purposes.

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
