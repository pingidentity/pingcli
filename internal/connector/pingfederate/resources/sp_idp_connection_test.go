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

func TestableResource_PingFederateSpIdpConnection(t *testing.T) *testutils_resource.TestableResource {
	t.Helper()

	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	return &testutils_resource.TestableResource{
		ClientInfo:         pingfederateClientInfo,
		ExportableResource: resources.AuthenticationApiApplication(pingfederateClientInfo),
		TestResource: testutils_resource.TestResource{
			Dependencies: nil,
			CreateFunc:   createSpIdpConnection,
			DeleteFunc:   deleteSpIdpConnection,
		},
	}
}

func Test_PingFederateSpIdpConnection(t *testing.T) {
	tr := TestableResource_PingFederateSpIdpConnection(t)

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

func createSpIdpConnection(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createSpIdpConnection(): %v", strArgs)
	}
	resourceType := strArgs[0]

	filedata, err := testutils.CreateX509Certificate()
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	request := clientInfo.PingFederateApiClient.SpIdpConnectionsAPI.CreateConnection(clientInfo.Context)
	result := client.IdpConnection{
		Connection: client.Connection{
			Active: utils.Pointer(true),
			Credentials: &client.ConnectionCredentials{
				Certs: []client.ConnectionCert{
					{
						ActiveVerificationCert: utils.Pointer(true),
						EncryptionCert:         utils.Pointer(false),
						X509File: client.X509File{
							FileData: filedata,
							Id:       utils.Pointer("testx509fileid"),
						},
					},
				},
			},
			EntityId:    "TestEntityId",
			Id:          utils.Pointer("TestSpIdpConnectionId"),
			LoggingMode: utils.Pointer("STANDARD"),
			Name:        "TestSpIdpConnectionName",
			Type:        utils.Pointer("IDP"),
		},
		WsTrust: &client.IdpWsTrust{
			AttributeContract: client.IdpWsTrustAttributeContract{
				CoreAttributes: []client.IdpWsTrustAttribute{
					{
						Masked: utils.Pointer(false),
						Name:   "TOKEN_SUBJECT",
					},
				},
			},
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   *resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteSpIdpConnection(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.SpIdpConnectionsAPI.DeleteConnection(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteConnection", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
