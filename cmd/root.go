// Copyright © 2025 Ping Identity Corporation

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-hclog"
	hplugin "github.com/hashicorp/go-plugin"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/pingidentity/pingcli/cmd/completion"
	"github.com/pingidentity/pingcli/cmd/config"
	"github.com/pingidentity/pingcli/cmd/platform"
	"github.com/pingidentity/pingcli/cmd/plugin"
	"github.com/pingidentity/pingcli/cmd/request"
	"github.com/pingidentity/pingcli/internal/autocompletion"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/shared"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
func NewRootCommand(version string, commit string) *cobra.Command {
	l := logger.Get()

	l.Debug().Msgf("Initializing Ping CLI options...")
	configuration.InitAllOptions()

	l.Debug().Msgf("Initializing Root command...")

	initKoanfProfile()

	cmd := &cobra.Command{
		Long:          "A CLI tool for managing the configuration of Ping Identity products.",
		Short:         "A CLI tool for managing the configuration of Ping Identity products.",
		SilenceErrors: true, // Upon error in RunE method, let output package in main.go handle error output
		Use:           "pingcli",
		Version:       fmt.Sprintf("%s (commit: %s)", version, commit),
	}

	cmd.AddCommand(
		// auth.NewAuthCommand(),
		completion.Command(),
		config.NewConfigCommand(),
		platform.NewPlatformCommand(),
		plugin.NewPluginCommand(),
		request.NewRequestCommand(),
	)

	err := addPluginCommands(cmd)
	if err != nil {
		output.SystemError(fmt.Sprintf("Failed to add plugin commands: %v", err), nil)
	}

	// FLAGS //
	// --config, -C
	cmd.PersistentFlags().AddFlag(options.RootConfigOption.Flag)

	// --detailed-exitcode, -D
	cmd.PersistentFlags().AddFlag(options.RootDetailedExitCodeOption.Flag)

	// --profile, -P
	cmd.PersistentFlags().AddFlag(options.RootProfileOption.Flag)
	// auto-completion
	err = cmd.RegisterFlagCompletionFunc(options.RootProfileOption.CobraParamName, autocompletion.RootProfileFunc)
	if err != nil {
		output.SystemError(fmt.Sprintf("Unable to register auto completion for pingcli global flag %s: %v", options.RootProfileOption.CobraParamName, err), nil)
	}

	// --no-color
	cmd.PersistentFlags().AddFlag(options.RootColorOption.Flag)

	// --output-format, -O
	cmd.PersistentFlags().AddFlag(options.RootOutputFormatOption.Flag)
	// auto-completion
	err = cmd.RegisterFlagCompletionFunc(options.RootOutputFormatOption.CobraParamName, autocompletion.RootOutputFormatFunc)
	if err != nil {
		output.SystemError(fmt.Sprintf("Unable to register auto completion for pingcli global flag %s: %v", options.RootOutputFormatOption.CobraParamName, err), nil)
	}

	// Make sure cobra is outputting to stdout and stderr
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

	return cmd
}

func initKoanfProfile() {
	l := logger.Get()

	cfgFile, err := profiles.GetOptionValue(options.RootConfigOption)
	if err != nil {
		output.SystemError(fmt.Sprintf("Failed to get configuration file location: %v", err), nil)
	}

	l.Debug().Msgf("Using configuration file location for initialization: %s", cfgFile)

	// Handle the config file location
	checkCfgFileLocation(cfgFile)

	l.Debug().Msgf("Validated configuration file location at: %s", cfgFile)

	// Configure the koanf instance
	initKoanf(cfgFile)

	userDefinedProfile, err := profiles.GetOptionValue(options.RootProfileOption)
	if err != nil {
		output.SystemError(fmt.Sprintf("Failed to get user-defined profile: %v", err), nil)
	}

	configFileActiveProfile, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		output.SystemError(fmt.Sprintf("Failed to get active profile from configuration file: %v", err), nil)
	}

	if userDefinedProfile != "" {
		l.Debug().Msgf("Using configuration profile: %s", userDefinedProfile)
	} else {
		l.Debug().Msgf("Using configuration profile: %s", configFileActiveProfile)
	}

	// Configure the profile koanf instance
	if err := profiles.GetKoanfConfig().ChangeActiveProfile(configFileActiveProfile); err != nil {
		output.UserFatal(fmt.Sprintf("Failed to set active profile: %v", err), nil)
	}

	// Validate the configuration
	if err := profiles.Validate(); err != nil {
		output.UserFatal(fmt.Sprintf("%v", err), nil)
	}
}

