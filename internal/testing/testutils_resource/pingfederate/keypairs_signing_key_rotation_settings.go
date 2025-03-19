// Copyright © 2025 Ping Identity Corporation

package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func TestableResource_PingFederateKeypairsSigningKeyRotationSettings(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createKeypairsSigningKeyRotationSettings,
		DeleteFunc: deleteKeypairsSigningKeyRotationSettings,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingFederateKeypairsSigningKey(t, clientInfo),
		},
		ExportableResource: resources.KeypairsSigningKeyRotationSettings(clientInfo),
	}
}

func createKeypairsSigningKeyRotationSettings(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createKeypairsSigningKeyRotationSettings(): %v", strArgs)
	}
	resourceType := strArgs[0]
	keyPairId := strArgs[1]

	request := clientInfo.PingFederateApiClient.KeyPairsSigningAPI.UpdateRotationSettings(clientInfo.PingFederateContext, keyPairId)
	clientStruct := client.KeyPairRotationSettings{
		ActivationBufferDays: 10,
		CreationBufferDays:   10,
		Id:                   utils.Pointer("TestRotationSettingsId"),
	}

	request = request.Body(clientStruct)

	_, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "UpdateRotationSettings", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	// Deletion of this resource is referenced by the keyPairId
	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID: keyPairId,
	}
}

func deleteKeypairsSigningKeyRotationSettings(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.KeyPairsSigningAPI.DeleteKeyPairRotationSettings(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteKeyPairRotationSettings", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
