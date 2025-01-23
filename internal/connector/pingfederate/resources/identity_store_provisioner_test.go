package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateIdentityStoreProvisioner_Export(t *testing.T) {
	pingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.IdentityStoreProvisioner(pingFederateClientInfo)

	identityStoreProvisionerId, identityStoreProvisionerName := createIdentityStoreProvisioner(t, pingFederateClientInfo, resource.ResourceType())
	defer deleteIdentityStoreProvisioner(t, pingFederateClientInfo, resource.ResourceType(), identityStoreProvisionerId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: identityStoreProvisionerName,
			ResourceID:   identityStoreProvisionerId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createIdentityStoreProvisioner(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.IdentityStoreProvisionersAPI.CreateIdentityStoreProvisioner(clientInfo.Context)
	result := client.IdentityStoreProvisioner{
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

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateIdentityStoreProvisioner", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteIdentityStoreProvisioner(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.IdentityStoreProvisionersAPI.DeleteIdentityStoreProvisioner(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteIdentityStoreProvisioner", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
