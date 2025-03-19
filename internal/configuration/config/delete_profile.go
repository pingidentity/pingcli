// Copyright © 2025 Ping Identity Corporation

package configuration_config

import (
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitConfigDeleteProfileOptions() {
	initDeleteAutoAcceptOption()
}

func initDeleteAutoAcceptOption() {
	cobraParamName := "yes"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)

	options.ConfigDeleteAutoAcceptOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "y",
			Usage: "Auto-accept the profile deletion confirmation prompt. " +
				"(default false)",
			Value:       cobraValue,
			NoOptDefVal: "true", // Make the flag a boolean flag
		},
		Sensitive: false,
		Type:      options.ENUM_STRING,
		ViperKey:  "", // No viper key
	}
}
