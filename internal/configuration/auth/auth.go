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
	initAuthMethodAuthorizationCodeOption()
	initAuthFileStorageOption()
	initAuthProviderOption()
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
			Usage:       "Use device authorization flow",
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
			Usage:       "Use client credentials flow",
			Value:       cobraValue,
			NoOptDefVal: "true", // Make this flag a boolean flag
		},
		Sensitive: false,
		Type:      options.BOOL,
		KoanfKey:  "", // No koanf key
	}
}

// initAuthMethodAuthorizationCodeOption initializes the --authorization-code authentication method flag
func initAuthMethodAuthorizationCodeOption() {
	cobraParamName := "authorization-code"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)

	options.AuthMethodAuthorizationCodeOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:        cobraParamName,
			Shorthand:   "a",
			Usage:       "Use authorization code flow",
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
			Usage:       "Store authentication tokens in local file storage only. Without this flag, keychain storage is attempted first with fallback to local file storage.",
			Value:       cobraValue,
			NoOptDefVal: "true", // Make this flag a boolean flag
		},
		Sensitive: false,
		Type:      options.BOOL,
		KoanfKey:  "login.fileStorage",
	}
}

// initAuthProviderOption initializes the --provider flag for specifying which provider to authenticate with
func initAuthProviderOption() {
	cobraParamName := "provider"
	cobraValue := new(customtypes.AuthProvider)
	defaultValue := customtypes.AuthProvider(customtypes.ENUM_AUTH_PROVIDER_PINGONE)

	options.AuthProviderOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "p",
			Usage:     "Authentication provider to use. Defaults to 'pingone' if not specified.",
			Value:     cobraValue,
		},
		Sensitive: false,
		Type:      options.STRING,
		KoanfKey:  "", // No koanf key
	}
}
