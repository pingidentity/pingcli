// Copyright © 2025 Ping Identity Corporation

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

func TestableResource_PingFederateIdpSpConnection(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createIdpSpConnection,
		DeleteFunc: deleteIdpSpConnection,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingFederateKeypairsSigningKey(t, clientInfo),
			TestableResource_PingFederateIdpTokenProcessor(t, clientInfo),
		},
		ExportableResource: resources.IdpSpConnection(clientInfo),
	}
}

func createIdpSpConnection(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 3 {
		t.Fatalf("Unexpected number of arguments provided to createIdpSpConnection(): %v", strArgs)
	}
	resourceType := strArgs[0]
	signingKeyPairId := strArgs[1]
	idpTokenProcessorId := strArgs[2]

	request := clientInfo.PingFederateApiClient.IdpSpConnectionsAPI.CreateSpConnection(clientInfo.PingFederateContext)
	clientStruct := client.SpConnection{
		Connection: client.Connection{
			Active: utils.Pointer(true),
			Credentials: &client.ConnectionCredentials{
				SigningSettings: &client.SigningSettings{
					Algorithm:                utils.Pointer("SHA256withRSA"),
					IncludeCertInSignature:   utils.Pointer(false),
					IncludeRawKeyInSignature: utils.Pointer(false),
					SigningKeyPairRef: client.ResourceLink{
						Id: signingKeyPairId,
					},
				},
			},
			EntityId:    "TestEntityId",
			Id:          utils.Pointer("TestSpConnectionId"),
			LoggingMode: utils.Pointer("STANDARD"),
			Name:        "TestSpConnectionName",
			Type:        utils.Pointer("SP"),
		},
		WsTrust: &client.SpWsTrust{
			AttributeContract: client.SpWsTrustAttributeContract{
				CoreAttributes: []client.SpWsTrustAttribute{
					{
						Name: "TOKEN_SUBJECT",
					},
				},
			},
			DefaultTokenType:       utils.Pointer("SAML20"),
			EncryptSaml2Assertion:  utils.Pointer(false),
			GenerateKey:            utils.Pointer(false),
			MinutesBefore:          utils.Pointer(int64(5)),
			MinutesAfter:           utils.Pointer(int64(30)),
			OAuthAssertionProfiles: utils.Pointer(false),
			PartnerServiceIds: []string{
				"TestIdentifier",
			},
			TokenProcessorMappings: []client.IdpTokenProcessorMapping{
				{
					AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
						"TOKEN_SUBJECT": {
							Source: client.SourceTypeIdKey{
								Type: "NO_MAPPING",
							},
						},
					},
					IdpTokenProcessorRef: client.ResourceLink{
						Id: idpTokenProcessorId,
					},
				},
			},
		},
		ConnectionTargetType: utils.Pointer("STANDARD"),
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateSpConnection", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   *resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteIdpSpConnection(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.IdpSpConnectionsAPI.DeleteSpConnection(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteSpConnection", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
