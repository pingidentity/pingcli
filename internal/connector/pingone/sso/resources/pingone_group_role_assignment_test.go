package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestGroupRoleAssignmentExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.GroupRoleAssignment(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_group_role_assignment",
			ResourceName: "testing_Client Application Developer_1db1accc-f63f-4f03-ab62-c767398fa730",
			ResourceID:   fmt.Sprintf("%s/b6924f30-73ca-4d3c-964b-90c77adce6a7/1db1accc-f63f-4f03-ab62-c767398fa730", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_group_role_assignment",
			ResourceName: "testing_Identity Data Read Only_53a88921-2a9f-44f1-958e-3db9be3f8c69",
			ResourceID:   fmt.Sprintf("%s/b6924f30-73ca-4d3c-964b-90c77adce6a7/53a88921-2a9f-44f1-958e-3db9be3f8c69", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
