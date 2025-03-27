// Copyright © 2025 Ping Identity Corporation

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pingidentity/pingcli/cmd/completion"
	"github.com/pingidentity/pingcli/cmd/config"
	"github.com/pingidentity/pingcli/cmd/feedback"
	"github.com/pingidentity/pingcli/cmd/platform"
	"github.com/pingidentity/pingcli/cmd/request"
	"github.com/pingidentity/pingcli/internal/autocompletion"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Ping CLI options...")
	configuration.InitAllOptions()

	l.Debug().Msgf("Initializing Root command...")
	cobra.OnInitialize(initViperProfile)
}

// rootCmd represents the base command when called without any subcommands
func NewRootCommand(version string, commit string) *cobra.Command {
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
		feedback.NewFeedbackCommand(),
		platform.NewPlatformCommand(),
		request.NewRequestCommand(),
	)

	// FLAGS //
	// --config, -C
	cmd.PersistentFlags().AddFlag(options.RootConfigOption.Flag)

	// --detailed-exitcode, -D
	cmd.PersistentFlags().AddFlag(options.RootDetailedExitCodeOption.Flag)

	// --profile, -P
	cmd.PersistentFlags().AddFlag(options.RootProfileOption.Flag)
	// auto-completion
	err := cmd.RegisterFlagCompletionFunc(options.RootProfileOption.CobraParamName, autocompletion.RootProfileFunc)
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

func initViperProfile() {
	l := logger.Get()

	cfgFile, err := profiles.GetOptionValue(options.RootConfigOption)
	if err != nil {
		output.SystemError(fmt.Sprintf("Failed to get configuration file location: %v", err), nil)
	}

	l.Debug().Msgf("Using configuration file location for initialization: %s", cfgFile)

	// Handle the config file location
	checkCfgFileLocation(cfgFile)

	l.Debug().Msgf("Validated configuration file location at: %s", cfgFile)

	// Configure the main viper instance
	initMainViper(cfgFile)

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

	// Configure the profile viper instance
	if err := profiles.GetMainConfig().ChangeActiveProfile(configFileActiveProfile); err != nil {
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

	tempViper := viper.New()
	tempViper.Set(options.RootActiveProfileOption.ViperKey, "default")
	tempViper.Set(fmt.Sprintf("default.%v", options.ProfileDescriptionOption.ViperKey), "Default profile created by Ping CLI")

	err = tempViper.WriteConfigAs(cfgFile)
	if err != nil {
		output.SystemError(fmt.Sprintf("Failed to create configuration file '%s': %v", cfgFile, err), nil)
	}
}

func initMainViper(cfgFile string) {
	l := logger.Get()

	loadMainViperConfig(cfgFile)

	// If there are no profiles in the configuration file, seed the default profile
	if len(profiles.GetMainConfig().ProfileNames()) == 0 {
		l.Debug().Msgf("No profiles found in configuration file. Creating default profile in configuration file '%s'", cfgFile)
		createConfigFile(cfgFile)
		loadMainViperConfig(cfgFile)
	}

	err := profiles.GetMainConfig().DefaultMissingViperKeys()
	if err != nil {
		output.SystemError(err.Error(), nil)
	}
}

func loadMainViperConfig(cfgFile string) {
	l := logger.Get()

	mainViper := profiles.GetMainConfig().ViperInstance()
	// Use config file from the flag.
	mainViper.SetConfigFile(cfgFile)
	mainViper.SetConfigType("yaml")

	// If a config file is found, read it in.
	if err := mainViper.ReadInConfig(); err != nil {
		output.SystemError(fmt.Sprintf("Failed to read configuration from file '%s': %v", cfgFile, err), nil)
	} else {
		l.Info().Msgf("Using configuration file: %s", mainViper.ConfigFileUsed())
	}
}
