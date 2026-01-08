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
	initPingOneAuthenticationAPIEnvironmentIDOption()
	initPingOneAuthenticationAuthorizationCodeClientIDOption()
	initPingOneAuthenticationAuthorizationCodeRedirectURIPathOption()
	initPingOneAuthenticationAuthorizationCodeRedirectURIPortOption()
	initPingOneAuthenticationClientCredentialsClientIDOption()
	initPingOneAuthenticationClientCredentialsClientSecretOption()
	initPingOneAuthenticationDeviceCodeClientIDOption()
	initPingOneAuthenticationTypeOption()
	initPingOneAuthenticationWorkerClientIDOption()
	initPingOneAuthenticationWorkerClientSecretOption()
	initPingOneAuthenticationWorkerEnvironmentIDOption()
	initPingOneRegionCodeOption()
}

func initPingOneAuthenticationAPIEnvironmentIDOption() {
	cobraParamName := "pingone-environment-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PINGCLI_PINGONE_ENVIRONMENT_ID"

	options.PingOneAuthenticationAPIEnvironmentIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The ID of the PingOne environment to use for authentication (used by all auth types).",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.UUID,
		KoanfKey:  "service.pingOne.authentication.environmentID",
	}
}

func initPingOneAuthenticationAuthorizationCodeClientIDOption() {
	cobraParamName := "pingone-authorization-code-client-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PINGCLI_PINGONE_AUTHORIZATION_CODE_CLIENT_ID"

	options.PingOneAuthenticationAuthorizationCodeClientIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The authorization code client ID used to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.UUID,
		KoanfKey:  "service.pingOne.authentication.authorizationCode.clientID",
	}
}

func initPingOneAuthenticationAuthorizationCodeRedirectURIPathOption() {
	cobraParamName := "pingone-authorization-code-redirect-uri-path"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGONE_AUTHORIZATION_CODE_REDIRECT_URI_PATH"

	options.PingOneAuthenticationAuthorizationCodeRedirectURIPathOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The redirect URI path to use when using the authorization code authorization grant type to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.STRING,
		KoanfKey:  "service.pingOne.authentication.authorizationCode.redirectURIPath",
	}
}

func initPingOneAuthenticationAuthorizationCodeRedirectURIPortOption() {
	cobraParamName := "pingone-authorization-code-redirect-uri-port"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGONE_AUTHORIZATION_CODE_REDIRECT_URI_PORT"

	options.PingOneAuthenticationAuthorizationCodeRedirectURIPortOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:  cobraParamName,
			Usage: "The redirect URI port to use when using the authorization code authorization grant type to authenticate to the PingOne management API.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.STRING,
		KoanfKey:  "service.pingOne.authentication.authorizationCode.redirectURIPort",
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
				"The authorization grant type to use to authenticate to the PingOne management API. (default %s)"+
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
			Usage: "DEPRECATED: Use --pingone-client-credentials-client-id instead. The worker client ID used to authenticate to the PingOne management API.",
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
			Usage: "DEPRECATED: Use --pingone-client-credentials-client-secret instead. The worker client secret used to authenticate to the PingOne management API.",
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
			Usage: "DEPRECATED: Use --pingone-environment-id instead. The ID of the PingOne environment that contains the worker client used to authenticate to " +
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
