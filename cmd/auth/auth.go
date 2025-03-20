// Copyright © 2025 Ping Identity Corporation

package auth

import (
	"github.com/spf13/cobra"
)

func NewAuthCommand() *cobra.Command {
	cmd := &cobra.Command{
		Long:  "Authenticate the CLI with configured Ping connections",
		Short: "Authenticate the CLI with configured Ping connections",
		Use:   "auth",
	}

	cmd.AddCommand(
		NewLoginCommand(),
		NewLogoutCommand(),
	)

	return cmd
}
