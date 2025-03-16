package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestAgreementExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.Agreement(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_agreement",
			ResourceName: "Test",
			ResourceID:   fmt.Sprintf("%s/37ab76b8-8eff-43ae-b499-a7dce9fe0e75", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_agreement",
			ResourceName: "Test2",
			ResourceID:   fmt.Sprintf("%s/38c0c463-b13d-4d22-8da5-f9fd8093d594", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
