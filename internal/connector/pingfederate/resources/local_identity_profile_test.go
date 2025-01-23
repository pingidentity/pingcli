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

func Test_PingFederateLocalIdentityProfile_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.LocalIdentityProfile(PingFederateClientInfo)

	localIdentityProfileId, localIdentityProfileName := createLocalIdentityProfile(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteLocalIdentityProfile(t, PingFederateClientInfo, resource.ResourceType(), localIdentityProfileId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: localIdentityProfileName,
			ResourceID:   localIdentityProfileId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createLocalIdentityProfile(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.LocalIdentityIdentityProfilesAPI.CreateIdentityProfile(clientInfo.Context)
	result := client.LocalIdentityProfile{}
	result.Id = utils.Pointer("TestLocalIdentityProfileId")
	result.Name = "TestLocalIdentityProfileName"
	result.ApcId = client.ResourceLink{
		Id: "",
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateIdentityProfile", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.Name
}

func deleteLocalIdentityProfile(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.LocalIdentityIdentityProfilesAPI.DeleteIdentityProfile(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteIdentityProfile", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
