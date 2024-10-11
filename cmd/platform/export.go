package platform

import (
	"github.com/pingidentity/pingcli/cmd/common"
	platform_internal "github.com/pingidentity/pingcli/internal/commands/platform"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/spf13/cobra"
)

const (
	commandExamples = `  pingcli platform export
  pingcli platform export --output-directory dir --overwrite
  pingcli platform export --export-format HCL
  pingcli platform export --services pingone-platform,pingone-sso
  pingcli platform export --services pingone-platform --pingone-client-environment-id envID --pingone-worker-client-id clientID --pingone-worker-client-secret clientSecret --pingone-region-code regionCode
  pingcli platform export --service pingfederate --pingfederate-username user --pingfederate-password password
  pingcli platform export --service pingfederate --pingfederate-client-id clientID --pingfederate-client-secret clientSecret --pingfederate-token-url tokenURL
  pingcli platform export --service pingfederate --pingfederate-access-token accessToken
  pingcli platform export --service pingfederate --x-bypass-external-validation=false --ca-certificate-pem-files "/path/to/cert.pem,/path/to/cert2.pem" --insecure-trust-all-tls=false`
)

func NewExportCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               commandExamples,
		Long:                  `Export configuration-as-code packages for the Ping Platform.`,
		Short:                 "Export configuration-as-code packages for the Ping Platform.",
		RunE:                  exportRunE,
		Use:                   "export [flags]",
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
	cmd.Flags().AddFlag(options.PlatformExportServiceOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportOutputDirectoryOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportOverwriteOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportPingOneEnvironmentIDOption.Flag)
}

func initPingOneExportFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PingOneAuthenticationWorkerEnvironmentIDOption.Flag)
	cmd.Flags().AddFlag(options.PingOneAuthenticationWorkerClientIDOption.Flag)
	cmd.Flags().AddFlag(options.PingOneAuthenticationWorkerClientSecretOption.Flag)
	cmd.Flags().AddFlag(options.PingOneRegionCodeOption.Flag)
	cmd.Flags().AddFlag(options.PingOneAuthenticationTypeOption.Flag)

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
