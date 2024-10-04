package platform

import (
	"fmt"

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

	profileConfigurationFormat = `Profile Configuration Format:
export:
	format: <Format>
	services:
		- <Service>
		- <Service>
	outputDirectory: <Filepath>
	overwrite: <true|false>
	pingone:
		environmentID: <ID>
service:
	pingfederate:
		httpsHost: <Host>
		adminAPIPath: <Path>
		x-bypass-external-validation: <true|false>
		ca-certificate-pem-files:
			- <Filepath>
			- <Filepath>
		insecure-trust-all-tls: <true|false>
		authentication:
			type: <Type>
			basicAuth:
				username: <Username>
				password: <Password>
			accessTokenAuth:
				accessToken: <Token>
			clientCredentialsAuth:
				clientID: <ID>
				clientSecret: <Secret>
				tokenURL: <URL>
				scopes:
					- <Scope>
					- <Scope>
    pingone:
        regionCode: <Code>
        authentication:
            type: <Type>
            worker:
                clientID: <ID>
                clientSecret: <Secret>
                environmentID: <ID>`
)

func NewExportCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               fmt.Sprintf("%s\n\n%s", commandExamples, profileConfigurationFormat),
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
	cmd.Flags().AddFlag(options.PlatformExportPingoneEnvironmentIDOption.Flag)
}

func initPingOneExportFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PingoneAuthenticationWorkerEnvironmentIDOption.Flag)
	cmd.Flags().AddFlag(options.PingoneAuthenticationWorkerClientIDOption.Flag)
	cmd.Flags().AddFlag(options.PingoneAuthenticationWorkerClientSecretOption.Flag)
	cmd.Flags().AddFlag(options.PingoneRegionCodeOption.Flag)
	cmd.Flags().AddFlag(options.PingoneAuthenticationTypeOption.Flag)

	cmd.MarkFlagsRequiredTogether(
		options.PingoneAuthenticationWorkerEnvironmentIDOption.CobraParamName,
		options.PingoneAuthenticationWorkerClientIDOption.CobraParamName,
		options.PingoneAuthenticationWorkerClientSecretOption.CobraParamName,
		options.PingoneRegionCodeOption.CobraParamName,
	)

}

func initPingFederateGeneralFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PingfederateHTTPSHostOption.Flag)
	cmd.Flags().AddFlag(options.PingfederateAdminAPIPathOption.Flag)

	cmd.MarkFlagsRequiredTogether(
		options.PingfederateHTTPSHostOption.CobraParamName,
		options.PingfederateAdminAPIPathOption.CobraParamName)

	cmd.Flags().AddFlag(options.PingfederateXBypassExternalValidationHeaderOption.Flag)
	cmd.Flags().AddFlag(options.PingfederateCACertificatePemFilesOption.Flag)
	cmd.Flags().AddFlag(options.PingfederateInsecureTrustAllTLSOption.Flag)
	cmd.Flags().AddFlag(options.PingfederateAuthenticationTypeOption.Flag)
}

func initPingFederateBasicAuthFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PingfederateBasicAuthUsernameOption.Flag)
	cmd.Flags().AddFlag(options.PingfederateBasicAuthPasswordOption.Flag)

	cmd.MarkFlagsRequiredTogether(
		options.PingfederateBasicAuthUsernameOption.CobraParamName,
		options.PingfederateBasicAuthPasswordOption.CobraParamName,
	)
}

func initPingFederateAccessTokenFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PingfederateAccessTokenAuthAccessTokenOption.Flag)
}

func initPingFederateClientCredentialsFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PingfederateClientCredentialsAuthClientIDOption.Flag)
	cmd.Flags().AddFlag(options.PingfederateClientCredentialsAuthClientSecretOption.Flag)
	cmd.Flags().AddFlag(options.PingfederateClientCredentialsAuthTokenURLOption.Flag)

	cmd.MarkFlagsRequiredTogether(
		options.PingfederateClientCredentialsAuthClientIDOption.CobraParamName,
		options.PingfederateClientCredentialsAuthClientSecretOption.CobraParamName,
		options.PingfederateClientCredentialsAuthTokenURLOption.CobraParamName)

	cmd.Flags().AddFlag(options.PingfederateClientCredentialsAuthScopesOption.Flag)
}
