// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"strconv"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/spf13/pflag"
)

var (
	intErrorPrefix = "custom type int error"
)

type Int int64

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*Int)(nil)

func (i *Int) Set(val string) error {
	if i == nil {
		return &errs.PingCLIError{Prefix: intErrorPrefix, Err: ErrCustomTypeNil}
	}

	parsedInt, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return &errs.PingCLIError{Prefix: intErrorPrefix, Err: fmt.Errorf("%w '%s': %w", ErrParseInt, val, err)}
	}
	*i = Int(parsedInt)

	return nil
}

func (i *Int) Type() string {
	return "int64"
}

func (i *Int) String() string {
	if i == nil {
		return "0"
	}

	return strconv.FormatInt(int64(*i), 10)
}

func (i *Int) Int64() int64 {
	if i == nil {
		return 0
	}

	return int64(*i)
}
