package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateServerSettingsLoggingExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.ServerSettingsLogging(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_server_settings_logging",
			ResourceName: "Server Settings Logging",
			ResourceID:   "server_settings_logging_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
