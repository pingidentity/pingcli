// Copyright © 2025 Ping Identity Corporation

package plugins

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/hashicorp/go-hclog"
	hplugin "github.com/hashicorp/go-plugin"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/shared/grpc"
	shared_logger "github.com/pingidentity/pingcli/shared/logger"
	"github.com/spf13/cobra"
)

func AddAllPluginToCmd(cmd *cobra.Command) error {
	l := logger.Get()

	// Plugin executables are stored in the profile configuration
	// via the command 'pingcli plugin add <plugin-executable>'
	pluginExecutables, err := profiles.GetOptionValue(options.PluginExecutablesOption)
	if err != nil {
		return fmt.Errorf("failed to get configured plugin executables: %w", err)
	}

	if pluginExecutables == "" {
		return nil
	}

	for pluginExecutable := range strings.SplitSeq(pluginExecutables, ",") {
		pluginExecutable = strings.TrimSpace(pluginExecutable)
		if pluginExecutable == "" {
			continue
		}

		conf, err := pluginConfiguration(pluginExecutable)
		if err != nil {
			return err
		}

		pluginCmd := &cobra.Command{
			Use:                   conf.Use,
			Short:                 conf.Short,
			Long:                  conf.Long,
			Example:               conf.Example,
			DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
			RunE:                  createCmdRunE(pluginExecutable),
			DisableFlagParsing:    true, // The plugin command will handle its own flags
		}

		cmd.AddCommand(pluginCmd)

		l.Info().Msgf("Loaded plugin executable: %s", pluginExecutable)
	}

	return nil
}

// createHPluginClient creates a new hplugin.Client for the given plugin executable.
// The caller is responsible for closing the client connection after use.
func createHPluginClient(pluginExecutable string) *hplugin.Client {
	// We use our own logger for the plugins to communicate to the user.
	// Discard any other plugin logging details to avoid user communication clutter.
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: io.Discard,
		Level:  hclog.Debug,
	})

	// Ping CLI is the host process. Start the plugin process
	client := hplugin.NewClient(&hplugin.ClientConfig{
		HandshakeConfig: grpc.HandshakeConfig,
		Plugins:         grpc.PluginMap,
		Cmd:             exec.Command(pluginExecutable),
		AllowedProtocols: []hplugin.Protocol{
			hplugin.ProtocolGRPC,
		},
		Logger: logger,
	})

	return client
}

// dispensePlugin connects to the plugin via RPC and dispenses the grpc.PingCliCommand interface.
// the caller is responsible for closing the client connection after use.
func dispensePlugin(client *hplugin.Client, pluginExecutable string) (grpc.PingCliCommand, error) {
	// Connect via RPC
	clientProtocol, err := client.Client()
	if err != nil {
		return nil, fmt.Errorf("failed to create Plugin RPC client: %w", err)
	}

	// All Ping CLI plugins are expected to serve the ENUM_PINGCLI_COMMAND_GRPC plugin via
	// the PluginMap within the plugin.Serve() method.
	// Non-Golang plugins unable to use the shared grpc module should supply the
	// raw value of ENUM_PINGCLI_COMMAND_GRPC "pingcli_command_grpc" for the PluginMap key.
	raw, err := clientProtocol.Dispense(grpc.ENUM_PINGCLI_COMMAND_GRPC)
	if err != nil {
		return nil, fmt.Errorf("the rpc client failed to dispense plugin executable '%s': %w", pluginExecutable, err)
	}

	// Cast the dispensed plugin to the interface we expect to work with: grpc.PingCliCommand.
	// However, this is not a normal interface, but rather implemeted over the RPC connection.
	plugin, ok := raw.(grpc.PingCliCommand)
	if !ok {
		return nil, fmt.Errorf("failed to cast plugin executable '%s' to grpc.PingCliCommand interface", pluginExecutable)
	}

	return plugin, nil
}

func pluginConfiguration(pluginExecutable string) (conf *grpc.PingCliCommandConfiguration, err error) {
	client := createHPluginClient(pluginExecutable)
	defer client.Kill()

	plugin, err := dispensePlugin(client, pluginExecutable)
	if err != nil {
		return nil, err
	}

	// The configuration method is defined by the protobuf definition.
	// The plugin should return relevant cobra command information
	// even if the plugin does not use cobra commands internally.
	// This allows the host process Ping CLI to present the plugin command
	// in the help output.
	resp, err := plugin.Configuration()
	if err != nil {
		return nil, fmt.Errorf("failed to run command from Plugin: %w", err)
	}

	return resp, nil
}

func createCmdRunE(pluginExecutable string) func(cmd *cobra.Command, args []string) (err error) {
	return func(cmd *cobra.Command, args []string) error {
		client := createHPluginClient(pluginExecutable)
		defer client.Kill()

		plugin, err := dispensePlugin(client, pluginExecutable)
		if err != nil {
			return err
		}

		err = plugin.Run(args, &shared_logger.SharedLogger{})
		if err != nil {
			return fmt.Errorf("failed to execute plugin command: %w", err)
		}

		return nil
	}
}
