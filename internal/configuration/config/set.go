package configuration_config

import (
	"github.com/pingidentity/pingctl/internal/configuration/options"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitConfigSetOptions() {
	initSetProfileOption()
}

func initSetProfileOption() {
	cobraParamName := "profile"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	options.ConfigSetProfileOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "p",
			Usage:     "The configuration profile used to set the configuration value.",
			Value:     cobraValue,
			DefValue:  "The active profile",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "", // No viper key
	}
}
