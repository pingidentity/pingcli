// Copyright © 2025 Ping Identity Corporation

package configuration_services

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitPingFederateServiceOptions() {
	initHTTPSHostOption()
	initAdminAPIPathOption()
	initXBypassExternalValidationHeaderOption()
	initCACertificatePemFilesOption()
	initInsecureTrustAllTLSOption()
	initUsernameOption()
	initPasswordOption()
	initAccessTokenOption()
	initClientIDOption()
	initClientSecretOption()
	initTokenURLOption()
	initScopesOption()
	initPingFederateAuthenticationTypeOption()
}

func initHTTPSHostOption() {
	cobraParamName := "pingfederate-https-host"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGFEDERATE_HTTPS_HOST"

	options.PingFederateHTTPSHostOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: "The PingFederate HTTPS host used to communicate with PingFederate's admin API." +
				"\nExample: 'https://pingfederate-admin.bxretail.org'",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_STRING,
		ViperKey:  "service.pingfederate.httpsHost",
	}
}

func initAdminAPIPathOption() {
	cobraParamName := "pingfederate-admin-api-path"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("/pf-admin-api/v1")
	envVar := "PINGCLI_PINGFEDERATE_ADMIN_API_PATH"

	options.PingFederateAdminAPIPathOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: "The PingFederate API URL path used to communicate with PingFederate's admin API. " +
				"(default /pf-admin-api/v1)",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_STRING,
		ViperKey:  "service.pingfederate.adminAPIPath",
	}
}

func initXBypassExternalValidationHeaderOption() {
	cobraParamName := "pingfederate-x-bypass-external-validation-header"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)
	envVar := "PINGCLI_PINGFEDERATE_X_BYPASS_EXTERNAL_VALIDATION_HEADER"

	options.PingFederateXBypassExternalValidationHeaderOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: "Bypass connection tests when configuring PingFederate (the X-BypassExternalValidation header " +
				"when using PingFederate's admin API). " +
				"(default false)",
			Value:       cobraValue,
			NoOptDefVal: "true", // Make this flag a boolean flag
		},
		Sensitive: false,
		Type:      options.ENUM_BOOL,
		ViperKey:  "service.pingfederate.xBypassExternalValidationHeader",
	}
}

func initCACertificatePemFilesOption() {
	cobraParamName := "pingfederate-ca-certificate-pem-files"
	cobraValue := new(customtypes.StringSlice)
	defaultValue := customtypes.StringSlice{}
	envVar := "PINGCLI_PINGFEDERATE_CA_CERTIFICATE_PEM_FILES"

	options.PingFederateCACertificatePemFilesOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: "Relative or full paths to PEM-encoded certificate files to be trusted as root CAs when " +
				"connecting to the PingFederate server over HTTPS. " +
				"(default [])" +
				"\nAccepts a comma-separated string to delimit multiple PEM files.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_STRING_SLICE,
		ViperKey:  "service.pingfederate.caCertificatePemFiles",
	}
}

func initInsecureTrustAllTLSOption() {
	cobraParamName := "pingfederate-insecure-trust-all-tls"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)
	envVar := "PINGCLI_PINGFEDERATE_INSECURE_TRUST_ALL_TLS"

	options.PingFederateInsecureTrustAllTLSOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: "Trust any certificate when connecting to the PingFederate server admin API. " +
				"(default false)" +
				"\nThis is insecure and should not be enabled outside of testing.",
			Value:       cobraValue,
			NoOptDefVal: "true", // Make this flag a boolean flag
		},
		Sensitive: false,
		Type:      options.ENUM_BOOL,
		ViperKey:  "service.pingfederate.insecureTrustAllTLS",
	}
}

func initUsernameOption() {
	cobraParamName := "pingfederate-username"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGFEDERATE_USERNAME"

	options.PingFederateBasicAuthUsernameOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: "The PingFederate username used to authenticate to the PingFederate admin API when using basic " +
				"authentication." +
				"\nExample: 'administrator'",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_STRING,
		ViperKey:  "service.pingfederate.authentication.basicAuth.username",
	}
}

