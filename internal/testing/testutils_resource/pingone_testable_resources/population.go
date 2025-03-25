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
	"github.com/pingidentity/pingcli/internal/utils"
)

func Population(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createPopulation,
		DeleteFunc: deletePopulation,
		Dependencies: []*testutils_resource.TestableResource{
			PasswordPolicy(t, clientInfo),
		},
		ExportableResource: resources.Population(clientInfo),
	}
}

func createPopulation(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Errorf("Unexpected number of arguments provided to createPopulation(): %v", strArgs)
		return testutils_resource.ResourceInfo{}
	}
	passwordPolicyId := strArgs[0]

	request := clientInfo.PingOneApiClient.ManagementAPIClient.PopulationsApi.CreatePopulation(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID)
	clientStruct := management.Population{
		Name:        "Test Population",
		Description: utils.Pointer("This is a test population"),
		Default:     utils.Pointer(true),
		PasswordPolicy: &management.PopulationPasswordPolicy{
			Id: passwordPolicyId,
		},
	}

	request = request.Population(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreatePopulation", resourceType)
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

func deletePopulation(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deletePopulation(): %v", ids)
		return
	}

	getRequest := clientInfo.PingOneApiClient.ManagementAPIClient.PopulationsApi.ReadOnePopulation(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0])
	populationClientStruct, response, err := getRequest.Execute()
	ok, err := common.HandleClientResponse(response, err, "ReadOnePopulation", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
		return
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
		return
	}

	updateRequest := clientInfo.PingOneApiClient.ManagementAPIClient.PopulationsApi.UpdatePopulation(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0])
	populationClientStruct.Default = utils.Pointer(false)

	updateRequest = updateRequest.Population(*populationClientStruct)
	_, response, err = updateRequest.Execute()
	ok, err = common.HandleClientResponse(response, err, "UpdatePopulation", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
		return
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
		return
	}

	deleteRequest := clientInfo.PingOneApiClient.ManagementAPIClient.PopulationsApi.DeletePopulation(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0])

	response, err = deleteRequest.Execute()
	ok, err = common.HandleClientResponse(response, err, "DeletePopulation", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
		return
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
		return
	}
}
