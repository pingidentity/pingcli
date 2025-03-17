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

func TestableResource_PingFederateIdpTokenProcessor(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createIdpTokenProcessor,
		DeleteFunc: deleteIdpTokenProcessor,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingFederatePasswordCredentialValidator(t, clientInfo),
		},
		ExportableResource: resources.IdpTokenProcessor(clientInfo),
	}
}

func createIdpTokenProcessor(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createIdpTokenProcessor(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testPCVId := strArgs[1]

	request := clientInfo.PingFederateApiClient.IdpTokenProcessorsAPI.CreateTokenProcessor(clientInfo.PingFederateContext)
	clientStruct := client.TokenProcessor{
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

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateTokenProcessor", resourceType)
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

func deleteIdpTokenProcessor(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.IdpTokenProcessorsAPI.DeleteTokenProcessor(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteTokenProcessor", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
