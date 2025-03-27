// Copyright © 2025 Ping Identity Corporation

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

func RangeArgs(minArgs, maxArgs int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < minArgs || len(args) > maxArgs {
			return fmt.Errorf("failed to execute '%s': command accepts %d to %d arg(s), received %d", cmd.CommandPath(), minArgs, maxArgs, len(args))
		}

		return nil
	}
}
