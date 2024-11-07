package config

import (
	"os"

	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	setActiveProfileCommandExamples = `  Set an active profile with an interactive prompt to select from an available profile.
    pingcli config set-active-profile

  Set an active profile with a specific profile name.
    pingcli config set-active-profile myprofile`
)

func NewConfigSetActiveProfileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.RangeArgs(0, 1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               setActiveProfileCommandExamples,
		Long:                  `Set a custom configuration profile as the in-use profile.`,
		RunE:                  configSetActiveProfileRunE,
		Short:                 "Set a custom configuration profile as the in-use profile.",
		Use:                   "set-active-profile [flags] [profile-name]",
	}

	return cmd
}

func configSetActiveProfileRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config set-active-profile Subcommand Called.")

	if err := config_internal.RunInternalConfigSetActiveProfile(args, os.Stdin); err != nil {
		return err
	}

	return nil
}
