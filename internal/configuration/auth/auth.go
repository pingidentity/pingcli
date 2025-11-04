// Copyright © 2025 Ping Identity Corporation

package configuration_auth

import (
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

// InitAuthOptions initializes all authentication-related configuration options
func InitAuthOptions() {
	initAuthMethodDeviceCodeOption()
	initAuthMethodClientCredentialsOption()
	initAuthMethodAuthCodeOption()
	initAuthFileStorageOption()
}

// initAuthMethodDeviceCodeOption initializes the --device-code authentication method flag
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

// initAuthMethodClientCredentialsOption initializes the --client-credentials authentication method flag
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

// initAuthMethodAuthCodeOption initializes the --auth-code authentication method flag
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

// initAuthFileStorageOption initializes the --file-storage flag for controlling file storage of auth tokens
func initAuthFileStorageOption() {
	cobraParamName := "file-storage"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)
	envVar := "PINGCLI_AUTH_FILE_STORAGE"

	options.AuthFileStorageOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:        cobraParamName,
			Usage:       "Store authentication tokens in file storage only, bypassing keychain. By default, keychain is attempted first with automatic fallback to file storage.",
			Value:       cobraValue,
			NoOptDefVal: "true", // Make this flag a boolean flag
		},
		Sensitive: false,
		Type:      options.BOOL,
		KoanfKey:  "auth.fileStorage",
	}
}
