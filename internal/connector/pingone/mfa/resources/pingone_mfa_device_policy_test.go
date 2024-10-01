package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/mfa/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

func TestMFAPolicyExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.MFADevicePolicy(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_mfa_device_policy",
			ResourceName: "Default MFA Policy",
			ResourceID:   fmt.Sprintf("%s/6adc6dfa-d883-08ed-37c5-ea8f61029ad9", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_mfa_device_policy",
			ResourceName: "Test MFA Policy",
			ResourceID:   fmt.Sprintf("%s/5ae2227f-cb5b-47c3-bb40-440db09a98e6", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
