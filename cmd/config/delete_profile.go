package config

import (
	"os"

	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	deleteProfileCommandExamples = `  Delete a configuration profile by selecting from the available profiles.
    pingcli config delete-profile`
)

func NewConfigDeleteProfileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               deleteProfileCommandExamples,
		Long: `Delete an existing custom configuration profile from the CLI.
		
The profile to delete will be removed from the CLI configuration file.`,
		RunE:  configDeleteProfileRunE,
		Short: "Delete a custom configuration profile.",
		Use:   "delete-profile",
	}

	return cmd
}

func configDeleteProfileRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config delete-profile Subcommand Called.")

	if err := config_internal.RunInternalConfigDeleteProfile(os.Stdin); err != nil {
		return err
	}

	return nil
}
