// Copyright © 2025 Ping Identity Corporation

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/pingidentity/pingcli/cmd/auth"
	"github.com/pingidentity/pingcli/cmd/completion"
	"github.com/pingidentity/pingcli/cmd/config"
	"github.com/pingidentity/pingcli/cmd/feedback"
	"github.com/pingidentity/pingcli/cmd/license"
	"github.com/pingidentity/pingcli/cmd/platform"
	"github.com/pingidentity/pingcli/cmd/plugin"
	"github.com/pingidentity/pingcli/cmd/request"
	"github.com/pingidentity/pingcli/internal/autocompletion"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/plugins"
	"github.com/pingidentity/pingcli/internal/profiles"
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
		auth.NewLoginCommand(),
		auth.NewLogoutCommand(),
		completion.Command(),
		config.NewConfigCommand(),
		feedback.NewFeedbackCommand(),
		platform.NewPlatformCommand(),
		plugin.NewPluginCommand(),
		request.NewRequestCommand(),
		license.NewLicenseCommand(),
	)

	err := plugins.AddAllPluginToCmd(cmd)
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

	cfgFile := ParseArgsForConfigFile(os.Args)
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

// ParseArgsForConfigFile parses the command line arguments for the configuration file flag.
func ParseArgsForConfigFile(args []string) string {
	for i, arg := range args {
		// Handle --config=value format
		if strings.HasPrefix(arg, fmt.Sprintf("--%s=", options.RootConfigOption.CobraParamName)) {
			return strings.TrimPrefix(arg, fmt.Sprintf("--%s=", options.RootConfigOption.CobraParamName))
		}
		// Handle -C=value format
		if strings.HasPrefix(arg, fmt.Sprintf("-%s=", options.RootConfigOption.Flag.Shorthand)) {
			return strings.TrimPrefix(arg, fmt.Sprintf("-%s=", options.RootConfigOption.Flag.Shorthand))
		}
		// Handle --config value format
		if arg == fmt.Sprintf("--%s", options.RootConfigOption.CobraParamName) && i+1 < len(args) {
			return args[i+1]
		}
		// Handle -C value format
		if arg == fmt.Sprintf("-%s", options.RootConfigOption.Flag.Shorthand) && i+1 < len(args) {
			return args[i+1]
		}
	}

	// No --config flag found, check environment variable
	if envValue := os.Getenv(options.RootConfigOption.EnvVar); envValue != "" {
		return envValue
	}

	// Fall back to default value
	return options.RootConfigOption.DefaultValue.String()
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

	koanfConfig := profiles.NewKoanfConfig(cfgFile)

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
