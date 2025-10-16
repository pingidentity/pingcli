// Copyright © 2025 Ping Identity Corporation

package configuration_auth

import (
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitAuthOptions() {
	initAuthMethodDeviceCodeOption()
	initAuthMethodClientCredentialsOption()
	initAuthMethodAuthCodeOption()
}

func initAuthMethodDeviceCodeOption() {
	cobraParamName := "device-code"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)

	options.AuthMethodDeviceCodeOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:        cobraParamName,
			Shorthand:   "d",
			Usage:       "Use device code authentication flow",
			Value:       cobraValue,
			NoOptDefVal: "true", // Make this flag a boolean flag
		},
		Sensitive: false,
		Type:      options.BOOL,
		KoanfKey:  "", // No koanf key
	}
}

func initAuthMethodClientCredentialsOption() {
	cobraParamName := "client-credentials"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)

	options.AuthMethodClientCredentialsOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:        cobraParamName,
			Shorthand:   "c",
			Usage:       "Use client credentials authentication flow",
			Value:       cobraValue,
			NoOptDefVal: "true", // Make this flag a boolean flag
		},
		Sensitive: false,
		Type:      options.BOOL,
		KoanfKey:  "", // No koanf key
	}
}

func initAuthMethodAuthCodeOption() {
	cobraParamName := "auth-code"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)

	options.AuthMethodAuthCodeOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:        cobraParamName,
			Shorthand:   "a",
			Usage:       "Use authorization code authentication flow",
			Value:       cobraValue,
			NoOptDefVal: "true", // Make this flag a boolean flag
		},
		Sensitive: false,
		Type:      options.BOOL,
		KoanfKey:  "", // No koanf key
	}
}
