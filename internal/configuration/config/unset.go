package configuration_config

import (
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitConfigUnsetOptions() {
	initUnsetProfileOption()
}

func initUnsetProfileOption() {
	cobraParamName := "profile-name"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	options.ConfigUnsetProfileOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "p",
			Usage:     "The name of the configuration profile to unset a configuration value from.",
			Value:     cobraValue,
			DefValue:  "The active profile",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "", // No viper key
	}
}
