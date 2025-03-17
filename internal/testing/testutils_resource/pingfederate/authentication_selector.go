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

func TestableResource_PingFederateAuthenticationSelector(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createAuthenticationSelector,
		DeleteFunc:         deleteAuthenticationSelector,
		Dependencies:       nil,
		ExportableResource: resources.AuthenticationSelector(clientInfo),
	}
}

func createAuthenticationSelector(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createAuthenticationSelector(): %v", strArgs)
	}
	resourceType := strArgs[0]

	request := clientInfo.PingFederateApiClient.AuthenticationSelectorsAPI.CreateAuthenticationSelector(clientInfo.PingFederateContext)
	clientStruct := client.AuthenticationSelector{
		Configuration: client.PluginConfiguration{
			Fields: []client.ConfigField{
				{
					Name:  "Header Name",
					Value: utils.Pointer("TestHeaderName"),
				},
			},
			Tables: []client.ConfigTable{
				{
					Name: "Results",
					Rows: []client.ConfigRow{
						{
							DefaultRow: utils.Pointer(false),
							Fields: []client.ConfigField{
								{
									Name:  "Match Expression",
									Value: utils.Pointer("TestMatchExpression"),
								},
							},
						},
					},
				},
			},
		},
		Id:   "TestAuthSelectorId",
		Name: "TestAuthSelectorName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "com.pingidentity.pf.selectors.http.HTTPHeaderAdapterSelector",
		},
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateAuthenticationSelector", resourceType)
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

func deleteAuthenticationSelector(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.AuthenticationSelectorsAPI.DeleteAuthenticationSelector(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteAuthenticationSelector", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
