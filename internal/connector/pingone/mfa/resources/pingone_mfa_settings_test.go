package resources_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/mfa/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

func TestMFASettingsExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.MFASettings(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_mfa_settings",
			ResourceName: "mfa_settings",
			ResourceID:   testutils.GetEnvironmentID(),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
