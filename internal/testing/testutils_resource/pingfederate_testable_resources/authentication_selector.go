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

func AuthenticationSelector(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createAuthenticationSelector,
		DeleteFunc:         deleteAuthenticationSelector,
		Dependencies:       nil,
		ExportableResource: resources.AuthenticationSelector(clientInfo),
	}
}

func createAuthenticationSelector(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 0 {
		t.Errorf("Unexpected number of arguments provided to createAuthenticationSelector(): %v", strArgs)

		return testutils_resource.ResourceInfo{}
	}

	request := clientInfo.PingFederateApiClient.AuthenticationSelectorsAPI.CreateAuthenticationSelector(clientInfo.PingFederateContext)
	clientStruct := client.AuthenticationSelector{
		Configuration: client.PluginConfiguration{
			Fields: []client.ConfigField{
				{
					Name:  "Header Name",
					Value: utils.Pointer("TestHeaderName"),
				},
			},
			Tables: []client.ConfigTable{
				{
					Name: "Results",
					Rows: []client.ConfigRow{
						{
							DefaultRow: utils.Pointer(false),
							Fields: []client.ConfigField{
								{
									Name:  "Match Expression",
									Value: utils.Pointer("TestMatchExpression"),
								},
							},
						},
					},
				},
			},
		},
		Id:   "TestAuthSelectorId",
		Name: "TestAuthSelectorName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "com.pingidentity.pf.selectors.http.HTTPHeaderAdapterSelector",
		},
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateAuthenticationSelector", resourceType)
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
			resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   resource.Id,
			testutils_resource.ENUM_NAME: resource.Name,
		},
	}
}

func deleteAuthenticationSelector(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteAuthenticationSelector(): %v", ids)

		return
	}

	request := clientInfo.PingFederateApiClient.AuthenticationSelectorsAPI.DeleteAuthenticationSelector(clientInfo.PingFederateContext, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteAuthenticationSelector", resourceType)
	if err != nil {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
