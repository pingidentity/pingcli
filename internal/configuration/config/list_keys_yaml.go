// Copyright © 2025 Ping Identity Corporation

package configuration_config

import (
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitConfigListKeyOptions() {
	initConfigListKeysYAMLOption()
}

func initConfigListKeysYAMLOption() {
	cobraParamName := "yaml"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)

	options.ConfigListKeysYamlOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "y",
			Usage: "Output configuration keys in YAML format. " +
				"(default false)",
			Value:       cobraValue,
			NoOptDefVal: "true", // Make this flag a boolean flag
		},
		Sensitive: false,
		Type:      options.ENUM_BOOL,
		ViperKey:  "", // No viper key
	}
}
