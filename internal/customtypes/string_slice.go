// Copyright © 2025 Ping Identity Corporation

package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

type StringSlice []string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*StringSlice)(nil)

func (ss *StringSlice) Set(val string) error {
	if ss == nil {
		return fmt.Errorf("failed to set StringSlice value: %s. StringSlice is nil", val)
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
		return false, fmt.Errorf("failed to remove StringSlice value: %s. StringSlice is nil", val)
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

func (ss StringSlice) Type() string {
	return "[]string"
}

func (ss StringSlice) String() string {
	return strings.Join(ss.StringSlice(), ",")
}

func (ss StringSlice) StringSlice() []string {
	if ss == nil {
		return []string{}
	}

	return []string(ss)
}
