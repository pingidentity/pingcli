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
	addProfilecommandExamples = `  pingcli config add-profile
  pingcli config add-profile --name myprofile --description "My Profile desc"
  pingcli config add-profile --set-active=true`
)

func NewConfigAddProfileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               addProfilecommandExamples,
		Long:                  `Add a new configuration profile to pingcli.`,
		RunE:                  configAddProfileRunE,
		Short:                 "Add a new configuration profile to pingcli.",
		Use:                   "add-profile [flags]",
	}

	cmd.Flags().AddFlag(options.ConfigAddProfileNameOption.Flag)
	cmd.Flags().AddFlag(options.ConfigAddProfileDescriptionOption.Flag)
	cmd.Flags().AddFlag(options.ConfigAddProfileSetActiveOption.Flag)

	return cmd
}

func configAddProfileRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config add-profile Subcommand Called.")

	if err := config_internal.RunInternalConfigAddProfile(os.Stdin); err != nil {
		return err
	}

	return nil
}
