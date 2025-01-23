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

func Test_PingFederateTokenProcessorToTokenGeneratorMapping_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.TokenProcessorToTokenGeneratorMapping(PingFederateClientInfo)

	testPCVId, _ := createPasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType())
	defer deletePasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType(), testPCVId)

	testTokenProcessorId, _ := createIdpTokenProcessor(t, PingFederateClientInfo, resource.ResourceType(), testPCVId)
	defer deleteIdpTokenProcessor(t, PingFederateClientInfo, resource.ResourceType(), testTokenProcessorId)

	testSigningKeyPairId, _, _ := createKeypairsSigningKey(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteKeypairsSigningKey(t, PingFederateClientInfo, resource.ResourceType(), testSigningKeyPairId)

	testTokenGeneratorId := createSpTokenGenerator(t, PingFederateClientInfo, resource.ResourceType(), testSigningKeyPairId)
	defer deleteSpTokenGenerator(t, PingFederateClientInfo, resource.ResourceType(), testTokenGeneratorId)

	tokenProcessorToTokenGeneratorMappingId, tokenProcessorToTokenGeneratorMappingSourceId, tokenProcessorToTokenGeneratorMappingTargetId := createTokenProcessorToTokenGeneratorMapping(t, PingFederateClientInfo, resource.ResourceType(), testTokenProcessorId, testTokenGeneratorId)
	defer deleteTokenProcessorToTokenGeneratorMapping(t, PingFederateClientInfo, resource.ResourceType(), tokenProcessorToTokenGeneratorMappingId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_to_%s", tokenProcessorToTokenGeneratorMappingSourceId, tokenProcessorToTokenGeneratorMappingTargetId),
			ResourceID:   tokenProcessorToTokenGeneratorMappingId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createTokenProcessorToTokenGeneratorMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, testTokenProcessorId, testTokenGeneratorId string) (string, string, string) {
	t.Helper()

	request := clientInfo.ApiClient.TokenProcessorToTokenGeneratorMappingsAPI.CreateTokenToTokenMapping(clientInfo.Context)
	result := client.TokenToTokenMapping{
		AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
			"SAML_SUBJECT": {
				Source: client.SourceTypeIdKey{
					Type: "NO_MAPPING",
				},
			},
		},
		Id:       utils.Pointer(testTokenProcessorId + "|" + testTokenGeneratorId),
		SourceId: testTokenProcessorId,
		TargetId: testTokenGeneratorId,
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateTokenToTokenMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.SourceId, resource.TargetId
}

func deleteTokenProcessorToTokenGeneratorMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.TokenProcessorToTokenGeneratorMappingsAPI.DeleteTokenToTokenMappingById(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteTokenToTokenMappingById", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
