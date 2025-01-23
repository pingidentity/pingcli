package resources_test

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateOauthCibaServerPolicyRequestPolicy_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.OauthCibaServerPolicyRequestPolicy(PingFederateClientInfo)

	gatewayId := createPingOnePingFederateGateway(t, PingOneClientInfo, resource.ResourceType())
	defer deletePingOnePingFederateGateway(t, PingOneClientInfo, resource.ResourceType(), gatewayId)

	credential := createPingOnePingFederateGatewayCredential(t, PingOneClientInfo, resource.ResourceType(), gatewayId)

	testPingOneConnectionId, _ := createPingoneConnection(t, PingFederateClientInfo, resource.ResourceType(), credential)
	defer deletePingoneConnection(t, PingFederateClientInfo, resource.ResourceType(), testPingOneConnectionId)

	testDeviceAuthApplicationId := createPingOneDeviceAuthApplication(t, PingOneClientInfo, resource.ResourceType())
	defer deletePingOneDeviceAuthApplication(t, PingOneClientInfo, resource.ResourceType(), testDeviceAuthApplicationId)

	testAuthenticatorId := createOutOfBandAuthPlugins(t, PingFederateClientInfo, resource.ResourceType(), testPingOneConnectionId, testDeviceAuthApplicationId)
	defer deleteOutOfBandAuthPlugins(t, PingFederateClientInfo, resource.ResourceType(), testAuthenticatorId)

	oauthCibaServerPolicyRequestPolicyId, oauthCibaServerPolicyRequestPolicyName := createOauthCibaServerPolicyRequestPolicy(t, PingFederateClientInfo, resource.ResourceType(), testAuthenticatorId)
	defer deleteOauthCibaServerPolicyRequestPolicy(t, PingFederateClientInfo, resource.ResourceType(), oauthCibaServerPolicyRequestPolicyId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: oauthCibaServerPolicyRequestPolicyName,
			ResourceID:   oauthCibaServerPolicyRequestPolicyId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createOauthCibaServerPolicyRequestPolicy(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, testAuthenticatorId string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthCibaServerPolicyAPI.CreateCibaServerPolicy(clientInfo.Context)
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
	err = common.HandleClientResponse(response, err, "CreateCibaServerPolicy", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteOauthCibaServerPolicyRequestPolicy(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthCibaServerPolicyAPI.DeleteCibaServerPolicy(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteCibaServerPolicy", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}

func createOutOfBandAuthPlugins(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, testPingOneConnectionId, testDeviceAuthApplicationId string) string {
	t.Helper()

	request := clientInfo.ApiClient.OauthOutOfBandAuthPluginsAPI.CreateOOBAuthenticator(clientInfo.Context)
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

	return resource.Id
}

func deleteOutOfBandAuthPlugins(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthOutOfBandAuthPluginsAPI.DeleteOOBAuthenticator(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteOOBAuthenticator", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}

func createPingOneDeviceAuthApplication(t *testing.T, clientInfo *connector.PingOneClientInfo, resourceType string) string {
	t.Helper()

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

	createApplication201Response, response, err := clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.CreateApplication(clientInfo.Context, testutils.GetEnvironmentID()).CreateApplicationRequest(result).Execute()
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

	return *appId
}

func deletePingOneDeviceAuthApplication(t *testing.T, clientInfo *connector.PingOneClientInfo, resourceType, id string) {
	t.Helper()

	response, err := clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.DeleteApplication(clientInfo.Context, testutils.GetEnvironmentID(), id).Execute()
	err = common.HandleClientResponse(response, err, "DeleteApplication", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
