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
			name:     "PingFederateCertificateCa",
			resource: resources.CertificateCa(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Value Length",
			},
		},
		{
			name:     "PingFederateCertificatesRevocationOcspCertificate",
			resource: resources.CertificatesRevocationOcspCertificate(PingFederateClientInfo),
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
			name:          "PingFederateDefaultUrls",
			resource:      resources.DefaultUrls(PingFederateClientInfo),
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
			name:          "PingFederateIdpAdapter",
			resource:      resources.IdpAdapter(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateIdpSpConnection",
			resource:      resources.IdpSpConnection(PingFederateClientInfo),
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
			name:          "PingFederateOauthAccessTokenManager",
			resource:      resources.OauthAccessTokenManager(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOauthAccessTokenManagerSettings",
			resource:      resources.OauthAccessTokenManagerSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOauthAccessTokenMapping",
			resource:      resources.OauthAccessTokenMapping(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOauthAuthenticationPolicyContractMapping",
			resource:      resources.OauthAuthenticationPolicyContractMapping(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOauthCibaServerPolicyRequestPolicy",
			resource:      resources.OauthCibaServerPolicyRequestPolicy(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOauthCIBAServerPolicySettings",
			resource:      resources.OauthCibaServerPolicySettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOauthClient",
			resource:      resources.OauthClient(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOauthClientRegistrationPolicy",
			resource:      resources.OauthClientRegistrationPolicy(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOauthClientSettings",
			resource:      resources.OauthClientSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOauthIdpAdapterMapping",
			resource:      resources.OauthIdpAdapterMapping(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOauthIssuer",
			resource:      resources.OauthIssuer(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOauthServerSettings",
			resource:      resources.OauthServerSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOauthTokenExchangeGeneratorSettings",
			resource:      resources.OauthTokenExchangeGeneratorSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOauthTokenExchangeTokenGeneratorMapping",
			resource:      resources.OauthTokenExchangeTokenGeneratorMapping(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOpenidConnectPolicy",
			resource:      resources.OpenidConnectPolicy(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOpenidConnectSettings",
			resource:      resources.OpenidConnectSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederatePasswordCredentialValidator",
			resource:      resources.PasswordCredentialValidator(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederatePingoneConnection",
			resource:      resources.PingoneConnection(PingFederateClientInfo),
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
			name:          "PingFederateSpAuthenticationPolicyContractMapping",
			resource:      resources.SpAuthenticationPolicyContractMapping(PingFederateClientInfo),
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
