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

func TestableResource_PingFederateSpIdpConnection(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createSpIdpConnection,
		DeleteFunc:         deleteSpIdpConnection,
		Dependencies:       nil,
		ExportableResource: resources.SpIdpConnection(clientInfo),
	}
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

	request := clientInfo.PingFederateApiClient.SpIdpConnectionsAPI.CreateConnection(clientInfo.PingFederateContext)
	clientStruct := client.IdpConnection{
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

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateConnection", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   *resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteSpIdpConnection(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.SpIdpConnectionsAPI.DeleteConnection(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteConnection", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
