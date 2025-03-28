// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingfederate_testable_resources"
)

func Test_PingFederateSessionAuthenticationPolicy(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingfederate_testable_resources.SessionAuthenticationPolicy(t, clientInfo)

	tr.CreateResource(t)
	defer tr.DeleteResource(t)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s_%s", tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID], tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_TYPE], tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_SOURCE_REF_ID]),
			ResourceID:   tr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}
