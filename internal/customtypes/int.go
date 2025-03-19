// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"strconv"

	"github.com/spf13/pflag"
)

type Int int64

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*Int)(nil)

func (i *Int) Set(val string) error {
	if i == nil {
		return fmt.Errorf("failed to set Int value: %s. Int is nil", val)
	}

	parsedInt, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return err
	}
	*i = Int(parsedInt)

	return nil
}

func (i Int) Type() string {
	return "int64"
}

func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int) Int64() int64 {
	return int64(i)
}
