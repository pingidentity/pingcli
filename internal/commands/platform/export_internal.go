// Copyright © 2025 Ping Identity Corporation

package platform_internal

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	pingoneGoClient "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate"
	"github.com/pingidentity/pingcli/internal/connector/pingone/authorize"
	"github.com/pingidentity/pingcli/internal/connector/pingone/mfa"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform"
	"github.com/pingidentity/pingcli/internal/connector/pingone/protect"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	pingfederateGoClient "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

var (
	pingfederateApiClient *pingfederateGoClient.APIClient
	pingfederateContext   context.Context

	pingoneApiClient   *pingoneGoClient.Client
	pingoneApiClientId string
	pingoneExportEnvID string
	pingoneContext     context.Context
)

func RunInternalExport(ctx context.Context, commandVersion string) (err error) {
	if ctx == nil {
		return fmt.Errorf("failed to run 'platform export' command. context is nil")
	}

	exportFormat, err := profiles.GetOptionValue(options.PlatformExportExportFormatOption)
	if err != nil {
		return err
	}
	exportServices, err := profiles.GetOptionValue(options.PlatformExportServiceOption)
	if err != nil {
		return err
	}
	outputDir, err := profiles.GetOptionValue(options.PlatformExportOutputDirectoryOption)
	if err != nil {
		return err
	}
	overwriteExport, err := profiles.GetOptionValue(options.PlatformExportOverwriteOption)
	if err != nil {
		return err
	}

	es := new(customtypes.ExportServices)
	if err = es.Set(exportServices); err != nil {
		return err
	}

	if es.ContainsPingOneService() {
		if err = initPingOneServices(ctx, commandVersion); err != nil {
			return err
		}
	}

	if es.ContainsPingFederateService() {
		if err = initPingFederateServices(ctx, commandVersion); err != nil {
			return err
		}
	}

	overwriteExportBool, err := strconv.ParseBool(overwriteExport)
	if err != nil {
		return err
	}
	if outputDir, err = createOrValidateOutputDir(outputDir, overwriteExportBool); err != nil {
		return err
	}

	exportableConnectors := getExportableConnectors(es)

	if err := exportConnectors(exportableConnectors, exportFormat, outputDir, overwriteExportBool); err != nil {
		return err
	}

	output.Success(fmt.Sprintf("Export to directory '%s' complete.", outputDir), nil)

	return nil
}

