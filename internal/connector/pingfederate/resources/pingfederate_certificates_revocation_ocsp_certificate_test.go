package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateCertificatesRevocationOcspCertificate_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.CertificatesRevocationOcspCertificate(PingFederateClientInfo)

	certificatesRevocationOcspCertificateId, certificatesRevocationOcspCertificateIssuerDn, certificatesRevocationOcspCertificateSerialNumber := createCertificatesRevocationOcspCertificate(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteCertificatesRevocationOcspCertificate(t, PingFederateClientInfo, resource.ResourceType(), certificatesRevocationOcspCertificateId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s", certificatesRevocationOcspCertificateIssuerDn, certificatesRevocationOcspCertificateSerialNumber),
			ResourceID:   certificatesRevocationOcspCertificateId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createCertificatesRevocationOcspCertificate(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string, string) {
	t.Helper()

	request := clientInfo.ApiClient.CertificatesRevocationAPI.ImportOcspCertificate(clientInfo.Context)
	result := client.X509File{}
	result.Id = utils.Pointer("testx509fileid")
	filedata, err := testutils.CreatePemCertificateCa()
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}
	result.FileData = filedata

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "ImportOcspCertificate", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, *resource.IssuerDN, *resource.SerialNumber
}

func deleteCertificatesRevocationOcspCertificate(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.CertificatesRevocationAPI.DeleteOcspCertificateById(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteOcspCertificateById", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
