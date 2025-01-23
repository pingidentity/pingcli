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

func Test_PingFederatePasswordCredentialValidator_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.PasswordCredentialValidator(PingFederateClientInfo)

	passwordCredentialValidatorId, passwordCredentialValidatorName := createPasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType())
	defer deletePasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType(), passwordCredentialValidatorId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: passwordCredentialValidatorName,
			ResourceID:   passwordCredentialValidatorId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createPasswordCredentialValidator(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.PasswordCredentialValidatorsAPI.CreatePasswordCredentialValidator(clientInfo.Context)
	result := client.PasswordCredentialValidator{
		Configuration: client.PluginConfiguration{
			Tables: []client.ConfigTable{
				{
					Name: "Users",
					Rows: []client.ConfigRow{
						{
							DefaultRow: utils.Pointer(false),
							Fields: []client.ConfigField{
								{
									Name:  "Username",
									Value: utils.Pointer("TestUsername"),
								},
								{
									Name:  "Password",
									Value: utils.Pointer("TestPassword1"),
								},
								{
									Name:  "Confirm Password",
									Value: utils.Pointer("TestPassword1"),
								},
								{
									Name:  "Relax Password Requirements",
									Value: utils.Pointer("false"),
								},
							},
						},
					},
				},
			},
		},
		Id:   "TestPCVId",
		Name: "TestPCVName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator",
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreatePasswordCredentialValidator", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deletePasswordCredentialValidator(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.PasswordCredentialValidatorsAPI.DeletePasswordCredentialValidator(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeletePasswordCredentialValidator", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
