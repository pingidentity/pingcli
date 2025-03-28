// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingfederate_testable_resources

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

func SpIdpConnection(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createSpIdpConnection,
		DeleteFunc:         deleteSpIdpConnection,
		Dependencies:       nil,
		ExportableResource: resources.SpIdpConnection(clientInfo),
	}
}

func createSpIdpConnection(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 0 {
		t.Fatalf("Unexpected number of arguments provided to createSpIdpConnection(): %v", strArgs)
	}

	filedata, err := testutils.CreateX509Certificate()
	if err != nil {
		t.Errorf("Failed to create test %s: %v", resourceType, err)

		return testutils_resource.ResourceInfo{}
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
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return testutils_resource.ResourceInfo{}
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return testutils_resource.ResourceInfo{}
	}

	return testutils_resource.ResourceInfo{
		DeletionIds: []string{
			*resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   *resource.Id,
			testutils_resource.ENUM_NAME: resource.Name,
		},
	}
}

func deleteSpIdpConnection(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteSpIdpConnection(): %v", ids)

		return
	}

	request := clientInfo.PingFederateApiClient.SpIdpConnectionsAPI.DeleteConnection(clientInfo.PingFederateContext, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteConnection", resourceType)
	if err != nil {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
