package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func Test_PingFederateSessionApplicationPolicy(t *testing.T) {
	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.SessionApplicationPolicy(pingfederateClientInfo)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: "Session Application Policy",
			ResourceID:   "session_application_policy_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
