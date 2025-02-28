package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/authorize/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestAuthorizeApplicationRolePermissionExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.AuthorizeApplicationRolePermission(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_authorize_application_role_permission",
			ResourceName: "test-role_test_permission1action1",
			ResourceID:   fmt.Sprintf("%s/f45cbcc7-2406-470b-93bc-ff477da0b8f7/080dd732-99ea-4730-a8a6-8da88a232131", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_authorize_application_role_permission",
			ResourceName: "test-role_test_permission1action1",
			ResourceID:   fmt.Sprintf("%s/f45cbcc7-2406-470b-93bc-ff477da0b8f7/05717cf9-3ce4-443a-8154-1986fe984780", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
