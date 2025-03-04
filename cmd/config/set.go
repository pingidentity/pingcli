package config

import (
	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	configSetCommandExamples = `  Set the color setting to true for the currently active profile.
    pingcli config set color=true

  Set the PingOne tenant region code setting to 'AP' for the profile named 'myProfile'.
    pingcli config set --profile myProfile service.pingone.regionCode=AP

  Set the PingFederate basic authentication password with unmasked output
    pingcli config set --profile myProfile --unmask-values service.pingfederate.authentication.basicAuth.password=1234`
)

func NewConfigSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               configSetCommandExamples,
		Long: "Set stored configuration settings for the CLI.\n\n" +
			"The `--profile` parameter can be used to set configuration settings for a specified custom configuration profile.\n" +
			"Where `--profile` is not specified, configuration settings will be set for the currently active profile.",
		RunE:      configSetRunE,
		Short:     "Set stored configuration settings for the CLI.",
		Use:       "set [flags] key=value",
		ValidArgs: configuration.ViperKeys(),
	}

	return cmd
}
func configSetRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config set Subcommand Called.")

	if err := config_internal.RunInternalConfigSet(args[0]); err != nil {
		return err
	}

	return nil
}
