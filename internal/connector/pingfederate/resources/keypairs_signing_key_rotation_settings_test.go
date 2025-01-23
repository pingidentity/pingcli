package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateKeypairsSigningKeyRotationSettings_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.KeypairsSigningKeyRotationSettings(PingFederateClientInfo)

	keypairsSigningKeyId, keypairsSigningKeyIssuerDn, keypairsSigningKeySerialNumber := createKeypairsSigningKey(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteKeypairsSigningKey(t, PingFederateClientInfo, resource.ResourceType(), keypairsSigningKeyId)

	createKeypairsSigningKeyRotationSettings(t, PingFederateClientInfo, resource.ResourceType(), keypairsSigningKeyId)
	defer deleteKeypairsSigningKeyRotationSettings(t, PingFederateClientInfo, resource.ResourceType(), keypairsSigningKeyId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s_rotation_settings", keypairsSigningKeyIssuerDn, keypairsSigningKeySerialNumber),
			ResourceID:   keypairsSigningKeyId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createKeypairsSigningKeyRotationSettings(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, keypairsSigningKeyId string) {
	t.Helper()

	request := clientInfo.ApiClient.KeyPairsSigningAPI.UpdateRotationSettings(clientInfo.Context, keypairsSigningKeyId)
	result := client.KeyPairRotationSettings{
		ActivationBufferDays: 10,
		CreationBufferDays:   10,
		Id:                   utils.Pointer("TestRotationSettingsId"),
	}

	request = request.Body(result)

	_, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "UpdateRotationSettings", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}
}

func deleteKeypairsSigningKeyRotationSettings(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.KeyPairsSigningAPI.DeleteKeyPairRotationSettings(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteKeyPairRotationSettings", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}

func createKeypairsSigningKey(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string, string) {
	t.Helper()

	request := clientInfo.ApiClient.KeyPairsSigningAPI.CreateSigningKeyPair(clientInfo.Context)
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

	return *resource.Id, *resource.IssuerDN, *resource.SerialNumber
}

func deleteKeypairsSigningKey(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.KeyPairsSigningAPI.DeleteSigningKeyPair(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteSigningKeyPair", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
