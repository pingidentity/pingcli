package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateIdpTokenProcessor_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.IdpTokenProcessor(PingFederateClientInfo)

	testPCVId, _ := createPasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType())
	defer deletePasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType(), testPCVId)

	idpTokenProcessorId, idpTokenProcessorName := createIdpTokenProcessor(t, PingFederateClientInfo, resource.ResourceType(), testPCVId)
	defer deleteIdpTokenProcessor(t, PingFederateClientInfo, resource.ResourceType(), idpTokenProcessorId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: idpTokenProcessorName,
			ResourceID:   idpTokenProcessorId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createIdpTokenProcessor(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, testPCVId string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.IdpTokenProcessorsAPI.CreateTokenProcessor(clientInfo.Context)
	result := client.TokenProcessor{
		AttributeContract: &client.TokenProcessorAttributeContract{
			CoreAttributes: []client.TokenProcessorAttribute{
				{
					Masked: utils.Pointer(false),
					Name:   "username",
				},
			},
			MaskOgnlValues: utils.Pointer(false),
		},
		Configuration: client.PluginConfiguration{
			Tables: []client.ConfigTable{
				{
					Name: "Credential Validators",
					Rows: []client.ConfigRow{
						{
							DefaultRow: utils.Pointer(false),
							Fields: []client.ConfigField{
								{
									Name:  "Password Credential Validator Instance",
									Value: &testPCVId,
								},
							},
						},
					},
				},
			},
		},
		Id:   "TestIdpTokenProcessorId",
		Name: "TestIdpTokenProcessorName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "com.pingidentity.pf.tokenprocessors.username.UsernameTokenProcessor",
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateTokenProcessor", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteIdpTokenProcessor(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.IdpTokenProcessorsAPI.DeleteTokenProcessor(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteTokenProcessor", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
