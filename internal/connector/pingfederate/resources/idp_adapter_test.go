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

func Test_PingFederateIdpAdapter_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.IdpAdapter(PingFederateClientInfo)

	passwordCredentialValidatorId, _ := createPasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType())
	defer deletePasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType(), passwordCredentialValidatorId)

	idpAdapterId, idpAdapterName := createIdpAdapter(t, PingFederateClientInfo, resource.ResourceType(), passwordCredentialValidatorId)
	defer deleteIdpAdapter(t, PingFederateClientInfo, resource.ResourceType(), idpAdapterId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: idpAdapterName,
			ResourceID:   idpAdapterId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createIdpAdapter(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, passwordCredentialValidatorId string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.IdpAdaptersAPI.CreateIdpAdapter(clientInfo.Context)
	result := client.IdpAdapter{
		AttributeContract: &client.IdpAdapterAttributeContract{
			CoreAttributes: []client.IdpAdapterAttribute{
				{
					Masked:    utils.Pointer(false),
					Name:      "username",
					Pseudonym: utils.Pointer(true),
				},
			},
		},
		Configuration: client.PluginConfiguration{
			Fields: []client.ConfigField{
				{
					Name:  "Realm",
					Value: utils.Pointer("TestAuthRealm"),
				},
			},
			Tables: []client.ConfigTable{
				{
					Name: "Credential Validators",
					Rows: []client.ConfigRow{
						{
							DefaultRow: utils.Pointer(false),
							Fields: []client.ConfigField{
								{
									Name:  "Password Credential Validator Instance",
									Value: utils.Pointer(passwordCredentialValidatorId),
								},
							},
						},
					},
				},
			},
		},
		Id:   "TestIdpAdapterId",
		Name: "TestIdpAdapterName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "com.pingidentity.adapters.httpbasic.idp.HttpBasicIdpAuthnAdapter",
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateIdpAdapter", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteIdpAdapter(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.IdpAdaptersAPI.DeleteIdpAdapter(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteIdpAdapter", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
