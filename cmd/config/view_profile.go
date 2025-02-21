package config

import (
	"github.com/pingidentity/pingcli/cmd/common"
	"github.com/pingidentity/pingcli/internal/autocompletion"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	viewProfileCommandExamples = `  View configuration for the currently active profile
    pingcli config view-profile

  View configuration for a specific profile
    pingcli config view-profile myprofile`
)

func NewConfigViewProfileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.RangeArgs(0, 1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               viewProfileCommandExamples,
		Long:                  `View the stored configuration of a custom configuration profile.`,
		RunE:                  configViewProfileRunE,
		Short:                 "View the stored configuration of a custom configuration profile.",
		Use:                   "view-profile [flags] [profile-name]",
		// Auto-completion function to return all valid profile names
		ValidArgsFunction: autocompletion.ConfigViewProfileFunc,
	}

	return cmd
}

func configViewProfileRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config view-profile Subcommand Called.")

	if err := config_internal.RunInternalConfigViewProfile(args); err != nil {
		return err
	}

	return nil
}
