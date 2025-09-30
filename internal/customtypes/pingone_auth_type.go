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
	ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER string = "worker"
)

var (
	pingOneAuthTypeErrorPrefix = "custom type pingone auth type error"
	ErrUnrecognizedPingOneAuth = errors.New("unrecognized pingone authentication type")
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
		ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER,
	}

	slices.Sort(types)

	return types
}
