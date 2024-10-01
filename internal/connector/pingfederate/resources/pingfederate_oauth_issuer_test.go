package resources_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

func TestPingFederateOAuthIssuerExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OAuthIssuer(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_oauth_issuer",
			ResourceName: "Test Issuer",
			ResourceID:   "BmoJwEmyzs4RSNMzVUlCs8qTPC",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
