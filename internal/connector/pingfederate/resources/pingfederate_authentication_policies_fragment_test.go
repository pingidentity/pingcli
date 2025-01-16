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

	pcvId := createPasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType())
	defer deletePasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType(), *pcvId)

	idpAdapterId := createIdpAdapter(t, PingFederateClientInfo, resource.ResourceType(), pcvId)
	defer deleteIdpAdapter(t, PingFederateClientInfo, resource.ResourceType(), idpAdapterId)

	fragmentId, fragmentName := createAuthenticationPoliciesFragment(t, PingFederateClientInfo, resource.ResourceType(), idpAdapterId)
	defer deleteAuthenticationPoliciesFragment(t, PingFederateClientInfo, resource.ResourceType(), fragmentId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fragmentName,
			ResourceID:   fragmentId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createAuthenticationPoliciesFragment(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, idpAdapterId string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.AuthenticationPoliciesAPI.CreateFragment(clientInfo.Context)
	result := client.AuthenticationPolicyFragment{}
	result.RootNode = &client.AuthenticationPolicyTreeNode{
		Action: client.PolicyActionAggregation{
			AuthnSourcePolicyAction: &client.AuthnSourcePolicyAction{
				PolicyAction: client.PolicyAction{
					Type: "AUTHN_SOURCE",
				},
				AuthenticationSource: client.AuthenticationSource{
					Type: "IDP_ADAPTER",
					SourceRef: client.ResourceLink{
						Id: idpAdapterId,
					},
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
	}
	result.Name = utils.Pointer("TestFragmentName")
	result.Id = utils.Pointer("TestFragmentId")

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateFragment", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, *resource.Name
}

func deleteAuthenticationPoliciesFragment(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string, id string) {
	t.Helper()

	request := clientInfo.ApiClient.AuthenticationPoliciesAPI.DeleteFragment(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteFragment", resourceType)
	if err != nil {
		t.Logf("Failed to delete test %s: %v", resourceType, err)
	}
}

func createIdpAdapter(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string, pcvId *string) string {
	t.Helper()

	request := clientInfo.ApiClient.IdpAdaptersAPI.CreateIdpAdapter(clientInfo.Context)
	result := client.IdpAdapter{}
	result.Id = "TestIdpAdapterId"
	result.Name = "TestIdpAdapterName"
	result.PluginDescriptorRef = client.ResourceLink{
		Id: "com.pingidentity.adapters.httpbasic.idp.HttpBasicIdpAuthnAdapter",
	}
	result.Configuration = client.PluginConfiguration{
		Tables: []client.ConfigTable{
			{
				Name: "Credential Validators",
				Rows: []client.ConfigRow{
					{
						DefaultRow: utils.Pointer(false),
						Fields: []client.ConfigField{
							{
								Name:  "Password Credential Validator Instance",
								Value: pcvId,
							},
						},
					},
				},
			},
		},
		Fields: []client.ConfigField{
			{
				Name:  "Realm",
				Value: utils.Pointer("testAuthenticationRealm"),
			},
			{
				Name:  "Challenge Retries",
				Value: utils.Pointer("3"),
			},
		},
	}
	result.AttributeContract = &client.IdpAdapterAttributeContract{
		CoreAttributes: []client.IdpAdapterAttribute{
			{
				Name:      "username",
				Pseudonym: utils.Pointer(true),
				Masked:    utils.Pointer(false),
			},
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateIdpAdapter", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id
}

func deleteIdpAdapter(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string, id string) {
	t.Helper()

	request := clientInfo.ApiClient.IdpAdaptersAPI.DeleteIdpAdapter(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteIdpAdapter", resourceType)
	if err != nil {
		t.Logf("Failed to delete test %s: %v", resourceType, err)
	}
}

func createPasswordCredentialValidator(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) *string {
	t.Helper()

	request := clientInfo.ApiClient.PasswordCredentialValidatorsAPI.CreatePasswordCredentialValidator(clientInfo.Context)
	result := client.PasswordCredentialValidator{}
	result.Id = "TestPCVId"
	result.Name = "TestPCVName"
	result.PluginDescriptorRef = client.ResourceLink{
		Id: "org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator",
	}
	result.Configuration = client.PluginConfiguration{
		Tables: []client.ConfigTable{
			{
				Name: "Users",
				Rows: []client.ConfigRow{
					{
						DefaultRow: utils.Pointer(true),
						Fields: []client.ConfigField{
							{
								Name:  "Username",
								Value: utils.Pointer("TestUser"),
							},
							{
								Name:  "Password",
								Value: utils.Pointer("TestPassword1"),
							},
							{
								Name:  "Confirm Password",
								Value: utils.Pointer("TestPassword1"),
							},
							{
								Name:  "Relax Password Requirements",
								Value: utils.Pointer("false"),
							},
						},
					},
				},
			},
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreatePasswordCredentialValidator", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return &resource.Id
}

func deletePasswordCredentialValidator(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string, id string) {
	t.Helper()

	request := clientInfo.ApiClient.PasswordCredentialValidatorsAPI.DeletePasswordCredentialValidator(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeletePasswordCredentialValidator", resourceType)
	if err != nil {
		t.Logf("Failed to delete test %s: %v", resourceType, err)
	}
}
