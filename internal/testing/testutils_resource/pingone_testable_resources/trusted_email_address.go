// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingone_testable_resources

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
)

func TrustedEmailAddress(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createTrustedEmailAddress,
		DeleteFunc: deleteTrustedEmailAddress,
		Dependencies: []*testutils_resource.TestableResource{
			TrustedEmailDomain(t, clientInfo),
		},
		ExportableResource: resources.TrustedEmailAddress(clientInfo),
	}
}

func createTrustedEmailAddress(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createTrustedEmailAddress(): %v", strArgs)
	}
	trustedEmailDomainId := strArgs[0]

	request := clientInfo.PingOneApiClient.ManagementAPIClient.TrustedEmailAddressesApi.CreateTrustedEmailAddress(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, trustedEmailDomainId)
	clientStruct := management.EmailDomainTrustedEmail{
		EmailAddress: "example@example.com",
	}

	request = request.EmailDomainTrustedEmail(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateTrustedEmailAddress", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		DepIds: []string{
			trustedEmailDomainId,
		},
		SelfInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   *resource.Id,
			testutils_resource.ENUM_NAME: resource.EmailAddress,
		},
	}
}

func deleteTrustedEmailAddress(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 2 {
		t.Fatalf("Unexpected number of arguments provided to deleteTrustedEmailAddress(): %v", ids)
	}
	trustedEmailDomainId := ids[0]
	TrustedEmailAddressId := ids[1]

	request := clientInfo.PingOneApiClient.ManagementAPIClient.TrustedEmailAddressesApi.DeleteTrustedEmailAddress(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, trustedEmailDomainId, TrustedEmailAddressId)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteTrustedEmailAddress", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
