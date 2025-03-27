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

func IdpStsRequestParametersContract(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createIdpStsRequestParametersContract,
		DeleteFunc:         deleteIdpStsRequestParametersContract,
		Dependencies:       nil,
		ExportableResource: resources.IdpStsRequestParametersContract(clientInfo),
	}
}

func createIdpStsRequestParametersContract(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 0 {
		t.Errorf("Unexpected number of arguments provided to createIdpStsRequestParametersContract(): %v", strArgs)

		return testutils_resource.ResourceInfo{}
	}

	request := clientInfo.PingFederateApiClient.IdpStsRequestParametersContractsAPI.CreateStsRequestParamContract(clientInfo.PingFederateContext)
	clientStruct := client.StsRequestParametersContract{
		Id:   "TestContractId",
		Name: "TestContractName",
		Parameters: []string{
			"TestParamName",
		},
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateStsRequestParamContract", resourceType)
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
			testutils_resource.ENUM_ID:   resource.Id,
			testutils_resource.ENUM_NAME: resource.Name,
		},
	}
}

func deleteIdpStsRequestParametersContract(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteIdpStsRequestParametersContract(): %v", ids)

		return
	}

	request := clientInfo.PingFederateApiClient.IdpStsRequestParametersContractsAPI.DeleteStsRequestParamContractById(clientInfo.PingFederateContext, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteStsRequestParamContractById", resourceType)
	if err != nil {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
