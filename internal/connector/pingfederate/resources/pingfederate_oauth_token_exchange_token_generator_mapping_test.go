package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateOAuthTokenExchangeTokenGeneratorMappingExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OAuthTokenExchangeTokenGeneratorMapping(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_oauth_token_exchange_token_generator_mapping",
			ResourceName: "tokenexchangeprocessorpolicy_to_tokengenerator",
			ResourceID:   "tokenexchangeprocessorpolicy|tokengenerator",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
