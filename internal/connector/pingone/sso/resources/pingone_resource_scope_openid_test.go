// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestResourceScopeOpenIdExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.ResourceScopeOpenId(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_resource_scope_openid",
			ResourceName: "openid_profile",
			ResourceID:   fmt.Sprintf("%s/5a2881ba-affc-4556-a9ff-ad662ea84e89", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_scope_openid",
			ResourceName: "openid_newscope2",
			ResourceID:   fmt.Sprintf("%s/5f07b021-5f0e-47d0-a62b-1e983bdff753", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_scope_openid",
			ResourceName: "openid_openid",
			ResourceID:   fmt.Sprintf("%s/6f095311-2cb9-4414-b30f-af8ee5e11e34", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_scope_openid",
			ResourceName: "openid_newscope",
			ResourceID:   fmt.Sprintf("%s/792fa804-8aae-43c8-bea7-ea2dbbb1ca88", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_scope_openid",
			ResourceName: "openid_email",
			ResourceID:   fmt.Sprintf("%s/a95eb903-b691-4aa9-91df-8b02d69816df", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_scope_openid",
			ResourceName: "openid_test",
			ResourceID:   fmt.Sprintf("%s/d4213f0d-e1fc-42db-bcc6-dfad730f7be7", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_scope_openid",
			ResourceName: "openid_phone",
			ResourceID:   fmt.Sprintf("%s/dad64f0c-187e-4991-a5b3-c4e53a4167e5", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_scope_openid",
			ResourceName: "openid_testing",
			ResourceID:   fmt.Sprintf("%s/eb7e9feb-6076-4a2e-9e9e-5c9c0a503606", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_scope_openid",
			ResourceName: "openid_address",
			ResourceID:   fmt.Sprintf("%s/fcd04665-fb97-4943-9c88-427331ebe930", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_scope_openid",
			ResourceName: "openid_offline_access",
			ResourceID:   fmt.Sprintf("%s/4ac3e5ac-7baf-4c05-80a5-01a40c813d1a", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
