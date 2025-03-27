// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingfederate_testable_resources"
)

func Test_PingFederateIdpAdapter(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingfederate_testable_resources.IdpAdapter(t, clientInfo)

	tr.CreateResource(t)
	defer tr.DeleteResource(t)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_NAME],
			ResourceID:   tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}
