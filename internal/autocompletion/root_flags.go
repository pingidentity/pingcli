// Copyright © 2025 Ping Identity Corporation

package autocompletion

import (
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/spf13/cobra"
)

func RootProfileFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return profiles.GetMainConfig().ProfileNames(), cobra.ShellCompDirectiveNoFileComp
}

func RootOutputFormatFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return customtypes.OutputFormatValidValues(), cobra.ShellCompDirectiveNoFileComp
}
