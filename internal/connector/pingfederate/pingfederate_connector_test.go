// Copyright © 2025 Ping Identity Corporation

package pingfederate_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingfederate"
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
			testableResource: pingfederate.TestableResource_PingFederateAuthenticationApiApplication(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateAuthenticationApiSettings",
			testableResource: pingfederate.TestableResource_PingFederateAuthenticationApiSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateAuthenticationPolicies",
			testableResource: pingfederate.TestableResource_PingFederateAuthenticationPolicies(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateAuthenticationPoliciesFragment",
			testableResource: pingfederate.TestableResource_PingFederateAuthenticationPoliciesFragment(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateAuthenticationPoliciesSettings",
			testableResource: pingfederate.TestableResource_PingFederateAuthenticationPoliciesSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateAuthenticationPolicyContract",
			testableResource: pingfederate.TestableResource_PingFederateAuthenticationPolicyContract(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateAuthenticationSelector",
			testableResource: pingfederate.TestableResource_PingFederateAuthenticationSelector(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateCaptchaProvider",
			testableResource: pingfederate.TestableResource_PingFederateCaptchaProvider(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateCaptchaProviderSettings",
			testableResource: pingfederate.TestableResource_PingFederateCaptchaProviderSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateCertificateCa",
			testableResource: pingfederate.TestableResource_PingFederateCertificateCa(t, pingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Value Length",
			},
		},
		{
			name:             "PingFederateCertificatesRevocationOcspCertificate",
			testableResource: pingfederate.TestableResource_PingFederateCertificatesRevocationOcspCertificate(t, pingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:             "PingFederateCertificatesRevocationSettings",
			testableResource: pingfederate.TestableResource_PingFederateCertificatesRevocationSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateClusterSettings",
			testableResource: pingfederate.TestableResource_PingFederateClusterSettings(t, pingFederateClientInfo),
			ignoredErrors: []string{
				"Error: PingFederate API error",
			},
		},
		{
			name:             "PingFederateConfigurationEncryptionKeysRotate",
			testableResource: pingfederate.TestableResource_PingFederateConfigurationEncryptionKeysRotate(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateDataStore",
			testableResource: pingfederate.TestableResource_PingFederateDataStore(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateDefaultUrls",
			testableResource: pingfederate.TestableResource_PingFederateDefaultUrls(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateExtendedProperties",
			testableResource: pingfederate.TestableResource_PingFederateExtendedProperties(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateIdentityStoreProvisioner",
			testableResource: pingfederate.TestableResource_PingFederateIdentityStoreProvisioner(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateIdpAdapter",
			testableResource: pingfederate.TestableResource_PingFederateIdpAdapter(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateIdpSpConnection",
			testableResource: pingfederate.TestableResource_PingFederateIdpSpConnection(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateIdpStsRequestParametersContract",
			testableResource: pingfederate.TestableResource_PingFederateIdpStsRequestParametersContract(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateIdpTokenProcessor",
			testableResource: pingfederate.TestableResource_PingFederateIdpTokenProcessor(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateIdpToSpAdapterMapping",
			testableResource: pingfederate.TestableResource_PingFederateIdpToSpAdapterMapping(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateIncomingProxySettings",
			testableResource: pingfederate.TestableResource_PingFederateIncomingProxySettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateKerberosRealm",
			testableResource: pingfederate.TestableResource_PingFederateKerberosRealm(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateKerberosRealmSettings",
			testableResource: pingfederate.TestableResource_PingFederateKerberosRealmSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateKeypairsOauthOpenidConnect",
			testableResource: pingfederate.TestableResource_PingFederateKeypairsOauthOpenidConnect(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateKeypairsOauthOpenidConnectAdditionalKeySet",
			testableResource: pingfederate.TestableResource_PingFederateKeypairsOauthOpenidConnectAdditionalKeySet(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateKeypairsSigningKeyRotationSettings",
			testableResource: pingfederate.TestableResource_PingFederateKeypairsSigningKeyRotationSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateKeypairsSslServerSettings",
			testableResource: pingfederate.TestableResource_PingFederateKeypairsSslServerSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateLocalIdentityProfile",
			testableResource: pingfederate.TestableResource_PingFederateLocalIdentityProfile(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateMetadataUrl",
			testableResource: pingfederate.TestableResource_PingFederateMetadataUrl(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateNotificationPublisher",
			testableResource: pingfederate.TestableResource_PingFederateNotificationPublisher(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateNotificationPublisherSettings",
			testableResource: pingfederate.TestableResource_PingFederateNotificationPublisherSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthAccessTokenManager",
			testableResource: pingfederate.TestableResource_PingFederateOauthAccessTokenManager(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthAccessTokenManagerSettings",
			testableResource: pingfederate.TestableResource_PingFederateOauthAccessTokenManagerSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthAccessTokenMapping",
			testableResource: pingfederate.TestableResource_PingFederateOauthAccessTokenMapping(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthAuthenticationPolicyContractMapping",
			testableResource: pingfederate.TestableResource_PingFederateOauthAuthenticationPolicyContractMapping(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		// TODO: Re-enable this test after PingFederate OOB Plugin API is triaged
		// {
		// 	name:             "PingFederateOauthCibaServerPolicyRequestPolicy",
		// 	testableResource: pingfederate.TestableResource_PingFederateOauthCibaServerPolicyRequestPolicy(t, pingFederateClientInfo),
		// 	ignoredErrors:    nil,
		// },
		{
			name:             "PingFederateOauthCibaServerPolicySettings",
			testableResource: pingfederate.TestableResource_PingFederateOauthCibaServerPolicySettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthClient",
			testableResource: pingfederate.TestableResource_PingFederateOauthClient(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthClientRegistrationPolicy",
			testableResource: pingfederate.TestableResource_PingFederateOauthClientRegistrationPolicy(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthClientSettings",
			testableResource: pingfederate.TestableResource_PingFederateOauthClientSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthIdpAdapterMapping",
			testableResource: pingfederate.TestableResource_PingFederateOauthIdpAdapterMapping(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthIssuer",
			testableResource: pingfederate.TestableResource_PingFederateOauthIssuer(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthServerSettings",
			testableResource: pingfederate.TestableResource_PingFederateOauthServerSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthTokenExchangeGeneratorSettings",
			testableResource: pingfederate.TestableResource_PingFederateOauthTokenExchangeGeneratorSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOauthTokenExchangeTokenGeneratorMapping",
			testableResource: pingfederate.TestableResource_PingFederateOauthTokenExchangeTokenGeneratorMapping(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOpenidConnectPolicy",
			testableResource: pingfederate.TestableResource_PingFederateOpenidConnectPolicy(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateOpenidConnectSettings",
			testableResource: pingfederate.TestableResource_PingFederateOpenidConnectSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederatePasswordCredentialValidator",
			testableResource: pingfederate.TestableResource_PingFederatePasswordCredentialValidator(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederatePingoneConnection",
			testableResource: pingfederate.TestableResource_PingFederatePingoneConnection(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateProtocolMetadataLifetimeSettings",
			testableResource: pingfederate.TestableResource_PingFederateProtocolMetadataLifetimeSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateProtocolMetadataSigningSettings",
			testableResource: pingfederate.TestableResource_PingFederateProtocolMetadataSigningSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateRedirectValidation",
			testableResource: pingfederate.TestableResource_PingFederateRedirectValidation(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSecretManager",
			testableResource: pingfederate.TestableResource_PingFederateSecretManager(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateServerSettings",
			testableResource: pingfederate.TestableResource_PingFederateServerSettings(t, pingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Value Length",
			},
		},
		{
			name:             "PingFederateServerSettingsGeneral",
			testableResource: pingfederate.TestableResource_PingFederateServerSettingsGeneral(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateServerSettingsLogging",
			testableResource: pingfederate.TestableResource_PingFederateServerSettingsLogging(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateServerSettingsSystemKeysRotate",
			testableResource: pingfederate.TestableResource_PingFederateServerSettingsSystemKeysRotate(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateServerSettingsWsTrustStsSettings",
			testableResource: pingfederate.TestableResource_PingFederateServerSettingsWsTrustStsSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateServerSettingsWsTrustStsSettingsIssuerCertificate",
			testableResource: pingfederate.TestableResource_PingFederateServerSettingsWsTrustStsSettingsIssuerCertificate(t, pingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:             "PingFederateServiceAuthentication",
			testableResource: pingfederate.TestableResource_PingFederateServiceAuthentication(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSessionApplicationPolicy",
			testableResource: pingfederate.TestableResource_PingFederateSessionApplicationPolicy(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSessionAuthenticationPoliciesGlobal",
			testableResource: pingfederate.TestableResource_PingFederateSessionAuthenticationPoliciesGlobal(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSessionAuthenticationPolicy",
			testableResource: pingfederate.TestableResource_PingFederateSessionAuthenticationPolicy(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSessionSettings",
			testableResource: pingfederate.TestableResource_PingFederateSessionSettings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSpAdapter",
			testableResource: pingfederate.TestableResource_PingFederateSpAdapter(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSpAuthenticationPolicyContractMapping",
			testableResource: pingfederate.TestableResource_PingFederateSpAuthenticationPolicyContractMapping(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSpIdpConnection",
			testableResource: pingfederate.TestableResource_PingFederateSpIdpConnection(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateSpTargetUrlMappings",
			testableResource: pingfederate.TestableResource_PingFederateSpTargetUrlMappings(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateTokenProcessorToTokenGeneratorMapping",
			testableResource: pingfederate.TestableResource_PingFederateTokenProcessorToTokenGeneratorMapping(t, pingFederateClientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PingFederateVirtualHostNames",
			testableResource: pingfederate.TestableResource_PingFederateVirtualHostNames(t, pingFederateClientInfo),
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
