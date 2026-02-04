// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/spf13/pflag"
)

const (
	ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS string = "client_credentials"
	ENUM_PINGONE_AUTHENTICATION_TYPE_AUTHORIZATION_CODE string = "authorization_code"
	ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE        string = "device_code"
	ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER             string = "worker"
)

var (
	pingOneAuthTypeErrorPrefix = "custom type pingone auth type error"
)

type PingOneAuthenticationType string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*PingOneAuthenticationType)(nil)

// Implement pflag.Value interface for custom type in cobra MultiService parameter
func (pat *PingOneAuthenticationType) Set(authType string) error {
	if pat == nil {
		return &errs.PingCLIError{Prefix: pingOneAuthTypeErrorPrefix, Err: ErrCustomTypeNil}
	}

	switch {
	case strings.EqualFold(authType, ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS):
		*pat = PingOneAuthenticationType(ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS)
	case strings.EqualFold(authType, ENUM_PINGONE_AUTHENTICATION_TYPE_AUTHORIZATION_CODE):
		*pat = PingOneAuthenticationType(ENUM_PINGONE_AUTHENTICATION_TYPE_AUTHORIZATION_CODE)
	case strings.EqualFold(authType, ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE):
		*pat = PingOneAuthenticationType(ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE)
	case strings.EqualFold(authType, ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER):
		*pat = PingOneAuthenticationType(ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER)
	case strings.EqualFold(authType, ""):
		*pat = PingOneAuthenticationType("")
	default:
		return &errs.PingCLIError{Prefix: pingOneAuthTypeErrorPrefix, Err: fmt.Errorf("%w: '%s'. Must be one of: %s", ErrUnrecognizedPingOneAuth, authType, strings.Join(PingOneAuthenticationTypeValidValues(), ", "))}
	}

	return nil
}

func (pat *PingOneAuthenticationType) Type() string {
	return "string"
}

func (pat *PingOneAuthenticationType) String() string {
	if pat == nil {
		return ""
	}

	return string(*pat)
}

func PingOneAuthenticationTypeValidValues() []string {
	types := []string{
		ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS,
		ENUM_PINGONE_AUTHENTICATION_TYPE_AUTHORIZATION_CODE,
		ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE,
		ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER,
	}

	slices.Sort(types)

	return types
}
