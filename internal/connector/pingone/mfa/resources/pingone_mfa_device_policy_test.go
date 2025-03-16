package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/mfa/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestMFAPolicyExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.MFADevicePolicy(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_mfa_device_policy",
			ResourceName: "Default MFA Policy",
			ResourceID:   fmt.Sprintf("%s/6adc6dfa-d883-08ed-37c5-ea8f61029ad9", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_mfa_device_policy",
			ResourceName: "Test MFA Policy",
			ResourceID:   fmt.Sprintf("%s/5ae2227f-cb5b-47c3-bb40-440db09a98e6", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
