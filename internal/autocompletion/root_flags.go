// Copyright © 2025 Ping Identity Corporation

package autocompletion

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/spf13/cobra"
)

func RootProfileFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		output.SystemError(fmt.Sprintf("Unable to get configuration: %v", err), nil)
	}

	return koanfConfig.ProfileNames(), cobra.ShellCompDirectiveNoFileComp
}

func RootOutputFormatFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return customtypes.OutputFormatValidValues(), cobra.ShellCompDirectiveNoFileComp
}
