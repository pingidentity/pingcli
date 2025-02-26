package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateMetadataUrlExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.MetadataUrl(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_metadata_url",
			ResourceName: "Test Metadata URL",
			ResourceID:   "i8uUHFDebYX7Z7gSfyhZ9yKUA",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
