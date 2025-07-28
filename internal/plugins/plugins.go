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
	"github.com/spf13/pflag"
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
			DisableFlagParsing:    true, // Let all flags pass through to plugin
			RunE:                  createCmdRunE(pluginExecutable),
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
		// Extract global flags before passing args to the plugin
		// This allows the host process to handle global flags and only pass plugin-specific args
		pluginArgs, err := filterRootFlags(args, cmd.Root().PersistentFlags())
		if err != nil {
			return fmt.Errorf("failed to execute plugin command: %w", err)
		}

		client := createHPluginClient(pluginExecutable)
		defer client.Kill()

		plugin, err := dispensePlugin(client, pluginExecutable)
		if err != nil {
			return err
		}

		err = plugin.Run(pluginArgs, &shared_logger.SharedLogger{})
		if err != nil {
			return fmt.Errorf("failed to execute plugin command: %w", err)
		}

		return nil
	}
}

// filterRootFlags filters out any flags that were parsed by the root command's persistent flags
// and processes them for the host application, returning only plugin-specific args.
func filterRootFlags(args []string, persistentFlags *pflag.FlagSet) ([]string, error) {
	pluginArgs := []string{}

	var (
		previousArgFlagName     = ""
		handlePreviousArgAsFlag = false
	)

	for _, arg := range args {
		switch {
		case handlePreviousArgAsFlag && previousArgFlagName != "":
			err := persistentFlags.Set(previousArgFlagName, arg)
			if err != nil {
				return nil, fmt.Errorf("failed to set persistent flag '%s' with value '%s': %w", previousArgFlagName, arg, err)
			}
			handlePreviousArgAsFlag = false
		case len(arg) > 0 && arg[0] == '-':
			// The argument is a flag, remove leading dashes
			flagArg := strings.TrimLeft(arg, "-")

			// Handle flags in the format --flag=value
			if strings.Contains(flagArg, "=") {
				parts := strings.SplitN(flagArg, "=", 2)
				flagName := parts[0]

				if flag := persistentFlags.Lookup(flagName); flag != nil {
					err := persistentFlags.Set(flagName, parts[1])
					if err != nil {
						return nil, fmt.Errorf("failed to set persistent flag '%s' with value '%s': %w", flagName, parts[1], err)
					}
				} else if len(flagName) == 1 && persistentFlags.ShorthandLookup(flagName) != nil {
					flag := persistentFlags.ShorthandLookup(flagName)
					err := persistentFlags.Set(flag.Name, parts[1])
					if err != nil {
						return nil, fmt.Errorf("failed to set persistent flag '%s' with value '%s': %w", flag.Name, parts[1], err)
					}
				} else {
					pluginArgs = append(pluginArgs, arg)
				}
			} else {
				if flag := persistentFlags.Lookup(flagArg); flag != nil {
					if flag.Value.Type() == "bool" {
						err := persistentFlags.Set(flagArg, "true")
						if err != nil {
							return nil, fmt.Errorf("failed to set persistent flag '%s' with value 'true': %w", flagArg, err)
						}
					} else {
						previousArgFlagName = flagArg
						handlePreviousArgAsFlag = true
					}
				} else if len(flagArg) == 1 && persistentFlags.ShorthandLookup(flagArg) != nil {
					flag := persistentFlags.ShorthandLookup(flagArg)
					if flag.Value.Type() == "bool" {
						err := persistentFlags.Set(flag.Name, "true")
						if err != nil {
							return nil, fmt.Errorf("failed to set persistent flag '%s' with value 'true': %w", flag.Name, err)
						}
					} else {
						previousArgFlagName = flag.Name
						handlePreviousArgAsFlag = true
					}
				} else {
					pluginArgs = append(pluginArgs, arg)
				}
			}
		default:
			pluginArgs = append(pluginArgs, arg)
		}
	}

	return pluginArgs, nil
}
