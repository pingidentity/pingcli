// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestResourceScopePingOneApiExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.ResourceScopePingOneApi(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:read:user",
			ResourceID:   fmt.Sprintf("%s/089adcde-be64-4e7e-9a5a-dda60ce38a9f", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:read:user:2",
			ResourceID:   fmt.Sprintf("%s/83d8ee1d-938f-4287-9792-aa808dc0cad9", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:update:user",
			ResourceID:   fmt.Sprintf("%s/d5bd66de-8044-41c5-aed2-278b6cf47dad", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
