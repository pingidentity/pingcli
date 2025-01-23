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

func Test_PingFederateCertificateCa_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.CertificateCa(PingFederateClientInfo)

	certificateCaId, certificateCaIssuerDn, certificateCaSerialNumber := createCertificateCa(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteCertificateCa(t, PingFederateClientInfo, resource.ResourceType(), certificateCaId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s", certificateCaIssuerDn, certificateCaSerialNumber),
			ResourceID:   certificateCaId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createCertificateCa(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string, string) {
	t.Helper()

	request := clientInfo.ApiClient.CertificatesCaAPI.ImportTrustedCA(clientInfo.Context)
	result := client.X509File{}
	result.Id = utils.Pointer("testx509fileid")
	filedata, err := testutils.CreateX509Certificate()
	if err != nil {
		t.Fatalf("Failed to create test pem certificate %s: %v", resourceType, err)
	}
	result.FileData = filedata

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "ImportTrustedCA", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, *resource.IssuerDN, *resource.SerialNumber
}

func deleteCertificateCa(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.CertificatesCaAPI.DeleteTrustedCA(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteTrustedCA", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
