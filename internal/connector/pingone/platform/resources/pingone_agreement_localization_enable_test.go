package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestAgreementLocalizationEnableExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.AgreementLocalizationEnable(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_agreement_localization_enable",
			ResourceName: "Test_fr_enable",
			ResourceID:   fmt.Sprintf("%s/37ab76b8-8eff-43ae-b499-a7dce9fe0e75/03cd7e69-2836-4bad-b69f-249684c42fd9", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_agreement_localization_enable",
			ResourceName: "Test_en_enable",
			ResourceID:   fmt.Sprintf("%s/37ab76b8-8eff-43ae-b499-a7dce9fe0e75/b5ceb6b5-025c-4896-951d-dd676c96d3c6", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
