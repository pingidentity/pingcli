// Copyright © 2025 Ping Identity Corporation

package configuration_root

import (
	"fmt"
	"os"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/pflag"
)

func InitRootOptions() {
	initActiveProfileOption()
	initProfileOption()
	initColorOption()
	initConfigOption()
	initDetailedExitCodeOption()
	initOutputFormatOption()
	initUnmaskSecretValuesOption()
}

func initActiveProfileOption() {
	defaultValue := customtypes.String("")

	options.RootActiveProfileOption = options.Option{
		CobraParamName:  "", // No cobra param
		CobraParamValue: nil,
		DefaultValue:    &defaultValue,
		EnvVar:          "",  // No env var
		Flag:            nil, // No flag
		Sensitive:       false,
		Type:            options.ENUM_STRING,
		ViperKey:        "activeProfile",
	}
}

func initProfileOption() {
	cobraParamName := "profile"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	options.RootProfileOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "PINGCLI_PROFILE",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "P",
			Usage:     "The name of a configuration profile to use.",
			Value:     cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_STRING,
		ViperKey:  "", // No viper key
	}
}

func initColorOption() {
	cobraParamName := "no-color"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)

	options.RootColorOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "PINGCLI_NO_COLOR",
		Flag: &pflag.Flag{
			Name:        cobraParamName,
			Usage:       "Disable text output in color. (default false)",
			Value:       cobraValue,
			NoOptDefVal: "true", // Make this flag a boolean flag
		},
		Sensitive: false,
		Type:      options.ENUM_BOOL,
		ViperKey:  "noColor",
	}
}

func initConfigOption() {
	cobraParamName := "config"
	cobraValue := new(customtypes.String)
	defaultValue := getDefaultConfigFilepath()

	options.RootConfigOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    defaultValue,
		EnvVar:          "PINGCLI_CONFIG",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "C",
			Usage: "The relative or full path to a custom Ping CLI configuration file. " +
				"(default $HOME/.pingcli/config.yaml)",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_STRING,
		ViperKey:  "", // No viper key
	}
}

func initDetailedExitCodeOption() {
	cobraParamName := "detailed-exitcode"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)

	options.RootDetailedExitCodeOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "PINGCLI_DETAILED_EXITCODE",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "D",
			Usage: "Enable detailed exit code output. (default false)" +
				"\n0 - pingcli command succeeded with no errors or warnings." +
				"\n1 - pingcli command failed with errors." +
				"\n2 - pingcli command succeeded with warnings.",
			Value:       cobraValue,
			NoOptDefVal: "true", // Make this flag a boolean flag
		},
		Sensitive: false,
		Type:      options.ENUM_BOOL,
		ViperKey:  "detailedExitCode",
	}
}

func initOutputFormatOption() {
	cobraParamName := "output-format"
	cobraValue := new(customtypes.OutputFormat)
	defaultValue := customtypes.OutputFormat(customtypes.ENUM_OUTPUT_FORMAT_TEXT)

	options.RootOutputFormatOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "PINGCLI_OUTPUT_FORMAT",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "O",
			Usage: fmt.Sprintf(
				"Specify the console output format. "+
					"(default %s)"+
					"\nOptions are: %s.",
				customtypes.ENUM_OUTPUT_FORMAT_TEXT,
				strings.Join(customtypes.OutputFormatValidValues(), ", "),
			),
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_OUTPUT_FORMAT,
		ViperKey:  "outputFormat",
	}
}

func initUnmaskSecretValuesOption() {
	cobraParamName := "unmask-values"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)

	options.ConfigUnmaskSecretValueOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No EnvVar
		Flag: &pflag.Flag{
			Name:        cobraParamName,
			Shorthand:   "U",
			Usage:       "Unmask secret values. (default false)",
			Value:       cobraValue,
			NoOptDefVal: "true", // Make this flag a boolean flag
		},
		Sensitive: false,
		Type:      options.ENUM_BOOL,
		ViperKey:  "", // No ViperKey
	}
}

func getDefaultConfigFilepath() (defaultConfigFilepath *customtypes.String) {
	l := logger.Get()

	defaultConfigFilepath = new(customtypes.String)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		l.Err(err).Msg("Failed to determine user's home directory")

		return nil
	}

	err = defaultConfigFilepath.Set(fmt.Sprintf("%s/.pingcli/config.yaml", homeDir))
	if err != nil {
		l.Err(err).Msg("Failed to set default config file path")

		return nil
	}

	return defaultConfigFilepath
}
