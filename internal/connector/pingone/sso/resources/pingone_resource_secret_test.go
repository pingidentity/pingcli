// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestResourceSecretExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.ResourceSecret(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_resource_secret",
			ResourceName: "Undeployed Test API Service_secret",
			ResourceID:   fmt.Sprintf("%s/a35fe5ea-084c-4245-80f1-85f9eaf4f063", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_secret",
			ResourceName: "authorize-api-service_secret",
			ResourceID:   fmt.Sprintf("%s/3c6001a0-6110-4934-9d34-fa8c4a2894c2", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_secret",
			ResourceName: "test_secret",
			ResourceID:   fmt.Sprintf("%s/4b9ef858-62ce-4bd0-9186-997b8527529d", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_secret",
			ResourceName: "testing_secret",
			ResourceID:   fmt.Sprintf("%s/52afd89f-f3c0-4c78-b896-432c0a07329b", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
