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
		EnvVar:          "PINGCLI_ACTIVE_PROFILE",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "P",
			Usage:     "The name of the stored custom configuration profile to use.",
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
		EnvVar:          "PINGCLI_COLOR",
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    "Show text output in color.",
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
		EnvVar:          "PINGCLI_CONFIG",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "C",
			Usage:     "The relative or full path to a custom Ping CLI configuration file.",
			Value:     cobraValue,
			DefValue:  "\"$HOME/.pingcli/config.yaml\"",
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
		EnvVar:          "PINGCLI_OUTPUT_FORMAT",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "O",
			Usage:     fmt.Sprintf("Specify the console output format. Options are: %s", strings.Join(customtypes.OutputFormatValidValues(), ", ")),
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

	err = defaultConfigFilepath.Set(fmt.Sprintf("%s/.pingcli/config.yaml", homeDir))
	if err != nil {
		l.Err(err).Msg("Failed to set default config file path")
		return nil
	}

	return defaultConfigFilepath
}
