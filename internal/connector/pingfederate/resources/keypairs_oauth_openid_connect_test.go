// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func Test_PingFederateKeypairsOauthOpenidConnect(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	resource := resources.KeypairsOauthOpenidConnect(clientInfo)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: "Keypairs Oauth Openid Connect",
			ResourceID:   "keypairs_oauth_openid_connect_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)

}
