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
	setActiveProfileCommandExamples = `  pingcli config set-active-profile
  pingcli config set-active-profile --profile myprofile`
)

func NewConfigSetActiveProfileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               setActiveProfileCommandExamples,
		Long:                  `Set a configuration profile as the in-use profile for pingcli.`,
		RunE:                  configSetActiveProfileRunE,
		Short:                 "Set a configuration profile as the in-use profile for pingcli.",
		Use:                   "set-active-profile [flags]",
	}

	cmd.Flags().AddFlag(options.ConfigSetActiveProfileOption.Flag)

	return cmd
}

func configSetActiveProfileRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config set-active-profile Subcommand Called.")

	if err := config_internal.RunInternalConfigSetActiveProfile(os.Stdin); err != nil {
		return err
	}

	return nil
}
