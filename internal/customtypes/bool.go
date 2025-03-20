// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"strconv"

	"github.com/spf13/pflag"
)

type Bool bool

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*Bool)(nil)

func (b *Bool) Set(val string) error {
	if b == nil {
		return fmt.Errorf("failed to set Bool value: %s. Bool is nil", val)
	}

	parsedBool, err := strconv.ParseBool(val)
	if err != nil {
		return err
	}
	*b = Bool(parsedBool)

	return nil
}

func (b Bool) Type() string {
	return "bool"
}

func (b Bool) String() string {
	return strconv.FormatBool(bool(b))
}

func (b Bool) Bool() bool {
	return bool(b)
}
