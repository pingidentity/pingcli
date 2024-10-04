package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateAuthenticationApiApplicationExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.AuthenticationApiApplication(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_authentication_api_application",
			ResourceName: "myauthenticationapiapplication",
			ResourceID:   "myauthenticationapiapplication",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
