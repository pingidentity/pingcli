package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestGatewayExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.Gateway(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_gateway",
			ResourceName: "random",
			ResourceID:   fmt.Sprintf("%s/0b1d882c-5c71-4600-a9fb-befdad921df2", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_gateway",
			ResourceName: "PingFederate LDAP Gateway",
			ResourceID:   fmt.Sprintf("%s/3b7b5d9d-1820-4b21-bb29-a5336af65352", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_gateway",
			ResourceName: "PF TF Provider",
			ResourceID:   fmt.Sprintf("%s/554257ac-76ca-447a-a210-722343328312", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_gateway",
			ResourceName: "Local Test",
			ResourceID:   fmt.Sprintf("%s/5cd3f6b7-35f0-4873-ac64-f32118bf3102", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_gateway",
			ResourceName: "TestGateway",
			ResourceID:   fmt.Sprintf("%s/bc37814f-b3a9-4149-b880-0ed457bbb5c5", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_gateway",
			ResourceName: "another connection for testing",
			ResourceID:   fmt.Sprintf("%s/8773b833-ade0-4883-9cad-05fe82b23135", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
