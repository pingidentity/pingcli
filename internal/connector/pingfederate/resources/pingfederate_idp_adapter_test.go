package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateIDPAdapterExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.IDPAdapter(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_idp_adapter",
			ResourceName: "OTIdPJava",
			ResourceID:   "OTIdPJava",
		},
		{
			ResourceType: "pingfederate_idp_adapter",
			ResourceName: "Employee HTML Form",
			ResourceID:   "htmlForm",
		},
		{
			ResourceType: "pingfederate_idp_adapter",
			ResourceName: "Identifier-First",
			ResourceID:   "IDFirst",
		},
		{
			ResourceType: "pingfederate_idp_adapter",
			ResourceName: "Customer HTML Form (PF)",
			ResourceID:   "ciamHtmlForm",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
