package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/mfa/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestMFAApplicationPushCredentialExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.MFAApplicationPushCredential(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_mfa_application_push_credential",
			ResourceName: "Test MFA_APNS",
			ResourceID:   fmt.Sprintf("%s/11cfc8c7-ec0c-43ff-b49a-64f5e243f932/7847f8a4-f81e-4994-a095-b4d579deaf52", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_mfa_application_push_credential",
			ResourceName: "Test MFA_FCM",
			ResourceID:   fmt.Sprintf("%s/11cfc8c7-ec0c-43ff-b49a-64f5e243f932/e22e0f8f-ed88-4bdd-a914-5a93202083d0", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_mfa_application_push_credential",
			ResourceName: "Test MFA_HMS",
			ResourceID:   fmt.Sprintf("%s/11cfc8c7-ec0c-43ff-b49a-64f5e243f932/e609b3a8-b112-4062-8031-e9ff0d87c9e9", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
