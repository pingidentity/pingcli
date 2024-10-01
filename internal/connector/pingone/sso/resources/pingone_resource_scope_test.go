package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

func TestResourceScopeExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.ResourceScope(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "test_testing",
			ResourceID:   fmt.Sprintf("%s/4b9ef858-62ce-4bd0-9186-997b8527529d/99bda6e7-f34b-4218-8fb0-221f5414e0db", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "test_oidc",
			ResourceID:   fmt.Sprintf("%s/4b9ef858-62ce-4bd0-9186-997b8527529d/9f2c9b87-a190-446e-bf6b-d97b7f8b1a70", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "testing_test",
			ResourceID:   fmt.Sprintf("%s/52afd89f-f3c0-4c78-b896-432c0a07329b/d9935d01-5baa-4843-970a-9df33b60439f", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
