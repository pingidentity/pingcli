package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederatePingoneConnectionExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.PingoneConnection(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_pingone_connection",
			ResourceName: "internal_brassteam_893438732",
			ResourceID:   "noeOvj5ltBnf4rcmtZAKdJ",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
