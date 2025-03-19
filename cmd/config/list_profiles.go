// Copyright © 2025 Ping Identity Corporation

package config

import (
	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	listProfilesCommandExamples = `  List all custom configuration profiles stored in the CLI configuration file.
    pingcli config list-profiles`
)

func NewConfigListProfilesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               listProfilesCommandExamples,
		Long:                  `List all custom configuration profiles stored in the CLI configuration file.`,
		RunE:                  configListProfilesRunE,
		Short:                 "List all custom configuration profiles.",
		Use:                   "list-profiles",
	}

	return cmd
}

func configListProfilesRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config list-profiles Subcommand Called.")

	err := config_internal.RunInternalConfigListProfiles()
	if err != nil {
		return err
	}

	return nil
}
