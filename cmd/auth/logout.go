package auth

import (
	"github.com/pingidentity/pingcli/cmd/common"
	"github.com/spf13/cobra"
)

func NewLogoutCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Long:                  "Logout user from the CLI",
		RunE:                  authLogoutRunE,
		Short:                 "Logout user from the CLI",
		Use:                   "logout [flags]",
	}

	return cmd
}

func authLogoutRunE(cmd *cobra.Command, args []string) error {
	// l := logger.Get()
	// l.Debug().Msgf("Auth Logout Subcommand Called.")

	return nil
}
