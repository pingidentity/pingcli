package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateOauthIdpAdapterMapping(t *testing.T) *testutils_resource.TestableResource {
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
							CreateFunc:   createPasswordCredentialValidator,
							DeleteFunc:   deletePasswordCredentialValidator,
						},
					},
					CreateFunc: createIdpAdapter,
					DeleteFunc: deleteIdpAdapter,
				},
			},
			CreateFunc: createOauthIdpAdapterMapping,
			DeleteFunc: deleteOauthIdpAdapterMapping,
		},
	}
}

func Test_PingFederateOauthIdpAdapterMapping(t *testing.T) {
	tr := TestableResource_PingFederateOauthIdpAdapterMapping(t)

	creationInfo := tr.CreateResource(t, tr.TestResource)
	defer tr.DeleteResource(t, tr.TestResource)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_mapping", creationInfo[testutils_resource.ENUM_ID]),
			ResourceID:   creationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}

func createOauthIdpAdapterMapping(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createOauthIdpAdapterMapping(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testIdpAdapterId := strArgs[1]

	request := clientInfo.PingFederateApiClient.OauthIdpAdapterMappingsAPI.CreateIdpAdapterMapping(clientInfo.Context)
	result := client.IdpAdapterMapping{
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

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID: resource.Id,
	}
}

func deleteOauthIdpAdapterMapping(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthIdpAdapterMappingsAPI.DeleteIdpAdapterMapping(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteIdpAdapterMapping", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
