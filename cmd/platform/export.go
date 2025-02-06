package platform

import (
	"github.com/pingidentity/pingcli/cmd/common"
	platform_internal "github.com/pingidentity/pingcli/internal/commands/platform"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	commandExamples = `  Export configuration-as-code for all products configured in the configuration file, applying default options.
    pingcli platform export

  Export configuration-as-code packages for all configured products to a specific directory, overwriting any previous export.
    pingcli platform export --output-directory /path/to/my/directory --overwrite

  Export configuration-as-code packages for all configured products, specifying the export format as Terraform HCL.
    pingcli platform export --format HCL

  Export configuration-as-code packages for PingOne (core platform and SSO services).
    pingcli platform export --services pingone-platform,pingone-sso

  Export configuration-as-code packages for PingOne (core platform), specifying the PingOne environment connection details.
    pingcli platform export --services pingone-platform --pingone-client-environment-id 3cf2... --pingone-worker-client-id a719... --pingone-worker-client-secret ey..... --pingone-region-code EU

  Export configuration-as-code packages for PingFederate, specifying the PingFederate connection details using basic authentication.
    pingcli platform export --services pingfederate --pingfederate-authentication-type basicAuth --pingfederate-username administrator --pingfederate-password 2FederateM0re --pingfederate-https-host https://pingfederate-admin.bxretail.org

  Export configuration-as-code packages for PingFederate, specifying the PingFederate connection details using OAuth 2.0 client credentials.
    pingcli platform export --services pingfederate --pingfederate-authentication-type clientCredentialsAuth --pingfederate-client-id clientID --pingfederate-client-secret clientSecret --pingfederate-token-url https://pingfederate-admin.bxretail.org/as/token.oauth2

  Export configuration-as-code packages for PingFederate, specifying optional connection properties
    pingcli platform export --services pingfederate --x-bypass-external-validation=false --ca-certificate-pem-files "/path/to/cert.pem,/path/to/cert2.pem" --insecure-trust-all-tls=false`
)

func NewExportCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               commandExamples,
		Long: "Export configuration-as-code packages for the Ping Platform.\n\n" +
			"The CLI can export Terraform HCL to use with released Terraform providers.\n" +
			"The Terraform HCL option generates `import {}` block statements for resources in the target environment.\n" +
			"Using Terraform `import {}` blocks, the platform's configuration can be generated and imported into state management.\n" +
			"More information can be found at https://developer.hashicorp.com/terraform/language/import",
		Short: "Export configuration-as-code packages for the Ping Platform.",
		RunE:  exportRunE,
		Use:   "export [flags]",
	}

	initGeneralExportFlags(cmd)
	initPingOneExportFlags(cmd)
	initPingFederateGeneralFlags(cmd)
	initPingFederateBasicAuthFlags(cmd)
	initPingFederateAccessTokenFlags(cmd)
	initPingFederateClientCredentialsFlags(cmd)

	return cmd
}

func exportRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()

	l.Debug().Msgf("Platform Export Subcommand Called.")

	return platform_internal.RunInternalExport(cmd.Context(), cmd.Root().Version)
}

func initGeneralExportFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PlatformExportExportFormatOption.Flag)
	// auto-completion
	cmd.RegisterFlagCompletionFunc("format", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		validExportFormats := customtypes.ExportFormatValidValues()
		return validExportFormats, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
	})

	cmd.Flags().AddFlag(options.PlatformExportServiceOption.Flag)
	// auto-completion
	cmd.RegisterFlagCompletionFunc("services", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		validServices := customtypes.ExportServicesValidValues()
		return validServices, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
	})

	cmd.Flags().AddFlag(options.PlatformExportOutputDirectoryOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportOverwriteOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportPingOneEnvironmentIDOption.Flag)
}

func initPingOneExportFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PingOneAuthenticationWorkerEnvironmentIDOption.Flag)
	cmd.Flags().AddFlag(options.PingOneAuthenticationWorkerClientIDOption.Flag)
	cmd.Flags().AddFlag(options.PingOneAuthenticationWorkerClientSecretOption.Flag)

	cmd.Flags().AddFlag(options.PingOneRegionCodeOption.Flag)
	// auto-completion
	cmd.RegisterFlagCompletionFunc("pingone-region-code", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		validRegionCodes := customtypes.PingOneRegionCodeValidValues()
		return validRegionCodes, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
	})

	cmd.Flags().AddFlag(options.PingOneAuthenticationTypeOption.Flag)
	// auto-completion
	cmd.RegisterFlagCompletionFunc("pingone-authentication-type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		validPingOneAuthTypes := customtypes.PingOneAuthenticationTypeValidValues()
		return validPingOneAuthTypes, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
	})

	cmd.MarkFlagsRequiredTogether(
		options.PingOneAuthenticationWorkerEnvironmentIDOption.CobraParamName,
		options.PingOneAuthenticationWorkerClientIDOption.CobraParamName,
		options.PingOneAuthenticationWorkerClientSecretOption.CobraParamName,
		options.PingOneRegionCodeOption.CobraParamName,
	)

}

func initPingFederateGeneralFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PingFederateHTTPSHostOption.Flag)
	cmd.Flags().AddFlag(options.PingFederateAdminAPIPathOption.Flag)

	cmd.MarkFlagsRequiredTogether(
		options.PingFederateHTTPSHostOption.CobraParamName,
		options.PingFederateAdminAPIPathOption.CobraParamName)

	cmd.Flags().AddFlag(options.PingFederateXBypassExternalValidationHeaderOption.Flag)
	cmd.Flags().AddFlag(options.PingFederateCACertificatePemFilesOption.Flag)
	cmd.Flags().AddFlag(options.PingFederateInsecureTrustAllTLSOption.Flag)

	cmd.Flags().AddFlag(options.PingFederateAuthenticationTypeOption.Flag)
	// auto-completion
	cmd.RegisterFlagCompletionFunc("pingfederate-authentication-type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		validPingFedAuthTypes := customtypes.PingFederateAuthenticationTypeValidValues()
		return validPingFedAuthTypes, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
	})
}

func initPingFederateBasicAuthFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PingFederateBasicAuthUsernameOption.Flag)
	cmd.Flags().AddFlag(options.PingFederateBasicAuthPasswordOption.Flag)

	cmd.MarkFlagsRequiredTogether(
		options.PingFederateBasicAuthUsernameOption.CobraParamName,
		options.PingFederateBasicAuthPasswordOption.CobraParamName,
	)
}

func initPingFederateAccessTokenFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PingFederateAccessTokenAuthAccessTokenOption.Flag)
}

func initPingFederateClientCredentialsFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PingFederateClientCredentialsAuthClientIDOption.Flag)
	cmd.Flags().AddFlag(options.PingFederateClientCredentialsAuthClientSecretOption.Flag)
	cmd.Flags().AddFlag(options.PingFederateClientCredentialsAuthTokenURLOption.Flag)

	cmd.MarkFlagsRequiredTogether(
		options.PingFederateClientCredentialsAuthClientIDOption.CobraParamName,
		options.PingFederateClientCredentialsAuthClientSecretOption.CobraParamName,
		options.PingFederateClientCredentialsAuthTokenURLOption.CobraParamName)

	cmd.Flags().AddFlag(options.PingFederateClientCredentialsAuthScopesOption.Flag)
}
