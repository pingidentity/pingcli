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

func Test_PingFederateOauthIssuer_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OauthIssuer(PingFederateClientInfo)

	oauthIssuerId, oauthIssuerName := createOauthIssuer(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteOauthIssuer(t, PingFederateClientInfo, resource.ResourceType(), oauthIssuerId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: oauthIssuerName,
			ResourceID:   oauthIssuerId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createOauthIssuer(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthIssuersAPI.AddOauthIssuer(clientInfo.Context)
	result := client.Issuer{
		Host: "TestIssuerHost",
		Id:   utils.Pointer("TestIssuerId"),
		Name: "TestIssuerName",
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "AddOauthIssuer", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.Name
}

func deleteOauthIssuer(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthIssuersAPI.DeleteOauthIssuer(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteOauthIssuer", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
