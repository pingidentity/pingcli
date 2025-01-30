package resources_test

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

func TestableResource_PingFederateServerSettingsWsTrustStsSettingsIssuerCertificate(t *testing.T) *testutils_resource.TestableResource {
	t.Helper()

	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	return &testutils_resource.TestableResource{
		ClientInfo:         pingfederateClientInfo,
		ExportableResource: resources.AuthenticationApiApplication(pingfederateClientInfo),
		TestResource: testutils_resource.TestResource{
			Dependencies: nil,
			CreateFunc:   createServerSettingsWsTrustStsSettingsIssuerCertificate,
			DeleteFunc:   deleteServerSettingsWsTrustStsSettingsIssuerCertificate,
		},
	}
}

func Test_PingFederateServerSettingsWsTrustStsSettingsIssuerCertificate(t *testing.T) {
	tr := TestableResource_PingFederateServerSettingsWsTrustStsSettingsIssuerCertificate(t)

	creationInfo := tr.CreateResource(t, tr.TestResource)
	defer tr.DeleteResource(t, tr.TestResource)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: creationInfo[testutils_resource.ENUM_NAME],
			ResourceID:   creationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
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
	result := client.X509File{
		FileData: fileData,
		Id:       utils.Pointer("testx509fileid"),
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
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
	err = common.HandleClientResponse(response, err, "DeleteCertificate", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
