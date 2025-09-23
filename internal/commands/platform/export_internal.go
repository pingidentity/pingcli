// Copyright © 2025 Ping Identity Corporation

package platform_internal

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
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
	"github.com/pingidentity/pingcli/internal/errs"
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

var (
	exportErrorPrefix           = "failed to export service(s)"
	ErrNilContext               = errors.New("context is nil")
	ErrReadCaCertPemFile        = errors.New("failed to read CA certificate PEM file")
	ErrAppendToCertPool         = errors.New("failed to append to certificate pool from PEM file")
	ErrBasicAuthEmpty           = errors.New("failed to initialize PingFederate service. Basic authentication username and/or password is not set")
	ErrAccessTokenEmpty         = errors.New("failed to initialize PingFederate service. Access token is not set")
	ErrClientCredentialsEmpty   = errors.New("failed to initialize PingFederate service. Client ID, Client Secret, and/or Token URL is not set")
	ErrPingFederateAuthType     = errors.New("failed to initialize PingFederate service. Unrecognized authentication type")
	ErrPingFederateInit         = errors.New("failed to initialize PingFederate service. Check authentication type and credentials")
	ErrHttpTransportNil         = errors.New("failed to initialize PingFederate service. http transport is nil")
	ErrHttpsHostEmpty           = errors.New("failed to initialize PingFederate service. HTTPS host is not set")
	ErrPingOneConfigValuesEmpty = errors.New("failed to initialize pingone API client. one of worker client ID, worker client secret, " +
		"pingone region code, and/or worker environment ID is not set. configure these properties via parameter flags, " +
		"environment variables, or the tool's configuration file (default: $HOME/.pingcli/config.yaml)")
	ErrPingOneInit = errors.New("failed to initialize pingone API client. Check worker client ID, worker client secret," +
		" worker environment ID, and pingone region code configuration values")
	ErrOutputDirectoryEmpty       = errors.New("output directory is not set")
	ErrGetPresentWorkingDirectory = errors.New("failed to get present working directory")
	ErrCreateOutputDirectory      = errors.New("failed to create output directory")
	ErrReadOutputDirectory        = errors.New("failed to read contents of output directory")
	ErrOutputDirectoryNotEmpty    = errors.New("output directory is not empty. use '--overwrite' to overwrite existing files and export data")
	ErrDeterminePingOneExportEnv  = errors.New("failed to determine pingone export environment ID")
	ErrPingOneClientNil           = errors.New("pingone API client is nil")
	ErrValidatePingOneEnvId       = errors.New("failed to validate pingone environment ID")
	ErrPingOneEnvNotExist         = errors.New("pingone environment does not exist")
	ErrConnectorListNil           = errors.New("exportable connectors list is nil")
	ErrExportService              = errors.New("failed to export service")
)

func RunInternalExport(ctx context.Context, commandVersion string) (err error) {
	if ctx == nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: ErrNilContext}
	}

	exportFormat, err := profiles.GetOptionValue(options.PlatformExportExportFormatOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}
	exportServiceGroup, err := profiles.GetOptionValue(options.PlatformExportServiceGroupOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}
	exportServices, err := profiles.GetOptionValue(options.PlatformExportServiceOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}
	outputDir, err := profiles.GetOptionValue(options.PlatformExportOutputDirectoryOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}
	overwriteExport, err := profiles.GetOptionValue(options.PlatformExportOverwriteOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	var exportableConnectors *[]connector.Exportable
	es := new(customtypes.ExportServices)
	if err = es.Set(exportServices); err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	esg := new(customtypes.ExportServiceGroup)
	if err = esg.Set(exportServiceGroup); err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	es2 := new(customtypes.ExportServices)
	if err = es2.SetServicesByServiceGroup(esg); err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	if err = es.Merge(*es2); err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	if es.ContainsPingOneService() {
		if err = initPingOneServices(ctx, commandVersion); err != nil {
			return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
		}
	}

	if es.ContainsPingFederateService() {
		if err = initPingFederateServices(ctx, commandVersion); err != nil {
			return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
		}
	}

	exportableConnectors = getExportableConnectors(es)

	overwriteExportBool, err := strconv.ParseBool(overwriteExport)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}
	if outputDir, err = createOrValidateOutputDir(outputDir, overwriteExportBool); err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	if err := exportConnectors(exportableConnectors, exportFormat, outputDir, overwriteExportBool); err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	output.Success(fmt.Sprintf("Export to directory '%s' complete.", outputDir), nil)

	return nil
}

