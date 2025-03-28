// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingfederate_testable_resources

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func IdpToSpAdapterMapping(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createIdpToSpAdapterMapping,
		DeleteFunc: deleteIdpToSpAdapterMapping,
		Dependencies: []*testutils_resource.TestableResource{
			IdpAdapter(t, clientInfo),
			SpAdapter(t, clientInfo),
		},
		ExportableResource: resources.IdpToSpAdapterMapping(clientInfo),
	}
}

func createIdpToSpAdapterMapping(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 3 {
		t.Errorf("Unexpected number of arguments provided to createIdpToSpAdapterMapping(): %v", strArgs)

		return testutils_resource.ResourceInfo{}
	}
	testIdpAdapterId := strArgs[1]
	testSpAdapterId := strArgs[2]

	request := clientInfo.PingFederateApiClient.IdpToSpAdapterMappingAPI.CreateIdpToSpAdapterMapping(clientInfo.PingFederateContext)
	clientStruct := client.IdpToSpAdapterMapping{
		AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
			"subject": {
				Source: client.SourceTypeIdKey{
					Type: "NO_MAPPING",
				},
			},
		},
		Id:       utils.Pointer(testIdpAdapterId + "|" + testSpAdapterId),
		SourceId: testIdpAdapterId,
		TargetId: testSpAdapterId,
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateIdpToSpAdapterMapping", resourceType)
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
			*resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:        *resource.Id,
			testutils_resource.ENUM_SOURCE_ID: resource.SourceId,
			testutils_resource.ENUM_TARGET_ID: resource.TargetId,
		},
	}
}

func deleteIdpToSpAdapterMapping(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteIdpToSpAdapterMapping(): %v", ids)

		return
	}

	request := clientInfo.PingFederateApiClient.IdpToSpAdapterMappingAPI.DeleteIdpToSpAdapterMappingsById(clientInfo.PingFederateContext, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteIdpToSpAdapterMappingsById", resourceType)
	if err != nil {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
