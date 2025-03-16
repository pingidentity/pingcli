package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func Test_PingFederateSessionAuthenticationPoliciesGlobal(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	resource := resources.SessionAuthenticationPoliciesGlobal(clientInfo)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: "Session Authentication Policies Global",
			ResourceID:   "session_authentication_policies_global_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)

}
