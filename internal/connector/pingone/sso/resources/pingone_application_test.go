package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestApplicationExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.Application(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_application",
			ResourceName: "Example OAuth App",
			ResourceID:   fmt.Sprintf("%s/2a7c1b5d-415b-4fb5-a6c0-1e290f776785", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application",
			ResourceName: "Getting Started Application",
			ResourceID:   fmt.Sprintf("%s/3da7aae6-92e5-4295-a37c-8515d1f2cd86", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application",
			ResourceName: "OAuth Worker App",
			ResourceID:   fmt.Sprintf("%s/9d6c443b-6329-4d3c-949e-880eda3b9599", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application",
			ResourceName: "PingOne DaVinci Connection",
			ResourceID:   fmt.Sprintf("%s/7b621870-7124-4426-b432-6c675642afcb", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application",
			ResourceName: "test app",
			ResourceID:   fmt.Sprintf("%s/a4cbf57e-fa2c-452f-bbc8-f40b551da0e2", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application",
			ResourceName: "Worker App",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application",
			ResourceName: "Test MFA",
			ResourceID:   fmt.Sprintf("%s/11cfc8c7-ec0c-43ff-b49a-64f5e243f932", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
