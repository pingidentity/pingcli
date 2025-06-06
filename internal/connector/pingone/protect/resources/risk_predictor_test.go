// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_protect_testable_resources"
)

func Test_RiskPredictor(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingone_protect_testable_resources.RiskPredictor(t, clientInfo)

	tr.CreateResource(t)
	defer tr.DeleteResource(t)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s", tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_TYPE], tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_NAME]),
			ResourceID:   fmt.Sprintf("%s/%s", clientInfo.PingOneExportEnvironmentID, tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID]),
		},
	}

	// Existing risk predictors are generated. test subset.
	testutils.ValidateImportBlockSubset(t, tr.ExportableResource, &expectedImportBlocks)
}
