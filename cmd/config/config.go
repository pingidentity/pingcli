package config

import (
	"github.com/spf13/cobra"
)

func NewConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Long: "Manage the configuration of the CLI, including Ping product connection parameters.\n\n" +
			"The Ping CLI supports the use of configuration profiles." +
			"\nConfiguration profiles can be used when connecting to multiple environments using the same Ping CLI " +
			"instance, such as when managing multiple development or demonstration environments.\n\n" +
			"A pre-defined default profile will be used to store the configuration of the CLI." +
			"\nAdditional custom profiles can be created using the `pingcli config add-profile` command." +
			"\nTo use a custom profile, the `--profile` flag can be used on supported commands to specify the " +
			"profile to use for that command." +
			"\nTo set a custom profile as the default, use the `pingcli config set-active-profile` command.",
		Short: "Manage the CLI configuration.",
		Use:   "config",
	}

	// Add subcommands
	cmd.AddCommand(
		NewConfigAddProfileCommand(),
		NewConfigDeleteProfileCommand(),
		NewConfigViewProfileCommand(),
		NewConfigListProfilesCommand(),
		NewConfigSetActiveProfileCommand(),
		NewConfigGetCommand(),
		NewConfigSetCommand(),
		NewConfigUnsetCommand(),
	)

	return cmd
}
