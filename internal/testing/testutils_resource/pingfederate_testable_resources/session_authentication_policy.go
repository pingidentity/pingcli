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

func SessionAuthenticationPolicy(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createSessionAuthenticationPolicy,
		DeleteFunc: deleteSessionAuthenticationPolicy,
		Dependencies: []*testutils_resource.TestableResource{
			IdpAdapter(t, clientInfo),
		},
		ExportableResource: resources.SessionAuthenticationPolicy(clientInfo),
	}
}

func createSessionAuthenticationPolicy(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Errorf("Unexpected number of arguments provided to createSessionAuthenticationPolicy(): %v", strArgs)

		return testutils_resource.ResourceInfo{}
	}
	testIdpAdapterId := strArgs[1]

	request := clientInfo.PingFederateApiClient.SessionAPI.CreateSourcePolicy(clientInfo.PingFederateContext)
	clientStruct := client.AuthenticationSessionPolicy{
		AuthenticationSource: client.AuthenticationSource{
			SourceRef: client.ResourceLink{
				Id: testIdpAdapterId,
			},
			Type: "IDP_ADAPTER",
		},
		Id: utils.Pointer("TestSessionAuthenticationPolicyId"),
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateSourcePolicy", resourceType)
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
			testutils_resource.ENUM_ID:            *resource.Id,
			testutils_resource.ENUM_TYPE:          resource.AuthenticationSource.Type,
			testutils_resource.ENUM_SOURCE_REF_ID: resource.AuthenticationSource.SourceRef.Id,
		},
	}
}

func deleteSessionAuthenticationPolicy(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteSessionAuthenticationPolicy(): %v", ids)

		return
	}

	request := clientInfo.PingFederateApiClient.SessionAPI.DeleteSourcePolicy(clientInfo.PingFederateContext, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteSourcePolicy", resourceType)
	if err != nil {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
