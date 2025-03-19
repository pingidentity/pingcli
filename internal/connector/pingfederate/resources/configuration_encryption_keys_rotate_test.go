// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func Test_PingFederateConfigurationEncryptionKeysRotate(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	resource := resources.ConfigurationEncryptionKeysRotate(clientInfo)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: "Configuration Encryption Keys Rotate",
			ResourceID:   "configuration_encryption_keys_rotate_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)

}
