// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_platform_testable_resources"
)

// Skipped - depends on TRIAGE-26955
func Test_Form(t *testing.T) {
	t.SkipNow()
	clientInfo := testutils.GetClientInfo(t)

	tr := pingone_platform_testable_resources.Form(t, clientInfo)

	tr.CreateResource(t)
	defer tr.DeleteResource(t)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_NAME],
			ResourceID:   fmt.Sprintf("%s/%s", clientInfo.PingOneExportEnvironmentID, tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID]),
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}
