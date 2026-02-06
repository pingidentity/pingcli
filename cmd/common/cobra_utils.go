// Copyright © 2026 Ping Identity Corporation

package common

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/spf13/cobra"
)

var (
	argsErrorPrefix = "failed to execute command"
)

func ExactArgs(numArgs int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) != numArgs {
			return &errs.PingCLIError{Prefix: argsErrorPrefix, Err: fmt.Errorf("%w: command accepts %d arg(s), received %d", ErrExactArgs, numArgs, len(args))}
		}

		return nil
	}
}

func RangeArgs(minArgs, maxArgs int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < minArgs || len(args) > maxArgs {
			return &errs.PingCLIError{Prefix: argsErrorPrefix, Err: fmt.Errorf("%w: command accepts %d to %d arg(s), received %d", ErrRangeArgs, minArgs, maxArgs, len(args))}
		}

		return nil
	}
}
