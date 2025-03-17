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

func TestableResource_PingFederatePasswordCredentialValidator(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createPasswordCredentialValidator,
		DeleteFunc:         deletePasswordCredentialValidator,
		Dependencies:       nil,
		ExportableResource: resources.PasswordCredentialValidator(clientInfo),
	}
}

func createPasswordCredentialValidator(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createPasswordCredentialValidator(): %v", strArgs)
	}
	resourceType := strArgs[0]

	request := clientInfo.PingFederateApiClient.PasswordCredentialValidatorsAPI.CreatePasswordCredentialValidator(clientInfo.PingFederateContext)
	clientStruct := client.PasswordCredentialValidator{
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

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreatePasswordCredentialValidator", resourceType)
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

func deletePasswordCredentialValidator(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.PasswordCredentialValidatorsAPI.DeletePasswordCredentialValidator(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeletePasswordCredentialValidator", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
