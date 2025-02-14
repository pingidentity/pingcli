package autocompletion_request_flags

import (
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/cobra"
)

func Data(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(toComplete) != 0 {
		return nil, cobra.ShellCompDirectiveDefault
	}
	return nil, cobra.ShellCompDirectiveNoFileComp
}

func HTTPMethod(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	validHTTPMethods := customtypes.HTTPMethodValidValues()
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	return validHTTPMethods, cobra.ShellCompDirectiveNoFileComp
}

func Service(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	validServices := customtypes.RequestServiceValidValues()
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	return validServices, cobra.ShellCompDirectiveNoFileComp
}
