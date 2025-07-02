// Copyright © 2025 Ping Identity Corporation

package configuration_request

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitRequestOptions() {
	initDataOption()
	initDataRawOption()
	initHeaderOption()
	initHTTPMethodOption()
	initServiceOption()
	initAccessTokenOption()
	initAccessTokenExpiryOption()
	initFailOption()
}

func initDataOption() {
	cobraParamName := "data"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_REQUEST_DATA"

	options.RequestDataOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: "The file containing data to send in the request. " +
				"\nExample: './data.json'",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_STRING,
		KoanfKey:  "", // No koanf key
	}
}

func initDataRawOption() {
	cobraParamName := "data-raw"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCLI_REQUEST_DATA_RAW"

	options.RequestDataRawOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name: cobraParamName,
			Usage: "The raw data to send in the request. " +
				"\nExample: '{\"name\": \"My environment\"}'",
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_STRING,
		KoanfKey:  "", // No koanf key
	}
}

func initHeaderOption() {
	cobraParamName := "header"
	cobraValue := new(customtypes.HeaderSlice)
	defaultValue := customtypes.HeaderSlice([]customtypes.Header{})

	options.RequestHeaderOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "r",
			Usage: fmt.Sprintf(
				"A custom header to send in the request." +
					"\nExample: --header \"Content-Type: application/vnd.pingidentity.user.import+json\"",
			),
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_HEADER,
		KoanfKey:  "", // No koanf key
	}
}

func initHTTPMethodOption() {
	cobraParamName := "http-method"
	cobraValue := new(customtypes.HTTPMethod)
	defaultValue := customtypes.HTTPMethod(customtypes.ENUM_HTTP_METHOD_GET)

	options.RequestHTTPMethodOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "m",
			Usage: fmt.Sprintf(
				"The HTTP method to use for the request. (default %s)"+
					"\nOptions are: `%s`."+
					"\nExample: `%s`",
				customtypes.ENUM_HTTP_METHOD_GET,
				strings.Join(customtypes.HTTPMethodValidValues(), ", "),
				customtypes.ENUM_HTTP_METHOD_POST,
			),
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_REQUEST_HTTP_METHOD,
		KoanfKey:  "", // No koanf key
	}
}

func initServiceOption() {
	cobraParamName := "service"
	cobraValue := new(customtypes.RequestService)
	defaultValue := customtypes.RequestService("")
	envVar := "PINGCLI_REQUEST_SERVICE"

	options.RequestServiceOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "s",
			Usage: fmt.Sprintf(
				"The Ping Identity service (configured in the active profile) to send the custom request to."+
					"\nOptions are: `%s`."+
					"\nExample: `%s`",
				strings.Join(customtypes.RequestServiceValidValues(), ", "),
				customtypes.ENUM_REQUEST_SERVICE_PINGONE,
			),
			Value: cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_REQUEST_SERVICE,
		KoanfKey:  "request.service",
	}
}

func initAccessTokenOption() {
	defaultValue := customtypes.String("")

	options.RequestAccessTokenOption = options.Option{
		CobraParamName:  "",  // No cobra param name
		CobraParamValue: nil, // No cobra param value
		DefaultValue:    &defaultValue,
		EnvVar:          "",  // No environment variable
		Flag:            nil, // No flag
		Sensitive:       true,
		Type:            options.ENUM_STRING,
		KoanfKey:        "request.accessToken",
	}
}

func initAccessTokenExpiryOption() {
	defaultValue := customtypes.Int(0)

	options.RequestAccessTokenExpiryOption = options.Option{
		CobraParamName:  "",  // No cobra param name
		CobraParamValue: nil, // No cobra param value
		DefaultValue:    &defaultValue,
		EnvVar:          "",  // No environment variable
		Flag:            nil, // No flag
		Sensitive:       false,
		Type:            options.ENUM_INT,
		KoanfKey:        "request.accessTokenExpiry",
	}
}

func initFailOption() {
	cobraParamName := "fail"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)

	options.RequestFailOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		Flag: &pflag.Flag{
			Name:        cobraParamName,
			NoOptDefVal: "true",
			Shorthand:   "f",
			Usage:       "Return non-zero exit code when HTTP request returns a failure status code.",
			Value:       cobraValue,
		},
		Sensitive: false,
		Type:      options.ENUM_BOOL,
		KoanfKey:  "request.fail",
	}
}
