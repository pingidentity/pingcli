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

func Test_PingFederateOauthAccessTokenMapping_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OauthAccessTokenMapping(PingFederateClientInfo)

	testKeyPairId, _, _ := createKeypairsSigningKey(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteKeypairsSigningKey(t, PingFederateClientInfo, resource.ResourceType(), testKeyPairId)

	testTokenManagerId, _ := createOauthAccessTokenManager(t, PingFederateClientInfo, resource.ResourceType(), testKeyPairId)
	defer deleteOauthAccessTokenManager(t, PingFederateClientInfo, resource.ResourceType(), testTokenManagerId)

	oauthAccessTokenMappingId, oauthAccessTokenMappingContextType := createOauthAccessTokenMapping(t, PingFederateClientInfo, resource.ResourceType(), testTokenManagerId)
	defer deleteOauthAccessTokenMapping(t, PingFederateClientInfo, resource.ResourceType(), oauthAccessTokenMappingId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s", oauthAccessTokenMappingContextType, oauthAccessTokenMappingId),
			ResourceID:   oauthAccessTokenMappingId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createOauthAccessTokenMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, testTokenManagerId string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthAccessTokenMappingsAPI.CreateMapping(clientInfo.Context)
	result := client.AccessTokenMapping{
		AccessTokenManagerRef: client.ResourceLink{
			Id: testTokenManagerId,
		},
		AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
			"testAttribute": {
				Source: client.SourceTypeIdKey{
					Type: "NO_MAPPING",
				},
			},
		},
		Context: client.AccessTokenMappingContext{
			Type: "DEFAULT",
		},
		Id: utils.Pointer("default|" + testTokenManagerId),
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.Context.Type
}

func deleteOauthAccessTokenMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthAccessTokenMappingsAPI.DeleteMapping(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteMapping", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
