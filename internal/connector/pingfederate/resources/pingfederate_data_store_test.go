package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func Test_PingFederateDataStore_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.DataStore(PingFederateClientInfo)

	// Data store already created from data.json

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: "JDBC",
			ResourceID:   "ProvisionerDS",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
