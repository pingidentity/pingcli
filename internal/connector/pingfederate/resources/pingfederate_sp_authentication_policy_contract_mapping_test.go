package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateSPAuthenticationPolicyContractMappingExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.SPAuthenticationPolicyContractMapping(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_sp_authentication_policy_contract_mapping",
			ResourceName: "wIdHhK789PmadmMS_to_spadapter",
			ResourceID:   "wIdHhK789PmadmMS|spadapter",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
