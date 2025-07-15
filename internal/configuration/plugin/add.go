// Copyright © 2025 Ping Identity Corporation

package configuration_plugin

import (
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
)

func InitPluginOptions() {
	initPluginExecutablesOption()
}

func initPluginExecutablesOption() {
	defaultValue := customtypes.StringSlice([]string{})

	options.PluginExecutablesOption = options.Option{
		CobraParamName:  "", // No cobra param
		CobraParamValue: nil,
		DefaultValue:    &defaultValue,
		EnvVar:          "",  // No env var
		Flag:            nil, // No flag
		Sensitive:       false,
		Type:            options.STRING_SLICE,
		KoanfKey:        "plugins",
	}
}
