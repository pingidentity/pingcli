package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateCertificateCa(t *testing.T) *testutils_resource.TestableResource {
	t.Helper()

	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	return &testutils_resource.TestableResource{
		ClientInfo:         pingfederateClientInfo,
		ExportableResource: resources.AuthenticationApiApplication(pingfederateClientInfo),
		TestResource: testutils_resource.TestResource{
			Dependencies: nil,
			CreateFunc:   createCertificateCa,
			DeleteFunc:   deleteCertificateCa,
		},
	}
}

func Test_PingFederateCertificateCa(t *testing.T) {
	tr := TestableResource_PingFederateCertificateCa(t)

	creationInfo := tr.CreateResource(t, tr.TestResource)
	defer tr.DeleteResource(t, tr.TestResource)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s", creationInfo[testutils_resource.ENUM_ISSUER_DN], creationInfo[testutils_resource.ENUM_SERIAL_NUMBER]),
			ResourceID:   creationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}

func createCertificateCa(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createCertificateCa(): %v", strArgs)
	}
	resourceType := strArgs[0]

	filedata, err := testutils.CreateX509Certificate()
	if err != nil {
		t.Fatalf("Failed to create test pem certificate %s: %v", resourceType, err)
	}

	request := clientInfo.PingFederateApiClient.CertificatesCaAPI.ImportTrustedCA(clientInfo.Context)
	result := client.X509File{
		FileData: filedata,
		Id:       utils.Pointer("testx509fileid"),
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
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
	err = common.HandleClientResponse(response, err, "DeleteTrustedCA", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
