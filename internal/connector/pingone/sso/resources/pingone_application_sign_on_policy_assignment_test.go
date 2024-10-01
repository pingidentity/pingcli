package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

func TestApplicationSignOnPolicyAssignmentExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.ApplicationSignOnPolicyAssignment(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_application_sign_on_policy_assignment",
			ResourceName: "Example OAuth App_Single_Factor",
			ResourceID:   fmt.Sprintf("%s/2a7c1b5d-415b-4fb5-a6c0-1e290f776785/056ed696-f2e9-44b1-8d2c-68e690cd1f24", testutils.GetEnvironmentID()),
		},

		{
			ResourceType: "pingone_application_sign_on_policy_assignment",
			ResourceName: "Test MFA_multi_factor",
			ResourceID:   fmt.Sprintf("%s/11cfc8c7-ec0c-43ff-b49a-64f5e243f932/b0ecdaab-9d7c-4c1f-ab0d-891cfdbc73b2", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
