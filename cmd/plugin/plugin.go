// Copyright © 2025 Ping Identity Corporation

package plugin

import (
	"github.com/spf13/cobra"
)

func NewPluginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Long:  `Manage PingCLI plugins.`,
		Short: "Manage PingCLI plugins.",
		Use:   "plugin",
	}

	cmd.AddCommand(NewPluginAddCommand())

	return cmd
}
