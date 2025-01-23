package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateIdpStsRequestParametersContract_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.IdpStsRequestParametersContract(PingFederateClientInfo)

	idpStsRequestParametersContractId, idpStsRequestParametersContractName := createIdpStsRequestParametersContract(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteIdpStsRequestParametersContract(t, PingFederateClientInfo, resource.ResourceType(), idpStsRequestParametersContractId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: idpStsRequestParametersContractName,
			ResourceID:   idpStsRequestParametersContractId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createIdpStsRequestParametersContract(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.IdpStsRequestParametersContractsAPI.CreateStsRequestParamContract(clientInfo.Context)
	result := client.StsRequestParametersContract{}
	result.Id = "TestIdpStsRequestParametersContractId"
	result.Name = "TestIdpStsRequestParametersContractName"
	result.Parameters = []string{}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateStsRequestParamContract", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteIdpStsRequestParametersContract(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.IdpStsRequestParametersContractsAPI.DeleteStsRequestParamContractById(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteStsRequestParamContractById", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
