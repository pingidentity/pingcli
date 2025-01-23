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

func Test_PingFederateAuthenticationPoliciesFragment_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.AuthenticationPoliciesFragment(PingFederateClientInfo)

	passwordCredentialValidatorId, _ := createPasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType())
	defer deletePasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType(), passwordCredentialValidatorId)

	idpAdapterId, _ := createIdpAdapter(t, PingFederateClientInfo, resource.ResourceType(), passwordCredentialValidatorId)
	defer deleteIdpAdapter(t, PingFederateClientInfo, resource.ResourceType(), idpAdapterId)

	authenticationPoliciesFragmentId, authenticationPoliciesFragmentName := createAuthenticationPoliciesFragment(t, PingFederateClientInfo, resource.ResourceType(), idpAdapterId)
	defer deleteAuthenticationPoliciesFragment(t, PingFederateClientInfo, resource.ResourceType(), authenticationPoliciesFragmentId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: authenticationPoliciesFragmentName,
			ResourceID:   authenticationPoliciesFragmentId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createAuthenticationPoliciesFragment(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, idpAdapterId string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.AuthenticationPoliciesAPI.CreateFragment(clientInfo.Context)
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
	err = common.HandleClientResponse(response, err, "CreateFragment", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, *resource.Name
}

func deleteAuthenticationPoliciesFragment(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.AuthenticationPoliciesAPI.DeleteFragment(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteFragment", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
