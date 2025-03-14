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

func TestableResource_PingFederateSessionAuthenticationPolicy(t *testing.T, clientInfo *connector.ClientInfo) testutils_resource.TestableResource {
	t.Helper()

	return testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createSessionAuthenticationPolicy,
		DeleteFunc: deleteSessionAuthenticationPolicy,
		Dependencies: []testutils_resource.TestableResource{
			TestableResource_PingFederateIdpAdapter(t, clientInfo),
		},
		ExportableResource: resources.SessionAuthenticationPolicy(clientInfo),
	}
}

func createSessionAuthenticationPolicy(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createSessionAuthenticationPolicy(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testIdpAdapterId := strArgs[1]

	request := clientInfo.PingFederateApiClient.SessionAPI.CreateSourcePolicy(clientInfo.Context)
	clientStruct := client.AuthenticationSessionPolicy{
		AuthenticationSource: client.AuthenticationSource{
			SourceRef: client.ResourceLink{
				Id: testIdpAdapterId,
			},
			Type: "IDP_ADAPTER",
		},
		Id: utils.Pointer("TestSessionAuthenticationPolicyId"),
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateSourcePolicy", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to create test %s: non-ok response", resourceType)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:            *resource.Id,
		testutils_resource.ENUM_TYPE:          resource.AuthenticationSource.Type,
		testutils_resource.ENUM_SOURCE_REF_ID: resource.AuthenticationSource.SourceRef.Id,
	}
}

func deleteSessionAuthenticationPolicy(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.SessionAPI.DeleteSourcePolicy(clientInfo.Context, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteSourcePolicy", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to delete test %s: non-ok response", resourceType)
	}
}
