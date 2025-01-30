package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateIdpAdapter(t *testing.T) *testutils_resource.TestableResource {
	t.Helper()

	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	return &testutils_resource.TestableResource{
		ClientInfo:         pingfederateClientInfo,
		ExportableResource: resources.AuthenticationApiApplication(pingfederateClientInfo),
		TestResource: testutils_resource.TestResource{
			Dependencies: []testutils_resource.TestResource{
				{
					Dependencies: nil,
					CreateFunc:   createPasswordCredentialValidator,
					DeleteFunc:   deletePasswordCredentialValidator,
				},
			},
			CreateFunc: createIdpAdapter,
			DeleteFunc: deleteIdpAdapter,
		},
	}
}

func Test_PingFederateIdpAdapter(t *testing.T) {
	tr := TestableResource_PingFederateIdpAdapter(t)

	creationInfo := tr.CreateResource(t, tr.TestResource)
	defer tr.DeleteResource(t, tr.TestResource)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: creationInfo[testutils_resource.ENUM_NAME],
			ResourceID:   creationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}

func createIdpAdapter(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createIdpAdapter(): %v", strArgs)
	}
	resourceType := strArgs[0]
	passwordCredentialValidatorId := strArgs[1]

	request := clientInfo.PingFederateApiClient.IdpAdaptersAPI.CreateIdpAdapter(clientInfo.Context)
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
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteIdpAdapter(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.IdpAdaptersAPI.DeleteIdpAdapter(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteIdpAdapter", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
