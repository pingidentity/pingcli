package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func Test_PingFederateServerSettingsWsTrustStsSettings(t *testing.T) {
	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.ServerSettingsWsTrustStsSettings(pingfederateClientInfo)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: "Server Settings Ws Trust Sts Settings",
			ResourceID:   "server_settings_ws_trust_sts_settings_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
