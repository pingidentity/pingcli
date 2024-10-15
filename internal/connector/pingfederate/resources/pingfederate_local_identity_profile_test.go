package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateLocalIdentityIdentityProfileExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.LocalIdentityProfile(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_local_identity_profile",
			ResourceName: "Admin Identity Profile",
			ResourceID:   "adminIdentityProfile",
		},
		{
			ResourceType: "pingfederate_local_identity_profile",
			ResourceName: "Registration Identity Profile",
			ResourceID:   "regIdentityProfile",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
