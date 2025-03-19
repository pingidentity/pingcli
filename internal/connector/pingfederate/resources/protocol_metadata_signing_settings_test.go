// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func Test_PingFederateProtocolMetadataSigningSettings(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	resource := resources.ProtocolMetadataSigningSettings(clientInfo)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: "Protocol Metadata Signing Settings",
			ResourceID:   "protocol_metadata_signing_settings_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)

}
