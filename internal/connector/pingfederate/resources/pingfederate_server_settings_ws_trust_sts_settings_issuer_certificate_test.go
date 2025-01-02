package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateServerSettingsWsTrustStsSettingsIssuerCertificateExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.ServerSettingsWsTrustStsSettingsIssuerCertificate(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_server_settings_ws_trust_sts_settings_issuer_certificate",
			ResourceName: "CN=test, O=Ping Identity Corporation, L=Denver, ST=CO, C=US_430421198347763948001683365009287878912609754790",
			ResourceID:   "ycrgw3j4ckw91gxdmd479qftb",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
