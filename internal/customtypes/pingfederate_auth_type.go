package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

const (
	ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC              string = "basicAuth"
	ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_ACCESS_TOKEN       string = "accessTokenAuth"
	ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS string = "clientCredentialsAuth"
)

type PingfederateAuthenticationType string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*PingfederateAuthenticationType)(nil)

// Implement pflag.Value interface for custom type in cobra MultiService parameter
func (pat *PingfederateAuthenticationType) Set(authType string) error {
	if pat == nil {
		return fmt.Errorf("failed to set Pingfederate Authentication Type value: %s. Pingfederate Authentication Type is nil", authType)
	}

	switch {
	case strings.EqualFold(authType, ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC):
		*pat = PingfederateAuthenticationType(ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC)
	case strings.EqualFold(authType, ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_ACCESS_TOKEN):
		*pat = PingfederateAuthenticationType(ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_ACCESS_TOKEN)
	case strings.EqualFold(authType, ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS):
		*pat = PingfederateAuthenticationType(ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS)
	case strings.EqualFold(authType, ""):
		*pat = PingfederateAuthenticationType("")
	default:
		return fmt.Errorf("unrecognized Pingfederate Authentication Type: '%s'. Must be one of: %s", authType, strings.Join(PingfederateAuthenticationTypeValidValues(), ", "))
	}
	return nil
}

func (pat PingfederateAuthenticationType) Type() string {
	return "string"
}

func (pat PingfederateAuthenticationType) String() string {
	return string(pat)
}

func PingfederateAuthenticationTypeValidValues() []string {
	types := []string{
		ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
		ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_ACCESS_TOKEN,
		ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS,
	}

	slices.Sort(types)

	return types
}
