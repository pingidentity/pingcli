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

func Test_AgreementEnable(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingone_platform_testable_resources.AgreementEnable(t, clientInfo)

	tr.CreateResource(t)
	defer tr.DeleteResource(t)

	agreementTr := tr.Dependencies[0]

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_enable", agreementTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_NAME]),
			ResourceID:   fmt.Sprintf("%s/%s", clientInfo.PingOneExportEnvironmentID, agreementTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID]),
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}
