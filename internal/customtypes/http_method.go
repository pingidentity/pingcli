// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

const (
	ENUM_HTTP_METHOD_GET    string = "GET"
	ENUM_HTTP_METHOD_POST   string = "POST"
	ENUM_HTTP_METHOD_PUT    string = "PUT"
	ENUM_HTTP_METHOD_DELETE string = "DELETE"
	ENUM_HTTP_METHOD_PATCH  string = "PATCH"
)

type HTTPMethod string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*HTTPMethod)(nil)

// Implement pflag.Value interface for custom type in cobra pingcli-output parameter

func (hm *HTTPMethod) Set(httpMethod string) error {
	if hm == nil {
		return fmt.Errorf("failed to set HTTP Method value: %s. HTTPMethod is nil", httpMethod)
	}

	switch {
	case strings.EqualFold(httpMethod, ENUM_HTTP_METHOD_GET):
		*hm = HTTPMethod(ENUM_HTTP_METHOD_GET)
	case strings.EqualFold(httpMethod, ENUM_HTTP_METHOD_POST):
		*hm = HTTPMethod(ENUM_HTTP_METHOD_POST)
	case strings.EqualFold(httpMethod, ENUM_HTTP_METHOD_PUT):
		*hm = HTTPMethod(ENUM_HTTP_METHOD_PUT)
	case strings.EqualFold(httpMethod, ENUM_HTTP_METHOD_DELETE):
		*hm = HTTPMethod(ENUM_HTTP_METHOD_DELETE)
	case strings.EqualFold(httpMethod, ENUM_HTTP_METHOD_PATCH):
		*hm = HTTPMethod(ENUM_HTTP_METHOD_PATCH)
	case strings.EqualFold(httpMethod, ""):
		*hm = HTTPMethod("")
	default:
		return fmt.Errorf("unrecognized HTTP Method: '%s'. Must be one of: %s", httpMethod, strings.Join(HTTPMethodValidValues(), ", "))
	}

	return nil
}

func (hm HTTPMethod) Type() string {
	return "string"
}

func (hm HTTPMethod) String() string {
	return string(hm)
}

func HTTPMethodValidValues() []string {
	methods := []string{
		ENUM_HTTP_METHOD_GET,
		ENUM_HTTP_METHOD_POST,
		ENUM_HTTP_METHOD_PUT,
		ENUM_HTTP_METHOD_DELETE,
		ENUM_HTTP_METHOD_PATCH,
	}

	slices.Sort(methods)

	return methods
}
