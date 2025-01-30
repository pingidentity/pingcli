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

func TestableResource_PingFederateOauthTokenExchangeTokenGeneratorMapping(t *testing.T) *testutils_resource.TestableResource {
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
									Dependencies: nil,
									CreateFunc:   createPasswordCredentialValidator,
									DeleteFunc:   deletePasswordCredentialValidator,
								},
							},
							CreateFunc: createIdpTokenProcessor,
							DeleteFunc: deleteIdpTokenProcessor,
						},
					},
					CreateFunc: createOauthTokenExchangeProcessorPolicy,
					DeleteFunc: deleteOauthTokenExchangeProcessorPolicy,
				},
				{
					Dependencies: []testutils_resource.TestResource{
						{
							Dependencies: nil,
							CreateFunc:   createKeypairsSigningKey,
							DeleteFunc:   deleteKeypairsSigningKey,
						},
					},
					CreateFunc: createSpTokenGenerator,
					DeleteFunc: deleteSpTokenGenerator,
				},
			},
			CreateFunc: createOauthTokenExchangeTokenGeneratorMapping,
			DeleteFunc: deleteOauthTokenExchangeTokenGeneratorMapping,
		},
	}
}

func Test_PingFederateOauthTokenExchangeTokenGeneratorMapping(t *testing.T) {
	tr := TestableResource_PingFederateOauthTokenExchangeTokenGeneratorMapping(t)

	creationInfo := tr.CreateResource(t, tr.TestResource)
	defer tr.DeleteResource(t, tr.TestResource)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_to_%s", creationInfo[testutils_resource.ENUM_SOURCE_ID], creationInfo[testutils_resource.ENUM_TARGET_ID]),
			ResourceID:   creationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}

func createOauthTokenExchangeTokenGeneratorMapping(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 3 {
		t.Fatalf("Unexpected number of arguments provided to createOauthTokenExchangeTokenGeneratorMapping(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testProcessorPolicyId := strArgs[1]
	testTokenGeneratorId := strArgs[2]

	request := clientInfo.PingFederateApiClient.OauthTokenExchangeTokenGeneratorMappingsAPI.CreateTokenGeneratorMapping(clientInfo.Context)
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
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:        *resource.Id,
		testutils_resource.ENUM_SOURCE_ID: testProcessorPolicyId,
		testutils_resource.ENUM_TARGET_ID: testTokenGeneratorId,
	}
}

func deleteOauthTokenExchangeTokenGeneratorMapping(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthTokenExchangeTokenGeneratorMappingsAPI.DeleteTokenGeneratorMappingById(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteTokenGeneratorMappingById", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}

func createOauthTokenExchangeProcessorPolicy(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createOauthTokenExchangeProcessorPolicy(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testTokenProcessorId := strArgs[1]

	request := clientInfo.PingFederateApiClient.OauthTokenExchangeProcessorAPI.CreateOauthTokenExchangeProcessorPolicy(clientInfo.Context)
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

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID: resource.Id,
	}
}

func deleteOauthTokenExchangeProcessorPolicy(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthTokenExchangeProcessorAPI.DeleteOauthTokenExchangeProcessorPolicyy(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteOauthTokenExchangeProcessorPolicy", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}

func createSpTokenGenerator(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createSpTokenGenerator(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testSigningKeyPairId := strArgs[1]

	request := clientInfo.PingFederateApiClient.SpTokenGeneratorsAPI.CreateTokenGenerator(clientInfo.Context)
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

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID: resource.Id,
	}
}

func deleteSpTokenGenerator(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.SpTokenGeneratorsAPI.DeleteTokenGenerator(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteTokenGenerator", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
