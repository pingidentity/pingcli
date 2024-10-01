package resources_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

func TestPingFederateServerSettingsGeneralExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.ServerSettingsGeneral(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_server_settings_general",
			ResourceName: "Server Settings General",
			ResourceID:   "pingfederate_server_settings_general_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
