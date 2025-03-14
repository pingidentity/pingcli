package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateCertificateCa(t *testing.T, clientInfo *connector.ClientInfo) testutils_resource.TestableResource {
	t.Helper()

	return testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createCertificateCa,
		DeleteFunc:         deleteCertificateCa,
		Dependencies:       nil, //TODO
		ExportableResource: resources.CertificateCa(clientInfo),
	}
}

func createCertificateCa(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 { //TODO
		t.Fatalf("Unexpected number of arguments provided to createCertificateCa(): %v", strArgs)
	}
	resourceType := strArgs[0]

	filedata, err := testutils.CreateX509Certificate()
	if err != nil {
		t.Fatalf("Failed to create test pem certificate %s: %v", resourceType, err)
	}

	request := clientInfo.PingFederateApiClient.CertificatesCaAPI.ImportTrustedCA(clientInfo.Context)
	clientStruct := client.X509File{
		FileData: filedata,
		Id:       utils.Pointer("testx509fileid"),
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "ImportTrustedCA", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to create test %s: non-ok response", resourceType)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:            *resource.Id,
		testutils_resource.ENUM_ISSUER_DN:     *resource.IssuerDN,
		testutils_resource.ENUM_SERIAL_NUMBER: *resource.SerialNumber,
	}
}

func deleteCertificateCa(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.CertificatesCaAPI.DeleteTrustedCA(clientInfo.Context, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteTrustedCA", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to delete test %s: non-ok response", resourceType)
	}
}
