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

func TestableResource_PingFederateKeypairsSigningKeyRotationSettings(t *testing.T) *testutils_resource.TestableResource {
	t.Helper()

	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	return &testutils_resource.TestableResource{
		ClientInfo:         pingfederateClientInfo,
		ExportableResource: resources.AuthenticationApiApplication(pingfederateClientInfo),
		TestResource: testutils_resource.TestResource{
			Dependencies: []testutils_resource.TestResource{
				{
					Dependencies: nil,
					CreateFunc:   createKeypairsSigningKey,
					DeleteFunc:   deleteKeypairsSigningKey,
				},
			},
			CreateFunc: createKeypairsSigningKeyRotationSettings,
			DeleteFunc: deleteKeypairsSigningKeyRotationSettings,
		},
	}
}

func Test_PingFederateKeypairsSigningKeyRotationSettings(t *testing.T) {
	tr := TestableResource_PingFederateKeypairsSigningKeyRotationSettings(t)

	_ = tr.CreateResource(t, tr.TestResource)
	defer tr.DeleteResource(t, tr.TestResource)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s_rotation_settings", tr.TestResource.Dependencies[0].CreationInfo[testutils_resource.ENUM_ISSUER_DN], tr.TestResource.Dependencies[0].CreationInfo[testutils_resource.ENUM_SERIAL_NUMBER]),
			ResourceID:   tr.TestResource.Dependencies[0].CreationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}

func createKeypairsSigningKeyRotationSettings(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createKeypairsSigningKeyRotationSettings(): %v", strArgs)
	}
	resourceType := strArgs[0]
	keyPairId := strArgs[1]

	request := clientInfo.PingFederateApiClient.KeyPairsSigningAPI.UpdateRotationSettings(clientInfo.Context, keyPairId)
	result := client.KeyPairRotationSettings{
		ActivationBufferDays: 10,
		CreationBufferDays:   10,
		Id:                   utils.Pointer("TestRotationSettingsId"),
	}

	request = request.Body(result)

	_, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	// Deletion of this resource is referenced by the keyPairId
	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID: keyPairId,
	}
}

func deleteKeypairsSigningKeyRotationSettings(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.KeyPairsSigningAPI.DeleteKeyPairRotationSettings(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteKeyPairRotationSettings", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}

func createKeypairsSigningKey(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createIdpAdapter(): %v", strArgs)
	}
	resourceType := strArgs[0]

	request := clientInfo.PingFederateApiClient.KeyPairsSigningAPI.CreateSigningKeyPair(clientInfo.Context)
	result := client.NewKeyPairSettings{
		City:               utils.Pointer("Denver"),
		CommonName:         "*.pingidentity.com",
		Country:            "US",
		Id:                 utils.Pointer("testkeypairid"),
		KeyAlgorithm:       "RSA",
		KeySize:            utils.Pointer(int64(2048)),
		Organization:       "Ping Identity Corporation",
		SignatureAlgorithm: utils.Pointer("SHA256withRSA"),
		State:              utils.Pointer("CO"),
		ValidDays:          365,
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateKeyPair", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:            *resource.Id,
		testutils_resource.ENUM_ISSUER_DN:     *resource.IssuerDN,
		testutils_resource.ENUM_SERIAL_NUMBER: *resource.SerialNumber,
	}
}

func deleteKeypairsSigningKey(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.KeyPairsSigningAPI.DeleteSigningKeyPair(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteSigningKeyPair", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
