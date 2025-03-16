package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/mfa/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestMFAFido2PolicyExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.MFAFido2Policy(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_mfa_fido2_policy",
			ResourceName: "Passkeys",
			ResourceID:   fmt.Sprintf("%s/0f9c510a-df48-4d56-9e44-17ac0bc78961", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_mfa_fido2_policy",
			ResourceName: "Security Keys",
			ResourceID:   fmt.Sprintf("%s/f8dc3094-cf9f-486f-9ca9-164e0856b0d8", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
