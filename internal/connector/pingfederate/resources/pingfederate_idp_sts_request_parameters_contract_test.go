package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateIdpStsRequestParametersContractExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.IdpStsRequestParametersContract(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_idp_sts_request_parameters_contract",
			ResourceName: "STS TestName",
			ResourceID:   "STSTestID",
		},
	}
	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
