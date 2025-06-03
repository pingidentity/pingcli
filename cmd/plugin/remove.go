// Copyright © 2025 Ping Identity Corporation

package plugin

import (
	"github.com/pingidentity/pingcli/cmd/common"
	plugin_internal "github.com/pingidentity/pingcli/internal/commands/plugin"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	removePluginCommandExamples = `  Remove a plugin from PingCLI.
    pingcli plugin remove plugin-executable`
)

func NewPluginRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               removePluginCommandExamples,
		Long:                  `Remove a plugin from PingCLI.`,
		RunE:                  pluginRemoveRunE,
		Short:                 "Remove a plugin from PingCLI",
		Use:                   "remove plugin-executable",
	}

	return cmd
}

func pluginRemoveRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Plugin Remove Subcommand Called.")

	if err := plugin_internal.RunInternalPluginRemove(args[0]); err != nil {
		return err
	}

	return nil
}
