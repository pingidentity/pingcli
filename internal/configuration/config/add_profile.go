package configuration_config

import (
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitConfigAddProfileOptions() {
	initAddProfileDescriptionOption()
	initAddProfileNameOption()
	initAddProfileSetActiveOption()
}

func initAddProfileDescriptionOption() {
	cobraParamName := "description"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	options.ConfigAddProfileDescriptionOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "d",
			Usage:     "The description of the new configuration profile.",
			Value:     cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_STRING,
		ViperKey:  "", // No viper key
	}
}

func initAddProfileNameOption() {
	cobraParamName := "name"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	options.ConfigAddProfileNameOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "n",
			Usage:     "The name of the new configuration profile.",
			Value:     cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_STRING,
		ViperKey:  "", // No viper key
	}
}

func initAddProfileSetActiveOption() {
	cobraParamName := "set-active"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)

	options.ConfigAddProfileSetActiveOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "s",
			Usage: "Set the new configuration profile as the active profile. " +
				"(default false)",
			Value:       cobraValue,
			NoOptDefVal: "true", // Make this flag a boolean flag
		},
		Sensitive: false,
		Type:      options.ENUM_BOOL,
		ViperKey:  "", // No viper key
	}
}
