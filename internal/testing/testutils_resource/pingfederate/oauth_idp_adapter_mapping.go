package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func TestableResource_PingFederateOauthIdpAdapterMapping(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createOauthIdpAdapterMapping,
		DeleteFunc: deleteOauthIdpAdapterMapping,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingFederateIdpAdapter(t, clientInfo),
		},
		ExportableResource: resources.OauthIdpAdapterMapping(clientInfo),
	}
}

func createOauthIdpAdapterMapping(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createOauthIdpAdapterMapping(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testIdpAdapterId := strArgs[1]

	request := clientInfo.PingFederateApiClient.OauthIdpAdapterMappingsAPI.CreateIdpAdapterMapping(clientInfo.PingFederateContext)
	clientStruct := client.IdpAdapterMapping{
		AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
			"USER_NAME": {
				Source: client.SourceTypeIdKey{
					Type: "NO_MAPPING",
				},
			},
			"USER_KEY": {
				Source: client.SourceTypeIdKey{
					Type: "NO_MAPPING",
				},
			},
		},
		Id: testIdpAdapterId,
		IdpAdapterRef: &client.ResourceLink{
			Id: testIdpAdapterId,
		},
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateIdpAdapterMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID: resource.Id,
	}
}

func deleteOauthIdpAdapterMapping(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthIdpAdapterMappingsAPI.DeleteIdpAdapterMapping(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteIdpAdapterMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
