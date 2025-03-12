package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPopulationDefaultIdpExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.PopulationDefaultIdp(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_population_default_identity_provider",
			ResourceName: "Default_default_identity_provider",
			ResourceID:   fmt.Sprintf("%s/720da2ce-4dd0-48d9-af75-aeadbda1860d", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_population_default_identity_provider",
			ResourceName: "LDAP Gateway Population_default_identity_provider",
			ResourceID:   fmt.Sprintf("%s/374fdb3c-4e94-4547-838a-0c200b9a7c70", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_population_default_identity_provider",
			ResourceName: "Test Default Idp Population_default_identity_provider",
			ResourceID:   fmt.Sprintf("%s/2814912d-4a0f-4104-a779-80c13b2a6dcd", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
