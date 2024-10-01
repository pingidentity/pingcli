package common

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ExactArgs(numArgs int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) != numArgs {
			return fmt.Errorf("failed to execute '%s': command accepts %d arg(s), received %d", cmd.CommandPath(), numArgs, len(args))
		}
		return nil
	}
}

func RangeArgs(min, max int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < min || len(args) > max {
			return fmt.Errorf("failed to execute '%s': command accepts %d to %d arg(s), received %d", cmd.CommandPath(), min, max, len(args))
		}
		return nil
	}
}