func checkCfgFileLocation(cfgFile string) {
	// Check existence of configuration file
	_, err := os.Stat(cfgFile)
	if os.IsNotExist(err) {
		// Only create a new configuration file if it is the default configuration file location
		if cfgFile == options.RootConfigOption.DefaultValue.String() {
			output.Message(fmt.Sprintf("Ping CLI configuration file '%s' does not exist.", cfgFile), nil)

			createConfigFile(options.RootConfigOption.DefaultValue.String())
		} else {
			output.UserFatal(fmt.Sprintf("Configuration file '%s' does not exist. Use the default configuration file location or specify a valid configuration file location with the --config flag.", cfgFile), nil)
		}
	} else if err != nil {
		output.SystemError(fmt.Sprintf("Failed to check if configuration file '%s' exists: %v", cfgFile, err), nil)
	}
}

func createConfigFile(cfgFile string) {
	output.Message(fmt.Sprintf("Creating new Ping CLI configuration file at: %s", cfgFile), nil)

	// MkdirAll does nothing if directories already exist. Create needed directories for config file location.
	err := os.MkdirAll(filepath.Dir(cfgFile), os.FileMode(0700))
	if err != nil {
		output.SystemError(fmt.Sprintf("Failed to make the directory for the new configuration file '%s': %v", cfgFile, err), nil)
	}

	tempKoanf := profiles.NewKoanfConfig(cfgFile)
	err = tempKoanf.KoanfInstance().Set(options.RootActiveProfileOption.KoanfKey, "default")
	if err != nil {
		output.SystemError(fmt.Sprintf("Failed to set active profile in new configuration file '%s': %v", cfgFile, err), nil)
	}

	err = tempKoanf.KoanfInstance().Set(fmt.Sprintf("default.%v", options.ProfileDescriptionOption.KoanfKey), "Default profile created by Ping CLI")
	if err != nil {
		output.SystemError(fmt.Sprintf("Failed to set default profile description in new configuration file '%s': %v", cfgFile, err), nil)
	}

	err = tempKoanf.WriteFile()
	if err != nil {
		output.SystemError(fmt.Sprintf("Failed to create new configuration file '%s': %v", cfgFile, err), nil)
	}
}

func initKoanf(cfgFile string) {
	l := logger.Get()

	loadKoanfConfig(cfgFile)

	// If there are no profiles in the configuration file, seed the default profile
	if len(profiles.GetKoanfConfig().ProfileNames()) == 0 {
		l.Debug().Msgf("No profiles found in configuration file. Creating default profile in configuration file '%s'", cfgFile)
		createConfigFile(cfgFile)
		loadKoanfConfig(cfgFile)
	}

	err := profiles.GetKoanfConfig().DefaultMissingKoanfKeys()
	if err != nil {
		output.SystemError(err.Error(), nil)
	}
}

func loadKoanfConfig(cfgFile string) {
	l := logger.Get()

	koanfConfig := profiles.GetKoanfConfig()
	koanfConfig.SetKoanfConfigFile(cfgFile)

	// Use config file from the flag.
	if err := koanfConfig.KoanfInstance().Load(file.Provider(cfgFile), yaml.Parser()); err != nil {
		output.SystemError(fmt.Sprintf("Failed to load configuration from file '%s': %v", cfgFile, err), nil)
	} else {
		l.Info().Msgf("Using configuration file: %s", cfgFile)
	}

	_, err := koanfConfig.KoanfInstance().Marshal(yaml.Parser())
	if err != nil {
		output.SystemError(fmt.Sprintf("Failed to marshal configuration file '%s': %v", cfgFile, err), nil)
	}
}

func addPluginCommands(cmd *cobra.Command) error {
	l := logger.Get()
	pluginExecutables, err := profiles.GetOptionValue(options.PluginExecutablesOption)
	if err != nil {
		return fmt.Errorf("failed to get configured plugin executables: %w", err)
	}

	if pluginExecutables == "" {
		return nil
	}

	pluginLogger := hclog.New(&hclog.LoggerOptions{
		Name:   "pingcli",
		Output: os.Stdout,
		Level:  hclog.Warn,
	})

	for _, pluginExecutable := range strings.Split(pluginExecutables, ",") {
		client := hplugin.NewClient(&hplugin.ClientConfig{
			HandshakeConfig: shared.HandshakeConfig,
			Plugins:         shared.PluginMap,
			Cmd:             exec.Command(pluginExecutable),
			AllowedProtocols: []hplugin.Protocol{
				hplugin.ProtocolGRPC,
			},
			Logger:     pluginLogger,
			SyncStdout: os.Stdout,
			SyncStderr: os.Stderr,
		})

		rpcClient, err := client.Client()
		if err != nil {
			return fmt.Errorf("failed to create Plugin RPC client: %w", err)
		}

		raw, err := rpcClient.Dispense(shared.ENUM_PINGCLI_COMMAND_GRPC)
		if err != nil {
			return fmt.Errorf("failed to dispense Plugin: %w", err)
		}

		plugin, ok := raw.(shared.PingCliCommand)
		if !ok {
			return fmt.Errorf("failed to cast Plugin to PingCliCommand Interface")
		}

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

				err := plugin.Run(args)
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
