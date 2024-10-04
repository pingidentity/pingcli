package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

const (
	ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER string = "worker"
)

type PingoneAuthenticationType string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*PingoneAuthenticationType)(nil)

// Implement pflag.Value interface for custom type in cobra MultiService parameter
func (pat *PingoneAuthenticationType) Set(authType string) error {
	if pat == nil {
		return fmt.Errorf("failed to set Pingone Authentication Type value: %s. Pingone Authentication Type is nil", authType)
	}

	switch {
	case strings.EqualFold(authType, ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER):
		*pat = PingoneAuthenticationType(ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER)
	case strings.EqualFold(authType, ""):
		*pat = PingoneAuthenticationType("")
	default:
		return fmt.Errorf("unrecognized Pingone Authentication Type: '%s'. Must be one of: %s", authType, strings.Join(PingoneAuthenticationTypeValidValues(), ", "))
	}
	return nil
}

func (pat PingoneAuthenticationType) Type() string {
	return "string"
}

func (pat PingoneAuthenticationType) String() string {
	return string(pat)
}

func PingoneAuthenticationTypeValidValues() []string {
	types := []string{
		ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER,
	}

	slices.Sort(types)

	return types
}
