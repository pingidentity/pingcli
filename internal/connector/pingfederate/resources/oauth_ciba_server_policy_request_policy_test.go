package resources_test

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateOauthCibaServerPolicyRequestPolicy(t *testing.T) *testutils_resource.TestableResource {
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
							Dependencies: []testutils_resource.TestResource{
								{
									Dependencies: []testutils_resource.TestResource{
										{
											Dependencies: nil,
											CreateFunc:   createPingOnePingFederateGateway,
											DeleteFunc:   deletePingOnePingFederateGateway,
										},
									},
									CreateFunc: createPingOnePingFederateGatewayCredential,
									DeleteFunc: nil,
								},
							},
							CreateFunc: createPingoneConnection,
							DeleteFunc: deletePingoneConnection,
						},
						{
							Dependencies: nil,
							CreateFunc:   createPingOneDeviceAuthApplication,
							DeleteFunc:   deletePingOneDeviceAuthApplication,
						},
					},
					CreateFunc: createOutOfBandAuthPlugins,
					DeleteFunc: deleteOutOfBandAuthPlugins,
				},
			},
			CreateFunc: createOauthCibaServerPolicyRequestPolicy,
			DeleteFunc: deleteOauthCibaServerPolicyRequestPolicy,
		},
	}
}

func Test_PingFederateOauthCibaServerPolicyRequestPolicy(t *testing.T) {
	tr := TestableResource_PingFederateOauthCibaServerPolicyRequestPolicy(t)

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

func createOauthCibaServerPolicyRequestPolicy(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createOauthCibaServerPolicyRequestPolicy(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testAuthenticatorId := strArgs[1]

	request := clientInfo.PingFederateApiClient.OauthCibaServerPolicyAPI.CreateCibaServerPolicy(clientInfo.Context)
	result := client.RequestPolicy{
		AllowUnsignedLoginHintToken: utils.Pointer(false),
		AuthenticatorRef: client.ResourceLink{
			Id: testAuthenticatorId,
		},
		Id: "TestRequestPolicyId",
		IdentityHintContract: client.IdentityHintContract{
			CoreAttributes: []client.IdentityHintAttribute{
				{
					Name: "IDENTITY_HINT_SUBJECT",
				},
			},
		},
		IdentityHintContractFulfillment: &client.AttributeMapping{
			AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
				"IDENTITY_HINT_SUBJECT": {
					Source: client.SourceTypeIdKey{
						Type: "REQUEST",
					},
					Value: "IDENTITY_HINT_SUBJECT",
				},
			},
		},
		IdentityHintMapping: &client.AttributeMapping{
			AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
				"subject": {
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
		},
		Name:                        "TestRequestPolicyName",
		RequireTokenForIdentityHint: utils.Pointer(false),
		TransactionLifetime:         utils.Pointer(int64(120)),
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteOauthCibaServerPolicyRequestPolicy(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthCibaServerPolicyAPI.DeleteCibaServerPolicy(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteCibaServerPolicy", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}

func createOutOfBandAuthPlugins(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 3 {
		t.Fatalf("Unexpected number of arguments provided to createOutOfBandAuthPlugins(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testPingOneConnectionId := strArgs[1]
	testDeviceAuthApplicationId := strArgs[2]

	request := clientInfo.PingFederateApiClient.OauthOutOfBandAuthPluginsAPI.CreateOOBAuthenticator(clientInfo.Context)
	result := client.OutOfBandAuthenticator{
		AttributeContract: &client.OutOfBandAuthAttributeContract{
			CoreAttributes: []client.OutOfBandAuthAttribute{
				{
					Name: "subject",
				},
			},
		},
		Configuration: client.PluginConfiguration{
			Fields: []client.ConfigField{
				{
					Name:  "PingOne Environment",
					Value: utils.Pointer(testPingOneConnectionId + "|" + testutils.GetEnvironmentID()),
				},
				{
					Name:  "Application",
					Value: &testDeviceAuthApplicationId,
				},
			},
		},
		Id:   "TestOOBAuthenticatorId",
		Name: "TestOOBAuthenticatorName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "com.pingidentity.oobauth.pingone.mfa.PingOneMfaCibaAuthenticator",
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateOutOfBandAuthPlugin", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID: resource.Id,
	}
}

func deleteOutOfBandAuthPlugins(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthOutOfBandAuthPluginsAPI.DeleteOOBAuthenticator(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteOOBAuthenticator", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}

func createPingOneDeviceAuthApplication(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createPingOneDeviceAuthApplication(): %v", strArgs)
	}
	resourceType := strArgs[0]

	result := management.CreateApplicationRequest{
		ApplicationOIDC: &management.ApplicationOIDC{
			Enabled: true,
			GrantTypes: []management.EnumApplicationOIDCGrantType{
				management.ENUMAPPLICATIONOIDCGRANTTYPE_DEVICE_CODE,
				management.ENUMAPPLICATIONOIDCGRANTTYPE_REFRESH_TOKEN,
			},
			Name:                    "TestDeviceAuthApplication",
			Protocol:                management.ENUMAPPLICATIONPROTOCOL_OPENID_CONNECT,
			TokenEndpointAuthMethod: management.ENUMAPPLICATIONOIDCTOKENAUTHMETHOD_NONE,
			Type:                    management.ENUMAPPLICATIONTYPE_CUSTOM_APP,
		},
	}

	createApplication201Response, response, err := clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationsApi.CreateApplication(clientInfo.Context, testutils.GetEnvironmentID()).CreateApplicationRequest(result).Execute()
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	if createApplication201Response == nil || createApplication201Response.ApplicationOIDC == nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	appId, appIdOk := createApplication201Response.ApplicationOIDC.GetIdOk()
	if !appIdOk {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID: *appId,
	}
}

func deletePingOneDeviceAuthApplication(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	response, err := clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationsApi.DeleteApplication(clientInfo.Context, testutils.GetEnvironmentID(), id).Execute()
	err = common.HandleClientResponse(response, err, "DeleteApplication", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
