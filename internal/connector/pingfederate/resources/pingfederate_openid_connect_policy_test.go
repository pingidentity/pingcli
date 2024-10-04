package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateOpenIDConnectPolicyExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OpenIDConnectPolicy(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_openid_connect_policy",
			ResourceName: "Test OpenID Connect Policy",
			ResourceID:   "test-openid-connect-policy",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
