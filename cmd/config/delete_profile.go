package config

import (
	"os"

	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	deleteProfileCommandExamples = `  pingcli config delete-profile
  pingcli config delete-profile --profile myprofile`
)

func NewConfigDeleteProfileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               deleteProfileCommandExamples,
		Long:                  `Delete a configuration profile from pingcli.`,
		RunE:                  configDeleteProfileRunE,
		Short:                 "Delete a configuration profile from pingcli.",
		Use:                   "delete-profile [flags]",
	}

	cmd.Flags().AddFlag(options.ConfigDeleteProfileOption.Flag)

	return cmd
}

func configDeleteProfileRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config delete-profile Subcommand Called.")

	if err := config_internal.RunInternalConfigDeleteProfile(os.Stdin); err != nil {
		return err
	}

	return nil
}
