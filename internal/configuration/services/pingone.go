// Copyright © 2025 Ping Identity Corporation

package configuration_services

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitPingOneServiceOptions() {
	initPingOneAuthenticationAuthCodeClientIDOption()
	initPingOneAuthenticationAuthCodeEnvironmentIDOption()
	initPingOneAuthenticationAuthCodeRedirectURIPathOption()
	initPingOneAuthenticationAuthCodeRedirectURIPortOption()
	initPingOneAuthenticationAuthCodeScopesOption()
	initPingOneAuthenticationClientCredentialsClientIDOption()
	initPingOneAuthenticationClientCredentialsClientSecretOption()
	initPingOneAuthenticationClientCredentialsEnvironmentIDOption()
	initPingOneAuthenticationClientCredentialsScopesOption()
	initPingOneAuthenticationDeviceCodeClientIDOption()
	initPingOneAuthenticationDeviceCodeEnvironmentIDOption()
	initPingOneAuthenticationDeviceCodeScopesOption()
	initPingOneAuthenticationTypeOption()
	initPingOneAuthenticationWorkerClientIDOption()
	initPingOneAuthenticationWorkerClientSecretOption()
	initPingOneAuthenticationWorkerEnvironmentIDOption()
	initPingOneRegionCodeOption()
}

func initPingOneAuthenticationAuthCodeClientIDOption() {
	cobraParamName := "pingone-oidc-auth-code-client-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PINGCLI_PINGONE_OIDC_AUTH_CODE_CLIENT_ID"

	options.PingOneAuthenticationAuthCodeClientIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The auth code client ID used to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.UUID,
		KoanfKey:  "service.pingOne.authentication.authCode.clientID",
	}
}

func initPingOneAuthenticationAuthCodeEnvironmentIDOption() {
	cobraParamName := "pingone-oidc-auth-code-environment-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PINGCLI_PINGONE_OIDC_AUTH_CODE_ENVIRONMENT_ID"

	options.PingOneAuthenticationAuthCodeEnvironmentIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The ID of the PingOne environment that contains the auth code client used to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.UUID,
		KoanfKey:  "service.pingOne.authentication.authCode.environmentID",
	}
}

func initPingOneAuthenticationAuthCodeRedirectURIPathOption() {
	cobraParamName := "pingone-oidc-auth-code-redirect-uri-path"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGONE_OIDC_AUTH_CODE_REDIRECT_URI_PATH"

	options.PingOneAuthenticationAuthCodeRedirectURIPathOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The redirect URI path to use when using the auth code authentication type to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.STRING,
		KoanfKey:  "service.pingOne.authentication.authCode.redirectURIPath",
	}
}

func initPingOneAuthenticationAuthCodeRedirectURIPortOption() {
	cobraParamName := "pingone-oidc-auth-code-redirect-uri-port"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGONE_OIDC_AUTH_CODE_REDIRECT_URI_PORT"

	options.PingOneAuthenticationAuthCodeRedirectURIPortOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The redirect URI port to use when using the auth code authentication type to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.STRING,
		KoanfKey:  "service.pingOne.authentication.authCode.redirectURIPort",
	}
}

func initPingOneAuthenticationAuthCodeScopesOption() {
	cobraParamName := "pingone-oidc-auth-code-scopes"
	cobraValue := new(customtypes.StringSlice)
	defaultValue := customtypes.StringSlice{}
	envVar := "PINGCLI_PINGONE_OIDC_AUTH_CODE_SCOPES"

	options.PingOneAuthenticationAuthCodeScopesOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The auth code scope(s) used to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.STRING_SLICE,
		KoanfKey:  "service.pingOne.authentication.authCode.scopes",
	}
}

func initPingOneAuthenticationClientCredentialsClientIDOption() {
	cobraParamName := "pingone-client-credentials-client-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PINGCLI_PINGONE_CLIENT_CREDENTIALS_CLIENT_ID"

	options.PingOneAuthenticationClientCredentialsClientIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The client credentials client ID used to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.UUID,
		KoanfKey:  "service.pingOne.authentication.clientCredentials.clientID",
	}
}

func initPingOneAuthenticationClientCredentialsClientSecretOption() {
	cobraParamName := "pingone-client-credentials-client-secret"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGONE_CLIENT_CREDENTIALS_CLIENT_SECRET"

	options.PingOneAuthenticationClientCredentialsClientSecretOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The client credentials client secret used to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: true,
		Type:      options.STRING,
		KoanfKey:  "service.pingOne.authentication.clientCredentials.clientSecret",
	}
}

