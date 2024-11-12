package config

import (
	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	configGetCommandExamples = `  Read all the configuration settings for the PingOne service in the active (or default) profile.
    pingcli config get pingone

  Read the color setting for the profile named 'myProfile'.
    pingcli config get --profile myProfile color

  Read the worker ID used to authenticate to the PingOne service management API.
    pingcli config get service.pingone.authentication.worker.environmentID`
)

func NewConfigGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               configGetCommandExamples,
		Long: "Read stored configuration settings for the CLI.\n\n" +
			"The `--profile` parameter can be used to read configuration settings for a specified custom configuration profile.\n" +
			"Where `--profile` is not specified, configuration settings will be read for the currently active profile.",
		RunE:  configGetRunE,
		Short: "Read stored configuration settings for the CLI.",
		Use:   "get [flags] key",
	}

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
