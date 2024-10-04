package configuration_services

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitPingoneServiceOptions() {
	initPingoneAuthenticationTypeOption()
	initAuthenticationWorkerClientIDOption()
	initAuthenticationWorkerClientSecretOption()
	initAuthenticationWorkerEnvironmentIDOption()
	initRegionCodeOption()

}

func initAuthenticationWorkerClientIDOption() {
	cobraParamName := "pingone-worker-client-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PINGCLI_PINGONE_WORKER_CLIENT_ID"

	options.PingoneAuthenticationWorkerClientIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The Pingone worker client ID used to authenticate. Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_UUID,
		ViperKey: "service.pingone.authentication.worker.clientID",
	}
}

func initAuthenticationWorkerClientSecretOption() {
	cobraParamName := "pingone-worker-client-secret"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_PINGONE_WORKER_CLIENT_SECRET"

	options.PingoneAuthenticationWorkerClientSecretOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The Pingone worker client secret used to authenticate. Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "service.pingone.authentication.worker.clientSecret",
	}
}

func initAuthenticationWorkerEnvironmentIDOption() {
	cobraParamName := "pingone-worker-environment-id"
	cobraValue := new(customtypes.UUID)
	defaultValue := customtypes.UUID("")
	envVar := "PINGCLI_PINGONE_WORKER_ENVIRONMENT_ID"

	options.PingoneAuthenticationWorkerEnvironmentIDOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The ID of the Pingone environment that contains the worker client used to authenticate. Also configurable via environment variable %s", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_UUID,
		ViperKey: "service.pingone.authentication.worker.environmentID",
	}
}

func initPingoneAuthenticationTypeOption() {
	cobraParamName := "pingone-authentication-type"
	cobraValue := new(customtypes.PingoneAuthenticationType)
	defaultValue := customtypes.PingoneAuthenticationType("")
	envVar := "PINGCLI_PINGONE_AUTHENTICATION_TYPE"

	options.PingoneAuthenticationTypeOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The authentication type to use with the Pingone service. Allowed: %s. Also configurable via environment variable %s", strings.Join(customtypes.PingoneAuthenticationTypeValidValues(), ", "), envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_PINGONE_AUTH_TYPE,
		ViperKey: "service.pingone.authentication.type",
	}
}

func initRegionCodeOption() {
	cobraParamName := "pingone-region-code"
	cobraValue := new(customtypes.PingoneRegionCode)
	defaultValue := customtypes.PingoneRegionCode("")
	envVar := "PINGCLI_PINGONE_REGION_CODE"

	options.PingoneRegionCodeOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The region code of the Pingone service. Allowed: %s. Also configurable via environment variable %s", strings.Join(customtypes.PingoneRegionCodeValidValues(), ", "), envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_PINGONE_REGION_CODE,
		ViperKey: "service.pingone.regionCode",
	}
}
