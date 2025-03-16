package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func Test_PingFederateDataStore(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	// Data store resource is already created, so no need to use a testable resource
	resource := resources.DataStore(clientInfo)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: "ProvisionerDS_JDBC",
			ResourceID:   "ProvisionerDS",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)

}
