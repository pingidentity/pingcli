// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingone_testable_resources

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
)

func ApplicationSignOnPolicyAssignment(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createApplicationSignOnPolicyAssignment,
		DeleteFunc: deleteApplicationSignOnPolicyAssignment,
		Dependencies: []*testutils_resource.TestableResource{
			ApplicationDeviceAuthorization(t, clientInfo),
			SignOnPolicy(t, clientInfo),
		},
		ExportableResource: resources.ApplicationSignOnPolicyAssignment(clientInfo),
	}
}

func createApplicationSignOnPolicyAssignment(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Errorf("Unexpected number of arguments provided to createApplicationSignOnPolicyAssignment(): %v", strArgs)
		return testutils_resource.ResourceInfo{}
	}
	applicationId := strArgs[0]
	policyId := strArgs[1]

	request := clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationSignOnPolicyAssignmentsApi.CreateSignOnPolicyAssignment(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, applicationId)
	clientStruct := management.SignOnPolicyAssignment{
		Priority: 1,
		SignOnPolicy: management.SignOnPolicyActionCommonSignOnPolicy{
			Id: policyId,
		},
	}

	request = request.SignOnPolicyAssignment(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateSignOnPolicyAssignment", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
		return testutils_resource.ResourceInfo{}
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
		return testutils_resource.ResourceInfo{}
	}

	signOnPolicy, response, err := clientInfo.PingOneApiClient.ManagementAPIClient.SignOnPoliciesApi.ReadOneSignOnPolicy(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, resource.SignOnPolicy.Id).Execute()
	ok, err = common.HandleClientResponse(response, err, "ReadOneSignOnPolicy", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
		return testutils_resource.ResourceInfo{}
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
		return testutils_resource.ResourceInfo{}
	}

	return testutils_resource.ResourceInfo{
		DeletionIds: []string{
			applicationId,
			*resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   *resource.Id,
			testutils_resource.ENUM_NAME: signOnPolicy.Name,
		},
	}
}

func deleteApplicationSignOnPolicyAssignment(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 2 {
		t.Errorf("Unexpected number of arguments provided to deleteApplicationSignOnPolicyAssignment(): %v", ids)
		return
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationSignOnPolicyAssignmentsApi.DeleteSignOnPolicyAssignment(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0], ids[1])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteSignOnPolicyAssignment", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
		return
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
		return
	}
}
