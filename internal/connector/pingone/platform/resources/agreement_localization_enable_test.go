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

func Test_AgreementLocalizationEnable(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingone_platform_testable_resources.AgreementLocalizationEnable(t, clientInfo)

	tr.CreateResource(t)
	defer tr.DeleteResource(t)

	agreementLocalizationTr := tr.Dependencies[0]
	agreementTr := agreementLocalizationTr.Dependencies[0]

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s_enable", agreementTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_NAME], agreementLocalizationTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_LOCALE]),
			ResourceID:   fmt.Sprintf("%s/%s/%s", clientInfo.PingOneExportEnvironmentID, agreementTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID], agreementLocalizationTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID]),
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}
