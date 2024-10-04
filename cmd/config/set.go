package config

import (
	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	configSetCommandExamples = `  pingcli config set color=true
  pingcli config set --profile myProfile service.pingone.regionCode=AP`
)

func NewConfigSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               configSetCommandExamples,
		Long:                  `Set pingcli configuration settings.`,
		RunE:                  configSetRunE,
		Short:                 "Set pingcli configuration settings.",
		Use:                   "set [flags] key=value",
	}

	cmd.Flags().AddFlag(options.ConfigSetProfileOption.Flag)

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
