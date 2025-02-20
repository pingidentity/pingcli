package autocompletion

import (
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/cobra"
)

func PlatformExportFormatFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return customtypes.ExportFormatValidValues(), cobra.ShellCompDirectiveNoFileComp
}

func PlatformExportPingFederateAuthenticationTypeFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return customtypes.PingFederateAuthenticationTypeValidValues(), cobra.ShellCompDirectiveNoFileComp
}

func PlatformExportPingOneAuthenticationTypeFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return customtypes.PingOneAuthenticationTypeValidValues(), cobra.ShellCompDirectiveNoFileComp
}

func PlatformExportPingOneRegionCodeFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return customtypes.PingOneRegionCodeValidValues(), cobra.ShellCompDirectiveNoFileComp
}

func PlatformExportServicesFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return customtypes.ExportServicesValidValues(), cobra.ShellCompDirectiveNoFileComp
}