func initPingOneAuthenticationClientCredentialsEnvironmentIDOption() {
	cobraParamName := "pingone-client-credentials-environment-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PINGCLI_PINGONE_CLIENT_CREDENTIALS_ENVIRONMENT_ID"

	options.PingOneAuthenticationClientCredentialsEnvironmentIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The ID of the PingOne environment that contains the client credentials client used to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.UUID,
		KoanfKey:  "service.pingOne.authentication.clientCredentials.environmentID",
	}
}

func initPingOneAuthenticationClientCredentialsScopesOption() {
	cobraParamName := "pingone-client-credentials-scopes"
	cobraValue := new(customtypes.StringSlice)
	defaultValue := customtypes.StringSlice{}
	envVar := "PINGCLI_PINGONE_CLIENT_CREDENTIALS_SCOPES"

	options.PingOneAuthenticationClientCredentialsScopesOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The scopes to request for the client credentials used to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.STRING_SLICE,
		KoanfKey:  "service.pingOne.authentication.clientCredentials.scopes",
	}
}

func initPingOneAuthenticationDeviceCodeClientIDOption() {
	cobraParamName := "pingone-device-code-client-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PINGCLI_PINGONE_DEVICE_CODE_CLIENT_ID"

	options.PingOneAuthenticationDeviceCodeClientIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The device code client ID used to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.UUID,
		KoanfKey:  "service.pingOne.authentication.deviceCode.clientID",
	}
}

func initPingOneAuthenticationDeviceCodeEnvironmentIDOption() {
	cobraParamName := "pingone-device-code-environment-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PINGCLI_PINGONE_DEVICE_CODE_ENVIRONMENT_ID"

	options.PingOneAuthenticationDeviceCodeEnvironmentIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The ID of the PingOne environment that contains the device code client used to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.UUID,
		KoanfKey:  "service.pingOne.authentication.deviceCode.environmentID",
	}
}

func initPingOneAuthenticationDeviceCodeScopesOption() {
	cobraParamName := "pingone-device-code-scopes"
	cobraValue := new(customtypes.StringSlice)
	defaultValue := customtypes.StringSlice{}
	envVar := "PINGCLI_PINGONE_DEVICE_CODE_SCOPES"

	options.PingOneAuthenticationDeviceCodeScopesOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The device code scope(s) used to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.STRING_SLICE,
		KoanfKey:  "service.pingOne.authentication.deviceCode.scopes",
	}
}

func initPingOneAuthenticationTypeOption() {
	cobraParamName := "pingone-authentication-type"
	cobraValue := new(customtypes.PingOneAuthenticationType)
	defaultValue := customtypes.PingOneAuthenticationType("")
	envVar := "PINGCLI_PINGONE_AUTHENTICATION_TYPE"

	options.PingOneAuthenticationTypeOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: fmt.Sprintf(
				"The authentication type to use to authenticate to the PingOne management API. (default %s)"+
					"\nOptions are: %s.",
				customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER,
				strings.Join(customtypes.PingOneAuthenticationTypeValidValues(), ", "),
			),
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.PINGONE_AUTH_TYPE,
		KoanfKey:  "service.pingOne.authentication.type",
	}
}

func initPingOneAuthenticationWorkerClientIDOption() {
	cobraParamName := "pingone-worker-client-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PINGCLI_PINGONE_WORKER_CLIENT_ID"

	options.PingOneAuthenticationWorkerClientIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The worker client ID used to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.UUID,
		KoanfKey:  "service.pingOne.authentication.worker.clientID",
	}
}

func initPingOneAuthenticationWorkerClientSecretOption() {
	cobraParamName := "pingone-worker-client-secret"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGONE_WORKER_CLIENT_SECRET"

	options.PingOneAuthenticationWorkerClientSecretOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The worker client secret used to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: true,
		Type:      options.STRING,
		KoanfKey:  "service.pingOne.authentication.worker.clientSecret",
	}
}

func initPingOneAuthenticationWorkerEnvironmentIDOption() {
	cobraParamName := "pingone-worker-environment-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PINGCLI_PINGONE_WORKER_ENVIRONMENT_ID"

	options.PingOneAuthenticationWorkerEnvironmentIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: "The ID of the PingOne environment that contains the worker client used to authenticate to " +
				"the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.UUID,
		KoanfKey:  "service.pingOne.authentication.worker.environmentID",
	}
}

func initPingOneRegionCodeOption() {
	cobraParamName := "pingone-region-code"
	cobraValue := new(customtypes.String)
	envVar := "PINGCLI_PINGONE_REGION_CODE"

	options.PingOneRegionCodeOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: fmt.Sprintf(
				"The region code of the PingOne tenant."+
					"\nOptions are: %s."+
					"\nExample: '%s'",
				strings.Join(customtypes.PingOneRegionCodeValidValues(), ", "),
				customtypes.ENUM_PINGONE_REGION_CODE_NA,
			),
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.PINGONE_REGION_CODE,
		KoanfKey:  "service.pingOne.regionCode",
	}
}
