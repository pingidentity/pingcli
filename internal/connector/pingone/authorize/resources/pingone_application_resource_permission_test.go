// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/authorize/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestAuthorizeApplicationResourcePermissionExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.AuthorizeApplicationResourcePermission(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_application_resource_permission",
			ResourceName: "test-permission1:action1",
			ResourceID:   fmt.Sprintf("%s/62b8a221-a530-44f4-ad02-cdb0d3b1395f/080dd732-99ea-4730-a8a6-8da88a232131", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_application_resource_permission",
			ResourceName: "test-permission1:action2",
			ResourceID:   fmt.Sprintf("%s/62b8a221-a530-44f4-ad02-cdb0d3b1395f/05717cf9-3ce4-443a-8154-1986fe984780", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
