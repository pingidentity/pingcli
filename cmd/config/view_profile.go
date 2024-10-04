package config

import (
	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	viewProfileCommandExamples = `  pingcli config view-profile
  pingcli config view-profile --profile myprofile`
)

func NewConfigViewProfileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               viewProfileCommandExamples,
		Long:                  `View a configuration profile from pingcli.`,
		RunE:                  configViewProfileRunE,
		Short:                 "View a configuration profile from pingcli.",
		Use:                   "view-profile [flags]",
	}

	cmd.Flags().AddFlag(options.ConfigViewProfileOption.Flag)

	return cmd
}

func configViewProfileRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config view-profile Subcommand Called.")

	if err := config_internal.RunInternalConfigViewProfile(); err != nil {
		return err
	}

	return nil
}
