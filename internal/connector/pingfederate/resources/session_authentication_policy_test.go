package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateSessionAuthenticationPolicy(t *testing.T) *testutils_resource.TestableResource {
	t.Helper()

	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	return &testutils_resource.TestableResource{
		ClientInfo:         pingfederateClientInfo,
		ExportableResource: resources.AuthenticationApiApplication(pingfederateClientInfo),
		TestResource: testutils_resource.TestResource{
			Dependencies: []testutils_resource.TestResource{
				{
					Dependencies: []testutils_resource.TestResource{
						{
							Dependencies: nil,
							CreateFunc:   createPasswordCredentialValidator,
							DeleteFunc:   deletePasswordCredentialValidator,
						},
					},
					CreateFunc: createIdpAdapter,
					DeleteFunc: deleteIdpAdapter,
				},
			},
			CreateFunc: createSessionAuthenticationPolicy,
			DeleteFunc: deleteSessionAuthenticationPolicy,
		},
	}
}

func Test_PingFederateSessionAuthenticationPolicy(t *testing.T) {
	tr := TestableResource_PingFederateSessionAuthenticationPolicy(t)

	creationInfo := tr.CreateResource(t, tr.TestResource)
	defer tr.DeleteResource(t, tr.TestResource)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s_%s", creationInfo[testutils_resource.ENUM_ID], creationInfo[testutils_resource.ENUM_TYPE], creationInfo[testutils_resource.ENUM_SOURCE_REF_ID]),
			ResourceID:   creationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}

func createSessionAuthenticationPolicy(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createSessionAuthenticationPolicy(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testIdpAdapterId := strArgs[1]

	request := clientInfo.PingFederateApiClient.SessionAPI.CreateSourcePolicy(clientInfo.Context)
	result := client.AuthenticationSessionPolicy{
		AuthenticationSource: client.AuthenticationSource{
			SourceRef: client.ResourceLink{
				Id: testIdpAdapterId,
			},
			Type: "IDP_ADAPTER",
		},
		Id: utils.Pointer("TestSessionAuthenticationPolicyId"),
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
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
	err = common.HandleClientResponse(response, err, "DeleteSourcePolicy", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
