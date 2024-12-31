package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateSessionAuthenticationPolicyExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.SessionAuthenticationPolicy(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_session_authentication_policy",
			ResourceName: "UfdnqYjWycSeo2vZZgSYB3gpw_IDP_ADAPTER_OTIdPJava",
			ResourceID:   "UfdnqYjWycSeo2vZZgSYB3gpw",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
