package configuration_config

import (
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitConfigOptions() {
	initConfigProfileOption()
	initConfigNameOption()
	initConfigDescriptionOption()
}

func initConfigProfileOption() {
	cobraParamName := "profile"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	options.ConfigProfileOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "p",
			Usage:     "The name of the profile to update.",
			Value:     cobraValue,
			DefValue:  "The active profile",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "", // No viper key
	}
}

func initConfigNameOption() {
	cobraParamName := "name"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	options.ConfigNameOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "n",
			Usage:     "The new name for the profile.",
			Value:     cobraValue,
			DefValue:  "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "", // No viper key
	}
}

func initConfigDescriptionOption() {
	cobraParamName := "description"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	options.ConfigDescriptionOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "d",
			Usage:     "The new description for the profile.",
			Value:     cobraValue,
			DefValue:  "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "", // No viper key
	}
}
