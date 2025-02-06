package config

import (
	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/profiles"
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
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			profileNames := profiles.GetMainConfig().ProfileNames()
			if len(args) != 0 {
				return nil, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
			}
			return profileNames, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
		},
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
