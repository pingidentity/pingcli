package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPopulationExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.Population(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_population",
			ResourceName: "Default",
			ResourceID:   fmt.Sprintf("%s/720da2ce-4dd0-48d9-af75-aeadbda1860d", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_population",
			ResourceName: "LDAP Gateway Population",
			ResourceID:   fmt.Sprintf("%s/374fdb3c-4e94-4547-838a-0c200b9a7c70", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_population",
			ResourceName: "Test Default Idp Population",
			ResourceID:   fmt.Sprintf("%s/2814912d-4a0f-4104-a779-80c13b2a6dcd", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
