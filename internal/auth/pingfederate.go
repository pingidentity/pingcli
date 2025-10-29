package auth_internal

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

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/profiles"
	pingfederateGoClient "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

// Use this file to SetPingFederateClient instead of pingfederateGoClient.NewAPIClient(pingfederateGoClient.NewConfiguration())

var (
	pingfederateApiClient *pingfederateGoClient.APIClient
)

// GetPingFederateClient returns the cached PingFederate API client instance or creates a new one
func GetPingFederateClient() *pingfederateGoClient.APIClient {
	if pingfederateApiClient == nil {
		pingfederateApiClient = pingfederateGoClient.NewAPIClient(pingfederateGoClient.NewConfiguration())
	}

	return pingfederateApiClient
}

// SetPingFederateClient initializes and configures a PingFederate API client with TLS settings and authentication
func SetPingFederateClient(ctx context.Context, client *pingfederateGoClient.APIClient, pingcliVersion string) (*pingfederateGoClient.APIClient, error) {
	httpsHost, err := profiles.GetOptionValue(options.PingFederateHTTPSHostOption)
	if err != nil {
		return nil, err
	}
	adminApiPath, err := profiles.GetOptionValue(options.PingFederateAdminAPIPathOption)
	if err != nil {
		return nil, err
	}

	if ctx == nil {
		return nil, ErrPingFederateContextNil
	}

	pfInsecureTrustAllTLS, err := profiles.GetOptionValue(options.PingFederateInsecureTrustAllTLSOption)
	if err != nil {
		return nil, err
	}
	caCertPemFiles, err := profiles.GetOptionValue(options.PingFederateCACertificatePemFilesOption)
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	for _, caCertPemFile := range strings.Split(caCertPemFiles, ",") {
		if caCertPemFile == "" {
			continue
		}
		caCertPemFile := filepath.Clean(caCertPemFile)
		caCert, err := os.ReadFile(caCertPemFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read CA certificate PEM file '%s': %w", caCertPemFile, err)
		}

		ok := caCertPool.AppendCertsFromPEM(caCert)
		if !ok {
			return nil, &errs.PingCLIError{
				Prefix: fmt.Sprintf("failed to parse CA certificate PEM file '%s'", caCertPemFile),
				Err:    ErrPingFederateCACertParse,
			}
		}
	}

	pfInsecureTrustAllTLSBool, err := strconv.ParseBool(pfInsecureTrustAllTLS)
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: pfInsecureTrustAllTLSBool, //#nosec G402 -- This is defined by the user (default false), and warned as inappropriate in production.
			RootCAs:            caCertPool,
		},
	}

	userAgent := fmt.Sprintf("pingcli/%s", pingcliVersion)

	if v := strings.TrimSpace(os.Getenv("PINGCLI_PINGFEDERATE_APPEND_USER_AGENT")); v != "" {
		userAgent = fmt.Sprintf("%s %s", userAgent, v)
	}

	xBypassExternalValidationHeader, err := profiles.GetOptionValue(options.PingFederateXBypassExternalValidationHeaderOption)
	if err != nil {
		return nil, err
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

	return pingfederateGoClient.NewAPIClient(pfClientConfig), nil
}
