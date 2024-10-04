package config

import (
	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	configGetCommandExamples = `  pingcli config get pingone
  pingcli config get --profile myProfile color
  pingcli config get service.pingone.authentication.worker.environmentID`
)

func NewConfigGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               configGetCommandExamples,
		Long:                  `Get pingcli configuration settings.`,
		RunE:                  configGetRunE,
		Short:                 "Get pingcli configuration settings.",
		Use:                   "get [flags] key",
	}

	cmd.Flags().AddFlag(options.ConfigGetProfileOption.Flag)

	return cmd
}

func configGetRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Get Subcommand Called.")

	if err := config_internal.RunInternalConfigGet(args[0]); err != nil {
		return err
	}

	return nil
}
