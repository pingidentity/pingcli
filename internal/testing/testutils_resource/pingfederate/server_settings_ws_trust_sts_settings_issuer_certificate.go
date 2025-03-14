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

func TestableResource_PingFederateServerSettingsWsTrustStsSettingsIssuerCertificate(t *testing.T, clientInfo *connector.ClientInfo) testutils_resource.TestableResource {
	t.Helper()

	return testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createServerSettingsWsTrustStsSettingsIssuerCertificate,
		DeleteFunc:         deleteServerSettingsWsTrustStsSettingsIssuerCertificate,
		Dependencies:       nil,
		ExportableResource: resources.ServerSettingsWsTrustStsSettingsIssuerCertificate(clientInfo),
	}
}

func createServerSettingsWsTrustStsSettingsIssuerCertificate(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createServerSettingsWsTrustStsSettingsIssuerCertificate(): %v", strArgs)
	}
	resourceType := strArgs[0]

	fileData, err := testutils.CreateX509Certificate()
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	request := clientInfo.PingFederateApiClient.ServerSettingsAPI.ImportCertificate(clientInfo.Context)
	clientStruct := client.X509File{
		FileData: fileData,
		Id:       utils.Pointer("testx509fileid"),
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "ImportCertificate", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to create test %s: non-ok response", resourceType)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:            *resource.CertView.Id,
		testutils_resource.ENUM_ISSUER_DN:     *resource.CertView.IssuerDN,
		testutils_resource.ENUM_SERIAL_NUMBER: *resource.CertView.SerialNumber,
	}
}

func deleteServerSettingsWsTrustStsSettingsIssuerCertificate(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.ServerSettingsAPI.DeleteCertificate(clientInfo.Context, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteCertificate", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to delete test %s: non-ok response", resourceType)
	}
}
