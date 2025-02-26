package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateOAuthCibaServerPolicyRequestPolicyExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OAuthCibaServerPolicyRequestPolicy(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_oauth_ciba_server_policy_request_policy",
			ResourceName: "exampleCibaReqPolicy",
			ResourceID:   "exampleCibaReqPolicy",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
