// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/spf13/pflag"
)

const (
	ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC              string = "basicAuth"
	ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_ACCESS_TOKEN       string = "accessTokenAuth"
	ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS string = "clientCredentialsAuth"
)

var (
	pingFederateAuthTypeErrorPrefix = "custom type pingfederate authentication type error"
	ErrUnrecognizedPingFederateAuth = errors.New("unrecognized pingfederate authentication type")
)

type PingFederateAuthenticationType string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*PingFederateAuthenticationType)(nil)

// Implement pflag.Value interface for custom type in cobra MultiService parameter
func (pat *PingFederateAuthenticationType) Set(authType string) error {
	if pat == nil {
		return &errs.PingCLIError{Prefix: pingFederateAuthTypeErrorPrefix, Err: ErrCustomTypeNil}
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
		return &errs.PingCLIError{Prefix: pingFederateAuthTypeErrorPrefix, Err: fmt.Errorf("%w: '%s'. Must be one of: %s", ErrUnrecognizedPingFederateAuth, authType, strings.Join(PingFederateAuthenticationTypeValidValues(), ", "))}
	}

	return nil
}

func (pat *PingFederateAuthenticationType) Type() string {
	return "string"
}

func (pat *PingFederateAuthenticationType) String() string {
	if pat == nil {
		return ""
	}

	return string(*pat)
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
