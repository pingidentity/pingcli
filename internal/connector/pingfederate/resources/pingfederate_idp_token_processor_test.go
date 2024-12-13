package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateIdpTokenProcessorExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.IdpTokenProcessor(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_idp_token_processor",
			ResourceName: "UsernameTokenProcessor",
			ResourceID:   "UsernameTokenProcessor",
		},
		{
			ResourceType: "pingfederate_idp_token_processor",
			ResourceName: "token processor",
			ResourceID:   "tokenprocessor",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
