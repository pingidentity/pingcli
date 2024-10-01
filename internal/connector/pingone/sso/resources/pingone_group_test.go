package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

func TestGroupExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.Group(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_group",
			ResourceName: "test group",
			ResourceID:   fmt.Sprintf("%s/ebdf1771-4f43-4fa6-bb9a-ec17333e5ca7", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_group",
			ResourceName: "testing",
			ResourceID:   fmt.Sprintf("%s/b6924f30-73ca-4d3c-964b-90c77adce6a7", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_group",
			ResourceName: "My parent group",
			ResourceID:   fmt.Sprintf("%s/298cf355-6806-4058-b87e-1ae92c7fb13b", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_group",
			ResourceName: "My nested group",
			ResourceID:   fmt.Sprintf("%s/d12ae346-c596-438c-95e3-3d76f364d527", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
