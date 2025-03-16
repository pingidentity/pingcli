package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func Test_PingFederateOauthAccessTokenManagerSettings(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	resource := resources.OauthAccessTokenManagerSettings(clientInfo)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: "Oauth Access Token Manager Settings",
			ResourceID:   "oauth_access_token_manager_settings_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)

}
