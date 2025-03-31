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

func Test_ResourceScopeOpenId(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingone_sso_testable_resources.ResourceScopeOpenId(t, clientInfo)

	tr.CreateResource(t)
	defer tr.DeleteResource(t)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s", "openid", tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_NAME]),
			ResourceID:   fmt.Sprintf("%s/%s", clientInfo.PingOneExportEnvironmentID, tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID]),
		},
	}

	// Existing scopes on the openid resource are generated. Test subset.
	testutils.ValidateImportBlockSubset(t, tr.ExportableResource, &expectedImportBlocks)
}
