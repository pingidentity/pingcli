// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingone_authorize_testable_resources

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/authorize/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_sso_testable_resources"
)

func AuthorizeApiService(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createAuthorizeApiService,
		DeleteFunc: deleteAuthorizeApiService,
		Dependencies: []*testutils_resource.TestableResource{
			pingone_sso_testable_resources.Resource(t, clientInfo),
		},
		ExportableResource: resources.AuthorizeApiService(clientInfo),
	}
}

func createAuthorizeApiService(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Errorf("Unexpected number of arguments provided to createAuthorizeApiService(): %v", strArgs)

		return testutils_resource.ResourceInfo{}
	}
	resourceId := strArgs[0]

	request := clientInfo.PingOneApiClient.AuthorizeAPIClient.APIServersApi.CreateAPIServer(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID)
	clientStruct := authorize.APIServer{
		Name: "Banking API-Advanced",
		BaseUrls: []string{
			"https://api.example.com/advbanking/v1",
			"https://example-api.cdn/advbanking/v1",
		},
		AuthorizationServer: authorize.APIServerAuthorizationServer{
			Resource: &authorize.APIServerAuthorizationServerResource{
				Id: resourceId,
			},
			Type: authorize.ENUMAPISERVERAUTHORIZATIONSERVERTYPE_PINGONE_SSO,
		},
	}

	request = request.APIServer(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateAPIServer", resourceType)
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
			*resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   *resource.Id,
			testutils_resource.ENUM_NAME: resource.Name,
		},
	}
}

func deleteAuthorizeApiService(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteAuthorizeApiService(): %v", ids)

		return
	}

	request := clientInfo.PingOneApiClient.AuthorizeAPIClient.APIServersApi.DeleteAPIServer(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteAPIServer", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
