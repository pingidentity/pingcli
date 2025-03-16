package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/protect/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestRiskPolicyExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.RiskPolicy(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_risk_policy",
			ResourceName: "Default Risk Policy",
			ResourceID:   fmt.Sprintf("%s/f277d6e2-e073-018c-1b78-8be4cd16d898", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_risk_policy",
			ResourceName: "Test Risk Polict",
			ResourceID:   fmt.Sprintf("%s/9964b80b-3140-4d70-9ed5-ff29baf8438f", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
