// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_testable_resources"
)

func Test_TrustedEmailAddress(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingone_testable_resources.TrustedEmailAddress(t, clientInfo)

	// TODO: Currently unable to create a trusted email address via API due to trust email domain verification requirement
	// trustedEmailDomainTr := tr.Dependencies[0]

	// tr.CreateResource(t)
	// defer tr.DeleteResource(t)

	// expectedImportBlocks := []connector.ImportBlock{
	// 	{
	// 		ResourceType: tr.ExportableResource.ResourceType(),
	// 		ResourceName: fmt.Sprintf("%s_%s", trustedEmailDomainTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_NAME], tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_NAME]),
	// 		ResourceID:   fmt.Sprintf("%s/%s/%s", clientInfo.PingOneExportEnvironmentID, trustedEmailDomainTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID], tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID]),
	// 	},
	// }

	expectedImportBlocks := []connector.ImportBlock{}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}
