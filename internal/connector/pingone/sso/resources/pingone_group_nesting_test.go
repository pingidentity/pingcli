package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestGroupNestingExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.GroupNesting(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_group_nesting",
			ResourceName: "My parent group_My nested group",
			ResourceID:   fmt.Sprintf("%s/298cf355-6806-4058-b87e-1ae92c7fb13b/d12ae346-c596-438c-95e3-3d76f364d527", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
