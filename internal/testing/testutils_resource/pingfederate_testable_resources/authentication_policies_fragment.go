// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingfederate_testable_resources

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func AuthenticationPoliciesFragment(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createAuthenticationPoliciesFragment,
		DeleteFunc: deleteAuthenticationPoliciesFragment,
		Dependencies: []*testutils_resource.TestableResource{
			IdpAdapter(t, clientInfo),
		},
		ExportableResource: resources.AuthenticationPoliciesFragment(clientInfo),
	}
}

func createAuthenticationPoliciesFragment(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Errorf("Unexpected number of arguments provided to createAuthenticationPoliciesFragment(): %v", strArgs)

		return testutils_resource.ResourceInfo{}
	}
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
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return testutils_resource.ResourceInfo{}
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return testutils_resource.ResourceInfo{}
	}

	return testutils_resource.ResourceInfo{
		DeletionIds: []string{
			*resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   *resource.Id,
			testutils_resource.ENUM_NAME: *resource.Name,
		},
	}
}

func deleteAuthenticationPoliciesFragment(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteAuthenticationPoliciesFragment(): %v", ids)

		return
	}

	request := clientInfo.PingFederateApiClient.AuthenticationPoliciesAPI.DeleteFragment(clientInfo.PingFederateContext, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteFragment", resourceType)
	if err != nil {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
