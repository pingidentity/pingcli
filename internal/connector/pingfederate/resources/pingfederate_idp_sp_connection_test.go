package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateIDPSPConnectionExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.IDPSPConnection(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_idp_sp_connection",
			ResourceName: "test",
			ResourceID:   "iIoQK.-GWcXI5kLp4KDNxQqAhDF",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
