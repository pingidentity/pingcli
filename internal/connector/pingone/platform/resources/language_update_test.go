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

func TestLanguageUpdateExport(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingone_platform_testable_resources.LanguageUpdate(t, clientInfo)

	tr.CreateResource(t)
	defer tr.DeleteResource(t)

	languageTr := tr.Dependencies[0]

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_update", languageTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_NAME]),
			ResourceID:   fmt.Sprintf("%s/%s", clientInfo.PingOneExportEnvironmentID, languageTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID]),
		},
	}

	// There are pre-configured languages in the environment, so only validate the import blocks as a subset.
	testutils.ValidateImportBlockSubset(t, tr.ExportableResource, &expectedImportBlocks)
}
