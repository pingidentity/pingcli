// Copyright © 2025 Ping Identity Corporation

package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func TestableResource_PingFederateSpTokenGenerator(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createSpTokenGenerator,
		DeleteFunc: deleteSpTokenGenerator,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingFederateKeypairsSigningKey(t, clientInfo),
		},
		ExportableResource: nil,
	}
}

func createSpTokenGenerator(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createSpTokenGenerator(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testSigningKeyPairId := strArgs[1]

	request := clientInfo.PingFederateApiClient.SpTokenGeneratorsAPI.CreateTokenGenerator(clientInfo.PingFederateContext)
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
	ok, err := common.HandleClientResponse(response, err, "CreateTokenGenerator", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID: resource.Id,
	}
}

func deleteSpTokenGenerator(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.SpTokenGeneratorsAPI.DeleteTokenGenerator(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteTokenGenerator", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
