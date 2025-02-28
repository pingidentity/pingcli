package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateCertificatesRevocationOCSPCertificateExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.CertificatesRevocationOCSPCertificate(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_certificates_revocation_ocsp_certificate",
			ResourceName: "CN=test, O=Ping Identity Corporation, L=Denver, ST=CO, C=US_430421198347763948001683365009287878912609754790",
			ResourceID:   "opcey20sf9djwvk8snv1actzq",
		},
	}
	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
