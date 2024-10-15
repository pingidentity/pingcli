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
			name:     "PingFederateAuthenticationPoliciesFragment",
			resource: resources.AuthenticationPoliciesFragment(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Reference to undeclared resource",
			},
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
			name:     "PingFederateCertificateCA",
			resource: resources.CertificateCA(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Value Length",
			},
		},
		{
			name:     "PingFederateDataStore",
			resource: resources.DataStore(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Invalid attribute configuration",
			},
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
			name:     "PingFederateIDPAdapter",
			resource: resources.IDPAdapter(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
				"Error: Reference to undeclared resource",
			},
		},
		{
			name:          "PingFederateIDPSPConnection",
			resource:      resources.IDPSPConnection(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateIncomingProxySettings",
			resource:      resources.IncomingProxySettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateKerberosRealm",
			resource: resources.KerberosRealm(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Invalid attribute configuration",
			},
		},
		{
			name:          "PingFederateLocalIdentityProfile",
			resource:      resources.LocalIdentityProfile(PingFederateClientInfo),
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
			name:          "PingFederateOAuthAccessTokenMapping",
			resource:      resources.OAuthAccessTokenMapping(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateOAuthCIBAServerPolicySettings",
			resource:      resources.OAuthCIBAServerPolicySettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateOAuthClient",
			resource: resources.OAuthClient(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Invalid attribute configuration",
			},
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
			name:     "PingFederatePasswordCredentialValidator",
			resource: resources.PasswordCredentialValidator(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:     "PingFederatePingOneConnection",
			resource: resources.PingOneConnection(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:          "PingFederateRedirectValidation",
			resource:      resources.RedirectValidation(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateServerSettings",
			resource: resources.ServerSettings(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Value Length",
			},
		},
		{
			name:          "PingFederateServerSettingsGeneral",
			resource:      resources.ServerSettingsGeneral(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateServerSettingsSystemKeysRotate",
			resource:      resources.ServerSettingsSystemKeysRotate(PingFederateClientInfo),
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
			name:          "PingFederateSessionSettings",
			resource:      resources.SessionSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateSPAuthenticationPolicyContractMapping",
			resource:      resources.SPAuthenticationPolicyContractMapping(PingFederateClientInfo),
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
