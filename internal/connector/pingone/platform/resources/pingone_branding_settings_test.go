package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestBrandingSettingsExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.BrandingSettings(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_branding_settings",
			ResourceName: "pingone_branding_settings",
			ResourceID:   clientInfo.PingOneExportEnvironmentID,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
