package configuration_config

import (
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitConfigGetOptions() {
	initGetProfileOption()
}

func initGetProfileOption() {
	cobraParamName := "profile"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	options.ConfigGetProfileOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "p",
			Usage:     "The name of the configuration profile used to get the configuration value from.",
			Value:     cobraValue,
			DefValue:  "The active profile",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "", // No viper key
	}
}
