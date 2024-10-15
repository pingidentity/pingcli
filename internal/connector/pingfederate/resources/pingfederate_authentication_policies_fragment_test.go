package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateAuthenticationPoliciesFragmentExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.AuthenticationPoliciesFragment(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_authentication_policies_fragment",
			ResourceName: "Internal AuthN",
			ResourceID:   "InternalAuthN",
		},
		{
			ResourceType: "pingfederate_authentication_policies_fragment",
			ResourceName: "Identify_First",
			ResourceID:   "Identify_First",
		},
		{
			ResourceType: "pingfederate_authentication_policies_fragment",
			ResourceName: "First_Factor",
			ResourceID:   "FirstFactor",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
