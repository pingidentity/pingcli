package resources_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

func TestBrandingSettingsExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.BrandingSettings(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_branding_settings",
			ResourceName: "branding_settings",
			ResourceID:   testutils.GetEnvironmentID(),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
