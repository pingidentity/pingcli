package configuration_config

import (
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitConfigSetOptions() {
	initSetProfileOption()
}

func initSetProfileOption() {
	cobraParamName := "profile-name"
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
			Usage: "The name of the configuration profile used to set the configuration value to. " +
				"(default The active profile)",
			Value: cobraValue,
		},
		Type:     options.ENUM_STRING,
		ViperKey: "", // No viper key
	}
}
