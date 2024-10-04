package configuration_services

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitPingfederateServiceOptions() {
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
	initPingfederateAuthenticationTypeOption()
}

func initHTTPSHostOption() {
	cobraParamName := "pingfederate-https-host"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGFEDERATE_HTTPS_HOST"

	options.PingfederateHTTPSHostOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate HTTPS host used to communicate with PingFederate's API.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "service.pingfederate.httpsHost",
	}
}

func initAdminAPIPathOption() {
	cobraParamName := "pingfederate-admin-api-path"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("/pf-admin-api/v1")
	envVar := "PINGCLI_PINGFEDERATE_ADMIN_API_PATH"

	options.PingfederateAdminAPIPathOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate API URL path used to communicate with PingFederate's API.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "/pf-admin-api/v1",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "service.pingfederate.adminAPIPath",
	}
}

func initXBypassExternalValidationHeaderOption() {
	cobraParamName := "pingfederate-x-bypass-external-validation-header"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)
	envVar := "PINGCLI_PINGFEDERATE_X_BYPASS_EXTERNAL_VALIDATION_HEADER"

	options.PingfederateXBypassExternalValidationHeaderOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("Header value in request for PingFederate. PingFederate's connection tests will be bypassed when set to true.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "false",
		},
		Type:     options.ENUM_BOOL,
		ViperKey: "service.pingfederate.xBypassExternalValidationHeader",
	}
}

func initCACertificatePemFilesOption() {
	cobraParamName := "pingfederate-ca-certificate-pem-files"
	cobraValue := new(customtypes.StringSlice)
	defaultValue := customtypes.StringSlice{}
	envVar := "PINGCLI_PINGFEDERATE_CA_CERTIFICATE_PEM_FILES"

	options.PingfederateCACertificatePemFilesOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("Paths to files containing PEM-encoded certificates to be trusted as root CAs when connecting to the PingFederate server over HTTPS. Accepts comma-separated string to delimit multiple PEM files.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "[]",
		},
		Type:     options.ENUM_STRING_SLICE,
		ViperKey: "service.pingfederate.caCertificatePemFiles",
	}
}

func initInsecureTrustAllTLSOption() {
	cobraParamName := "pingfederate-insecure-trust-all-tls"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)
	envVar := "PINGCLI_PINGFEDERATE_INSECURE_TRUST_ALL_TLS"

	options.PingfederateInsecureTrustAllTLSOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("Set to true to trust any certificate when connecting to the PingFederate server. This is insecure and should not be enabled outside of testing.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "false",
		},
		Type:     options.ENUM_BOOL,
		ViperKey: "service.pingfederate.insecureTrustAllTLS",
	}
}

func initUsernameOption() {
	cobraParamName := "pingfederate-username"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGFEDERATE_USERNAME"

	options.PingfederateBasicAuthUsernameOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate username used to authenticate.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "service.pingfederate.authentication.basicAuth.username",
	}
}

func initPasswordOption() {
	cobraParamName := "pingfederate-password"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGFEDERATE_PASSWORD"

	options.PingfederateBasicAuthPasswordOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate password used to authenticate.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "service.pingfederate.authentication.basicAuth.password",
	}
}

func initAccessTokenOption() {
	cobraParamName := "pingfederate-access-token"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGFEDERATE_ACCESS_TOKEN"

	options.PingfederateAccessTokenAuthAccessTokenOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate access token used to authenticate.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "service.pingfederate.authentication.accessTokenAuth.accessToken",
	}
}

func initClientIDOption() {
	cobraParamName := "pingfederate-client-id"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGFEDERATE_CLIENT_ID"

	options.PingfederateClientCredentialsAuthClientIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate OAuth client ID used to authenticate.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "service.pingfederate.authentication.clientCredentialsAuth.clientID",
	}
}

func initClientSecretOption() {
	cobraParamName := "pingfederate-client-secret"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGFEDERATE_CLIENT_SECRET"

	options.PingfederateClientCredentialsAuthClientSecretOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate OAuth client secret used to authenticate.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "service.pingfederate.authentication.clientCredentialsAuth.clientSecret",
	}
}

func initTokenURLOption() {
	cobraParamName := "pingfederate-token-url"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGFEDERATE_TOKEN_URL"

	options.PingfederateClientCredentialsAuthTokenURLOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate OAuth token URL used to authenticate.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "service.pingfederate.authentication.clientCredentialsAuth.tokenURL",
	}
}

func initScopesOption() {
	cobraParamName := "pingfederate-scopes"
	cobraValue := new(customtypes.StringSlice)
	defaultValue := customtypes.StringSlice{}
	envVar := "PINGCLI_PINGFEDERATE_SCOPES"

	options.PingfederateClientCredentialsAuthScopesOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The PingFederate OAuth scopes used to authenticate. Accepts comma-separated string to delimit multiple scopes.  Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "[]",
		},
		Type:     options.ENUM_STRING_SLICE,
		ViperKey: "service.pingfederate.authentication.clientCredentialsAuth.scopes",
	}
}

func initPingfederateAuthenticationTypeOption() {
	cobraParamName := "pingfederate-authentication-type"
	cobraValue := new(customtypes.PingfederateAuthenticationType)
	defaultValue := customtypes.PingfederateAuthenticationType("")
	envVar := "PINGCLI_PINGFEDERATE_AUTHENTICATION_TYPE"

	options.PingfederateAuthenticationTypeOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The authentication type to use with the PingFederate service. Allowed: %s.  Also configurable via environment variable %s", strings.Join(customtypes.PingfederateAuthenticationTypeValidValues(), ", "), envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_PINGFEDERATE_AUTH_TYPE,
		ViperKey: "service.pingfederate.authentication.type",
	}
}
