package resources_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

func TestPingFederateServerSettingsSystemKeysExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.ServerSettingsSystemKeys(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_server_settings_system_keys",
			ResourceName: "Server Settings System Keys",
			ResourceID:   "server_settings_system_keys_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
