package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateTokenProcessorToTokenGeneratorMappingExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.TokenProcessorToTokenGeneratorMapping(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_token_processor_to_token_generator_mapping",
			ResourceName: "tokenprocessor_to_tokengenerator",
			ResourceID:   "tokenprocessor|tokengenerator",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
