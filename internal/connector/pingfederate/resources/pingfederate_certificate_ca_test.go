package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateCertificateCAExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.CertificateCA(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_certificate_ca",
			ResourceName: "C=US, O=CDR, OU=PING, L=AUSTIN, ST=TEXAS_38647788523832031312085637263346848131",
			ResourceID:   "sslservercert",
		},
	}
	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
