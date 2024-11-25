package config

import (
	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	configUnsetCommandExamples = `  Unset the color setting for the currently active profile.
    pingcli config unset color

  Unset the PingOne tenant region code setting for the profile named 'myProfile'.
    pingcli config unset --profile myProfile service.pingone.regionCode`
)

func NewConfigUnsetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               configUnsetCommandExamples,
		Long: "Unset stored configuration settings for the CLI.\n\n" +
			"The `--profile` parameter can be used to unset configuration settings for a specified custom configuration profile.\n" +
			"Where `--profile` is not specified, configuration settings will be unset for the currently active profile.",
		RunE:  configUnsetRunE,
		Short: "Unset stored configuration settings for the CLI.",
		Use:   "unset [flags] key",
	}

	return cmd
}
func configUnsetRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config unset Subcommand Called.")

	if err := config_internal.RunInternalConfigUnset(args[0]); err != nil {
		return err
	}

	return nil
}
