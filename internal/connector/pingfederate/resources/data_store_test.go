package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func Test_PingFederateDataStore(t *testing.T) {
	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.DataStore(pingfederateClientInfo)

	// Data store is already configured in the PingFederate instance

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: "JDBC_ProvisionerDS",
			ResourceID:   "ProvisionerDS",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
