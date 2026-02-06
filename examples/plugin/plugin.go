// Copyright © 2025 Ping Identity Corporation

// Package 'plugin' provides an example implementation of a Ping CLI command plugin.
//
// It demonstrates the required structure and interfaces for building a new
// command that can be dynamically loaded and executed by the main pingcli
// application. This includes implementing the PingCliCommand interface and
// serving it over gRPC using Hashicorp's `go-plugin“ library.
package main

import (
	"fmt"

	"github.com/hashicorp/go-plugin"
	"github.com/pingidentity/pingcli/shared/grpc"
)

// These variables define the command's metadata, which is sent to the pingcli
// host process. This information is used by the host's command-line framework
// (Cobra) to display help text, usage, and examples, making the plugin feel
// like a native command.
var (
	// Example provides one or more usage examples for the command.
	Example = "pingcli pingcli-example-plugin --flag value"

	// Long provides a detailed description of the command. It's shown in the
	// help text when a user runs `pingcli help pingcli-example-plugin`.
	Long = `This command is an example of a plugin command that can be used with pingcli. 
	It demonstrates how to implement a custom command that can be executed by the pingcli host process`

	// Short provides a brief, one-line description of the command.
	Short = "An example plugin command for pingcli"

	// Use defines the command's name and its arguments/flags syntax.
	Use = "pingcli-example-plugin [flags]"
)

// PingCliCommand is the implementation of the grpc.PingCliCommand interface.
// It encapsulates the logic for the custom command provided by this plugin.
type PingCliCommand struct{}

// A compile-time check to ensure PingCliCommand correctly implements the
// grpc.PingCliCommand interface.
var _ grpc.PingCliCommand = (*PingCliCommand)(nil)

// Configuration is called by the pingcli host to retrieve the command's
// metadata, such as its name, description, and usage examples. This allows
// the host to integrate the plugin's command into its own help and usage
// displays without executing the plugin's main logic.
func (c *PingCliCommand) Configuration() (*grpc.PingCliCommandConfiguration, error) {
	cmdConfig := &grpc.PingCliCommandConfiguration{
		Example: Example,
		Long:    Long,
		Short:   Short,
		Use:     Use,
	}

	return cmdConfig, nil
}

// Run is the execution entry point for the plugin command. The pingcli host
// calls this method when a user invokes the plugin command.
//
// The `args` parameter contains all command-line arguments and flags passed
// after the command's name (as defined in the `Use` variable). For example,
// if a user runs `pingcli plugin-command add --flag value`, the `args` slice
// will be `["add", "--flag", "value"]`.
//
// The `logger` parameter is a gRPC client that allows the plugin to send log
// messages back to the host process, ensuring that all output is displayed
// consistently through the main pingcli interface.
//
// The `auth` parameter is a gRPC client that allows the plugin to request
// an authentication token from the host process.
func (c *PingCliCommand) Run(args []string, logger grpc.Logger, auth grpc.Authentication) error {
	// 1. Process arguments
	if len(args) > 0 {
		if err := logger.Message(fmt.Sprintf("Received arguments: %v", args), nil); err != nil {
			return err
		}
	} else {
		if err := logger.Message("No arguments provided. Usage: pingcli plugin-command [args]", nil); err != nil {
			return err
		}
	}

	// 2. Obtain Authentication Token
	// Plugins delegate authentication to the host. The host manages the active session
	// and provides a valid access token for the currently selected profile.
	token, err := auth.GetToken()
	if err != nil {
		// Log the error concisely as a plugin error for debugging
		if logErr := logger.PluginError(fmt.Sprintf("Authentication failed: %v", err), nil); logErr != nil {
			return logErr
		}
		// Return the error to the host to stop execution
		return fmt.Errorf("authentication required: %w", err)
	}

	if resultErr := logger.Success("Successfully authenticated!", nil); resultErr != nil {
		return resultErr
	}

	// 3. Command Logic (Example)
	// Use the token to make an API call. For demonstration, we just log it.
	// In a real plugin, you would use this token in the Authorization header:
	// req.Header.Set("Authorization", "Bearer " + token)
	if err := logger.Message(fmt.Sprintf("Token (truncated): %s...", token[:10]), nil); err != nil {
		return err
	}

	if err := logger.Warn("This is an example plugin command.", nil); err != nil {
		return err
	}

	return nil
}

// main is the entry point for the plugin's executable. When the pingcli host
// launches this plugin, this function starts a gRPC server that serves the
// PingCliCommand implementation. The go-plugin library handles the handshake
// and communication between the host and the plugin process.
func main() {
	plugin.Serve(&plugin.ServeConfig{
		// HandshakeConfig is a shared configuration used to verify that the host
		// and plugin are compatible.
		HandshakeConfig: grpc.HandshakeConfig,

		// Plugins defines the set of services this plugin serves. The key is a
		// unique name for the plugin service, and the value is an implementation
		// of the plugin.Plugin interface.
		Plugins: map[string]plugin.Plugin{
			grpc.ENUM_PINGCLI_COMMAND_GRPC: &grpc.PingCliCommandGrpcPlugin{
				Impl: &PingCliCommand{},
			},
		},

		// GRPCServer specifies the gRPC server implementation to use.
		// plugin.DefaultGRPCServer is a sane default provided by the library.
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
