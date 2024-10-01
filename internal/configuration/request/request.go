package configuration_request

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingctl/internal/configuration/options"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/spf13/pflag"
)

func InitRequestOptions() {
	initDataOption()
	initHTTPMethodOption()
	initServiceOption()
	initAccessTokenOption()
	initAccessTokenExpiryOption()

}

func initDataOption() {
	cobraParamName := "data"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")
	envVar := "PINGCTL_REQUEST_DATA"

	options.RequestDataOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    fmt.Sprintf("The data to send in the request. Use prefix '@' to specify data filepath instead of raw data. Also configurable via environment variable %s.", envVar),
			Value:    cobraValue,
			DefValue: "",
		},
		Type:     options.ENUM_STRING,
		ViperKey: "", // No viper key
	}
}

func initHTTPMethodOption() {
	cobraParamName := "http-method"
	cobraValue := new(customtypes.HTTPMethod)
	defaultValue := customtypes.HTTPMethod(customtypes.ENUM_HTTP_METHOD_GET)
	envVar := "PINGCTL_REQUEST_HTTP_METHOD"

	options.RequestHTTPMethodOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "m",
			Usage:     fmt.Sprintf("The HTTP method to use for the request. Allowed: %s. Also configurable via environment variable %s.", strings.Join(customtypes.HTTPMethodValidValues(), ", "), envVar),
			Value:     cobraValue,
			DefValue:  customtypes.ENUM_HTTP_METHOD_GET,
		},
		Type:     options.ENUM_REQUEST_HTTP_METHOD,
		ViperKey: "request.httpMethod",
	}
}

func initServiceOption() {
	cobraParamName := "service"
	cobraValue := new(customtypes.RequestService)
	defaultValue := customtypes.RequestService("")
	envVar := "PINGCTL_REQUEST_SERVICE"

	options.RequestServiceOption = options.Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          envVar,
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "s",
			Usage:     fmt.Sprintf("The service to send a custom request. Allowed: %s. Also configurable via environment variable %s.", strings.Join(customtypes.RequestServiceValidValues(), ", "), envVar),
			Value:     cobraValue,
			DefValue:  "",
		},
		Type:     options.ENUM_REQUEST_SERVICE,
		ViperKey: "request.service",
	}
}

func initAccessTokenOption() {
	defaultValue := customtypes.String("")

	options.RequestAccessTokenOption = options.Option{
		CobraParamName:  "",            // No cobra param name
		CobraParamValue: nil,           // No cobra param value
		DefaultValue:    &defaultValue, // No default value
		EnvVar:          "",            // No environment variable
		Flag:            nil,
		Type:            options.ENUM_STRING,
		ViperKey:        "request.accessToken",
	}
}

func initAccessTokenExpiryOption() {
	defaultValue := customtypes.Int(0)

	options.RequestAccessTokenExpiryOption = options.Option{
		CobraParamName:  "",            // No cobra param name
		CobraParamValue: nil,           // No cobra param value
		DefaultValue:    &defaultValue, // No default value
		EnvVar:          "",            // No environment variable
		Flag:            nil,           // No flag
		Type:            options.ENUM_INT,
		ViperKey:        "request.accessTokenExpiry",
	}
}
