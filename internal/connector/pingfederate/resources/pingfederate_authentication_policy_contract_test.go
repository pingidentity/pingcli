package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateAuthenticationPolicyContractExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.AuthenticationPolicyContract(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_authentication_policy_contract",
			ResourceName: "Default",
			ResourceID:   "default",
		},
		{
			ResourceType: "pingfederate_authentication_policy_contract",
			ResourceName: "Fragment - Form",
			ResourceID:   "wIdHhK789PmadmMS",
		},
		{
			ResourceType: "pingfederate_authentication_policy_contract",
			ResourceName: "Fragment - Subject",
			ResourceID:   "DkhZxRcZchsed90U",
		},
		{
			ResourceType: "pingfederate_authentication_policy_contract",
			ResourceName: "Sample Policy Contract",
			ResourceID:   "samplePolicyContract",
		},
		{
			ResourceType: "pingfederate_authentication_policy_contract",
			ResourceName: "apc",
			ResourceID:   "QGxlec5CX693lBQL",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
