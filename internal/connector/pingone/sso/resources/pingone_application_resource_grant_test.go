package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestApplicationResourceGrantExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.ApplicationResourceGrant(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_application_resource_grant",
			ResourceName: "PingOne Application Portal_openid",
			ResourceID:   fmt.Sprintf("%s/92a3765c-e135-4afa-8b12-4469672ac8a9/7e1e25cd-a29e-43b3-bf4a-317ffaabb49c", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_resource_grant",
			ResourceName: "PingOne Application Portal_PingOne API",
			ResourceID:   fmt.Sprintf("%s/92a3765c-e135-4afa-8b12-4469672ac8a9/cf7c2b8e-718c-4ccc-ad1e-1612724baf8e", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_resource_grant",
			ResourceName: "PingOne Self-Service - MyAccount_PingOne API",
			ResourceID:   fmt.Sprintf("%s/4ce54d01-5138-4c56-8175-4f02f69278f5/78d28a77-127d-434b-ae30-71bc18c97902", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_resource_grant",
			ResourceName: "PingOne Self-Service - MyAccount_openid",
			ResourceID:   fmt.Sprintf("%s/4ce54d01-5138-4c56-8175-4f02f69278f5/88063562-7b01-4dbc-b638-119435f74860", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
