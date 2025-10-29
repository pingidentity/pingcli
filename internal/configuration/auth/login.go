// Copyright © 2025 Ping Identity Corporation

package configuration_auth

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

// InitAuthOptions initializes all authentication-related configuration options
func InitAuthOptions() {
	initAuthMethodDeviceCodeOption()
	initAuthMethodClientCredentialsOption()
	initAuthMethodAuthCodeOption()
	initAuthServiceOption()
	initAuthUseKeychainOption()
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

// initAuthServiceOption initializes the --service flag for specifying which services to authenticate
func initAuthServiceOption() {
	cobraParamName := "service"
	cobraValue := new(customtypes.AuthServices)
	defaultValue := customtypes.AuthServices([]string{})
	envVar := "PINGCLI_AUTH_SERVICE"

	options.AuthServiceOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "s",
			Usage: fmt.Sprintf(
				"Specifies the service(s) to authenticate. Accepts a comma-separated string to delimit multiple services. "+
					"\nOptions are: %s."+
					"\nExample: '%s,%s'",
				strings.Join(customtypes.AuthServicesValidValues(), ", "),
				customtypes.ENUM_AUTH_SERVICE_PINGONE,
				customtypes.ENUM_AUTH_SERVICE_PINGFEDERATE,
			),
			Value:    cobraValue,
			DefValue: "",
		},
		Sensitive: false,
		Type:      options.AUTH_SERVICES,
		KoanfKey:  "auth.services",
	}
}

// initAuthUseKeychainOption initializes the --use-keychain flag for controlling keychain storage
func initAuthUseKeychainOption() {
	cobraParamName := "use-keychain"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(true)
	envVar := "PINGCLI_AUTH_USE_KEYCHAIN"

	options.AuthUseKeychainOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "Use system keychain for storing authentication tokens. If false or keychain is unavailable, tokens will be stored in ~/.pingcli/credentials/. (default true)",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.BOOL,
		KoanfKey:  "auth.useKeychain",
	}
}
