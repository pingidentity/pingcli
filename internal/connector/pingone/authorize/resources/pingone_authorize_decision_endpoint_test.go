package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/authorize/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

func TestAuthorizeDecisionEndpointExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.AuthorizeDecisionEndpoint(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_authorize_decision_endpoint",
			ResourceName: "Test Authorize Decision Endpoint",
			ResourceID:   fmt.Sprintf("%s/5ae2227f-cb5b-47c3-bb40-440db09a98e6", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
