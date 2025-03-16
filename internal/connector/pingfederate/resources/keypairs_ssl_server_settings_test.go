package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func Test_PingFederateKeypairsSslServerSettings(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	resource := resources.KeypairsSslServerSettings(clientInfo)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: "Keypairs Ssl Server Settings",
			ResourceID:   "keypairs_ssl_server_settings_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)

}
