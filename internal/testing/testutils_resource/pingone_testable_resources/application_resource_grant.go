// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingone_testable_resources

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
)

func ApplicationResourceGrant(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createApplicationResourceGrant,
		DeleteFunc: deleteApplicationResourceGrant,
		Dependencies: []*testutils_resource.TestableResource{
			ApplicationDeviceAuthorization(t, clientInfo),
			Resource(t, clientInfo),
		},
		ExportableResource: resources.ApplicationResourceGrant(clientInfo),
	}
}

func createApplicationResourceGrant(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createApplicationResourceGrant(): %v", strArgs)
	}
	applicationId := strArgs[0]
	resourceId := strArgs[1]

	resReq := clientInfo.PingOneApiClient.ManagementAPIClient.ResourceScopesApi.CreateResourceScope(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, resourceId)

	resClientStruct := management.ResourceScope{
		Name: "testCustomScope",
	}
	resReq = resReq.ResourceScope(resClientStruct)
	scope, response, err := resReq.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateResourceScope", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationResourceGrantsApi.CreateApplicationGrant(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, applicationId)
	clientStruct := management.ApplicationResourceGrant{
		Resource: management.ApplicationResourceGrantResource{
			Id: resourceId,
		},
		Scopes: []management.ApplicationResourceGrantScopesInner{
			{
				Id: *scope.Id,
			},
		},
	}

	request = request.ApplicationResourceGrant(clientStruct)

	resource, response, err := request.Execute()
	ok, err = common.HandleClientResponse(response, err, "CreateApplicationGrant", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	resObj, resResponse, resErr := clientInfo.PingOneApiClient.ManagementAPIClient.ResourcesApi.ReadOneResource(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, resourceId).Execute()
	ok, err = common.HandleClientResponse(resResponse, resErr, "ReadOneResource", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", resResponse.Status, resResponse.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", resResponse.Status, resResponse.Body)
	}

	return testutils_resource.ResourceInfo{
		DeletionIds: []string{
			applicationId,
			*resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   *resource.Id,
			testutils_resource.ENUM_NAME: resObj.Name,
		},
	}
}

func deleteApplicationResourceGrant(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 2 {
		t.Fatalf("Unexpected number of arguments provided to deleteApplicationResourceGrant(): %v", ids)
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationResourceGrantsApi.DeleteApplicationGrant(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0], ids[1])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteApplicationGrant", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
