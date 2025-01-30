package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateIdpSpConnection(t *testing.T) *testutils_resource.TestableResource {
	t.Helper()

	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	return &testutils_resource.TestableResource{
		ClientInfo:         pingfederateClientInfo,
		ExportableResource: resources.AuthenticationApiApplication(pingfederateClientInfo),
		TestResource: testutils_resource.TestResource{
			Dependencies: []testutils_resource.TestResource{
				{
					Dependencies: nil,
					CreateFunc:   createKeypairsSigningKey,
					DeleteFunc:   deleteKeypairsSigningKey,
				},
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
			CreateFunc: createIdpSpConnection,
			DeleteFunc: deleteIdpSpConnection,
		},
	}
}

func Test_PingFederateIdpSpConnection(t *testing.T) {
	tr := TestableResource_PingFederateIdpSpConnection(t)

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

func createIdpSpConnection(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 3 {
		t.Fatalf("Unexpected number of arguments provided to createIdpSpConnection(): %v", strArgs)
	}
	resourceType := strArgs[0]
	signingKeyPairId := strArgs[1]
	idpTokenProcessorId := strArgs[2]

	request := clientInfo.PingFederateApiClient.IdpSpConnectionsAPI.CreateSpConnection(clientInfo.Context)
	result := client.SpConnection{
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

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   *resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteIdpSpConnection(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.IdpSpConnectionsAPI.DeleteSpConnection(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteSpConnection", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
