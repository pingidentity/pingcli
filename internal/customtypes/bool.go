// Copyright © 2026 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"strconv"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/spf13/pflag"
)

var (
	boolErrorPrefix = "custom type bool error"
)

type Bool bool

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*Bool)(nil)

func (b *Bool) Set(val string) error {
	if b == nil {
		return &errs.PingCLIError{Prefix: boolErrorPrefix, Err: ErrCustomTypeNil}
	}

	parsedBool, err := strconv.ParseBool(val)
	if err != nil {
		return &errs.PingCLIError{Prefix: boolErrorPrefix, Err: fmt.Errorf("%w '%s': %w", ErrParseBool, val, err)}
	}
	*b = Bool(parsedBool)

	return nil
}

func (b *Bool) Type() string {
	return "bool"
}

func (b *Bool) String() string {
	if b == nil {
		return "false"
	}

	return strconv.FormatBool(bool(*b))
}

func (b *Bool) Bool() bool {
	if b == nil {
		return false
	}

	return bool(*b)
}
