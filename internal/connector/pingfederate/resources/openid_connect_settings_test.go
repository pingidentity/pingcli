package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func Test_PingFederateOpenidConnectSettings(t *testing.T) {
	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OpenidConnectSettings(pingfederateClientInfo)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: "Openid Connect Settings",
			ResourceID:   "openid_connect_settings_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
