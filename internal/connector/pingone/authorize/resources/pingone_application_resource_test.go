package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/authorize/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestApplicationResourceExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.ApplicationResource(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_application_resource",
			ResourceName: "authorize-api-service_test-permission1",
			ResourceID:   fmt.Sprintf("%s/3c6001a0-6110-4934-9d34-fa8c4a2894c2/62b8a221-a530-44f4-ad02-cdb0d3b1395f", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
