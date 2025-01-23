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

func Test_PingFederateIdpSpConnection_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.IdpSpConnection(PingFederateClientInfo)

	testPCVId, _ := createPasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType())
	defer deletePasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType(), testPCVId)

	idpTokenProcessorId, _ := createIdpTokenProcessor(t, PingFederateClientInfo, resource.ResourceType(), testPCVId)
	defer deleteIdpTokenProcessor(t, PingFederateClientInfo, resource.ResourceType(), idpTokenProcessorId)

	signingKeyPairId, _, _ := createKeypairsSigningKey(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteKeypairsSigningKey(t, PingFederateClientInfo, resource.ResourceType(), signingKeyPairId)

	idpSpConnectionId, idpSpConnectionName := createIdpSpConnection(t, PingFederateClientInfo, resource.ResourceType(), idpTokenProcessorId, signingKeyPairId)
	defer deleteIdpSpConnection(t, PingFederateClientInfo, resource.ResourceType(), idpSpConnectionId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: idpSpConnectionName,
			ResourceID:   idpSpConnectionId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createIdpSpConnection(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, idpTokenProcessorId, signingKeyPairId string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.IdpSpConnectionsAPI.CreateSpConnection(clientInfo.Context)
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
	err = common.HandleClientResponse(response, err, "CreateSpConnection", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.Name
}

func deleteIdpSpConnection(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.IdpSpConnectionsAPI.DeleteSpConnection(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteSpConnection", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
