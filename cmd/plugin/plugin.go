// Copyright © 2026 Ping Identity Corporation

package plugin

import (
	"github.com/spf13/cobra"
)

func NewPluginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Long:  `Manage Ping CLI plugins.`,
		Short: "Manage Ping CLI plugins.",
		Use:   "plugin",
	}

	cmd.AddCommand(
		NewPluginAddCommand(),
		NewPluginListCommand(),
		NewPluginRemoveCommand(),
	)

	return cmd
}
