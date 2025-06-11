// Copyright © 2025 Ping Identity Corporation

package plugin

import (
	"github.com/pingidentity/pingcli/cmd/common"
	plugin_internal "github.com/pingidentity/pingcli/internal/commands/plugin"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	addPluginCommandExamples = `  Add a plugin to use with PingCLI.
    pingcli plugin add pingcli-plugin-executable`
)

func NewPluginAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               addPluginCommandExamples,
		Long:                  `Add a plugin to use with PingCLI.`,
		RunE:                  pluginAddRunE,
		Short:                 "Add a plugin to use with PingCLI",
		Use:                   "add plugin-executable",
	}

	return cmd
}

func pluginAddRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Plugin Add Subcommand Called.")

	if err := plugin_internal.RunInternalPluginAdd(args[0]); err != nil {
		return err
	}

	return nil
}
