// Copyright © 2025 Ping Identity Corporation

package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func TestableResource_PingFederateCertificatesRevocationOcspCertificate(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createCertificatesRevocationOcspCertificate,
		DeleteFunc:         deleteCertificatesRevocationOcspCertificate,
		Dependencies:       nil,
		ExportableResource: resources.CertificatesRevocationOcspCertificate(clientInfo),
	}
}

func createCertificatesRevocationOcspCertificate(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createCertificatesRevocationOcspCertificate(): %v", strArgs)
	}
	resourceType := strArgs[0]

	filedata, err := testutils.CreateX509Certificate()
	if err != nil {
		t.Fatalf("Failed to create test pem certificate %s: %v", resourceType, err)
	}

	request := clientInfo.PingFederateApiClient.CertificatesRevocationAPI.ImportOcspCertificate(clientInfo.PingFederateContext)
	clientStruct := client.X509File{
		FileData: filedata,
		Id:       utils.Pointer("testx509fileid"),
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "ImportOcspCertificate", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:            *resource.Id,
		testutils_resource.ENUM_ISSUER_DN:     *resource.IssuerDN,
		testutils_resource.ENUM_SERIAL_NUMBER: *resource.SerialNumber,
	}
}

func deleteCertificatesRevocationOcspCertificate(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.CertificatesRevocationAPI.DeleteOcspCertificateById(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteOcspCertificateById", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
