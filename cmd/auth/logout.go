// Copyright © 2026 Ping Identity Corporation

package auth

import (
	"github.com/pingidentity/pingcli/cmd/common"
	auth_internal "github.com/pingidentity/pingcli/internal/commands/auth"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

// NewLogoutCommand creates a new logout command that clears stored credentials
func NewLogoutCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Long:                  "Logout user from the CLI by clearing stored credentials. Credentials are cleared from both keychain and file storage. By default, uses the authentication method configured in the active profile. You can specify a different authentication method using the grant type flags.",
		RunE:                  authLogoutRunE,
		Short:                 "Logout user from the CLI",
		Use:                   "logout [flags]",
	}

	// Add the same grant type flags as login command
	cmd.Flags().AddFlag(options.AuthMethodAuthorizationCodeOption.Flag)
	cmd.Flags().AddFlag(options.AuthMethodClientCredentialsOption.Flag)
	cmd.Flags().AddFlag(options.AuthMethodDeviceCodeOption.Flag)

	// These flags are mutually exclusive - only one can be specified
	cmd.MarkFlagsMutuallyExclusive(
		options.AuthMethodAuthorizationCodeOption.Flag.Name,
		options.AuthMethodClientCredentialsOption.Flag.Name,
		options.AuthMethodDeviceCodeOption.Flag.Name,
	)

	return cmd
}

func authLogoutRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config logout Subcommand Called.")

	if err := auth_internal.AuthLogoutRunE(cmd, args); err != nil {
		return &errs.PingCLIError{Prefix: "", Err: err}
	}

	return nil
}
