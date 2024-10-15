# Ping CLI

The Ping CLI is a unified command line interface for configuring and managing Ping Identity Services.

## Install

<!-- TODO -->

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

Run `pingcli config view-profile` to generate a set of empty configuration.

```text
$ pingcli config view-profile
Profile: dev

color: true
description: configuration for development environment
export:
    format: HCL
    outputdirectory: /Users/samirgandhi/projects/terraform/UAT/pingcli-uat/export
    overwrite: false
    pingone:
        environmentid: ""
# additional configuration trimmed
```

The available settings can be configured by editing the empty fields within the generated `~/.pingcli/config.yaml`. General Ping Identity service connection settings are found under the `service` key, and settings relevant to individual commands are found under their command names e.g. `export` and `request`.

To see descriptions for the settings the `export` command's help option.
```
pingcli platform export --help
```

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
