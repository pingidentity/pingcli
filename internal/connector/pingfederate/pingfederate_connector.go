package pingfederate

import (
	"context"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/logger"
	pingfederateGoClient "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
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
	clientInfo connector.ClientInfo
}

// Utility method for creating a PingFederateConnector
func PFConnector(ctx context.Context, apiClient *pingfederateGoClient.APIClient) *PingFederateConnector {
	return &PingFederateConnector{
		clientInfo: connector.ClientInfo{
			PingFederateApiClient: apiClient,
			PingFederateContext:   ctx,
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
		resources.CertificateCa(&c.clientInfo),
		resources.CertificatesRevocationOcspCertificate(&c.clientInfo),
		resources.CertificatesRevocationSettings(&c.clientInfo),
		resources.ClusterSettings(&c.clientInfo),
		resources.ConfigurationEncryptionKeysRotate(&c.clientInfo),
		resources.DataStore(&c.clientInfo),
		resources.DefaultUrls(&c.clientInfo),
		resources.ExtendedProperties(&c.clientInfo),
		resources.IdentityStoreProvisioner(&c.clientInfo),
		resources.IdpAdapter(&c.clientInfo),
		resources.IdpSpConnection(&c.clientInfo),
		resources.IdpStsRequestParametersContract(&c.clientInfo),
		resources.IdpTokenProcessor(&c.clientInfo),
		resources.IdpToSpAdapterMapping(&c.clientInfo),
		resources.IncomingProxySettings(&c.clientInfo),
		resources.KerberosRealm(&c.clientInfo),
		resources.KerberosRealmSettings(&c.clientInfo),
		resources.KeypairsOauthOpenidConnect(&c.clientInfo),
		resources.KeypairsOauthOpenidConnectAdditionalKeySet(&c.clientInfo),
		resources.KeypairsSigningKeyRotationSettings(&c.clientInfo),
		resources.KeypairsSslServerSettings(&c.clientInfo),
		resources.LocalIdentityProfile(&c.clientInfo),
		resources.MetadataUrl(&c.clientInfo),
		resources.NotificationPublisher(&c.clientInfo),
		resources.NotificationPublisherSettings(&c.clientInfo),
		resources.OauthAccessTokenManager(&c.clientInfo),
		resources.OauthAccessTokenManagerSettings(&c.clientInfo),
		resources.OauthAccessTokenMapping(&c.clientInfo),
		resources.OauthAuthenticationPolicyContractMapping(&c.clientInfo),
		resources.OauthCibaServerPolicyRequestPolicy(&c.clientInfo),
		resources.OauthCibaServerPolicySettings(&c.clientInfo),
		resources.OauthClient(&c.clientInfo),
		resources.OauthClientRegistrationPolicy(&c.clientInfo),
		resources.OauthClientSettings(&c.clientInfo),
		resources.OauthIdpAdapterMapping(&c.clientInfo),
		resources.OauthIssuer(&c.clientInfo),
		resources.OauthServerSettings(&c.clientInfo),
		resources.OauthTokenExchangeGeneratorSettings(&c.clientInfo),
		resources.OauthTokenExchangeTokenGeneratorMapping(&c.clientInfo),
		resources.OpenidConnectPolicy(&c.clientInfo),
		resources.OpenidConnectSettings(&c.clientInfo),
		resources.PasswordCredentialValidator(&c.clientInfo),
		resources.PingoneConnection(&c.clientInfo),
		resources.ProtocolMetadataLifetimeSettings(&c.clientInfo),
		resources.ProtocolMetadataSigningSettings(&c.clientInfo),
		resources.RedirectValidation(&c.clientInfo),
		resources.SecretManager(&c.clientInfo),
		resources.ServerSettings(&c.clientInfo),
		resources.ServerSettingsGeneral(&c.clientInfo),
		resources.ServerSettingsLogging(&c.clientInfo),
		resources.ServerSettingsSystemKeysRotate(&c.clientInfo),
		resources.ServerSettingsWsTrustStsSettings(&c.clientInfo),
		resources.ServerSettingsWsTrustStsSettingsIssuerCertificate(&c.clientInfo),
		resources.ServiceAuthentication(&c.clientInfo),
		resources.SessionApplicationPolicy(&c.clientInfo),
		resources.SessionAuthenticationPoliciesGlobal(&c.clientInfo),
		resources.SessionAuthenticationPolicy(&c.clientInfo),
		resources.SessionSettings(&c.clientInfo),
		resources.SpAdapter(&c.clientInfo),
		resources.SpAuthenticationPolicyContractMapping(&c.clientInfo),
		resources.SpIdpConnection(&c.clientInfo),
		resources.SpTargetUrlMappings(&c.clientInfo),
		resources.TokenProcessorToTokenGeneratorMapping(&c.clientInfo),
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
