package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestResourceExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.Resource(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_resource",
			ResourceName: "authorize-api-service",
			ResourceID:   fmt.Sprintf("%s/3c6001a0-6110-4934-9d34-fa8c4a2894c2", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource",
			ResourceName: "test",
			ResourceID:   fmt.Sprintf("%s/4b9ef858-62ce-4bd0-9186-997b8527529d", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource",
			ResourceName: "testing",
			ResourceID:   fmt.Sprintf("%s/52afd89f-f3c0-4c78-b896-432c0a07329b", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource",
			ResourceName: "PingOne API",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource",
			ResourceName: "openid",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource",
			ResourceName: "Undeployed Test API Service",
			ResourceID:   fmt.Sprintf("%s/a35fe5ea-084c-4245-80f1-85f9eaf4f063", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
