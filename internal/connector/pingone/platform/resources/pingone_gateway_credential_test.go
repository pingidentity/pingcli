package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestGatewayCredentialExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.GatewayCredential(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_gateway_credential",
			ResourceName: "random_credential_1",
			ResourceID:   fmt.Sprintf("%s/0b1d882c-5c71-4600-a9fb-befdad921df2/932c1ca6-da29-4a0e-b19c-d012f5b6014f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_gateway_credential",
			ResourceName: "PingFederate LDAP Gateway_credential_1",
			ResourceID:   fmt.Sprintf("%s/3b7b5d9d-1820-4b21-bb29-a5336af65352/fa809636-4796-4a25-8693-2b786eed4f71", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_gateway_credential",
			ResourceName: "PF TF Provider_credential_1",
			ResourceID:   fmt.Sprintf("%s/554257ac-76ca-447a-a210-722343328312/971b5d20-0955-4030-b49b-7e349b3b9b1e", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_gateway_credential",
			ResourceName: "Local Test_credential_1",
			ResourceID:   fmt.Sprintf("%s/5cd3f6b7-35f0-4873-ac64-f32118bf3102/bd2307d8-2a5e-4c11-a397-cfb991179f3f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_gateway_credential",
			ResourceName: "TestGateway_credential_1",
			ResourceID:   fmt.Sprintf("%s/bc37814f-b3a9-4149-b880-0ed457bbb5c5/2e2ab72c-6dcf-4ec2-96be-1a5ba2e66f4a", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_gateway_credential",
			ResourceName: "TestGateway_credential_2",
			ResourceID:   fmt.Sprintf("%s/bc37814f-b3a9-4149-b880-0ed457bbb5c5/5aa73594-66a3-4175-ad69-67fa38b5e307", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_gateway_credential",
			ResourceName: "TestGateway_credential_3",
			ResourceID:   fmt.Sprintf("%s/bc37814f-b3a9-4149-b880-0ed457bbb5c5/ed648842-d109-4a40-97ba-ef4f8ce8eabe", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_gateway_credential",
			ResourceName: "another connection for testing_credential_1",
			ResourceID:   fmt.Sprintf("%s/8773b833-ade0-4883-9cad-05fe82b23135/98f9946c-3a78-4b4b-8645-a425f89c7ab5", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
