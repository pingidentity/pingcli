package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateOAuthClientRegistrationPolicyExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OAuthClientRegistrationPolicy(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_oauth_client_registration_policy",
			ResourceName: "Test Registration Policy",
			ResourceID:   "testRegistrationPolicy",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
