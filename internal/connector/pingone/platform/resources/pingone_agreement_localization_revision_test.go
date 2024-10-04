package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestAgreementLocalizationRevisionExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.AgreementLocalizationRevision(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_agreement_localization_revision",
			ResourceName: "Test_fr_1",
			ResourceID:   fmt.Sprintf("%s/37ab76b8-8eff-43ae-b499-a7dce9fe0e75/03cd7e69-2836-4bad-b69f-249684c42fd9/08c49433-b84b-4aeb-860b-02d58336e309", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_agreement_localization_revision",
			ResourceName: "Test_en_1",
			ResourceID:   fmt.Sprintf("%s/37ab76b8-8eff-43ae-b499-a7dce9fe0e75/b5ceb6b5-025c-4896-951d-dd676c96d3c6/2c2a98d7-8c11-4887-b35b-7f5358f75ec1", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_agreement_localization_revision",
			ResourceName: "Test_en_2",
			ResourceID:   fmt.Sprintf("%s/37ab76b8-8eff-43ae-b499-a7dce9fe0e75/b5ceb6b5-025c-4896-951d-dd676c96d3c6/82cff258-1af7-4f2d-94f0-a0ebbbc84a5e", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
