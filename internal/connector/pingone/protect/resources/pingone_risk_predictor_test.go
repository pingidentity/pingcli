package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/protect/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestRiskPredictorExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.RiskPredictor(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_risk_predictor",
			ResourceName: "USER_RISK_BEHAVIOR_User Risk Behavior",
			ResourceID:   fmt.Sprintf("%s/b7a259a3-f762-03df-1c0c-4c558a94e783", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_risk_predictor",
			ResourceName: "VELOCITY_IP Velocity",
			ResourceID:   fmt.Sprintf("%s/eaf75445-0fa4-07f9-31f6-806a6b513b59", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_risk_predictor",
			ResourceName: "VELOCITY_User Velocity",
			ResourceID:   fmt.Sprintf("%s/ab6c4119-90c4-0f07-0708-bf7802182a70", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_risk_predictor",
			ResourceName: "USER_RISK_BEHAVIOR_User-Based Risk Behavior",
			ResourceID:   fmt.Sprintf("%s/f6e64983-2ae1-0b02-2af4-73389ce879fa", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_risk_predictor",
			ResourceName: "DEVICE_New Device",
			ResourceID:   fmt.Sprintf("%s/b5339087-4e8c-08da-0c51-c826d67ca317", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_risk_predictor",
			ResourceName: "ANONYMOUS_NETWORK_Anonymous Network Detection",
			ResourceID:   fmt.Sprintf("%s/1bb86739-b74d-0882-1fbb-4da12fc3afc9", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_risk_predictor",
			ResourceName: "IP_REPUTATION_IP Reputation",
			ResourceID:   fmt.Sprintf("%s/c2753cc6-21b0-0490-3526-de9082c47fac", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_risk_predictor",
			ResourceName: "GEO_VELOCITY_Geovelocity Anomaly",
			ResourceID:   fmt.Sprintf("%s/13c1f49f-0f98-0870-2b3c-cf68549e9a4a", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_risk_predictor",
			ResourceName: "USER_LOCATION_ANOMALY_User Location Anomaly",
			ResourceID:   fmt.Sprintf("%s/6c2a6e1f-f345-07a3-348e-8d820577338f", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_risk_predictor",
			ResourceName: "BOT_Bot Detection",
			ResourceID:   fmt.Sprintf("%s/818db5ee-209f-0371-18cf-f440ecbe3982", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_risk_predictor",
			ResourceName: "DEVICE_Suspicious Device",
			ResourceID:   fmt.Sprintf("%s/dd53d209-c3eb-0980-3acb-0d1d4ecc10d1", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_risk_predictor",
			ResourceName: "ADVERSARY_IN_THE_MIDDLE_Adversary In The Middle",
			ResourceID:   fmt.Sprintf("%s/8e9a0b6f-61a8-06a1-2bdb-fc168000a55d", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_risk_predictor",
			ResourceName: "EMAIL_REPUTATION_Email Reputation",
			ResourceID:   fmt.Sprintf("%s/0a59b68e-e772-0eed-213a-a351f275f418", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
