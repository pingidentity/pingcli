package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateOAuthAccessTokenMappingExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OAuthAccessTokenMapping(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_oauth_access_token_mapping",
			ResourceName: "default|jwt_DEFAULT",
			ResourceID:   "default|jwt",
		},
		{
			ResourceType: "pingfederate_oauth_access_token_mapping",
			ResourceName: "client_credentials|jwt_CLIENT_CREDENTIALS",
			ResourceID:   "client_credentials|jwt",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
