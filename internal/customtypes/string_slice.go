// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"slices"
	"strings"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/spf13/pflag"
)

var (
	stringSliceErrorPrefix = "custom type string slice error"
)

type StringSlice []string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*StringSlice)(nil)

func (ss *StringSlice) Set(val string) error {
	if ss == nil {
		return &errs.PingCLIError{Prefix: stringSliceErrorPrefix, Err: ErrCustomTypeNil}
	}

	if val == "" || val == "[]" {
		return nil
	} else {
		valSs := strings.Split(val, ",")
		*ss = append(*ss, valSs...)
	}

	return nil
}

func (ss *StringSlice) Remove(val string) (bool, error) {
	if ss == nil {
		return false, &errs.PingCLIError{Prefix: stringSliceErrorPrefix, Err: ErrCustomTypeNil}
	}

	if val == "" || val == "[]" {
		return false, nil
	}

	for i, v := range *ss {
		if v == val {
			*ss = slices.Delete(*ss, i, i+1)

			return true, nil
		}
	}

	return false, nil
}

func (ss *StringSlice) Type() string {
	return "[]string"
}

func (ss *StringSlice) String() string {
	if ss == nil {
		return ""
	}

	return strings.Join(ss.StringSlice(), ",")
}

func (ss *StringSlice) StringSlice() []string {
	if ss == nil {
		return []string{}
	}

	return []string(*ss)
}
