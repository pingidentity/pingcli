package pingfederate_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_terraform"
)

func TestPingFederateTerraformPlan(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)

	testutils_terraform.InitPingFederateTerraform(t)

	testCases := []struct {
		name          string
		resource      connector.ExportableResource
		ignoredErrors []string
	}{
		{
			name:          "PingFederateAuthenticationApiApplication",
			resource:      resources.AuthenticationApiApplication(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateAuthenticationApiSettings",
			resource:      resources.AuthenticationApiSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateAuthenticationPolicies",
			resource:      resources.AuthenticationPolicies(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateAuthenticationPoliciesFragment",
			resource:      resources.AuthenticationPoliciesFragment(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateAuthenticationPoliciesSettings",
			resource:      resources.AuthenticationPoliciesSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateAuthenticationPolicyContract",
			resource:      resources.AuthenticationPolicyContract(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateAuthenticationSelector",
			resource:      resources.AuthenticationSelector(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateCaptchaProvider",
			resource:      resources.CaptchaProvider(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateCaptchaProviderSettings",
			resource:      resources.CaptchaProviderSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateCertificateCA",
			resource: resources.CertificateCA(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Value Length",
			},
		},
		{
			name:     "PingFederateCertificatesRevocationOCSPCertificate",
			resource: resources.CertificatesRevocationOCSPCertificate(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:          "PingFederateCertificatesRevocationSettings",
			resource:      resources.CertificatesRevocationSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateClusterSettings",
			resource: resources.ClusterSettings(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: PingFederate API error",
			},
		},
		{
			name:          "PingFederateConfigurationEncryptionKeysRotate",
			resource:      resources.ConfigurationEncryptionKeysRotate(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateDataStore",
			resource:      resources.DataStore(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateDefaultURLs",
			resource:      resources.DefaultURLs(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateExtendedProperties",
			resource:      resources.ExtendedProperties(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateIdentityStoreProvisioner",
			resource:      resources.IdentityStoreProvisioner(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateIDPAdapter",
			resource:      resources.IDPAdapter(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateIDPSPConnection",
			resource:      resources.IDPSPConnection(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateIdpStsRequestParametersContract",
			resource:      resources.IdpStsRequestParametersContract(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateIdpTokenProcessor",
			resource:      resources.IdpTokenProcessor(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateIdpToSpAdapterMapping",
			resource:      resources.IdpToSpAdapterMapping(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateIncomingProxySettings",
			resource:      resources.IncomingProxySettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateKerberosRealm",
			resource:      resources.KerberosRealm(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateKerberosRealmSettings",
			resource:      resources.KerberosRealmSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateKeypairsOauthOpenidConnect",
			resource:      resources.KeypairsOauthOpenidConnect(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateKeypairsOauthOpenidConnectAdditionalKeySet",
			resource:      resources.KeypairsOauthOpenidConnectAdditionalKeySet(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateKeypairsSigningKeyRotationSettings",
			resource:      resources.KeypairsSigningKeyRotationSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateKeypairsSslServerSettings",
			resource:      resources.KeypairsSslServerSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateLocalIdentityProfile",
			resource:      resources.LocalIdentityProfile(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateMetadataUrl",
			resource:      resources.MetadataUrl(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateNotificationPublisher",
			resource:      resources.NotificationPublisher(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateNotificationPublishersSettings",
			resource:      resources.NotificationPublisherSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOAuthAccessTokenManager",
			resource:      resources.OAuthAccessTokenManager(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOAuthAccessTokenManagerSettings",
			resource:      resources.OAuthAccessTokenManagerSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOAuthAccessTokenMapping",
			resource:      resources.OAuthAccessTokenMapping(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOAuthAuthenticationPolicyContractMapping",
			resource:      resources.OAuthAuthenticationPolicyContractMapping(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOAuthCibaServerPolicyRequestPolicy",
			resource:      resources.OAuthCibaServerPolicyRequestPolicy(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOAuthCIBAServerPolicySettings",
			resource:      resources.OAuthCIBAServerPolicySettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOAuthClient",
			resource:      resources.OAuthClient(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOAuthClientRegistrationPolicy",
			resource:      resources.OAuthClientRegistrationPolicy(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOAuthClientSettings",
			resource:      resources.OAuthClientSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOAuthIdpAdapterMapping",
			resource:      resources.OAuthIdpAdapterMapping(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOAuthIssuer",
			resource:      resources.OAuthIssuer(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOAuthServerSettings",
			resource:      resources.OAuthServerSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOAuthTokenExchangeGeneratorSettings",
			resource:      resources.OAuthTokenExchangeGeneratorSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOAuthTokenExchangeTokenGeneratorMapping",
			resource:      resources.OAuthTokenExchangeTokenGeneratorMapping(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOpenIDConnectPolicy",
			resource:      resources.OpenIDConnectPolicy(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOpenIDConnectSettings",
			resource:      resources.OpenIDConnectSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederatePasswordCredentialValidator",
			resource:      resources.PasswordCredentialValidator(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederatePingOneConnection",
			resource:      resources.PingOneConnection(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateProtocolMetadataLifetimeSettings",
			resource:      resources.ProtocolMetadataLifetimeSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateProtocolMetadataSigningSettings",
			resource:      resources.ProtocolMetadataSigningSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateRedirectValidation",
			resource:      resources.RedirectValidation(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateSecretManager",
			resource:      resources.SecretManager(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateServerSettings",
			resource:      resources.ServerSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateServerSettingsGeneral",
			resource:      resources.ServerSettingsGeneral(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateServerSettingsLogging",
			resource:      resources.ServerSettingsLogging(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateServerSettingsSystemKeysRotate",
			resource:      resources.ServerSettingsSystemKeysRotate(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateServerSettingsWsTrustStsSettings",
			resource:      resources.ServerSettingsWsTrustStsSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateServerSettingsWsTrustStsSettingsIssuerCertificate",
			resource: resources.ServerSettingsWsTrustStsSettingsIssuerCertificate(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:          "PingFederateServiceAuthentication",
			resource:      resources.ServiceAuthentication(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateSessionApplicationPolicy",
			resource:      resources.SessionApplicationPolicy(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateSessionAuthenticationPoliciesGlobal",
			resource:      resources.SessionAuthenticationPoliciesGlobal(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateSessionAuthenticationPolicy",
			resource:      resources.SessionAuthenticationPolicy(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateSessionSettings",
			resource:      resources.SessionSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateSpAdapter",
			resource:      resources.SpAdapter(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateSPAuthenticationPolicyContractMapping",
			resource:      resources.SPAuthenticationPolicyContractMapping(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateSpIdpConnection",
			resource: resources.SpIdpConnection(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Object Attribute Type",
			},
		},
		{
			name:          "PingFederateSpTargetUrlMappings",
			resource:      resources.SpTargetUrlMappings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateTokenProcessorToTokenGeneratorMapping",
			resource:      resources.TokenProcessorToTokenGeneratorMapping(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateVirtualHostNames",
			resource:      resources.VirtualHostNames(PingFederateClientInfo),
			ignoredErrors: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_terraform.ValidateTerraformPlan(t, tc.resource, tc.ignoredErrors)
		})
	}
}
