package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateIdpToSpAdapterMappingExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.IdpToSpAdapterMapping(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_idp_to_sp_adapter_mapping",
			ResourceName: "ciamHtmlForm_to_spadapter",
			ResourceID:   "ciamHtmlForm|spadapter",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
