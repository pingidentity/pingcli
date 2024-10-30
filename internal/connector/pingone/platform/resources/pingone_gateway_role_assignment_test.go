package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestGatewayRoleAssignmentExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.GatewayRoleAssignment(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_gateway_role_assignment",
			ResourceName: "PF TF Provider_Identity Data Admin_1c5549f9-95f5-4380-b975-d0165aadd9d2",
			ResourceID:   fmt.Sprintf("%s/554257ac-76ca-447a-a210-722343328312/1c5549f9-95f5-4380-b975-d0165aadd9d2", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_gateway_role_assignment",
			ResourceName: "PF TF Provider_Environment Admin_1cf8fca5-f14f-4a64-a521-60efc7891e7e",
			ResourceID:   fmt.Sprintf("%s/554257ac-76ca-447a-a210-722343328312/1cf8fca5-f14f-4a64-a521-60efc7891e7e", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_gateway_role_assignment",
			ResourceName: "Local Test_Identity Data Admin_e424fff4-a8ca-4a75-a169-3376dd2aad96",
			ResourceID:   fmt.Sprintf("%s/5cd3f6b7-35f0-4873-ac64-f32118bf3102/e424fff4-a8ca-4a75-a169-3376dd2aad96", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_gateway_role_assignment",
			ResourceName: "Local Test_Environment Admin_393d4c4e-6642-432d-bc11-1638948d6dd2",
			ResourceID:   fmt.Sprintf("%s/5cd3f6b7-35f0-4873-ac64-f32118bf3102/393d4c4e-6642-432d-bc11-1638948d6dd2", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_gateway_role_assignment",
			ResourceName: "another connection for testing_Identity Data Admin_239579d0-fc0b-4b50-ba03-dfe80e2bb6d0",
			ResourceID:   fmt.Sprintf("%s/8773b833-ade0-4883-9cad-05fe82b23135/239579d0-fc0b-4b50-ba03-dfe80e2bb6d0", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_gateway_role_assignment",
			ResourceName: "another connection for testing_Environment Admin_07ed5801-4d44-4578-9d2f-c6ef6d537e83",
			ResourceID:   fmt.Sprintf("%s/8773b833-ade0-4883-9cad-05fe82b23135/07ed5801-4d44-4578-9d2f-c6ef6d537e83", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
