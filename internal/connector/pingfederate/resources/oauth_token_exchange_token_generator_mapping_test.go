package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateOauthTokenExchangeTokenGeneratorMapping_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OauthTokenExchangeTokenGeneratorMapping(PingFederateClientInfo)

	testPCVId, _ := createPasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType())
	defer deletePasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType(), testPCVId)

	testTokenProcessorId, _ := createIdpTokenProcessor(t, PingFederateClientInfo, resource.ResourceType(), testPCVId)
	defer deleteIdpTokenProcessor(t, PingFederateClientInfo, resource.ResourceType(), testTokenProcessorId)

	testProcessorPolicyId := createOauthTokenExchangeProcessorPolicy(t, PingFederateClientInfo, resource.ResourceType(), testTokenProcessorId)
	defer deleteOauthTokenExchangeProcessorPolicy(t, PingFederateClientInfo, resource.ResourceType(), testProcessorPolicyId)

	testSigningKeyPairId, _, _ := createKeypairsSigningKey(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteKeypairsSigningKey(t, PingFederateClientInfo, resource.ResourceType(), testSigningKeyPairId)

	testTokenGeneratorId := createSpTokenGenerator(t, PingFederateClientInfo, resource.ResourceType(), testSigningKeyPairId)
	defer deleteSpTokenGenerator(t, PingFederateClientInfo, resource.ResourceType(), testTokenGeneratorId)

	oauthTokenExchangeTokenGeneratorMappingId, oauthTokenExchangeTokenGeneratorMappingSourceId, oauthTokenExchangeTokenGeneratorMappingTargetId := createOauthTokenExchangeTokenGeneratorMapping(t, PingFederateClientInfo, resource.ResourceType(), testProcessorPolicyId, testTokenGeneratorId)
	defer deleteOauthTokenExchangeTokenGeneratorMapping(t, PingFederateClientInfo, resource.ResourceType(), oauthTokenExchangeTokenGeneratorMappingId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_to_%s", oauthTokenExchangeTokenGeneratorMappingSourceId, oauthTokenExchangeTokenGeneratorMappingTargetId),
			ResourceID:   oauthTokenExchangeTokenGeneratorMappingId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createOauthTokenExchangeTokenGeneratorMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, testProcessorPolicyId, testTokenGeneratorId string) (string, string, string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthTokenExchangeTokenGeneratorMappingsAPI.CreateTokenGeneratorMapping(clientInfo.Context)
	result := client.ProcessorPolicyToGeneratorMapping{
		AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
			"SAML_SUBJECT": {
				Source: client.SourceTypeIdKey{
					Type: "NO_MAPPING",
				},
			},
		},
		Id:       utils.Pointer(testProcessorPolicyId + "|" + testTokenGeneratorId),
		SourceId: testProcessorPolicyId,
		TargetId: testTokenGeneratorId,
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateTokenGeneratorMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.SourceId, resource.TargetId
}

func deleteOauthTokenExchangeTokenGeneratorMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthTokenExchangeTokenGeneratorMappingsAPI.DeleteTokenGeneratorMappingById(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteTokenGeneratorMappingById", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}

func createOauthTokenExchangeProcessorPolicy(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, testTokenProcessorId string) string {
	t.Helper()

	request := clientInfo.ApiClient.OauthTokenExchangeProcessorAPI.CreateOauthTokenExchangeProcessorPolicy(clientInfo.Context)
	result := client.TokenExchangeProcessorPolicy{
		ActorTokenRequired: utils.Pointer(false),
		AttributeContract: client.TokenExchangeProcessorAttributeContract{
			CoreAttributes: []client.TokenExchangeProcessorAttribute{
				{
					Name: "subject",
				},
			},
		},
		Id:   "TestProcessorPolicyId",
		Name: "TestProcessorPolicyName",
		ProcessorMappings: []client.TokenExchangeProcessorMapping{
			{
				AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
					"subject": {
						Source: client.SourceTypeIdKey{
							Type: "NO_MAPPING",
						},
					},
				},
				SubjectTokenType: "TestTokenType",
				SubjectTokenProcessor: client.ResourceLink{
					Id: testTokenProcessorId,
				},
			},
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateOauthTokenExchangeProcessorPolicy", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id
}

func deleteOauthTokenExchangeProcessorPolicy(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthTokenExchangeProcessorAPI.DeleteOauthTokenExchangeProcessorPolicyy(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteOauthTokenExchangeProcessorPolicy", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}

func createSpTokenGenerator(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, testSigningKeyPairId string) string {
	t.Helper()

	request := clientInfo.ApiClient.SpTokenGeneratorsAPI.CreateTokenGenerator(clientInfo.Context)
	result := client.TokenGenerator{
		AttributeContract: &client.TokenGeneratorAttributeContract{
			CoreAttributes: []client.TokenGeneratorAttribute{
				{
					Name: "SAML_SUBJECT",
				},
			},
		},
		Configuration: client.PluginConfiguration{
			Fields: []client.ConfigField{
				{
					Name:  "Minutes Before",
					Value: utils.Pointer("10"),
				},
				{
					Name:  "Minutes After",
					Value: utils.Pointer("10"),
				},
				{
					Name:  "Issuer",
					Value: utils.Pointer("issuerIdentifier"),
				},
				{
					Name:  "Signing Certificate",
					Value: &testSigningKeyPairId,
				},
				{
					Name:  "Signing Algorithm",
					Value: utils.Pointer("RSA_SHA256"),
				},
				{
					Name:  "Audience",
					Value: utils.Pointer("www.example.com"),
				},
			},
		},
		Id:   "TestTokenGeneratorId",
		Name: "TestTokenGeneratorName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "org.sourceid.wstrust.generator.saml.Saml20TokenGenerator",
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateTokenGenerator", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id
}

func deleteSpTokenGenerator(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.SpTokenGeneratorsAPI.DeleteTokenGenerator(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteTokenGenerator", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
