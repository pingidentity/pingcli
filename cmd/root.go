package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pingidentity/pingctl/cmd/config"
	"github.com/pingidentity/pingctl/cmd/feedback"
	"github.com/pingidentity/pingctl/cmd/platform"
	"github.com/pingidentity/pingctl/cmd/request"
	"github.com/pingidentity/pingctl/internal/configuration"
	"github.com/pingidentity/pingctl/internal/configuration/options"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	ConfigurationFileFormat = `Configuration File Format:
activeProfile: <ProfileName>

<ProfileName>:
	color: <true|false>
	outputFormat: <Format>
	...`
)

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Pingctl options...")
	configuration.InitAllOptions()

	l.Debug().Msgf("Initializing Root command...")
	cobra.OnInitialize(initViperProfile)
}

// rootCmd represents the base command when called without any subcommands
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Example:       ConfigurationFileFormat,
		Long:          `A CLI tool for managing Ping Identity products.`,
		Short:         "A CLI tool for managing Ping Identity products.",
		SilenceErrors: true, // Upon error in RunE method, let output package in main.go handle error output
		Use:           "pingctl",
		Version:       "v2.0.0-alpha.4",
	}

	cmd.AddCommand(
		// auth.NewAuthCommand(),
		config.NewConfigCommand(),
		feedback.NewFeedbackCommand(),
		platform.NewPlatformCommand(),
		request.NewRequestCommand(),
	)

	cmd.PersistentFlags().AddFlag(options.RootConfigOption.Flag)
	cmd.PersistentFlags().AddFlag(options.RootActiveProfileOption.Flag)
	cmd.PersistentFlags().AddFlag(options.RootOutputFormatOption.Flag)
	cmd.PersistentFlags().AddFlag(options.RootColorOption.Flag)

	// Make sure cobra is outputting to stdout and stderr
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

	return cmd
}

func initViperProfile() {
	l := logger.Get()

	cfgFile, err := profiles.GetOptionValue(options.RootConfigOption)
	if err != nil {
		output.Print(output.Opts{
			Message:      "Failed to get configuration file location",
			Result:       output.ENUM_RESULT_FAILURE,
			FatalMessage: err.Error(),
		})
	}

	l.Debug().Msgf("Using configuration file location for initialization: %s", cfgFile)

	// Handle the config file location
	checkCfgFileLocation(cfgFile)

	l.Debug().Msgf("Validated configuration file location at: %s", cfgFile)

	//Configure the main viper instance
	initMainViper(cfgFile)

	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		output.Print(output.Opts{
			Message:      "Failed to get active profile",
			Result:       output.ENUM_RESULT_FAILURE,
			FatalMessage: err.Error(),
		})
	}

	l.Debug().Msgf("Using configuration profile: %s", profileName)

	// Configure the profile viper instance
	if err := profiles.GetMainConfig().ChangeActiveProfile(profileName); err != nil {
		output.Print(output.Opts{
			Message:      "Failed to set profile viper",
			Result:       output.ENUM_RESULT_FAILURE,
			FatalMessage: err.Error(),
		})
	}

	// Validate the configuration
	if err := profiles.Validate(); err != nil {
		output.Print(output.Opts{
			Message:      "Failed to validate pingctl configuration",
			Result:       output.ENUM_RESULT_FAILURE,
			FatalMessage: err.Error(),
		})
	}
}

func checkCfgFileLocation(cfgFile string) {
	// Check existence of configuration file
	_, err := os.Stat(cfgFile)
	if os.IsNotExist(err) {
		// Only create a new configuration file if it is the default configuration file location
		if cfgFile == options.RootConfigOption.DefaultValue.String() {
			output.Print(output.Opts{
				Message: fmt.Sprintf("Pingctl configuration file '%s' does not exist.", cfgFile),
				Result:  output.ENUM_RESULT_NOACTION_WARN,
			})

			createConfigFile(options.RootConfigOption.DefaultValue.String())
		} else {
			output.Print(output.Opts{
				Message:      fmt.Sprintf("Configuration file '%s' does not exist.", cfgFile),
				Result:       output.ENUM_RESULT_FAILURE,
				FatalMessage: fmt.Sprintf("Configuration file '%s' does not exist. Use the default configuration file location or specify a valid configuration file location with the --config flag.", cfgFile),
			})
		}
	} else if err != nil {
		output.Print(output.Opts{
			Message:      fmt.Sprintf("Failed to check if configuration file '%s' exists", cfgFile),
			Result:       output.ENUM_RESULT_FAILURE,
			FatalMessage: err.Error(),
		})
	}

}

func createConfigFile(cfgFile string) {
	l := logger.Get()
	l.Debug().Msgf("Creating new pingctl configuration file at: %s", cfgFile)

	// MkdirAll does nothing if directories already exist. Create needed directories for config file location.
	err := os.MkdirAll(filepath.Dir(cfgFile), os.ModePerm)
	if err != nil {
		output.Print(output.Opts{
			Message:      fmt.Sprintf("Failed to make directories needed for new pingctl configuration file: %s", cfgFile),
			Result:       output.ENUM_RESULT_FAILURE,
			FatalMessage: err.Error(),
		})
	}

	tempViper := viper.New()
	tempViper.Set(options.RootActiveProfileOption.ViperKey, options.RootActiveProfileOption.DefaultValue)
	tempViper.Set(fmt.Sprintf("%s.%v", options.RootActiveProfileOption.DefaultValue.String(), options.ProfileDescriptionOption.ViperKey), "Default profile created by pingctl")

	err = tempViper.WriteConfigAs(cfgFile)
	if err != nil {
		output.Print(output.Opts{
			Message:      fmt.Sprintf("Failed to create configuration file at: %s", cfgFile),
			Result:       output.ENUM_RESULT_FAILURE,
			FatalMessage: err.Error(),
		})
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
}

func loadMainViperConfig(cfgFile string) {
	l := logger.Get()

	mainViper := profiles.GetMainConfig().ViperInstance()
	// Use config file from the flag.
	mainViper.SetConfigFile(cfgFile)

	// If a config file is found, read it in.
	if err := mainViper.ReadInConfig(); err != nil {
		output.Print(output.Opts{
			Message:      fmt.Sprintf("Failed to read configuration from file: %s", cfgFile),
			Result:       output.ENUM_RESULT_FAILURE,
			FatalMessage: err.Error(),
		})
	} else {
		l.Info().Msgf("Using configuration file: %s", mainViper.ConfigFileUsed())
	}
}
