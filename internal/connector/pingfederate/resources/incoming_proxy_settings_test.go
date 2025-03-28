// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingfederate_testable_resources"
)

func Test_PingFederateIncomingProxySettings(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingfederate_testable_resources.IncomingProxySettings(t, clientInfo)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: "Incoming Proxy Settings",
			ResourceID:   "incoming_proxy_settings_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}
