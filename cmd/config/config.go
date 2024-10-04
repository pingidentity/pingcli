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
		Long: `A set of command for profile configuration management.`,
		// RunE:                  configRunE,
		Short: "A set of command for profile configuration management.",
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
