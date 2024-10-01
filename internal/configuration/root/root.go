package configuration_root

import (
	"fmt"
	"os"
	"strings"

	"github.com/pingidentity/pingctl/internal/configuration/options"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/pflag"
)

func InitRootOptions() {
	initActiveProfileOption()
	initColorOption()
	initConfigOption()
	initOutputFormatOption()
}

func initActiveProfileOption() {
	cobraParamName := "active-profile"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("default")

	options.RootActiveProfileOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "PINGCTL_ACTIVE_PROFILE",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "P",
			Usage:     "Profile to use from configuration file",
			Value:     cobraValue,
			DefValue:  "default",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "activeProfile",
	}
}

func initColorOption() {
	cobraParamName := "color"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(true)

	options.RootColorOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "PINGCTL_COLOR",
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    "Use colorized output",
			Value:    cobraValue,
			DefValue: "true",
		},
		Type:     options.ENUM_BOOL,
		ViperKey: "color",
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
		EnvVar:          "PINGCTL_CONFIG",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "C",
			Usage:     "Configuration file location",
			Value:     cobraValue,
			DefValue:  "\"$HOME/.pingctl/config.yaml\"",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "", // No viper key
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
		EnvVar:          "PINGCTL_OUTPUT_FORMAT",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "O",
			Usage:     fmt.Sprintf("Specifies pingctl's console output format. Allowed: %s", strings.Join(customtypes.OutputFormatValidValues(), ", ")),
			Value:     cobraValue,
			DefValue:  customtypes.ENUM_OUTPUT_FORMAT_TEXT,
		},
		Type:     options.ENUM_OUTPUT_FORMAT,
		ViperKey: "outputFormat",
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

	err = defaultConfigFilepath.Set(fmt.Sprintf("%s/.pingctl/config.yaml", homeDir))
	if err != nil {
		l.Err(err).Msg("Failed to set default config file path")
		return nil
	}

	return defaultConfigFilepath
}
