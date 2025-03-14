package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateLocalIdentityProfile(t *testing.T, clientInfo *connector.ClientInfo) testutils_resource.TestableResource {
	t.Helper()

	return testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createLocalIdentityProfile,
		DeleteFunc: deleteLocalIdentityProfile,
		Dependencies: []testutils_resource.TestableResource{
			TestableResource_PingFederateAuthenticationPolicyContract(t, clientInfo),
		},
		ExportableResource: resources.LocalIdentityProfile(clientInfo),
	}
}

func createLocalIdentityProfile(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createLocalIdentityProfile(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testApcId := strArgs[1]

	request := clientInfo.PingFederateApiClient.LocalIdentityIdentityProfilesAPI.CreateIdentityProfile(clientInfo.Context)
	clientStruct := client.LocalIdentityProfile{
		ApcId: client.ResourceLink{
			Id: testApcId,
		},
		Id:   utils.Pointer("TestLocalIdentityProfileId"),
		Name: "TestLocalIdentityProfileName",
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateIdentityProfile", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to create test %s: non-ok response", resourceType)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   *resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteLocalIdentityProfile(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.LocalIdentityIdentityProfilesAPI.DeleteIdentityProfile(clientInfo.Context, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteIdentityProfile", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to delete test %s: non-ok response", resourceType)
	}
}