func initPingFederateServices(ctx context.Context, pingcliVersion string) (err error) {
	if ctx == nil {
		return fmt.Errorf("failed to initialize PingFederate services. context is nil")
	}

	pfInsecureTrustAllTLS, err := profiles.GetOptionValue(options.PingFederateInsecureTrustAllTLSOption)
	if err != nil {
		return err
	}
	caCertPemFiles, err := profiles.GetOptionValue(options.PingFederateCACertificatePemFilesOption)
	if err != nil {
		return err
	}

	caCertPool := x509.NewCertPool()
	for _, caCertPemFile := range strings.Split(caCertPemFiles, ",") {
		if caCertPemFile == "" {
			continue
		}
		caCertPemFile := filepath.Clean(caCertPemFile)
		caCert, err := os.ReadFile(caCertPemFile)
		if err != nil {
			return fmt.Errorf("failed to read CA certificate PEM file '%s': %v", caCertPemFile, err)
		}

		ok := caCertPool.AppendCertsFromPEM(caCert)
		if !ok {
			return fmt.Errorf("failed to parse CA certificate PEM file '%s' to certificate pool", caCertPemFile)
		}
	}

	pfInsecureTrustAllTLSBool, err := strconv.ParseBool(pfInsecureTrustAllTLS)
	if err != nil {
		return err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: pfInsecureTrustAllTLSBool, //#nosec G402 -- This is defined by the user (default false), and warned as inappropriate in production.
			RootCAs:            caCertPool,
		},
	}

	if err = initPingFederateApiClient(tr, pingcliVersion); err != nil {
		return err
	}

	// Create context based on pingfederate authentication type
	authType, err := profiles.GetOptionValue(options.PingFederateAuthenticationTypeOption)
	if err != nil {
		return err
	}

	switch {
	case strings.EqualFold(authType, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC):
		pfUsername, err := profiles.GetOptionValue(options.PingFederateBasicAuthUsernameOption)
		if err != nil {
			return err
		}
		pfPassword, err := profiles.GetOptionValue(options.PingFederateBasicAuthPasswordOption)
		if err != nil {
			return err
		}

		if pfUsername == "" || pfPassword == "" {
			return fmt.Errorf("failed to initialize PingFederate services. Basic authentication username or password is empty")
		}

		pingfederateContext = context.WithValue(ctx, pingfederateGoClient.ContextBasicAuth, pingfederateGoClient.BasicAuth{
			UserName: pfUsername,
			Password: pfPassword,
		})
	case strings.EqualFold(authType, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_ACCESS_TOKEN):
		pfAccessToken, err := profiles.GetOptionValue(options.PingFederateAccessTokenAuthAccessTokenOption)
		if err != nil {
			return err
		}

		if pfAccessToken == "" {
			return fmt.Errorf("failed to initialize PingFederate services. Access token is empty")
		}

		pingfederateContext = context.WithValue(ctx, pingfederateGoClient.ContextAccessToken, pfAccessToken)
	case strings.EqualFold(authType, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS):
		pfClientID, err := profiles.GetOptionValue(options.PingFederateClientCredentialsAuthClientIDOption)
		if err != nil {
			return err
		}
		pfClientSecret, err := profiles.GetOptionValue(options.PingFederateClientCredentialsAuthClientSecretOption)
		if err != nil {
			return err
		}
		pfTokenUrl, err := profiles.GetOptionValue(options.PingFederateClientCredentialsAuthTokenURLOption)
		if err != nil {
			return err
		}
		pfScopes, err := profiles.GetOptionValue(options.PingFederateClientCredentialsAuthScopesOption)
		if err != nil {
			return err
		}

		if pfClientID == "" || pfClientSecret == "" || pfTokenUrl == "" {
			return fmt.Errorf("failed to initialize PingFederate services. Client ID, Client Secret, or Token URL is empty")
		}

		pingfederateContext = context.WithValue(ctx, pingfederateGoClient.ContextOAuth2, pingfederateGoClient.OAuthValues{
			Transport:    tr,
			TokenUrl:     pfTokenUrl,
			ClientId:     pfClientID,
			ClientSecret: pfClientSecret,
			Scopes:       strings.Split(pfScopes, ","),
		})
	default:
		return fmt.Errorf("failed to initialize PingFederate services. unrecognized authentication type '%s'", authType)
	}

	// Test PF API client with create Context Auth
	_, response, err := pingfederateApiClient.VersionAPI.GetVersion(pingfederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetVersion", "pingfederate_client_init")
	if err != nil || !ok {
		return fmt.Errorf("failed to initialize PingFederate Go Client. Check authentication type and credentials")
	}

	return nil
}

func initPingOneServices(ctx context.Context, cmdVersion string) (err error) {
	if err = initPingOneApiClient(ctx, cmdVersion); err != nil {
		return err
	}

	if err = getPingOneExportEnvID(); err != nil {
		return err
	}

	if err := validatePingOneExportEnvID(ctx); err != nil {
		return err
	}

	pingoneContext = ctx

	return nil
}

func initPingFederateApiClient(tr *http.Transport, pingcliVersion string) (err error) {
	l := logger.Get()
	l.Debug().Msgf("Initializing PingFederate API client.")

	if tr == nil {
		return fmt.Errorf("failed to initialize pingfederate API client. http transport is nil")
	}

	httpsHost, err := profiles.GetOptionValue(options.PingFederateHTTPSHostOption)
	if err != nil {
		return err
	}
	adminApiPath, err := profiles.GetOptionValue(options.PingFederateAdminAPIPathOption)
	if err != nil {
		return err
	}
	xBypassExternalValidationHeader, err := profiles.GetOptionValue(options.PingFederateXBypassExternalValidationHeaderOption)
	if err != nil {
		return err
	}

	// default adminApiPath to /pf-admin-api/v1 if not set
	if adminApiPath == "" {
		adminApiPath = "/pf-admin-api/v1"
	}

	if httpsHost == "" {
		return fmt.Errorf(`failed to initialize pingfederate API client. the pingfederate https host configuration value is not set: configure this property via parameter flags, environment variables, or the tool's configuration file (default: $HOME/.pingcli/config.yaml)`)
	}

	userAgent := fmt.Sprintf("pingcli/%s", pingcliVersion)

	if v := strings.TrimSpace(os.Getenv("PINGCLI_PINGFEDERATE_APPEND_USER_AGENT")); v != "" {
		userAgent = fmt.Sprintf("%s %s", userAgent, v)
	}

	pfClientConfig := pingfederateGoClient.NewConfiguration()
	pfClientConfig.UserAgentSuffix = &userAgent
	pfClientConfig.DefaultHeader["X-Xsrf-Header"] = "PingFederate"
	pfClientConfig.DefaultHeader["X-BypassExternalValidation"] = xBypassExternalValidationHeader
	pfClientConfig.Servers = pingfederateGoClient.ServerConfigurations{
		{
			URL: httpsHost + adminApiPath,
		},
	}
	httpClient := &http.Client{Transport: tr}
	pfClientConfig.HTTPClient = httpClient

	pingfederateApiClient = pingfederateGoClient.NewAPIClient(pfClientConfig)

	return nil
}

