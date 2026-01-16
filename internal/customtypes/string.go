// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/spf13/pflag"
)

var (
	stringErrorPrefix = "custom type string error"
)

type String string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*String)(nil)

func (s *String) Set(val string) error {
	if s == nil {
		return &errs.PingCLIError{Prefix: stringErrorPrefix, Err: ErrCustomTypeNil}
	}

	*s = String(val)

	return nil
}

func (s *String) Type() string {
	return "string"
}

func (s *String) String() string {
	if s == nil {
		return ""
	}

	return string(*s)
}

func StringPtr(val string) *String {
	s := String(val)
	return &s
}
