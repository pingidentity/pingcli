package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateKeypairsSigningKeyRotationSettingsExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.KeypairsSigningKeyRotationSettings(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_keypairs_signing_key_rotation_settings",
			ResourceName: "CN=rotationTest, O=pingidentity, L=Denver, ST=CO, C=US_1735851845119_rotation_settings",
			ResourceID:   "9vgmnd36wykte1l2nm8s8uead",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
