package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateDataStoreExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.DataStore(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_data_store",
			ResourceName: "ProvisionerDS_JDBC",
			ResourceID:   "ProvisionerDS",
		},
		{
			ResourceType: "pingfederate_data_store",
			ResourceName: "LDAP-PingDirectory_LDAP",
			ResourceID:   "LDAP-PingDirectory",
		},
		{
			ResourceType: "pingfederate_data_store",
			ResourceName: "pingdirectory_LDAP",
			ResourceID:   "pingdirectory",
		},
	}
	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
