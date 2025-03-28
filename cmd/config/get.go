// Copyright © 2025 Ping Identity Corporation

package config

import (
	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	configGetCommandExamples = `  Read all the configuration settings for the PingOne service in the active (or default) profile.
    pingcli config get pingone

  Read the noColor setting for the profile named 'myprofile'.
    pingcli config get --profile myprofile noColor

  Read the worker ID used to authenticate to the PingOne service management API.
    pingcli config get service.pingone.authentication.worker.environmentID
	
  Read the unmasked value for the request access token.
    pingcli config get --unmask-values request.accessToken`
)

func NewConfigGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               configGetCommandExamples,
		Long: "Read stored configuration settings for the CLI.\n\n" +
			"The `--profile` parameter can be used to read configuration settings for a specified custom configuration profile.\n" +
			"Where `--profile` is not specified, configuration settings will be read for the currently active profile.",
		RunE:      configGetRunE,
		Short:     "Read stored configuration settings for the CLI.",
		Use:       "get [flags] key",
		ValidArgs: configuration.ViperKeys(),
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
