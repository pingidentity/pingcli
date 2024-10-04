package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

const (
	ENUM_REQUEST_SERVICE_PINGONE string = "pingone"
)

type RequestService string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*RequestService)(nil)

// Implement pflag.Value interface for custom type in cobra MultiService parameter
func (rs *RequestService) Set(service string) error {
	if rs == nil {
		return fmt.Errorf("failed to set RequestService value: %s. RequestService is nil", service)
	}

	switch {
	case strings.EqualFold(service, ENUM_REQUEST_SERVICE_PINGONE):
		*rs = RequestService(ENUM_REQUEST_SERVICE_PINGONE)
	case strings.EqualFold(service, ""):
		*rs = RequestService("")
	default:
		return fmt.Errorf("unrecognized Request Service: '%s'. Must be one of: %s", service, strings.Join(RequestServiceValidValues(), ", "))
	}
	return nil
}

func (rs RequestService) Type() string {
	return "string"
}

func (rs RequestService) String() string {
	return string(rs)
}

func RequestServiceValidValues() []string {
	allServices := []string{
		ENUM_REQUEST_SERVICE_PINGONE,
	}

	slices.Sort(allServices)

	return allServices
}
