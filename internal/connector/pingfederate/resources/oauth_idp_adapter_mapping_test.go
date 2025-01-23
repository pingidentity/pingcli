package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateOauthIdpAdapterMapping_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OauthIdpAdapterMapping(PingFederateClientInfo)

	testPCVId, _ := createPasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType())
	defer deletePasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType(), testPCVId)

	testIdpAdapterId, _ := createIdpAdapter(t, PingFederateClientInfo, resource.ResourceType(), testPCVId)
	defer deleteIdpAdapter(t, PingFederateClientInfo, resource.ResourceType(), testIdpAdapterId)

	oauthIdpAdapterMappingId := createOauthIdpAdapterMapping(t, PingFederateClientInfo, resource.ResourceType(), testIdpAdapterId)
	defer deleteOauthIdpAdapterMapping(t, PingFederateClientInfo, resource.ResourceType(), oauthIdpAdapterMappingId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_mapping", oauthIdpAdapterMappingId),
			ResourceID:   oauthIdpAdapterMappingId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createOauthIdpAdapterMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, testIdpAdapterId string) string {
	t.Helper()

	request := clientInfo.ApiClient.OauthIdpAdapterMappingsAPI.CreateIdpAdapterMapping(clientInfo.Context)
	result := client.IdpAdapterMapping{
		AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
			"USER_NAME": {
				Source: client.SourceTypeIdKey{
					Type: "NO_MAPPING",
				},
			},
			"USER_KEY": {
				Source: client.SourceTypeIdKey{
					Type: "NO_MAPPING",
				},
			},
		},
		Id: testIdpAdapterId,
		IdpAdapterRef: &client.ResourceLink{
			Id: testIdpAdapterId,
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateIdpAdapterMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id
}

func deleteOauthIdpAdapterMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthIdpAdapterMappingsAPI.DeleteIdpAdapterMapping(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteIdpAdapterMapping", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
