package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func TestableResource_PingFederateIdentityStoreProvisioner(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createIdentityStoreProvisioner,
		DeleteFunc:         deleteIdentityStoreProvisioner,
		Dependencies:       nil,
		ExportableResource: resources.IdentityStoreProvisioner(clientInfo),
	}
}

func createIdentityStoreProvisioner(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createIdentityStoreProvisioner(): %v", strArgs)
	}
	resourceType := strArgs[0]

	request := clientInfo.PingFederateApiClient.IdentityStoreProvisionersAPI.CreateIdentityStoreProvisioner(clientInfo.PingFederateContext)
	clientStruct := client.IdentityStoreProvisioner{
		AttributeContract: &client.IdentityStoreProvisionerAttributeContract{
			CoreAttributes: []client.Attribute{
				{
					Name: "username",
				},
			},
		},
		GroupAttributeContract: &client.IdentityStoreProvisionerGroupAttributeContract{
			CoreAttributes: []client.GroupAttribute{
				{
					Name: "groupname",
				},
			},
		},
		Id:   "TestISPId",
		Name: "TestISPName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "com.pingidentity.identitystoreprovisioners.sample.SampleIdentityStoreProvisioner",
		},
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateIdentityStoreProvisioner", resourceType)
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

func deleteIdentityStoreProvisioner(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.IdentityStoreProvisionersAPI.DeleteIdentityStoreProvisioner(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteIdentityStoreProvisioner", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
