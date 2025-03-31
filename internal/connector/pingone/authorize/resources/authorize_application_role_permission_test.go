// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_authorize_testable_resources"
)

func Test_AuthorizeApplicationRolePermission(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingone_authorize_testable_resources.AuthorizeApplicationRolePermission(t, clientInfo)

	tr.CreateResource(t)
	defer tr.DeleteResource(t)

	applicationRoleTr := tr.Dependencies[0]

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s", applicationRoleTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_NAME], tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_NAME]),
			ResourceID:   fmt.Sprintf("%s/%s/%s", clientInfo.PingOneExportEnvironmentID, applicationRoleTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID], tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID]),
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}