func initPasswordOption() {
	cobraParamName := "pingfederate-password"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGFEDERATE_PASSWORD"

	options.PingFederateBasicAuthPasswordOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: "The PingFederate password used to authenticate to the PingFederate admin API when using basic " +
				"authentication.",
			Value: cobraValue,
		},
		Sensitive: true,
		Type:      options.ENUM_STRING,
		ViperKey:  "service.pingfederate.authentication.basicAuth.password",
	}
}

func initAccessTokenOption() {
	cobraParamName := "pingfederate-access-token"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGFEDERATE_ACCESS_TOKEN"

	options.PingFederateAccessTokenAuthAccessTokenOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: "The PingFederate access token used to authenticate to the PingFederate admin API when using a " +
				"custom OAuth 2.0 token method.",
			Value: cobraValue,
		},
		Sensitive: true,
		Type:      options.ENUM_STRING,
		ViperKey:  "service.pingfederate.authentication.accessTokenAuth.accessToken",
	}
}

func initClientIDOption() {
	cobraParamName := "pingfederate-client-id"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGFEDERATE_CLIENT_ID"

	options.PingFederateClientCredentialsAuthClientIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: "The PingFederate OAuth client ID used to authenticate to the PingFederate admin API when using " +
				"the OAuth 2.0 client credentials grant type.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_STRING,
		ViperKey:  "service.pingfederate.authentication.clientCredentialsAuth.clientID",
	}
}

func initClientSecretOption() {
	cobraParamName := "pingfederate-client-secret"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGFEDERATE_CLIENT_SECRET"

	options.PingFederateClientCredentialsAuthClientSecretOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: "The PingFederate OAuth client secret used to authenticate to the PingFederate admin API when " +
				"using the OAuth 2.0 client credentials grant type.",
			Value: cobraValue,
		},
		Sensitive: true,
		Type:      options.ENUM_STRING,
		ViperKey:  "service.pingfederate.authentication.clientCredentialsAuth.clientSecret",
	}
}

func initTokenURLOption() {
	cobraParamName := "pingfederate-token-url"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGFEDERATE_TOKEN_URL"

	options.PingFederateClientCredentialsAuthTokenURLOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: "The PingFederate OAuth token URL used to authenticate to the PingFederate admin API when using " +
				"the OAuth 2.0 client credentials grant type.",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_STRING,
		ViperKey:  "service.pingfederate.authentication.clientCredentialsAuth.tokenURL",
	}
}

func initScopesOption() {
	cobraParamName := "pingfederate-scopes"
	cobraValue := new(customtypes.StringSlice)
	defaultValue := customtypes.StringSlice{}
	envVar := "PINGCLI_PINGFEDERATE_SCOPES"

	options.PingFederateClientCredentialsAuthScopesOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: "The PingFederate OAuth scopes used to authenticate to the PingFederate admin API when using " +
				"the OAuth 2.0 client credentials grant type. " +
				"(default [])" +
				"\nAccepts a comma-separated string to delimit multiple scopes." +
				"\nExample: 'openid,profile'",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_STRING_SLICE,
		ViperKey:  "service.pingfederate.authentication.clientCredentialsAuth.scopes",
	}
}

func initPingFederateAuthenticationTypeOption() {
	cobraParamName := "pingfederate-authentication-type"
	cobraValue := new(customtypes.PingFederateAuthenticationType)
	defaultValue := customtypes.PingFederateAuthenticationType("")
	envVar := "PINGCLI_PINGFEDERATE_AUTHENTICATION_TYPE"

	options.PingFederateAuthenticationTypeOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: fmt.Sprintf(
				"The authentication type to use when connecting to the PingFederate admin API."+
					"\nOptions are: %s."+
					"\nExample: '%s'",
				strings.Join(customtypes.PingFederateAuthenticationTypeValidValues(), ", "),
				customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
			),
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_PINGFEDERATE_AUTH_TYPE,
		ViperKey:  "service.pingfederate.authentication.type",
	}
}
