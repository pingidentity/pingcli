package pingfederate

import (
	"context"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/logger"
	pingfederateGoClient "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

const (
	serviceName = "pingfederate"
)

// Verify that the connector satisfies the expected interfaces
var (
	_ connector.Exportable      = &PingFederateConnector{}
	_ connector.Authenticatable = &PingFederateConnector{}
)

type PingFederateConnector struct {
	clientInfo connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateConnector
func PFConnector(ctx context.Context, apiClient *pingfederateGoClient.APIClient) *PingFederateConnector {
	return &PingFederateConnector{
		clientInfo: connector.PingFederateClientInfo{
			ApiClient: apiClient,
			Context:   ctx,
		},
	}
}

func (c *PingFederateConnector) Export(format, outputDir string, overwriteExport bool) error {
	l := logger.Get()

	l.Debug().Msgf("Exporting all PingFederate Resources...")

	exportableResources := []connector.ExportableResource{
		resources.AuthenticationApiApplication(&c.clientInfo),
		resources.AuthenticationApiSettings(&c.clientInfo),
		resources.AuthenticationPolicies(&c.clientInfo),
		resources.AuthenticationPoliciesFragment(&c.clientInfo),
		resources.AuthenticationPoliciesSettings(&c.clientInfo),
		resources.AuthenticationPolicyContract(&c.clientInfo),
		resources.AuthenticationSelector(&c.clientInfo),
		resources.CaptchaProvider(&c.clientInfo),
		resources.CaptchaProviderSettings(&c.clientInfo),
		resources.CertificateCA(&c.clientInfo),
		resources.CertificatesRevocationOCSPCertificate(&c.clientInfo),
		resources.CertificatesRevocationSettings(&c.clientInfo),
		resources.ClusterSettings(&c.clientInfo),
		resources.ConfigurationEncryptionKeysRotate(&c.clientInfo),
		resources.DataStore(&c.clientInfo),
		resources.DefaultURLs(&c.clientInfo),
		resources.ExtendedProperties(&c.clientInfo),
		resources.IdentityStoreProvisioner(&c.clientInfo),
		resources.IDPAdapter(&c.clientInfo),
		resources.IDPSPConnection(&c.clientInfo),
		resources.IdpStsRequestParametersContract(&c.clientInfo),
		resources.IncomingProxySettings(&c.clientInfo),
		resources.KerberosRealm(&c.clientInfo),
		resources.LocalIdentityProfile(&c.clientInfo),
		resources.NotificationPublisherSettings(&c.clientInfo),
		resources.OAuthAccessTokenManager(&c.clientInfo),
		resources.OAuthAccessTokenMapping(&c.clientInfo),
		resources.OAuthCIBAServerPolicySettings(&c.clientInfo),
		resources.OAuthClient(&c.clientInfo),
		resources.OAuthIssuer(&c.clientInfo),
		resources.OAuthServerSettings(&c.clientInfo),
		resources.OpenIDConnectPolicy(&c.clientInfo),
		resources.OpenIDConnectSettings(&c.clientInfo),
		resources.PasswordCredentialValidator(&c.clientInfo),
		resources.PingOneConnection(&c.clientInfo),
		resources.RedirectValidation(&c.clientInfo),
		resources.ServerSettings(&c.clientInfo),
		resources.ServerSettingsGeneral(&c.clientInfo),
		resources.ServerSettingsSystemKeysRotate(&c.clientInfo),
		resources.SessionApplicationPolicy(&c.clientInfo),
		resources.SessionAuthenticationPoliciesGlobal(&c.clientInfo),
		resources.SessionSettings(&c.clientInfo),
		resources.SPAuthenticationPolicyContractMapping(&c.clientInfo),
		resources.VirtualHostNames(&c.clientInfo),
	}

	return common.WriteFiles(exportableResources, format, outputDir, overwriteExport)
}

func (c *PingFederateConnector) ConnectorServiceName() string {
	return serviceName
}

func (c *PingFederateConnector) Login() error {
	return nil
}

func (c *PingFederateConnector) Logout() error {
	return nil
}
