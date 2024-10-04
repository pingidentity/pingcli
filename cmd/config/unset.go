package config

import (
	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	configUnsetCommandExamples = `  pingcli config unset color
  pingcli config unset --profile myProfile service.pingone.regionCode`
)

func NewConfigUnsetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               configUnsetCommandExamples,
		Long:                  `Unset pingcli configuration settings.`,
		RunE:                  configUnsetRunE,
		Short:                 "Unset pingcli configuration settings.",
		Use:                   "unset [flags] key",
	}

	cmd.Flags().AddFlag(options.ConfigUnsetProfileOption.Flag)

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
