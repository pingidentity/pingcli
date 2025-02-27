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
	addProfilecommandExamples = `  Add a new configuration profile with a guided experience.
    pingcli config add-profile

  Add a new configuration profile with a specific name and description.
    pingcli config add-profile --name myprofile --description "My awesome new profile for my development environment"

  Add a new configuration profile with a guided experience and set it as the active profile.
    pingcli config add-profile --set-active`
)

func NewConfigAddProfileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               addProfilecommandExamples,
		Long: `Add a new custom configuration profile to the CLI.

The new configuration profile will be stored in the CLI configuration file.`,
		RunE:  configAddProfileRunE,
		Short: "Add a new custom configuration profile.",
		Use:   "add-profile [flags]",
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
