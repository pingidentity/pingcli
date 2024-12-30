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
			ResourceName: "CN=common, O=org, C=US_1696532438981_rotation_settings",
			ResourceID:   "419x9yg43rlawqwq9v6az997k",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
