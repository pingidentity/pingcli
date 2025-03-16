package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingfederate"
)

func Test_PingFederateTokenProcessorToTokenGeneratorMapping(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingfederate.TestableResource_PingFederateTokenProcessorToTokenGeneratorMapping(t, clientInfo)

	creationInfo := tr.CreateResource(t)
	defer tr.DeleteResource(t)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_to_%s", creationInfo[testutils_resource.ENUM_SOURCE_ID], creationInfo[testutils_resource.ENUM_TARGET_ID]),
			ResourceID:   creationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)

}
