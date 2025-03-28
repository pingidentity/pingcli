// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingfederate_testable_resources

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func KeypairsOauthOpenidConnectAdditionalKeySet(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createKeypairsOauthOpenidConnectAdditionalKeySet,
		DeleteFunc: deleteKeypairsOauthOpenidConnectAdditionalKeySet,
		Dependencies: []*testutils_resource.TestableResource{
			OauthIssuer(t, clientInfo),
			KeypairsSigningKey(t, clientInfo),
		},
		ExportableResource: resources.KeypairsOauthOpenidConnectAdditionalKeySet(clientInfo),
	}
}

func createKeypairsOauthOpenidConnectAdditionalKeySet(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Errorf("Unexpected number of arguments provided to createKeypairsOauthOpenidConnectAdditionalKeySet(): %v", strArgs)

		return testutils_resource.ResourceInfo{}
	}
	testOauthIssuerId := strArgs[0]
	testKeyPairId := strArgs[1]

	request := clientInfo.PingFederateApiClient.KeyPairsOauthOpenIdConnectAPI.CreateKeySet(clientInfo.PingFederateContext)
	clientStruct := client.AdditionalKeySet{
		Id: utils.Pointer("TestAdditionalKeySetId"),
		Issuers: []client.ResourceLink{
			{
				Id: testOauthIssuerId,
			},
		},
		Name: "TestAdditionalKeySetName",
		SigningKeys: client.SigningKeys{
			RsaActiveCertRef: &client.ResourceLink{
				Id: testKeyPairId,
			},
		},
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateKeySet", resourceType)
	if err != nil {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return testutils_resource.ResourceInfo{}
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return testutils_resource.ResourceInfo{}
	}

	return testutils_resource.ResourceInfo{
		DeletionIds: []string{
			*resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   *resource.Id,
			testutils_resource.ENUM_NAME: resource.Name,
		},
	}
}

func deleteKeypairsOauthOpenidConnectAdditionalKeySet(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteKeypairsOauthOpenidConnectAdditionalKeySet(): %v", ids)

		return
	}

	request := clientInfo.PingFederateApiClient.KeyPairsOauthOpenIdConnectAPI.DeleteKeySet(clientInfo.PingFederateContext, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteKeySet", resourceType)
	if err != nil {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