func initPingOneApiClient(ctx context.Context, pingcliVersion string) (err error) {
	l := logger.Get()
	l.Debug().Msgf("Initializing PingOne API client.")

	if ctx == nil {
		return fmt.Errorf("failed to initialize pingone API client. context is nil")
	}

	pingoneApiClientId, err = profiles.GetOptionValue(options.PingOneAuthenticationWorkerClientIDOption)
	if err != nil {
		return err
	}
	clientSecret, err := profiles.GetOptionValue(options.PingOneAuthenticationWorkerClientSecretOption)
	if err != nil {
		return err
	}
	environmentID, err := profiles.GetOptionValue(options.PingOneAuthenticationWorkerEnvironmentIDOption)
	if err != nil {
		return err
	}
	regionCode, err := profiles.GetOptionValue(options.PingOneRegionCodeOption)
	if err != nil {
		return err
	}

	if pingoneApiClientId == "" || clientSecret == "" || environmentID == "" || regionCode == "" {
		return fmt.Errorf("failed to initialize pingone API client. one of worker client ID, worker client secret, " +
			"pingone region code, and/or worker environment ID is empty. configure these properties via parameter flags, " +
			"environment variables, or the tool's configuration file (default: $HOME/.pingcli/config.yaml)")
	}

	userAgent := fmt.Sprintf("pingcli/%s", pingcliVersion)

	if v := strings.TrimSpace(os.Getenv("PINGCLI_PINGONE_APPEND_USER_AGENT")); v != "" {
		userAgent = fmt.Sprintf("%s %s", userAgent, v)
	}

	enumRegionCode := management.EnumRegionCode(regionCode)

	apiConfig := &pingoneGoClient.Config{
		ClientID:        &pingoneApiClientId,
		ClientSecret:    &clientSecret,
		EnvironmentID:   &environmentID,
		RegionCode:      &enumRegionCode,
		UserAgentSuffix: &userAgent,
	}

	pingoneApiClient, err = apiConfig.APIClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to initialize pingone API client. Check worker client ID, worker client secret,"+
			" worker environment ID, and pingone region code configuration values. %v", err)
	}

	return nil
}

func createOrValidateOutputDir(outputDir string, overwriteExport bool) (resolvedOutputDir string, err error) {
	l := logger.Get()

	// Check if outputDir is empty
	if outputDir == "" {
		return "", fmt.Errorf("failed to export services. The output directory is not set. Specify the output directory "+
			"via the '--%s' flag, '%s' environment variable, or key '%s' in the configuration file",
			options.PlatformExportOutputDirectoryOption.CobraParamName,
			options.PlatformExportOutputDirectoryOption.EnvVar,
			options.PlatformExportOutputDirectoryOption.ViperKey)
	}

	// Check if path is absolute. If not, make it absolute using the present working directory
	if !filepath.IsAbs(outputDir) {
		pwd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("failed to get present working directory: %v", err)
		}

		outputDir = filepath.Join(pwd, outputDir)
	}

	// Check if outputDir exists
	// If not, create the directory
	l.Debug().Msgf("Validating export output directory '%s'", outputDir)
	_, err = os.Stat(outputDir)
	if err != nil {
		output.Warn(fmt.Sprintf("Output directory does not exist. Creating the directory at filepath '%s'", outputDir), nil)

		err = os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("failed to create output directory '%s': %s", outputDir, err.Error())
		}

		output.Success(fmt.Sprintf("Output directory '%s' created", outputDir), nil)
	} else {
		// Check if the output directory is empty
		// If not, default behavior is to exit and not overwrite.
		// This can be changed with the --overwrite export parameter
		if !overwriteExport {
			dirEntries, err := os.ReadDir(outputDir)
			if err != nil {
				return "", fmt.Errorf("failed to read contents of output directory '%s': %v", outputDir, err)
			}

			if len(dirEntries) > 0 {
				return "", fmt.Errorf("output directory '%s' is not empty. Use --overwrite to overwrite existing export data", outputDir)
			}
		}
	}

	return outputDir, nil
}

