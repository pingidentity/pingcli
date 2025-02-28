package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/authorize/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestAuthorizeAPIServiceExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.AuthorizeAPIService(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_authorize_api_service",
			ResourceName: "Test API Service",
			ResourceID:   fmt.Sprintf("%s/cee5d5a9-49aa-478d-816e-ec47a2b5aede", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_authorize_api_service",
			ResourceName: "Undeployed Test API Service",
			ResourceID:   fmt.Sprintf("%s/5558f5ab-46b2-40ef-ac78-9a32a07e31c3", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
