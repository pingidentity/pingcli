// Copyright © 2026 Ping Identity Corporation

package configuration_auth

import (
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingone-go-client/config"
	"github.com/spf13/pflag"
)

// InitAuthOptions initializes all authentication-related configuration options
func InitAuthOptions() {
	initAuthMethodDeviceCodeOption()
	initAuthMethodClientCredentialsOption()
	initAuthMethodAuthorizationCodeOption()
	initAuthStorageOption()
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

// initAuthStorageOption initializes the --storage-type flag for controlling file storage of auth tokens
func initAuthStorageOption() {
	cobraParamName := "storage-type"
	// Use custom type wrapper compatible with pflag.Value
	cobraValue := new(customtypes.StorageType)
	// Default to secure local (keychain) storage when not specified
	defaultValue := customtypes.StorageType(config.StorageTypeSecureLocal)
	envVar := "PINGCLI_LOGIN_STORAGE_TYPE"

	options.AuthStorageOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: "Auth token storage (default: secure_local)\n" +
				"  secure_local  - Use OS keychain (default)\n" +
				"  file_system   - Store tokens in ~/.pingcli/credentials\n" +
				"  none          - Do not persist tokens",
			Value: cobraValue,
			// Require an explicit value to avoid noisy help like string[=...] output
			NoOptDefVal: "",
		},
		Sensitive: false,
		Type:      options.STORAGE_TYPE,
		KoanfKey:  "login.storage.type",
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
