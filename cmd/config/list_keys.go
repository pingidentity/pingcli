package config

import (
	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	listKeysCommandExamples = `  List all configuration keys stored in the CLI configuration file.
  pingcli config list-keys
	pingcli config list-keys --yaml`
)

func NewConfigListKeysCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               listKeysCommandExamples,
		Long: `View the complete list of available configuration options. These attributes can be saved via the set and unset config subcommands or stored in a profile within the config file.
For details on the configuration options visit: https://github.com/pingidentity/pingcli/blob/main/docs/tool-configuration/configuration-key.md`,
		RunE:  configListKeysRunE,
		Short: "List all configuration keys.",
		Use:   "list-keys [flags]",
	}

	cmd.Flags().AddFlag(options.ConfigListKeysYamlOption.Flag)

	return cmd
}

func configListKeysRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config list-keys Subcommand Called.")

	err := config_internal.RunInternalConfigListKeys()
	if err != nil {
		return err
	}

	return nil
}
