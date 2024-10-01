package resources_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

func TestPingFederateOAuthCIBAServerPolicySettingsExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OAuthCIBAServerPolicySettings(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_oauth_ciba_server_policy_settings",
			ResourceName: "OAuth CIBA Server Policy Settings",
			ResourceID:   "oauth_ciba_server_policy_settings_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
