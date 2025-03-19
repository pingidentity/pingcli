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
	initPingOneAuthenticationTypeOption()
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
		Type:      options.ENUM_UUID,
		ViperKey:  "service.pingone.authentication.worker.clientID",
	}
}

func initAuthenticationWorkerClientSecretOption() {
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
		Type:      options.ENUM_STRING,
		ViperKey:  "service.pingone.authentication.worker.clientSecret",
	}
}

func initAuthenticationWorkerEnvironmentIDOption() {
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
		Type:      options.ENUM_UUID,
		ViperKey:  "service.pingone.authentication.worker.environmentID",
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
		Type:      options.ENUM_PINGONE_AUTH_TYPE,
		ViperKey:  "service.pingone.authentication.type",
	}
}

func initRegionCodeOption() {
	cobraParamName := "pingone-region-code"
	cobraValue := new(customtypes.PingOneRegionCode)
	defaultValue := customtypes.PingOneRegionCode("")
	envVar := "PINGCLI_PINGONE_REGION_CODE"

	options.PingOneRegionCodeOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
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
		Type:      options.ENUM_PINGONE_REGION_CODE,
		ViperKey:  "service.pingone.regionCode",
	}
}
