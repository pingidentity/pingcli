package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pingidentity/pingcli/cmd/config"
	"github.com/pingidentity/pingcli/cmd/feedback"
	"github.com/pingidentity/pingcli/cmd/platform"
	"github.com/pingidentity/pingcli/cmd/request"
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
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Long:          "A CLI tool for managing the configuration of Ping Identity products.",
		Short:         "A CLI tool for managing the configuration of Ping Identity products.",
		SilenceErrors: true, // Upon error in RunE method, let output package in main.go handle error output
		Use:           "pingcli",
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
			Message:      "Failed to validate Ping CLI configuration",
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
				Message: fmt.Sprintf("Ping CLI configuration file '%s' does not exist.", cfgFile),
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
	output.Print(output.Opts{
		Message: fmt.Sprintf("Creating new Ping CLI configuration file at: %s", cfgFile),
		Result:  output.ENUM_RESULT_NIL,
	})

	// MkdirAll does nothing if directories already exist. Create needed directories for config file location.
	err := os.MkdirAll(filepath.Dir(cfgFile), os.ModePerm)
	if err != nil {
		output.Print(output.Opts{
			Message:      fmt.Sprintf("Failed to make directories needed for new Ping CLI configuration file: %s", cfgFile),
			Result:       output.ENUM_RESULT_FAILURE,
			FatalMessage: err.Error(),
		})
	}

	tempViper := viper.New()
	tempViper.Set(options.RootActiveProfileOption.ViperKey, options.RootActiveProfileOption.DefaultValue)
	tempViper.Set(fmt.Sprintf("%s.%v", options.RootActiveProfileOption.DefaultValue.String(), options.ProfileDescriptionOption.ViperKey), "Default profile created by Ping CLI")

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

	// For each profile, if a viper key from an option doesn't exist, set it to nil
	for _, pName := range profiles.GetMainConfig().ProfileNames() {
		subViper := profiles.GetMainConfig().ViperInstance().Sub(pName)
		for _, opt := range options.Options() {
			if opt.ViperKey == "" || opt.ViperKey == options.RootActiveProfileOption.ViperKey {
				continue
			}
			if !subViper.IsSet(opt.ViperKey) {
				subViper.Set(opt.ViperKey, opt.DefaultValue)
			}
		}
		err := profiles.GetMainConfig().SaveProfile(pName, subViper)
		if err != nil {
			output.Print(output.Opts{
				Message: fmt.Sprintf("Error: Failed to save profile %s.", pName),
				Result:  output.ENUM_RESULT_FAILURE,
			})
		}
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
