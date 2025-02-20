package autocompletion

import (
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/spf13/cobra"
)

func Profile(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return profiles.GetMainConfig().ProfileNames(), cobra.ShellCompDirectiveNoFileComp
}

func OutputFormat(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return customtypes.OutputFormatValidValues(), cobra.ShellCompDirectiveNoFileComp
}