func initPingFederateServices(ctx context.Context, pingcliVersion string) (err error) {
	if ctx == nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: ErrNilContext}
	}

	pfInsecureTrustAllTLS, err := profiles.GetOptionValue(options.PingFederateInsecureTrustAllTLSOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}
	caCertPemFiles, err := profiles.GetOptionValue(options.PingFederateCACertificatePemFilesOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	caCertPool := x509.NewCertPool()
	for _, caCertPemFile := range strings.Split(caCertPemFiles, ",") {
		if caCertPemFile == "" {
			continue
		}
		caCertPemFile := filepath.Clean(caCertPemFile)
		caCert, err := os.ReadFile(caCertPemFile)
		if err != nil {
			return &errs.PingCLIError{
				Prefix: exportErrorPrefix,
				Err:    fmt.Errorf("%w '%s': %v", ErrReadCaCertPemFile, caCertPemFile, err),
			}
		}

		ok := caCertPool.AppendCertsFromPEM(caCert)
		if !ok {
			return &errs.PingCLIError{
				Prefix: exportErrorPrefix,
				Err:    fmt.Errorf("%w '%s': %v", ErrAppendToCertPool, caCertPemFile, err),
			}
		}
	}

	pfInsecureTrustAllTLSBool, err := strconv.ParseBool(pfInsecureTrustAllTLS)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: pfInsecureTrustAllTLSBool, //#nosec G402 -- This is defined by the user (default false), and warned as inappropriate in production.
			RootCAs:            caCertPool,
		},
	}

	if err = initPingFederateApiClient(tr, pingcliVersion); err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	// Create context based on pingfederate authentication type
	authType, err := profiles.GetOptionValue(options.PingFederateAuthenticationTypeOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	switch {
	case strings.EqualFold(authType, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC):
		pfUsername, err := profiles.GetOptionValue(options.PingFederateBasicAuthUsernameOption)
		if err != nil {
			return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
		}
		pfPassword, err := profiles.GetOptionValue(options.PingFederateBasicAuthPasswordOption)
		if err != nil {
			return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
		}

		if pfUsername == "" || pfPassword == "" {
			return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: ErrBasicAuthEmpty}
		}

		pingfederateContext = context.WithValue(ctx, pingfederateGoClient.ContextBasicAuth, pingfederateGoClient.BasicAuth{
			UserName: pfUsername,
			Password: pfPassword,
		})
	case strings.EqualFold(authType, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_ACCESS_TOKEN):
		pfAccessToken, err := profiles.GetOptionValue(options.PingFederateAccessTokenAuthAccessTokenOption)
		if err != nil {
			return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
		}

		if pfAccessToken == "" {
			return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: ErrAccessTokenEmpty}
		}

		pingfederateContext = context.WithValue(ctx, pingfederateGoClient.ContextAccessToken, pfAccessToken)
	case strings.EqualFold(authType, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS):
		pfClientID, err := profiles.GetOptionValue(options.PingFederateClientCredentialsAuthClientIDOption)
		if err != nil {
			return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
		}
		pfClientSecret, err := profiles.GetOptionValue(options.PingFederateClientCredentialsAuthClientSecretOption)
		if err != nil {
			return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
		}
		pfTokenUrl, err := profiles.GetOptionValue(options.PingFederateClientCredentialsAuthTokenURLOption)
		if err != nil {
			return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
		}
		pfScopes, err := profiles.GetOptionValue(options.PingFederateClientCredentialsAuthScopesOption)
		if err != nil {
			return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
		}

		if pfClientID == "" || pfClientSecret == "" || pfTokenUrl == "" {
			return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: ErrClientCredentialsEmpty}
		}

		pingfederateContext = context.WithValue(ctx, pingfederateGoClient.ContextOAuth2, pingfederateGoClient.OAuthValues{
			Transport:    tr,
			TokenUrl:     pfTokenUrl,
			ClientId:     pfClientID,
			ClientSecret: pfClientSecret,
			Scopes:       strings.Split(pfScopes, ","),
		})
	default:
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: fmt.Errorf("%w '%s'", ErrPingFederateAuthType, authType)}
	}

	// Test PF API client with create Context Auth
	_, response, err := pingfederateApiClient.VersionAPI.GetVersion(pingfederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetVersion", "pingfederate_client_init")
	if err != nil || !ok {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: ErrPingFederateInit}
	}

	return nil
}

func initPingOneServices(ctx context.Context, cmdVersion string) (err error) {
	if err = initPingOneApiClient(ctx, cmdVersion); err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	if err = getPingOneExportEnvID(); err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	if err := validatePingOneExportEnvID(ctx); err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	pingoneContext = ctx

	return nil
}

func initPingFederateApiClient(tr *http.Transport, pingcliVersion string) (err error) {
	l := logger.Get()
	l.Debug().Msgf("Initializing PingFederate API client.")

	if tr == nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: ErrHttpTransportNil}
	}

	httpsHost, err := profiles.GetOptionValue(options.PingFederateHTTPSHostOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}
	adminApiPath, err := profiles.GetOptionValue(options.PingFederateAdminAPIPathOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}
	xBypassExternalValidationHeader, err := profiles.GetOptionValue(options.PingFederateXBypassExternalValidationHeaderOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	// default adminApiPath to /pf-admin-api/v1 if not set
	if adminApiPath == "" {
		adminApiPath = "/pf-admin-api/v1"
	}

	if httpsHost == "" {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: ErrHttpsHostEmpty}
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
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: ErrNilContext}
	}

	pingoneApiClientId, err = profiles.GetOptionValue(options.PingOneAuthenticationWorkerClientIDOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}
	clientSecret, err := profiles.GetOptionValue(options.PingOneAuthenticationWorkerClientSecretOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}
	environmentID, err := profiles.GetOptionValue(options.PingOneAuthenticationWorkerEnvironmentIDOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}
	regionCode, err := profiles.GetOptionValue(options.PingOneRegionCodeOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	if pingoneApiClientId == "" || clientSecret == "" || environmentID == "" || regionCode == "" {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: ErrPingOneConfigValuesEmpty}
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
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: fmt.Errorf("%w: %v", ErrPingOneInit, err)}
	}

	return nil
}

