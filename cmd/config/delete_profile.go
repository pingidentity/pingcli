package config

import (
	"os"

	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/spf13/cobra"
)

const (
	deleteProfileCommandExamples = `  Delete a configuration profile by selecting from the available profiles.
    pingcli config delete-profile

  Delete a configuration profile by specifying the name of an existing configured profile.
    pingcli config delete-profile MyDeveloperEnv
	
  Delete a configuration profile by auto-accepting the deletion.
    pingcli config delete-profile --yes MyDeveloperEnv`
)

func NewConfigDeleteProfileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.RangeArgs(0, 1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               deleteProfileCommandExamples,
		Long: `Delete an existing custom configuration profile from the CLI.
		
The profile to delete will be removed from the CLI configuration file.`,
		RunE:  configDeleteProfileRunE,
		Short: "Delete a custom configuration profile.",
		Use:   "delete-profile [flags] [profile-name]",
		// Auto-completion function to return all valid profile names
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			profileNames := profiles.GetMainConfig().ProfileNames()
			if len(args) != 0 {
				return nil, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
			}
			return profileNames, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
		},
	}

	cmd.Flags().AddFlag(options.ConfigDeleteAutoAcceptOption.Flag)

	return cmd
}

func configDeleteProfileRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config delete-profile Subcommand Called.")

	if err := config_internal.RunInternalConfigDeleteProfile(args, os.Stdin); err != nil {
		return err
	}

	return nil
}
