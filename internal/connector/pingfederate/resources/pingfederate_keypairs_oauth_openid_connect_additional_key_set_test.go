package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateKeypairsOauthOpenidConnectAdditionalKeySetExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.KeypairsOauthOpenidConnectAdditionalKeySet(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_keypairs_oauth_openid_connect_additional_key_set",
			ResourceName: "testName",
			ResourceID:   "testID",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