func createOrValidateOutputDir(outputDir string, overwriteExport bool) (resolvedOutputDir string, err error) {
	l := logger.Get()

	// Check if outputDir is empty
	if outputDir == "" {
		return resolvedOutputDir, &errs.PingCLIError{Prefix: exportErrorPrefix, Err: ErrOutputDirectoryEmpty}
	}

	// Check if path is absolute. If not, make it absolute using the present working directory
	if !filepath.IsAbs(outputDir) {
		pwd, err := os.Getwd()
		if err != nil {
			return resolvedOutputDir, &errs.PingCLIError{Prefix: exportErrorPrefix, Err: fmt.Errorf("%w: %v", ErrGetPresentWorkingDirectory, err)}
		}

		outputDir = filepath.Join(pwd, outputDir)
	}

	// Check if outputDir exists
	// If not, create the directory
	l.Debug().Msgf("Validating export output directory '%s'", outputDir)
	_, err = os.Stat(outputDir)
	if err != nil {
		output.Message(fmt.Sprintf("Output directory does not exist. Creating the directory at filepath '%s'", outputDir), nil)

		err = os.MkdirAll(outputDir, os.FileMode(0700))
		if err != nil {
			return resolvedOutputDir, &errs.PingCLIError{Prefix: exportErrorPrefix, Err: fmt.Errorf("%w '%s': %v", ErrCreateOutputDirectory, outputDir, err)}
		}

		output.Success(fmt.Sprintf("Output directory '%s' created", outputDir), nil)
	} else if !overwriteExport {
		// Check if the output directory is empty
		// If not, default behavior is to exit and not overwrite.
		// This can be changed with the --overwrite export parameter
		dirEntries, err := os.ReadDir(outputDir)
		if err != nil {
			return resolvedOutputDir, &errs.PingCLIError{Prefix: exportErrorPrefix, Err: fmt.Errorf("%w '%s': %v", ErrReadOutputDirectory, outputDir, err)}
		}

		if len(dirEntries) > 0 {
			return resolvedOutputDir, &errs.PingCLIError{Prefix: exportErrorPrefix, Err: ErrOutputDirectoryNotEmpty}
		}
	}

	return outputDir, nil
}

func getPingOneExportEnvID() (err error) {
	pingoneExportEnvID, err = profiles.GetOptionValue(options.PlatformExportPingOneEnvironmentIDOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}

	if pingoneExportEnvID == "" {
		pingoneExportEnvID, err = profiles.GetOptionValue(options.PingOneAuthenticationWorkerEnvironmentIDOption)
		if err != nil {
			return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
		}
		if pingoneExportEnvID == "" {
			return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: ErrDeterminePingOneExportEnv}
		}

		output.Message("No target PingOne export environment ID specified. Defaulting export environment ID to the Worker App environment ID.", nil)
	}

	return nil
}

func validatePingOneExportEnvID(ctx context.Context) (err error) {
	l := logger.Get()
	l.Debug().Msgf("Validating export environment ID...")

	if ctx == nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: fmt.Errorf("failed to validate pingone environment ID '%s'. %w", pingoneExportEnvID, ErrNilContext)}
	}

	if pingoneApiClient == nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: fmt.Errorf("failed to validate pingone environment ID '%s'. %w", pingoneExportEnvID, ErrPingOneClientNil)}
	}

	environment, response, err := pingoneApiClient.ManagementAPIClient.EnvironmentsApi.ReadOneEnvironment(ctx, pingoneExportEnvID).Execute()
	ok, err := common.HandleClientResponse(response, err, "ReadOneEnvironment", "pingone_environment")
	if err != nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: err}
	}
	if !ok {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: fmt.Errorf("%w '%s'", ErrValidatePingOneEnvId, pingoneExportEnvID)}
	}

	if environment == nil {
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: fmt.Errorf("%w: '%s'", ErrPingOneEnvNotExist, pingoneExportEnvID)}
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
		return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: ErrConnectorListNil}
	}

	// Loop through user defined exportable connectors and export them
	for _, connector := range *exportableConnectors {
		output.Message(fmt.Sprintf("Exporting %s service...", connector.ConnectorServiceName()), nil)

		err := connector.Export(exportFormat, outputDir, overwriteExport)
		if err != nil {
			return &errs.PingCLIError{Prefix: exportErrorPrefix, Err: fmt.Errorf("%w '%s': %v", ErrExportService, connector.ConnectorServiceName(), err)}
		}
	}

	return nil
}
