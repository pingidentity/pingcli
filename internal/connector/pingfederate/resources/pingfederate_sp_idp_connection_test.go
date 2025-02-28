package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateSpIdpConnectionExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.SpIdpConnection(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_sp_idp_connection",
			ResourceName: "testConnection",
			ResourceID:   "n26SCl49a8lB_ifAaLF_MyUbquv",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
