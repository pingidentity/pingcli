package autocompletion_export_flags

import (
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/spf13/cobra"
)

func Format(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	validExportFormats := customtypes.ExportFormatValidValues()
	return validExportFormats, cobra.ShellCompDirectiveNoFileComp
}

func PingFederateAuthenticationType(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	validPingFedAuthTypes := customtypes.PingFederateAuthenticationTypeValidValues()
	return validPingFedAuthTypes, cobra.ShellCompDirectiveNoFileComp
}

func PingOneAuthenticationType(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	validPingOneAuthTypes := customtypes.PingOneAuthenticationTypeValidValues()
	return validPingOneAuthTypes, cobra.ShellCompDirectiveNoFileComp
}

func PingOneRegionCode(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	validRegionCodes := customtypes.PingOneRegionCodeValidValues()
	return validRegionCodes, cobra.ShellCompDirectiveNoFileComp
}

func Services(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	validServices := customtypes.ExportServicesValidValues()
	return validServices, cobra.ShellCompDirectiveNoFileComp
}
