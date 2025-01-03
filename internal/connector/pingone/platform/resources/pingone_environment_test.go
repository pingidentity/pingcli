package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestEnvironmentExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.Environment(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_environment",
			ResourceName: "pingone_environment",
			ResourceID:   testutils.GetEnvironmentID(),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
