package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPopulationDefaultExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.PopulationDefault(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_population_default",
			ResourceName: "Default_population_default",
			ResourceID:   clientInfo.PingOneExportEnvironmentID,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
