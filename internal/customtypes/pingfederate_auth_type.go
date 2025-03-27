// Copyright © 2025 Ping Identity Corporation

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

type PingFederateAuthenticationType string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*PingFederateAuthenticationType)(nil)

// Implement pflag.Value interface for custom type in cobra MultiService parameter
func (pat *PingFederateAuthenticationType) Set(authType string) error {
	if pat == nil {
		return fmt.Errorf("failed to set PingFederate Authentication Type value: %s. PingFederate Authentication Type is nil", authType)
	}

	switch {
	case strings.EqualFold(authType, ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC):
		*pat = PingFederateAuthenticationType(ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC)
	case strings.EqualFold(authType, ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_ACCESS_TOKEN):
		*pat = PingFederateAuthenticationType(ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_ACCESS_TOKEN)
	case strings.EqualFold(authType, ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS):
		*pat = PingFederateAuthenticationType(ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS)
	case strings.EqualFold(authType, ""):
		*pat = PingFederateAuthenticationType("")
	default:
		return fmt.Errorf("unrecognized PingFederate Authentication Type: '%s'. Must be one of: %s", authType, strings.Join(PingFederateAuthenticationTypeValidValues(), ", "))
	}

	return nil
}

func (pat PingFederateAuthenticationType) Type() string {
	return "string"
}

func (pat PingFederateAuthenticationType) String() string {
	return string(pat)
}

func PingFederateAuthenticationTypeValidValues() []string {
	types := []string{
		ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
		ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_ACCESS_TOKEN,
		ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS,
	}

	slices.Sort(types)

	return types
}
