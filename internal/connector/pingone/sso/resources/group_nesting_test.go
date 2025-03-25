// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_sso_testable_resources"
)

func Test_GroupNesting(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingone_sso_testable_resources.GroupNesting(t, clientInfo)

	tr.CreateResource(t)
	defer tr.DeleteResource(t)

	groupTr := tr.Dependencies[0]
	nestedGroupTr := tr.Dependencies[1]

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s", groupTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_NAME], nestedGroupTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_NAME]),
			ResourceID:   fmt.Sprintf("%s/%s/%s", clientInfo.PingOneExportEnvironmentID, groupTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID], tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID]),
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}
