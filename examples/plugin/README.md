# Ping CLI Plugin Development Guide

Welcome to the developer guide for creating `pingcli` plugins! This document provides all the information you need to build, test, and distribute your own custom commands to extend the functionality of the `pingcli` tool.

## Table of Contents

- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [How Plugins Work](#how-plugins-work)
- [Authentication](#authentication)
- [Building a Plugin](#building-a-plugin)
- [Registering and Managing Plugins](#registering-and-managing-plugins)
  - [Adding a Plugin](#adding-a-plugin)
  - [Listing Plugins](#listing-plugins)
  - [Removing a Plugin](#removing-a-plugin)
- [Plugin Command Interface](#plugin-command-interface)
  - [Configuration](#configuration-pingclicommandconfiguration-error)
  - [Run](#runargs-string-logger-grpclogger-error)
- [Logging from Plugins](#logging-from-plugins)
- [Troubleshooting](#troubleshooting)
- [Further Reading](#further-reading)

## Introduction

The `pingcli` plugin system allows developers to create new commands that integrate seamlessly into the main application. Each plugin is a standalone executable that communicates with the `pingcli` host process over gRPC. This architecture ensures that plugins are isolated and secure. Currently, the plugin framework only supports plugins written in Go.

## Prerequisites

- **Go 1.24+** (for building Go plugins)
- [HashiCorp go-plugin](https://github.com/hashicorp/go-plugin) (used by both host and plugin)
- **Ping CLI v0.7.0+** installed and configured

## Quick Start
You can quickly build and test the example plugin provided in this directory.

1.  **Build the plugin**:
    Run this command from the root of the repository:
    ```bash
    go build -o "$HOME/go/bin/pingcli-example-plugin" ./examples/plugin
    ```
    *Note: This assumes `$HOME/go/bin` is in your system PATH, which is standard for Go development.*

2.  **Register the plugin**:
    ```bash
    pingcli plugin add pingcli-example-plugin
    ```

3.  **Run the command**:
    ```bash
    pingcli pingcli-example-plugin --hello world
    ```

## How Plugins Work

1.  **Discovery**: When `pingcli` starts, it loads a list of registered plugin executables from its configuration profile. This list is managed by the `pingcli plugin` command. Plugins must be on your system `PATH`
2.  **Handshake**: For each registered plugin, `pingcli` launches the executable as a child process. A secure handshake is performed to verify that the child process is a valid plugin and is compatible with the host.
3.  **Communication**: Once the handshake is complete, the host and plugin communicate over gRPC. The host can call functions defined in the plugin (like `Run`), and the plugin can send log messages back to the host.
4.  **Execution**: When a user runs a command provided by a plugin, `pingcli` invokes the corresponding gRPC method in the plugin process, passing along any arguments and flags.
5.  **Compatibility**: The Handshake process includes a ProtocolVersion check. This ensures that the plugin and the pingcli host are compatible, preventing issues if the underlying gRPC interface changes in future versions of pingcli.

## Authentication

Plugins can leverage the host `pingcli`'s authenticated session to make API calls to supported services. This avoids the need for plugins to manage credentials or perform OAuth flows themselves.

For detailed documentation and usage examples, see [AUTHENTICATION.md](AUTHENTICATION.md).

## Building a Plugin

1. **Clone or create your plugin source code.**  
   See [`plugin.go`](plugin.go) for a complete example.

2. **Build the plugin binary:**
   ```sh
   go build -o my-plugin
   ```

3. **Place the binary in a directory on your `PATH`:**
   ```sh
   mv my-plugin ~/go/bin/
   # or any directory in your $PATH
   ```

## Registering and Managing Plugins

`pingcli` provides the `plugin` command to manage the lifecycle of your plugins.

### Adding a Plugin

To add a new plugin, use the add subcommand. Crucially, the plugin executable must first be placed in a directory that is part of your system's PATH environment variable. pingcli relies on the system's PATH to find the executable to run.

```bash
pingcli plugin add <executable-name>
```

### Listing Plugins

To see a list of all currently registered plugins, use the `list` subcommand.

```bash
pingcli plugin list
```

### Removing a Plugin

To unregister a plugin from `pingcli`, use the `remove` subcommand.

```bash
pingcli plugin remove <executable-name>
```

## Plugin Command Interface

To create a valid plugin, you must implement the `grpc.PingCliCommand` interface. This interface has two methods:

#### `Configuration() (*grpc.PingCliCommandConfiguration, error)`

This method is called by the `pingcli` host to get metadata about your command. This allows `pingcli` to display your command in the help text (`pingcli --help`).

The `PingCliCommandConfiguration` struct has the following fields, which correspond directly to properties of a [Cobra](https://github.com/spf13/cobra) command:

-   `Use`: The one-line usage message for the command (e.g., `my-command [flags]`).
-   `Short`: A short description of the command.
-   `Long`: A longer, more detailed description of the command.
-   `Example`: One or more examples of how to use the command.

By providing this metadata, pingcli can present your plugin in a manner that is consistent and feels native to the main application.

#### `Run(args []string, logger grpc.Logger) error`

This is the main entry point for your command's logic. It is executed when a user runs your command.

-   `args []string`: A slice of strings containing all the command-line arguments and flags that were passed to your command. For example, if a user runs `pingcli my-command first-arg --verbose`, the `args` slice will be `["first-arg", "--verbose"]`.
-   `logger grpc.Logger`: A gRPC client that allows your plugin to send log messages back to the `pingcli` host. **This is the only way your plugin should produce output.**
-   `auth grpc.Authentication`: A gRPC client that allows your plugin to request an authentication token from the host.

## Logging from Plugins

Plugins must not write directly to `stdout` or `stderr`. Instead, they must use the provided `logger` object in the `Run` method. This ensures that all output is managed by the host and presented to the user in a consistent format.

The `logger` interface provides several methods for different log levels:

-   `logger.Message(message string, fields map[string]string)`
-   `logger.Warn(message string, fields map[string]string)`
-   `logger.PluginError(message string, fields map[string]string)`
-   `logger.Success(message string, fields map[string]string)`
-   `logger.UserError(message string, fields map[string]string)`
-   `logger.UserFatal(message string, fields map[string]string)`

## Troubleshooting

- **Plugin not found:**  
  Ensure the binary is on your `PATH` and registered with `pingcli plugin add`.

- **Handshake failed:**  
  Check that both host and plugin use compatible protocol versions.
  
  *Note for macOS users*: In some environments, simply running the plugin process might require setting `export GOMAXPROCS=1` before running `pingcli` if you encounter stability issues during handshake, though this is rare.

- **gRPC errors:**  
  Ensure your plugin implements the correct interface and uses the expected gRPC protocol.

- **No output:**  
  All output is expected to go through the provided `logger`.

## Further Reading

- [HashiCorp go-plugin documentation](https://github.com/hashicorp/go-plugin)
- [`pingcli` main documentation](../../README.md)
- [Cobra CLI framework](https://github.com/spf13/cobra)
