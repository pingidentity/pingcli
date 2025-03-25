// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_platform_testable_resources"
)

func TestAlertChannelExport(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingone_platform_testable_resources.AlertChannel(t, clientInfo)

	tr.CreateResource(t)
	defer tr.DeleteResource(t)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_NAME],
			ResourceID:   fmt.Sprintf("%s/%s", clientInfo.PingOneExportEnvironmentID, tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID]),
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}
