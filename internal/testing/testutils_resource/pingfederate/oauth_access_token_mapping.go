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

func TestableResource_PingFederateOauthAccessTokenMapping(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createOauthAccessTokenMapping,
		DeleteFunc: deleteOauthAccessTokenMapping,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingFederateOauthAccessTokenManager(t, clientInfo),
		},
		ExportableResource: resources.OauthAccessTokenMapping(clientInfo),
	}
}

func createOauthAccessTokenMapping(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createOauthAccessTokenMapping(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testTokenManagerId := strArgs[1]

	request := clientInfo.PingFederateApiClient.OauthAccessTokenMappingsAPI.CreateMapping(clientInfo.PingFederateContext)
	clientStruct := client.AccessTokenMapping{
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

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:           *resource.Id,
		testutils_resource.ENUM_CONTEXT_TYPE: resource.Context.Type,
	}
}

func deleteOauthAccessTokenMapping(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthAccessTokenMappingsAPI.DeleteMapping(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
