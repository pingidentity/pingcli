package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/authorize/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestAuthorizeDecisionEndpointExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.AuthorizeDecisionEndpoint(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_authorize_decision_endpoint",
			ResourceName: "DEV",
			ResourceID:   fmt.Sprintf("%s/f8660b46-b96e-457c-8d8f-8ee455e4baa3", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_authorize_decision_endpoint",
			ResourceName: "PROD",
			ResourceID:   fmt.Sprintf("%s/07a4f450-d99f-439f-834a-46b8332a3e31", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_authorize_decision_endpoint",
			ResourceName: "TEST",
			ResourceID:   fmt.Sprintf("%s/3368886d-7d57-4aa8-a8f6-7d24dffa4b3c", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_authorize_decision_endpoint",
			ResourceName: "CLI",
			ResourceID:   fmt.Sprintf("%s/6f4cf36d-fdc1-445c-a1df-37c8e3305eaf", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
