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

func TestableResource_PingFederateOauthAccessTokenMapping(t *testing.T) *testutils_resource.TestableResource {
	t.Helper()

	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	return &testutils_resource.TestableResource{
		ClientInfo:         pingfederateClientInfo,
		ExportableResource: resources.AuthenticationApiApplication(pingfederateClientInfo),
		TestResource: testutils_resource.TestResource{
			Dependencies: []testutils_resource.TestResource{
				{
					Dependencies: []testutils_resource.TestResource{
						{
							Dependencies: nil,
							CreateFunc:   createKeypairsSigningKey,
							DeleteFunc:   deleteKeypairsSigningKey,
						},
					},
					CreateFunc: createOauthAccessTokenManager,
					DeleteFunc: deleteOauthAccessTokenManager,
				},
			},
			CreateFunc: createOauthAccessTokenMapping,
			DeleteFunc: deleteOauthAccessTokenMapping,
		},
	}
}

func Test_PingFederateOauthAccessTokenMapping(t *testing.T) {
	tr := TestableResource_PingFederateOauthAccessTokenMapping(t)

	creationInfo := tr.CreateResource(t, tr.TestResource)
	defer tr.DeleteResource(t, tr.TestResource)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s", creationInfo[testutils_resource.ENUM_CONTEXT_TYPE], creationInfo[testutils_resource.ENUM_ID]),
			ResourceID:   creationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}

func createOauthAccessTokenMapping(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createOauthAccessTokenMapping(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testTokenManagerId := strArgs[1]

	request := clientInfo.PingFederateApiClient.OauthAccessTokenMappingsAPI.CreateMapping(clientInfo.Context)
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
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:           *resource.Id,
		testutils_resource.ENUM_CONTEXT_TYPE: resource.Context.Type,
	}
}

func deleteOauthAccessTokenMapping(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthAccessTokenMappingsAPI.DeleteMapping(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteMapping", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
