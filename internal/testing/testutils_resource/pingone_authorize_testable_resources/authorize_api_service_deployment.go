// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingone_authorize_testable_resources

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/authorize/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
)

func AuthorizeApiServiceDeployment(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createAuthorizeApiServiceDeployment,
		DeleteFunc: nil,
		Dependencies: []*testutils_resource.TestableResource{
			AuthorizeApiService(t, clientInfo),
		},
		ExportableResource: resources.AuthorizeApiServiceDeployment(clientInfo),
	}
}

func createAuthorizeApiServiceDeployment(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Errorf("Unexpected number of arguments provided to createAuthorizeApiServiceDeployment(): %v", strArgs)
		return testutils_resource.ResourceInfo{}
	}
	authorizeApiServiceId := strArgs[0]

	request := clientInfo.PingOneApiClient.AuthorizeAPIClient.APIServerDeploymentApi.DeployAPIServer(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, authorizeApiServiceId)
	request = request.ContentType("application/vnd.pingidentity.apiserver.deploy+json")

	_, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateApplicationPermission", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
		return testutils_resource.ResourceInfo{}
	}
	if !ok {
		if response != nil {
			t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
		} else {
			t.Errorf("Failed to execute PingOne client function")
		}
		return testutils_resource.ResourceInfo{}
	}

	return testutils_resource.ResourceInfo{}
}
