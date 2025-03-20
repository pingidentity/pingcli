// Copyright © 2025 Ping Identity Corporation

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
	clientInfo := testutils.GetClientInfo(t)

	tr := pingfederate.TestableResource_PingFederateKeypairsSigningKeyRotationSettings(t, clientInfo)

	_ = tr.CreateResource(t)
	defer tr.DeleteResource(t)

	// Test_PingFederateKeypairsSigningKeyRotationSettings only has one dependency
	keypairTr := tr.Dependencies[0]

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s_rotation_settings", keypairTr.CreationInfo[testutils_resource.ENUM_ISSUER_DN], keypairTr.CreationInfo[testutils_resource.ENUM_SERIAL_NUMBER]),
			ResourceID:   keypairTr.CreationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)

}
