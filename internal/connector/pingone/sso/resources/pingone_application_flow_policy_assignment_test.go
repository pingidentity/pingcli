package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestApplicationFlowPolicyAssignmentExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.ApplicationFlowPolicyAssignment(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_application_flow_policy_assignment",
			ResourceName: "Getting Started Application_PingOne - Sign On and Registration",
			ResourceID:   fmt.Sprintf("%s/3da7aae6-92e5-4295-a37c-8515d1f2cd86/0b08c0c3-db40-4be2-aa5b-eb0e17396a75", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_flow_policy_assignment",
			ResourceName: "test app_PingOne - Sign On and Registration",
			ResourceID:   fmt.Sprintf("%s/a4cbf57e-fa2c-452f-bbc8-f40b551da0e2/87a6045e-fa59-41fd-9a06-867ef8cc7a0c", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
