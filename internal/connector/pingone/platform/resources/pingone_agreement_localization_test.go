package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestAgreementLocalizationExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.AgreementLocalization(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_agreement_localization",
			ResourceName: "Test_fr",
			ResourceID:   fmt.Sprintf("%s/37ab76b8-8eff-43ae-b499-a7dce9fe0e75/03cd7e69-2836-4bad-b69f-249684c42fd9", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_agreement_localization",
			ResourceName: "Test_en",
			ResourceID:   fmt.Sprintf("%s/37ab76b8-8eff-43ae-b499-a7dce9fe0e75/b5ceb6b5-025c-4896-951d-dd676c96d3c6", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
