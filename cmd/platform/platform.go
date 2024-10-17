package platform

import (
	"github.com/spf13/cobra"
)

func NewPlatformCommand() *cobra.Command {
	cmd := &cobra.Command{
		Long: `Administer and manage the Ping integrated platform.
		
When multiple products are configured in the CLI, the platform command can be used to manage one or more products collectively.

The --profile command switch can be used to specify the profile of Ping products to be managed.`,
		Short: "Administer and manage the Ping integrated platform.",
		Use:   "platform",
	}

	cmd.AddCommand(NewExportCommand())

	return cmd
}
