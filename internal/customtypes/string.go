// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"

	"github.com/spf13/pflag"
)

type String string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*String)(nil)

func (s *String) Set(val string) error {
	if s == nil {
		return fmt.Errorf("failed to set String value: %s. String is nil", val)
	}

	*s = String(val)

	return nil
}

func (s String) Type() string {
	return "string"
}

func (s String) String() string {
	return string(s)
}
