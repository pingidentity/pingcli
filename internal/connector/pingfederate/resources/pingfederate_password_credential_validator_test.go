package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederatePasswordCredentialValidatorExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.PasswordCredentialValidator(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_password_credential_validator",
			ResourceName: "pingdirectory",
			ResourceID:   "pingdirectory",
		},
		{
			ResourceType: "pingfederate_password_credential_validator",
			ResourceName: "simple",
			ResourceID:   "simple",
		},
		{
			ResourceType: "pingfederate_password_credential_validator",
			ResourceName: "PD PCV",
			ResourceID:   "PDPCV",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
