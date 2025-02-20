package autocompletion

import (
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/spf13/cobra"
)

func Profile(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	validProfileNames := profiles.GetMainConfig().ProfileNames()
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	return validProfileNames, cobra.ShellCompDirectiveNoFileComp
}

func OutputFormat(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	validOutputFormats := customtypes.OutputFormatValidValues()
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	return validOutputFormats, cobra.ShellCompDirectiveNoFileComp
}
