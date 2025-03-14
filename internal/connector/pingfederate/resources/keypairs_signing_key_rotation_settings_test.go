package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingfederate"
)

func Test_PingFederateKeypairsSigningKeyRotationSettings(t *testing.T) {
	clientInfo := testutils.GetPingFederateClientInfo(t)

	tr := pingfederate.TestableResource_PingFederateKeypairsSigningKeyRotationSettings(t, clientInfo)

	creationInfo := tr.CreateResource(t)
	defer tr.DeleteResource(t)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s_rotation_settings", creationInfo[testutils_resource.ENUM_ISSUER_DN], creationInfo[testutils_resource.ENUM_SERIAL_NUMBER]),
			ResourceID:   creationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)

}
