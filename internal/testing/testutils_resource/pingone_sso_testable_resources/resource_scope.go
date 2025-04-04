// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingone_sso_testable_resources

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
)

func ResourceScope(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createResourceScope,
		DeleteFunc: deleteResourceScope,
		Dependencies: []*testutils_resource.TestableResource{
			Resource(t, clientInfo),
		},
		ExportableResource: resources.ResourceScope(clientInfo),
	}
}

func createResourceScope(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Errorf("Unexpected number of arguments provided to createResourceScope(): %v", strArgs)

		return testutils_resource.ResourceInfo{}
	}
	resourceId := strArgs[0]

	request := clientInfo.PingOneApiClient.ManagementAPIClient.ResourceScopesApi.CreateResourceScope(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, resourceId)
	clientStruct := management.ResourceScope{
		Name:        "CustomScope",
		Description: utils.Pointer("This is a custom scope"),
	}

	request = request.ResourceScope(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateResourceScope", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return testutils_resource.ResourceInfo{}
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return testutils_resource.ResourceInfo{}
	}

	return testutils_resource.ResourceInfo{
		DeletionIds: []string{
			resourceId,
			*resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   *resource.Id,
			testutils_resource.ENUM_NAME: resource.Name,
		},
	}
}

func deleteResourceScope(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 2 {
		t.Errorf("Unexpected number of arguments provided to deleteResourceScope(): %v", ids)

		return
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.ResourceScopesApi.DeleteResourceScope(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0], ids[1])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteResourceScope", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
