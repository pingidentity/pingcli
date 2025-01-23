package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateOpenidConnectPolicy_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OpenidConnectPolicy(PingFederateClientInfo)

	openidConnectPolicyId, openidConnectPolicyName := createOpenidConnectPolicy(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteOpenidConnectPolicy(t, PingFederateClientInfo, resource.ResourceType(), openidConnectPolicyId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: openidConnectPolicyName,
			ResourceID:   openidConnectPolicyId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createOpenidConnectPolicy(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthOpenIdConnectAPI.CreateOIDCPolicy(clientInfo.Context)
	result := client.OpenIdConnectPolicy{
		AccessTokenManagerRef: client.ResourceLink{
			Id: "",
		},
		AttributeContract: client.OpenIdConnectAttributeContract{},
		AttributeMapping:  client.AttributeMapping{},
		Id:                "TestOpenidConnectPolicyId",
		Name:              "TestOpenidConnectPolicyName",
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateOIDCPolicy", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteOpenidConnectPolicy(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthOpenIdConnectAPI.DeleteOIDCPolicy(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteOIDCPolicy", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
