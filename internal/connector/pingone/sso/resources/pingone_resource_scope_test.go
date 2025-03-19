// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestResourceScopeExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.ResourceScope(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "authorize-api-service_apiscope",
			ResourceID:   fmt.Sprintf("%s/3c6001a0-6110-4934-9d34-fa8c4a2894c2/97b9c81c-56a3-4727-8626-9c55826f98c0", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "authorize-api-service_testing",
			ResourceID:   fmt.Sprintf("%s/3c6001a0-6110-4934-9d34-fa8c4a2894c2/6aa03c9d-7003-4ddb-9395-b176d4bde6d6", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "test_testing",
			ResourceID:   fmt.Sprintf("%s/4b9ef858-62ce-4bd0-9186-997b8527529d/99bda6e7-f34b-4218-8fb0-221f5414e0db", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "test_oidc",
			ResourceID:   fmt.Sprintf("%s/4b9ef858-62ce-4bd0-9186-997b8527529d/9f2c9b87-a190-446e-bf6b-d97b7f8b1a70", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "testing_test",
			ResourceID:   fmt.Sprintf("%s/52afd89f-f3c0-4c78-b896-432c0a07329b/d9935d01-5baa-4843-970a-9df33b60439f", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
