// Copyright © 2025 Ping Identity Corporation

package platform

import (
	"fmt"

	"github.com/pingidentity/pingcli/cmd/common"
	"github.com/pingidentity/pingcli/internal/autocompletion"
	platform_internal "github.com/pingidentity/pingcli/internal/commands/platform"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/spf13/cobra"
)

const (
	commandExamples = `  Export Configuration as Code for all products configured in the configuration file, applying default options.
    pingcli platform export

  Export Configuration as Code packages for all configured products to a specific directory, overwriting any previous export.
    pingcli platform export --output-directory /path/to/my/directory --overwrite

  Export Configuration as Code packages for all configured products, specifying the export format as Terraform HCL.
    pingcli platform export --format HCL

  Export Configuration as Code packages for PingOne (core platform and SSO services).
    pingcli platform export --services pingone-platform,pingone-sso

  Export all Configuration as Code packages for PingOne. The --service-group flag can be used instead of listing all pingone-* packages in --services flag.
    pingcli platform export --service-group pingone

  Export Configuration as Code packages for PingOne (core platform), specifying the PingOne environment connection details.
    pingcli platform export --services pingone-platform --pingone-client-environment-id 3cf2... --pingone-worker-client-id a719... --pingone-worker-client-secret ey..... --pingone-region-code EU

  Export Configuration as Code packages for PingFederate, specifying the PingFederate connection details using basic authentication.
    pingcli platform export --services pingfederate --pingfederate-authentication-type basicAuth --pingfederate-username administrator --pingfederate-password 2FederateM0re --pingfederate-https-host https://pingfederate-admin.bxretail.org

  Export Configuration as Code packages for PingFederate, specifying the PingFederate connection details using OAuth 2.0 client credentials.
    pingcli platform export --services pingfederate --pingfederate-authentication-type clientCredentialsAuth --pingfederate-client-id clientID --pingfederate-client-secret clientSecret --pingfederate-token-url https://pingfederate-admin.bxretail.org/as/token.oauth2

  Export Configuration as Code packages for PingFederate, specifying optional connection properties
    pingcli platform export --services pingfederate --x-bypass-external-validation=false --ca-certificate-pem-files "/path/to/cert.pem,/path/to/cert2.pem" --insecure-trust-all-tls=false`
)

func NewExportCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               commandExamples,
		Long: "Export Configuration as Code packages for the Ping Platform.\n\n" +
			"The CLI can export Terraform HCL to use with released Terraform providers.\n" +
			"The Terraform HCL option generates `import {}` block statements for resources in the target environment.\n" +
			"Using Terraform `import {}` blocks, the platform's configuration can be generated and imported into state management.\n" +
			"More information can be found at https://developer.hashicorp.com/terraform/language/import",
		Short: "Export Configuration as Code packages for the Ping Platform.",
		RunE:  exportRunE,
		Use:   "export [flags]",
	}

	initGeneralExportFlags(cmd)
	initPingOneExportFlags(cmd)
	initPingFederateGeneralFlags(cmd)
	initPingFederateBasicAuthFlags(cmd)
	initPingFederateAccessTokenFlags(cmd)
	initPingFederateClientCredentialsFlags(cmd)

	// auto-completion
	err := cmd.RegisterFlagCompletionFunc(options.PlatformExportExportFormatOption.CobraParamName, autocompletion.PlatformExportFormatFunc)
	if err != nil {
		output.SystemError(fmt.Sprintf("Unable to register auto completion for platform export flag %s: %v", options.PlatformExportExportFormatOption.CobraParamName, err), nil)
	}

	err = cmd.RegisterFlagCompletionFunc(options.PlatformExportServiceOption.CobraParamName, autocompletion.PlatformExportServicesFunc)
	if err != nil {
		output.SystemError(fmt.Sprintf("Unable to register auto completion for platform export flag %s: %v", options.PlatformExportServiceOption.CobraParamName, err), nil)
	}

	err = cmd.RegisterFlagCompletionFunc(options.PingOneAuthenticationTypeOption.CobraParamName, autocompletion.PlatformExportPingOneAuthenticationTypeFunc)
	if err != nil {
		output.SystemError(fmt.Sprintf("Unable to register auto completion for platform export flag %s: %v", options.PingOneAuthenticationTypeOption.CobraParamName, err), nil)
	}

	return cmd
}

func exportRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()

	l.Debug().Msgf("Platform Export Subcommand Called.")

	return platform_internal.RunInternalExport(cmd.Context(), cmd.Root().Version)
}

func initGeneralExportFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PlatformExportExportFormatOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportServiceOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportServiceGroupOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportOutputDirectoryOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportOverwriteOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportPingOneEnvironmentIDOption.Flag)
}

func initPingOneExportFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PingOneAuthenticationWorkerEnvironmentIDOption.Flag)
	cmd.Flags().AddFlag(options.PingOneAuthenticationWorkerClientIDOption.Flag)
	cmd.Flags().AddFlag(options.PingOneAuthenticationWorkerClientSecretOption.Flag)
	cmd.Flags().AddFlag(options.PingOneAuthenticationTypeOption.Flag)
	cmd.Flags().AddFlag(options.PingOneRegionCodeOption.Flag)

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
