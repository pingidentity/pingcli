package customtypes

import (
	"fmt"
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
		*ss = StringSlice([]string{})
	} else {
		valSs := strings.Split(val, ",")
		*ss = StringSlice(valSs)
	}

	return nil
}

func (ss StringSlice) Type() string {
	return "[]string"
}

func (ss StringSlice) String() string {
	if ss == nil {
		return ""
	}

	return strings.Join(ss, ",")
}

func (ss StringSlice) StringSlice() []string {
	if ss == nil {
		return []string{}
	}

	return []string(ss)
}
