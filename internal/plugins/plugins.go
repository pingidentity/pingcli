// Copyright © 2025 Ping Identity Corporation

package plugins

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/hashicorp/go-hclog"
	hplugin "github.com/hashicorp/go-plugin"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/shared/grpc"
	shared_logger "github.com/pingidentity/pingcli/shared/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	pluginsErrorPrefix      = "plugins error"
	ErrGetPluginExecutables = errors.New("failed to get configured plugin executables")
	ErrCreateRPCClient      = errors.New("failed to create plugin rpc client")
	ErrDispensePlugin       = errors.New("the rpc client failed to dispense plugin executable")
	ErrCastPluginInterface  = errors.New("failed to cast plugin executable to grpc.PingCliCommand interface")
	ErrPluginConfiguration  = errors.New("failed to get plugin configuration")
	ErrExecutePlugin        = errors.New("failed to execute plugin command")
)

func AddAllPluginToCmd(cmd *cobra.Command) error {
	l := logger.Get()

	// Plugin executables are stored in the profile configuration
	// via the command 'pingcli plugin add <plugin-executable>'
	pluginExecutables, err := profiles.GetOptionValue(options.PluginExecutablesOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: pluginsErrorPrefix, Err: fmt.Errorf("%w: %w", ErrGetPluginExecutables, err)}
	}

	if pluginExecutables == "" {
		return nil
	}

	for pluginExecutable := range strings.SplitSeq(pluginExecutables, ",") {
		pluginExecutable = strings.TrimSpace(pluginExecutable)
		if pluginExecutable == "" {
			continue
		}

		ctx := cmd.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		conf, err := pluginConfiguration(ctx, pluginExecutable)
		if err != nil {
			return &errs.PingCLIError{Prefix: pluginsErrorPrefix, Err: err}
		}

		pluginCmd := &cobra.Command{
			DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
			DisableFlagParsing:    true, // Let all flags pass through to plugin
			Example:               conf.Example,
			Long:                  conf.Long,
			RunE:                  createCmdRunE(pluginExecutable),
			Short:                 conf.Short,
			Use:                   conf.Use,
		}

		cmd.AddCommand(pluginCmd)

		l.Info().Msgf("Loaded plugin executable: %s", pluginExecutable)
	}

	return nil
}

// createHPluginClient creates a new hplugin.Client for the given plugin executable.
// The caller is responsible for closing the client connection after use.
func createHPluginClient(ctx context.Context, pluginExecutable string) *hplugin.Client {
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
		Cmd:             exec.CommandContext(ctx, pluginExecutable),
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
		return nil, &errs.PingCLIError{Prefix: pluginsErrorPrefix, Err: fmt.Errorf("%w: %w", ErrCreateRPCClient, err)}
	}

	// All Ping CLI plugins are expected to serve the ENUM_PINGCLI_COMMAND_GRPC plugin via
	// the PluginMap within the plugin.Serve() method.
	// Non-Golang plugins unable to use the shared grpc module should supply the
	// raw value of ENUM_PINGCLI_COMMAND_GRPC "pingcli_command_grpc" for the PluginMap key.
	raw, err := clientProtocol.Dispense(grpc.ENUM_PINGCLI_COMMAND_GRPC)
	if err != nil {
		return nil, &errs.PingCLIError{Prefix: pluginsErrorPrefix, Err: fmt.Errorf("%w '%s': %w", ErrDispensePlugin, pluginExecutable, err)}
	}

	// Cast the dispensed plugin to the interface we expect to work with: grpc.PingCliCommand.
	// However, this is not a normal interface, but rather implemeted over the RPC connection.
	plugin, ok := raw.(grpc.PingCliCommand)
	if !ok {
		return nil, &errs.PingCLIError{Prefix: pluginsErrorPrefix, Err: fmt.Errorf("%w '%s'", ErrCastPluginInterface, pluginExecutable)}
	}

	return plugin, nil
}

func pluginConfiguration(ctx context.Context, pluginExecutable string) (conf *grpc.PingCliCommandConfiguration, err error) {
	client := createHPluginClient(ctx, pluginExecutable)
	defer client.Kill()

	plugin, err := dispensePlugin(client, pluginExecutable)
	if err != nil {
		return nil, &errs.PingCLIError{Prefix: pluginsErrorPrefix, Err: err}
	}

	// The configuration method is defined by the protobuf definition.
	// The plugin should return relevant cobra command information
	// even if the plugin does not use cobra commands internally.
	// This allows the host process Ping CLI to present the plugin command
	// in the help output.
	resp, err := plugin.Configuration()
	if err != nil {
		return nil, &errs.PingCLIError{Prefix: pluginsErrorPrefix, Err: fmt.Errorf("%w: %w", ErrPluginConfiguration, err)}
	}

	return resp, nil
}

func createCmdRunE(pluginExecutable string) func(cmd *cobra.Command, args []string) (err error) {
	return func(cmd *cobra.Command, args []string) error {
		// Because DisableFlagParsing is true, `args` contains all arguments after the command name.
		// We need to filter out the persistent flags that belong to the root command.
		pluginArgs := filterRootFlags(cmd, args)

		client := createHPluginClient(cmd.Context(), pluginExecutable)
		defer client.Kill()

		plugin, err := dispensePlugin(client, pluginExecutable)
		if err != nil {
			return &errs.PingCLIError{Prefix: pluginsErrorPrefix, Err: err}
		}

		err = plugin.Run(pluginArgs, &shared_logger.SharedLogger{})
		if err != nil {
			return &errs.PingCLIError{Prefix: pluginsErrorPrefix, Err: fmt.Errorf("%w: %w", ErrExecutePlugin, err)}
		}

		return nil
	}
}

// filterRootFlags filters out any flags that were parsed by the root command's persistent flags
// and processes them for the host application, returning only plugin-specific args.
func filterRootFlags(cmd *cobra.Command, args []string) []string {
	pluginArgs := make([]string, 0) // Initialize as an empty slice
	rootFlags := cmd.Root().PersistentFlags()

	// isRootFlag checks if a given argument (like "--profile") is a known persistent flag on the root command.
	isRootFlag := func(arg string) *pflag.Flag {
		// Positional arguments don't start with a hyphen, so they can't be flags.
		if !strings.HasPrefix(arg, "-") {
			return nil
		}

		name := strings.SplitN(strings.TrimLeft(arg, "-"), "=", 2)[0]

		if strings.HasPrefix(arg, "--") {
			return rootFlags.Lookup(name)
		}

		// It's a shorthand flag. pflag panics if the lookup key is > 1 character,
		// so we must ensure the name is a single character.
		// NOTE: This does not handle stacked short flags (e.g., `-vp`). The entire
		// stacked flag group will be passed to the plugin as a single argument.
		if len(name) == 1 {
			return rootFlags.ShorthandLookup(name)
		}

		return nil
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		flag := isRootFlag(arg)

		if flag == nil {
			// If it's not a recognized root flag, it must be for the plugin.
			pluginArgs = append(pluginArgs, arg)
			continue
		}

		// It is a root flag. We need to skip it and, if necessary, its value.
		// If the flag is a boolean, it has no separate value, so we just skip the flag itself.
		// If it's a non-boolean flag and the value is attached with '=', we also just skip this one argument.
		if flag.Value.Type() == "bool" || strings.Contains(arg, "=") {
			continue
		}

		// It's a non-boolean flag in the form "--flag value". We need to skip both.
		i++
	}
	return pluginArgs
}