func getPingOneExportEnvID() (err error) {
	pingoneExportEnvID, err = profiles.GetOptionValue(options.PlatformExportPingOneEnvironmentIDOption)
	if err != nil {
		return err
	}

	if pingoneExportEnvID == "" {
		pingoneExportEnvID, err = profiles.GetOptionValue(options.PingOneAuthenticationWorkerEnvironmentIDOption)
		if err != nil {
			return err
		}
		if pingoneExportEnvID == "" {
			return fmt.Errorf("failed to determine pingone export environment ID")
		}

		output.Warn("No target PingOne export environment ID specified. Defaulting export environment ID to the Worker App environment ID.", nil)
	}

	return nil
}

func validatePingOneExportEnvID(ctx context.Context) (err error) {
	l := logger.Get()
	l.Debug().Msgf("Validating export environment ID...")

	if ctx == nil {
		return fmt.Errorf("failed to validate pingone environment ID '%s'. context is nil", pingoneExportEnvID)
	}

	if pingoneApiClient == nil {
		return fmt.Errorf("failed to validate pingone environment ID '%s'. apiClient is nil", pingoneExportEnvID)
	}

	environment, response, err := pingoneApiClient.ManagementAPIClient.EnvironmentsApi.ReadOneEnvironment(ctx, pingoneExportEnvID).Execute()
	ok, err := common.HandleClientResponse(response, err, "ReadOneEnvironment", "pingone_environment")
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("failed to validate pingone environment ID '%s'", pingoneExportEnvID)
	}

	if environment == nil {
		return fmt.Errorf("failed to validate pingone environment ID '%s'. environment matching ID does not exist", pingoneExportEnvID)
	}

	return nil
}

func getExportableConnectors(exportServices *customtypes.ExportServices) (exportableConnectors *[]connector.Exportable) {
	// Using the --service parameter(s) provided by user, build list of connectors to export
	connectors := []connector.Exportable{}

	if exportServices == nil {
		return &connectors
	}

	for _, service := range exportServices.GetServices() {
		switch service {
		case customtypes.ENUM_EXPORT_SERVICE_PINGONE_PLATFORM:
			connectors = append(connectors, platform.PlatformConnector(pingoneContext, pingoneApiClient, &pingoneApiClientId, pingoneExportEnvID))
		case customtypes.ENUM_EXPORT_SERVICE_PINGONE_AUTHORIZE:
			connectors = append(connectors, authorize.AuthorizeConnector(pingoneContext, pingoneApiClient, &pingoneApiClientId, pingoneExportEnvID))
		case customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO:
			connectors = append(connectors, sso.SSOConnector(pingoneContext, pingoneApiClient, &pingoneApiClientId, pingoneExportEnvID))
		case customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA:
			connectors = append(connectors, mfa.MFAConnector(pingoneContext, pingoneApiClient, &pingoneApiClientId, pingoneExportEnvID))
		case customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT:
			connectors = append(connectors, protect.ProtectConnector(pingoneContext, pingoneApiClient, &pingoneApiClientId, pingoneExportEnvID))
		case customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE:
			connectors = append(connectors, pingfederate.PFConnector(pingfederateContext, pingfederateApiClient))
			// default:
			// This unrecognized service condition is handled by cobra with the custom type MultiService
		}
	}

	return &connectors
}

func exportConnectors(exportableConnectors *[]connector.Exportable, exportFormat, outputDir string, overwriteExport bool) (err error) {
	if exportableConnectors == nil {
		return fmt.Errorf("failed to export services. exportable connectors list is nil")
	}

	// Loop through user defined exportable connectors and export them
	for _, connector := range *exportableConnectors {
		output.Message(fmt.Sprintf("Exporting %s service...", connector.ConnectorServiceName()), nil)

		err := connector.Export(exportFormat, outputDir, overwriteExport)
		if err != nil {
			return fmt.Errorf("failed to export '%s' service: %s", connector.ConnectorServiceName(), err.Error())
		}
	}

	return nil
}
