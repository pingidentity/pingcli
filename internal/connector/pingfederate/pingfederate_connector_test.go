// Copyright © 2025 Ping Identity Corporation

package pingfederate_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingfederate_testable_resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_terraform"
)

func TestPingFederateTerraformPlan(t *testing.T) {
	pingFederateClientInfo := testutils.GetClientInfo(t)

	testutils_terraform.InitPingFederateTerraform(t)

	testCases := []struct {
		name             string
		testableResource *testutils_resource.TestableResource
		ignoredErrors    []string
	}{
		{
			name:             "PingFederateAuthenticationApiApplication",
			testableResource: pingfederate_testable_resources.AuthenticationApiApplication(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateAuthenticationApiSettings",
			testableResource: pingfederate_testable_resources.AuthenticationApiSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateAuthenticationPolicies",
			testableResource: pingfederate_testable_resources.AuthenticationPolicies(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateAuthenticationPoliciesFragment",
			testableResource: pingfederate_testable_resources.AuthenticationPoliciesFragment(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateAuthenticationPoliciesSettings",
			testableResource: pingfederate_testable_resources.AuthenticationPoliciesSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateAuthenticationPolicyContract",
			testableResource: pingfederate_testable_resources.AuthenticationPolicyContract(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateAuthenticationSelector",
			testableResource: pingfederate_testable_resources.AuthenticationSelector(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateCaptchaProvider",
			testableResource: pingfederate_testable_resources.CaptchaProvider(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateCaptchaProviderSettings",
			testableResource: pingfederate_testable_resources.CaptchaProviderSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateCertificateCa",
			testableResource: pingfederate_testable_resources.CertificateCa(t, pingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Value Length",
			},
		},
		{
			name:             "PingFederateCertificatesRevocationOcspCertificate",
			testableResource: pingfederate_testable_resources.CertificatesRevocationOcspCertificate(t, pingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:             "PingFederateCertificatesRevocationSettings",
			testableResource: pingfederate_testable_resources.CertificatesRevocationSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateClusterSettings",
			testableResource: pingfederate_testable_resources.ClusterSettings(t, pingFederateClientInfo),
			ignoredErrors: []string{
				"Error: PingFederate API error",
			},
		},
		{
			name:             "PingFederateConfigurationEncryptionKeysRotate",
			testableResource: pingfederate_testable_resources.ConfigurationEncryptionKeysRotate(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateDataStore",
			testableResource: pingfederate_testable_resources.DataStore(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateDefaultUrls",
			testableResource: pingfederate_testable_resources.DefaultUrls(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateExtendedProperties",
			testableResource: pingfederate_testable_resources.ExtendedProperties(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateIdentityStoreProvisioner",
			testableResource: pingfederate_testable_resources.IdentityStoreProvisioner(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateIdpAdapter",
			testableResource: pingfederate_testable_resources.IdpAdapter(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateIdpSpConnection",
			testableResource: pingfederate_testable_resources.IdpSpConnection(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateIdpStsRequestParametersContract",
			testableResource: pingfederate_testable_resources.IdpStsRequestParametersContract(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateIdpTokenProcessor",
			testableResource: pingfederate_testable_resources.IdpTokenProcessor(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateIdpToSpAdapterMapping",
			testableResource: pingfederate_testable_resources.IdpToSpAdapterMapping(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateIncomingProxySettings",
			testableResource: pingfederate_testable_resources.IncomingProxySettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateKerberosRealm",
			testableResource: pingfederate_testable_resources.KerberosRealm(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateKerberosRealmSettings",
			testableResource: pingfederate_testable_resources.KerberosRealmSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateKeypairsOauthOpenidConnect",
			testableResource: pingfederate_testable_resources.KeypairsOauthOpenidConnect(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateKeypairsOauthOpenidConnectAdditionalKeySet",
			testableResource: pingfederate_testable_resources.KeypairsOauthOpenidConnectAdditionalKeySet(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateKeypairsSigningKeyRotationSettings",
			testableResource: pingfederate_testable_resources.KeypairsSigningKeyRotationSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateKeypairsSslServerSettings",
			testableResource: pingfederate_testable_resources.KeypairsSslServerSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateLocalIdentityProfile",
			testableResource: pingfederate_testable_resources.LocalIdentityProfile(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateMetadataUrl",
			testableResource: pingfederate_testable_resources.MetadataUrl(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateNotificationPublisher",
			testableResource: pingfederate_testable_resources.NotificationPublisher(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateNotificationPublisherSettings",
			testableResource: pingfederate_testable_resources.NotificationPublisherSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthAccessTokenManager",
			testableResource: pingfederate_testable_resources.OauthAccessTokenManager(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthAccessTokenManagerSettings",
			testableResource: pingfederate_testable_resources.OauthAccessTokenManagerSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthAccessTokenMapping",
			testableResource: pingfederate_testable_resources.OauthAccessTokenMapping(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthAuthenticationPolicyContractMapping",
			testableResource: pingfederate_testable_resources.OauthAuthenticationPolicyContractMapping(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		// TODO: Re-enable this test after PingFederate OOB Plugin API is triaged
		// {
		// 	name:             "PingFederateOauthCibaServerPolicyRequestPolicy",
		// 	testableResource: pingfederate_testable_resources.OauthCibaServerPolicyRequestPolicy(t, pingFederateClientInfo),
		// 	ignoredErrors:    nil,
		// },
		{
			name:             "PingFederateOauthCibaServerPolicySettings",
			testableResource: pingfederate_testable_resources.OauthCibaServerPolicySettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthClient",
			testableResource: pingfederate_testable_resources.OauthClient(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthClientRegistrationPolicy",
			testableResource: pingfederate_testable_resources.OauthClientRegistrationPolicy(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthClientSettings",
			testableResource: pingfederate_testable_resources.OauthClientSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthIdpAdapterMapping",
			testableResource: pingfederate_testable_resources.OauthIdpAdapterMapping(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthIssuer",
			testableResource: pingfederate_testable_resources.OauthIssuer(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthServerSettings",
			testableResource: pingfederate_testable_resources.OauthServerSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthTokenExchangeGeneratorSettings",
			testableResource: pingfederate_testable_resources.OauthTokenExchangeGeneratorSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthTokenExchangeTokenGeneratorMapping",
			testableResource: pingfederate_testable_resources.OauthTokenExchangeTokenGeneratorMapping(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOpenidConnectPolicy",
			testableResource: pingfederate_testable_resources.OpenidConnectPolicy(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOpenidConnectSettings",
			testableResource: pingfederate_testable_resources.OpenidConnectSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederatePasswordCredentialValidator",
			testableResource: pingfederate_testable_resources.PasswordCredentialValidator(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederatePingoneConnection",
			testableResource: pingfederate_testable_resources.PingoneConnection(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateProtocolMetadataLifetimeSettings",
			testableResource: pingfederate_testable_resources.ProtocolMetadataLifetimeSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateProtocolMetadataSigningSettings",
			testableResource: pingfederate_testable_resources.ProtocolMetadataSigningSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateRedirectValidation",
			testableResource: pingfederate_testable_resources.RedirectValidation(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSecretManager",
			testableResource: pingfederate_testable_resources.SecretManager(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateServerSettings",
			testableResource: pingfederate_testable_resources.ServerSettings(t, pingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Value Length",
			},
		},
		{
			name:             "PingFederateServerSettingsGeneral",
			testableResource: pingfederate_testable_resources.ServerSettingsGeneral(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateServerSettingsLogging",
			testableResource: pingfederate_testable_resources.ServerSettingsLogging(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateServerSettingsSystemKeysRotate",
			testableResource: pingfederate_testable_resources.ServerSettingsSystemKeysRotate(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateServerSettingsWsTrustStsSettings",
			testableResource: pingfederate_testable_resources.ServerSettingsWsTrustStsSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateServerSettingsWsTrustStsSettingsIssuerCertificate",
			testableResource: pingfederate_testable_resources.ServerSettingsWsTrustStsSettingsIssuerCertificate(t, pingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:             "PingFederateServiceAuthentication",
			testableResource: pingfederate_testable_resources.ServiceAuthentication(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSessionApplicationPolicy",
			testableResource: pingfederate_testable_resources.SessionApplicationPolicy(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSessionAuthenticationPoliciesGlobal",
			testableResource: pingfederate_testable_resources.SessionAuthenticationPoliciesGlobal(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSessionAuthenticationPolicy",
			testableResource: pingfederate_testable_resources.SessionAuthenticationPolicy(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSessionSettings",
			testableResource: pingfederate_testable_resources.SessionSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSpAdapter",
			testableResource: pingfederate_testable_resources.SpAdapter(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSpAuthenticationPolicyContractMapping",
			testableResource: pingfederate_testable_resources.SpAuthenticationPolicyContractMapping(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSpIdpConnection",
			testableResource: pingfederate_testable_resources.SpIdpConnection(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSpTargetUrlMappings",
			testableResource: pingfederate_testable_resources.SpTargetUrlMappings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateTokenProcessorToTokenGeneratorMapping",
			testableResource: pingfederate_testable_resources.TokenProcessorToTokenGeneratorMapping(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateVirtualHostNames",
			testableResource: pingfederate_testable_resources.VirtualHostNames(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testableResource.CreateResource(t)
			defer tc.testableResource.DeleteResource(t)

			testutils_terraform.ValidateTerraformPlan(t, tc.testableResource.ExportableResource, tc.ignoredErrors)
		})
	}
}
