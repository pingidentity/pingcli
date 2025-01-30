package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateAuthenticationPoliciesFragment(t *testing.T) *testutils_resource.TestableResource {
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
			CreateFunc: createAuthenticationPoliciesFragment,
			DeleteFunc: deleteAuthenticationPoliciesFragment,
		},
	}
}

func Test_PingFederateAuthenticationPoliciesFragment(t *testing.T) {
	tr := TestableResource_PingFederateAuthenticationPoliciesFragment(t)

	creationInfo := tr.CreateResource(t, tr.TestResource)
	defer tr.DeleteResource(t, tr.TestResource)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: creationInfo[testutils_resource.ENUM_NAME],
			ResourceID:   creationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}

func createAuthenticationPoliciesFragment(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createAuthenticationPoliciesFragment(): %v", strArgs)
	}
	resourceType := strArgs[0]
	idpAdapterId := strArgs[1]

	request := clientInfo.PingFederateApiClient.AuthenticationPoliciesAPI.CreateFragment(clientInfo.Context)
	result := client.AuthenticationPolicyFragment{
		Id:   utils.Pointer("TestFragmentId"),
		Name: utils.Pointer("TestFragmentName"),
		RootNode: &client.AuthenticationPolicyTreeNode{
			Action: client.PolicyActionAggregation{
				AuthnSourcePolicyAction: &client.AuthnSourcePolicyAction{
					PolicyAction: client.PolicyAction{
						Type: "AUTHN_SOURCE",
					},
					AuthenticationSource: client.AuthenticationSource{
						SourceRef: client.ResourceLink{
							Id: idpAdapterId,
						},
						Type: "IDP_ADAPTER",
					},
				},
			},
			Children: []client.AuthenticationPolicyTreeNode{
				{
					Action: client.PolicyActionAggregation{
						DonePolicyAction: &client.DonePolicyAction{
							PolicyAction: client.PolicyAction{
								Type:    "DONE",
								Context: utils.Pointer("Fail"),
							},
						},
					},
				},
				{
					Action: client.PolicyActionAggregation{
						DonePolicyAction: &client.DonePolicyAction{
							PolicyAction: client.PolicyAction{
								Type:    "DONE",
								Context: utils.Pointer("Success"),
							},
						},
					},
				},
			},
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   *resource.Id,
		testutils_resource.ENUM_NAME: *resource.Name,
	}
}

func deleteAuthenticationPoliciesFragment(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.AuthenticationPoliciesAPI.DeleteFragment(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteFragment", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
