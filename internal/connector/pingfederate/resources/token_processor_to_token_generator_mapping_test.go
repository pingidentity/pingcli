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

func TestableResource_PingFederateTokenProcessorToTokenGeneratorMapping(t *testing.T) *testutils_resource.TestableResource {
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
							Dependencies: nil,
							CreateFunc:   createPasswordCredentialValidator,
							DeleteFunc:   deletePasswordCredentialValidator,
						},
					},
					CreateFunc: createIdpTokenProcessor,
					DeleteFunc: deleteIdpTokenProcessor,
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
			CreateFunc: createTokenProcessorToTokenGeneratorMapping,
			DeleteFunc: deleteTokenProcessorToTokenGeneratorMapping,
		},
	}
}

func Test_PingFederateTokenProcessorToTokenGeneratorMapping(t *testing.T) {
	tr := TestableResource_PingFederateTokenProcessorToTokenGeneratorMapping(t)

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

func createTokenProcessorToTokenGeneratorMapping(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 3 {
		t.Fatalf("Unexpected number of arguments provided to createTokenProcessorToTokenGeneratorMapping(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testTokenProcessorId := strArgs[1]
	testTokenGeneratorId := strArgs[2]

	request := clientInfo.PingFederateApiClient.TokenProcessorToTokenGeneratorMappingsAPI.CreateTokenToTokenMapping(clientInfo.Context)
	result := client.TokenToTokenMapping{
		AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
			"SAML_SUBJECT": {
				Source: client.SourceTypeIdKey{
					Type: "NO_MAPPING",
				},
			},
		},
		Id:       utils.Pointer(testTokenProcessorId + "|" + testTokenGeneratorId),
		SourceId: testTokenProcessorId,
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
		testutils_resource.ENUM_SOURCE_ID: resource.SourceId,
		testutils_resource.ENUM_TARGET_ID: resource.TargetId,
	}
}

func deleteTokenProcessorToTokenGeneratorMapping(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.TokenProcessorToTokenGeneratorMappingsAPI.DeleteTokenToTokenMappingById(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteTokenToTokenMappingById", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
