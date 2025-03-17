package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func TestableResource_PingFederateIdpAdapter(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createIdpAdapter,
		DeleteFunc: deleteIdpAdapter,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingFederatePasswordCredentialValidator(t, clientInfo),
		},
		ExportableResource: resources.IdpAdapter(clientInfo),
	}
}

func createIdpAdapter(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createIdpAdapter(): %v", strArgs)
	}
	resourceType := strArgs[0]
	pcvId := strArgs[1]

	request := clientInfo.PingFederateApiClient.IdpAdaptersAPI.CreateIdpAdapter(clientInfo.PingFederateContext)
	clientStruct := client.IdpAdapter{
		Id:   "testIdpAdapterId",
		Name: "testIdpAdapterName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "com.pingidentity.adapters.httpbasic.idp.HttpBasicIdpAuthnAdapter",
		},
		Configuration: client.PluginConfiguration{
			Tables: []client.ConfigTable{
				{
					Name: "Credential Validators",
					Rows: []client.ConfigRow{
						{
							Fields: []client.ConfigField{
								{
									Name:  "Password Credential Validator Instance",
									Value: utils.Pointer(pcvId),
								},
							},
							DefaultRow: utils.Pointer(false),
						},
					},
				},
			},
			Fields: []client.ConfigField{
				{
					Name:  "Realm",
					Value: utils.Pointer("testRealmName"),
				},
			},
		},
		AttributeContract: &client.IdpAdapterAttributeContract{
			CoreAttributes: []client.IdpAdapterAttribute{
				{
					Name:      "username",
					Masked:    utils.Pointer(false),
					Pseudonym: utils.Pointer(true),
				},
			},
		},
		AttributeMapping: &client.IdpAdapterContractMapping{
			AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
				"username": {
					Source: client.SourceTypeIdKey{
						Type: "ADAPTER",
					},
					Value: "username",
				},
			},
		},
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateIdpAdapter", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteIdpAdapter(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.IdpAdaptersAPI.DeleteIdpAdapter(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteIdpAdapter", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
