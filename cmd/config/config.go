package config

import (
	"github.com/spf13/cobra"
)

// const (
// 	configCommandExamples = `  pingcli config
//   pingcli config --profile myprofile
//   pingcli config --name myprofile --description "My Profile"`
// )

func NewConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		// Args:                  common.ExactArgs(0),
		// DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		// Example:               configCommandExamples,
		Long: "Manage the configuration of the CLI, including Ping product connection parameters.\n\n" +
			"The Ping CLI supports the use of configuration profiles.\nConfiguration profiles can be used when connecting to multiple environments using the same Ping CLI instance, such as when managing multiple development or demonstration environments.\n\n" +
			"A pre-defined default profile will be used to store the configuration of the CLI.\nAdditional custom profiles can be created using the `pingcli config add-profile` command.\nTo use a custom profile, the `--profile` flag can be used on supported commands to specify the profile to use for that command.\nTo set a custom profile as the default, use the `pingcli config set-active-profile` command.",
		// RunE:                  configRunE,
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

	// cmd.Flags().AddFlag(options.ConfigProfileOption.Flag)
	// cmd.Flags().AddFlag(options.ConfigNameOption.Flag)
	// cmd.Flags().AddFlag(options.ConfigDescriptionOption.Flag)

	return cmd
}

// func configRunE(cmd *cobra.Command, args []string) error {
// 	l := logger.Get()
// 	l.Debug().Msgf("Config Subcommand Called.")

// 	if err := config_internal.RunInternalConfig(os.Stdin); err != nil {
// 		return err
// 	}

// 	return nil
// }
