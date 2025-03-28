// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingfederate_testable_resources

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func OauthIdpAdapterMapping(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createOauthIdpAdapterMapping,
		DeleteFunc: deleteOauthIdpAdapterMapping,
		Dependencies: []*testutils_resource.TestableResource{
			IdpAdapter(t, clientInfo),
		},
		ExportableResource: resources.OauthIdpAdapterMapping(clientInfo),
	}
}

func createOauthIdpAdapterMapping(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Errorf("Unexpected number of arguments provided to createOauthIdpAdapterMapping(): %v", strArgs)

		return testutils_resource.ResourceInfo{}
	}
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
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return testutils_resource.ResourceInfo{}
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return testutils_resource.ResourceInfo{}
	}

	return testutils_resource.ResourceInfo{
		DeletionIds: []string{
			resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID: resource.Id,
		},
	}
}

func deleteOauthIdpAdapterMapping(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteOauthIdpAdapterMapping(): %v", ids)

		return
	}

	request := clientInfo.PingFederateApiClient.OauthIdpAdapterMappingsAPI.DeleteIdpAdapterMapping(clientInfo.PingFederateContext, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteIdpAdapterMapping", resourceType)
	if err != nil {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
