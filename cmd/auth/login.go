package auth

import (
	"github.com/pingidentity/pingcli/cmd/common"
	"github.com/spf13/cobra"
)

func NewLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Long:                  "Login user to the CLI",
		RunE:                  authLoginRunE,
		Short:                 "Login user to the CLI",
		Use:                   "login [flags]",
	}

	return cmd
}

func authLoginRunE(cmd *cobra.Command, args []string) error {
	// l := logger.Get()
	// l.Debug().Msgf("Auth Login Subcommand Called.")

	return nil
}
