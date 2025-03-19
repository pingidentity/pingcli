// Copyright © 2025 Ping Identity Corporation

package autocompletion

import (
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/cobra"
)

func RequestHTTPMethodFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return customtypes.HTTPMethodValidValues(), cobra.ShellCompDirectiveNoFileComp
}

func RequestServiceFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return customtypes.RequestServiceValidValues(), cobra.ShellCompDirectiveNoFileComp
}
