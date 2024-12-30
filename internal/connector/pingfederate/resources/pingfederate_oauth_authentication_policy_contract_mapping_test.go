package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateOAuthAuthenticationPolicyContractMappingExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OAuthAuthenticationPolicyContractMapping(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_oauth_authentication_policy_contract_mapping",
			ResourceName: "QGxlec5CX693lBQL_mapping",
			ResourceID:   "QGxlec5CX693lBQL",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
