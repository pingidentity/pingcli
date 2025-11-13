// Copyright © 2025 Ping Identity Corporation

package auth

import (
	"fmt"

	"github.com/pingidentity/pingcli/cmd/common"
	auth_internal "github.com/pingidentity/pingcli/internal/commands/auth"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

var (
	// ErrUnknownAuthMethod is returned when an unknown authentication method is specified
	ErrUnknownAuthMethod = fmt.Errorf("unknown authentication method")
)

// NewLoginCommand creates a new login command that authenticates users using one of the supported
// authentication flows: device code, authorization code, or client credentials
func NewLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Long:                  "Login user to the CLI using one of the supported authentication flows: device code, authorization code, or client credentials",
		RunE:                  authLoginRunE,
		Short:                 "Login user to the CLI",
		Use:                   "login [flags]",
	}

	cmd.Flags().AddFlag(options.AuthMethodAuthorizationCodeOption.Flag)
	cmd.Flags().AddFlag(options.AuthMethodClientCredentialsOption.Flag)
	cmd.Flags().AddFlag(options.AuthMethodDeviceCodeOption.Flag)
	cmd.Flags().AddFlag(options.AuthFileStorageOption.Flag)

	// Enforce that exactly one authentication method must be specified
	cmd.MarkFlagsMutuallyExclusive(
		options.AuthMethodAuthorizationCodeOption.Flag.Name,
		options.AuthMethodClientCredentialsOption.Flag.Name,
		options.AuthMethodDeviceCodeOption.Flag.Name,
	)

	return cmd
}

func authLoginRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config login Subcommand Called.")

	if err := auth_internal.AuthLoginRunE(cmd, args); err != nil {
		return &errs.PingCLIError{Prefix: "", Err: err}
	}

	return nil
}
