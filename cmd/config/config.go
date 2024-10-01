package config

import (
	"os"

	"github.com/pingidentity/pingctl/cmd/common"
	config_internal "github.com/pingidentity/pingctl/internal/commands/config"
	"github.com/pingidentity/pingctl/internal/configuration/options"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

const (
	configCommandExamples = `  pingctl config
  pingctl config --profile myprofile
  pingctl config --name myprofile --description "My Profile"`
)

func NewConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               configCommandExamples,
		Long:                  `Update an existing configuration profile's name and description. See subcommands for more profile configuration management options.`,
		RunE:                  configRunE,
		Short:                 "Update an existing configuration profile's name and description. See subcommands for more profile configuration management options.",
		Use:                   "config [flags]",
	}

	// Add subcommands
	cmd.AddCommand(
		NewConfigAddProfileCommand(),
		NewConfigDeleteProfileCommand(),
		NewConfigViewProfileCommand(),
		NewConfigListProfilesCommand(),
		NewConfigSetActiveProfileCommand(),
		NewConfigGetCommand(),
		NewConfigSetCommand(),
		NewConfigUnsetCommand(),
	)

	cmd.Flags().AddFlag(options.ConfigProfileOption.Flag)
	cmd.Flags().AddFlag(options.ConfigNameOption.Flag)
	cmd.Flags().AddFlag(options.ConfigDescriptionOption.Flag)

	return cmd
}

func configRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Subcommand Called.")

	if err := config_internal.RunInternalConfig(os.Stdin); err != nil {
		return err
	}

	return nil
}
