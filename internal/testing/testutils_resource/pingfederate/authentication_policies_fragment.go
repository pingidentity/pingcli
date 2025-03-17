package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func TestableResource_PingFederateAuthenticationPoliciesFragment(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createAuthenticationPoliciesFragment,
		DeleteFunc: deleteAuthenticationPoliciesFragment,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingFederateIdpAdapter(t, clientInfo),
		},
		ExportableResource: resources.AuthenticationPoliciesFragment(clientInfo),
	}
}

func createAuthenticationPoliciesFragment(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createAuthenticationPoliciesFragment(): %v", strArgs)
	}
	resourceType := strArgs[0]
	idpAdapterId := strArgs[1]

	request := clientInfo.PingFederateApiClient.AuthenticationPoliciesAPI.CreateFragment(clientInfo.PingFederateContext)
	clientStruct := client.AuthenticationPolicyFragment{
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

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateFragment", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   *resource.Id,
		testutils_resource.ENUM_NAME: *resource.Name,
	}
}

func deleteAuthenticationPoliciesFragment(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.AuthenticationPoliciesAPI.DeleteFragment(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteFragment", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
