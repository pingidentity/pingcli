package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateServerSettingsSystemKeysExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.ServerSettingsSystemKeysRotate(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_server_settings_system_keys_rotate",
			ResourceName: "Server Settings System Keys Rotate",
			ResourceID:   "server_settings_system_keys_rotate_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
