// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

const (
	ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS string = "client_credentials"
	ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE          string = "auth_code"
	ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE        string = "device_code"
	ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER             string = "worker"
)

type PingOneAuthenticationType string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*PingOneAuthenticationType)(nil)

// Implement pflag.Value interface for custom type in cobra MultiService parameter
func (pat *PingOneAuthenticationType) Set(authType string) error {
	if pat == nil {
		return fmt.Errorf("failed to set PingOne Authentication Type value: %s. PingOne Authentication Type is nil", authType)
	}

	switch {
	case strings.EqualFold(authType, ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS):
		*pat = PingOneAuthenticationType(ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS)
	case strings.EqualFold(authType, ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE):
		*pat = PingOneAuthenticationType(ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE)
	case strings.EqualFold(authType, ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE):
		*pat = PingOneAuthenticationType(ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE)
	case strings.EqualFold(authType, ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER):
		*pat = PingOneAuthenticationType(ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER)
	case strings.EqualFold(authType, ""):
		*pat = PingOneAuthenticationType("")
	default:
		return fmt.Errorf("unrecognized PingOne Authentication Type: '%s'. Must be one of: %s", authType, strings.Join(PingOneAuthenticationTypeValidValues(), ", "))
	}

	return nil
}

func (pat PingOneAuthenticationType) Type() string {
	return "string"
}

func (pat PingOneAuthenticationType) String() string {
	return string(pat)
}

func PingOneAuthenticationTypeValidValues() []string {
	types := []string{
		ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS,
		ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE,
		ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE,
		ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER,
	}

	slices.Sort(types)

	return types
}
