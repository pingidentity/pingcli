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

	// We use our own logger for the plugins to communicate to the user.
	// Discard any other plugin logging details to avoid user communication clutter.
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: io.Discard,
		Level:  hclog.Debug,
	})

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
		// PingCLI is the host process. Start the plugin process
		client := hplugin.NewClient(&hplugin.ClientConfig{
			HandshakeConfig: grpc.HandshakeConfig,
			Plugins:         grpc.PluginMap,
			Cmd:             exec.Command(pluginExecutable),
			AllowedProtocols: []hplugin.Protocol{
				hplugin.ProtocolGRPC,
			},
			Logger: logger,
		})

		// Connect via RPC
		rpcClient, err := client.Client()
		if err != nil {
			return fmt.Errorf("failed to create Plugin RPC client: %w", err)
		}

		// All PingCLI plugins are expected to serve the ENUM_PINGCLI_COMMAND_GRPC plugin via
		// the PluginMap within the plugin.Serve() method.
		// Non-Golang plugins unable to use the shared grpc module should supply the
		// raw value of ENUM_PINGCLI_COMMAND_GRPC "pingcli_command_grpc" for the PluginMap key.
		raw, err := rpcClient.Dispense(grpc.ENUM_PINGCLI_COMMAND_GRPC)
		if err != nil {
			return fmt.Errorf("the rpc client failed to dispense plugin executable '%s': %w", pluginExecutable, err)
		}

		// Cast the dispensed plugin to the interface we expect to work with: grpc.PingCliCommand.
		// However, this is not a normal interface, but rather implemeted over the RPC connection.
		plugin, ok := raw.(grpc.PingCliCommand)
		if !ok {
			return fmt.Errorf("failed to cast plugin executable '%s' to grpc.PingCliCommand interface", pluginExecutable)
		}

		// The configuration method is defined by the protobuf definition.
		// The plugin should return relevant cobra command information
		// even if the plugin does not use cobra commands internally.
		// This allows the host process PingCLI to present the plugin command
		// in the help output.
		resp, err := plugin.Configuration()
		if err != nil {
			return fmt.Errorf("failed to run command from Plugin: %w", err)
		}

		pluginCmd := &cobra.Command{
			Use:                   resp.Use,
			Short:                 resp.Short,
			Long:                  resp.Long,
			Example:               resp.Example,
			DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
			RunE: func(cmd *cobra.Command, args []string) error {
				// TODO: Right now, this means we only cleanup the plugin after the command is run
				// TODO: This leaves all other plugins added uncleaned up
				defer client.Kill()

				err := plugin.Run(args, &shared_logger.SharedLogger{})
				if err != nil {
					return fmt.Errorf("failed to execute plugin command: %w", err)
				}

				return nil
			},
		}

		cmd.AddCommand(pluginCmd)

		l.Info().Msgf("Loaded plugin executable: %s", pluginExecutable)
	}

	return nil
}
