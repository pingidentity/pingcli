package configuration_config

import (
	"github.com/pingidentity/pingctl/internal/configuration/options"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitConfigSetActiveProfileOptions() {
	initSetActiveProfileOption()
}

func initSetActiveProfileOption() {
	cobraParamName := "profile"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	options.ConfigSetActiveProfileOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "p",
			Usage:     "The configuration profile to set as the active profile.",
			Value:     cobraValue,
			DefValue:  "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "", // No viper key
	}
}
