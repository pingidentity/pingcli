// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func Test_PingFederateOauthCibaServerPolicySettings(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	resource := resources.OauthCibaServerPolicySettings(clientInfo)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: "Oauth Ciba Server Policy Settings",
			ResourceID:   "oauth_ciba_server_policy_settings_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)

}
