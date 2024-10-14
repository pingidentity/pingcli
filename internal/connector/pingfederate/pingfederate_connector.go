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
	_ connector.Exportable      = &PingfederateConnector{}
	_ connector.Authenticatable = &PingfederateConnector{}
)

type PingfederateConnector struct {
	clientInfo connector.PingFederateClientInfo
}

// Utility method for creating a PingfederateConnector
func PFConnector(ctx context.Context, apiClient *pingfederateGoClient.APIClient) *PingfederateConnector {
	return &PingfederateConnector{
		clientInfo: connector.PingFederateClientInfo{
			ApiClient: apiClient,
			Context:   ctx,
		},
	}
}

func (c *PingfederateConnector) Export(format, outputDir string, overwriteExport bool) error {
	l := logger.Get()

	l.Debug().Msgf("Exporting all Pingfederate Resources...")

	exportableResources := []connector.ExportableResource{
		resources.AuthenticationApiApplication(&c.clientInfo),
		resources.AuthenticationApiSettings(&c.clientInfo),
		resources.AuthenticationPolicies(&c.clientInfo),
		resources.AuthenticationPoliciesFragment(&c.clientInfo),
		resources.AuthenticationPoliciesSettings(&c.clientInfo),
		resources.AuthenticationPolicyContract(&c.clientInfo),
		resources.AuthenticationSelector(&c.clientInfo),
		resources.CertificateCA(&c.clientInfo),
		resources.DataStore(&c.clientInfo),
		resources.DefaultURLs(&c.clientInfo),
		resources.ExtendedProperties(&c.clientInfo),
		resources.IDPAdapter(&c.clientInfo),
		resources.IDPSPConnection(&c.clientInfo),
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
		resources.PingoneConnection(&c.clientInfo),
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

func (c *PingfederateConnector) ConnectorServiceName() string {
	return serviceName
}

func (c *PingfederateConnector) Login() error {
	return nil
}

func (c *PingfederateConnector) Logout() error {
	return nil
}
