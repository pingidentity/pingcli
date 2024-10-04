package config

import (
	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	listProfilesCommandExamples = `  pingcli config list-profiles`
)

func NewConfigListProfilesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               listProfilesCommandExamples,
		Long:                  `List all configuration profiles from pingcli.`,
		RunE:                  configListProfilesRunE,
		Short:                 "List all configuration profiles from pingcli.",
		Use:                   "list-profiles",
	}

	return cmd
}

func configListProfilesRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config list-profiles Subcommand Called.")

	config_internal.RunInternalConfigListProfiles()

	return nil
}
