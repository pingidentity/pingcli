// Copyright © 2026 Ping Identity Corporation

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
	// ErrUnknownAuthMethod is returned when an unknown authorization grant type is specified
	ErrUnknownAuthMethod = fmt.Errorf("unknown authorization grant type")
)

// NewLoginCommand creates a new login command that authenticates users using one of the supported
// authentication flows: device code, authorization code, or client credentials
func NewLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Long:                  "Authenticate the CLI to a supported provider, using one of the supported authorization grant types.",
		RunE:                  authLoginRunE,
		Short:                 "Authenticate a supported provider",
		Use:                   "login [flags]",
	}

	cmd.Flags().AddFlag(options.AuthMethodAuthorizationCodeOption.Flag)
	cmd.Flags().AddFlag(options.AuthMethodClientCredentialsOption.Flag)
	cmd.Flags().AddFlag(options.AuthMethodDeviceCodeOption.Flag)
	cmd.Flags().AddFlag(options.AuthStorageOption.Flag)
	cmd.Flags().AddFlag(options.AuthProviderOption.Flag)

	// Enforce that exactly one authorization grant type must be specified
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
