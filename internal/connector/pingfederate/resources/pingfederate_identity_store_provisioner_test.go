package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateIdentityStoreProvisionerExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.IdentityStoreProvisioner(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_identity_store_provisioner",
			ResourceName: "ISP TestName",
			ResourceID:   "ISPTestID",
		},
	}
	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
