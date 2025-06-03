// Copyright © 2025 Ping Identity Corporation

package plugin

import (
	"github.com/pingidentity/pingcli/cmd/common"
	plugin_internal "github.com/pingidentity/pingcli/internal/commands/plugin"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	listPluginCommandExamples = `  List all plugins currently in use with PingCLI.
    pingcli plugin list`
)

func NewPluginListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               listPluginCommandExamples,
		Long:                  `List all plugins currently in use with PingCLI.`,
		RunE:                  pluginListRunE,
		Short:                 "List all plugins currently in use with PingCLI",
		Use:                   "list",
	}

	return cmd
}

func pluginListRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Plugin List Subcommand Called.")

	if err := plugin_internal.RunInternalPluginList(); err != nil {
		return err
	}

	return nil
}
