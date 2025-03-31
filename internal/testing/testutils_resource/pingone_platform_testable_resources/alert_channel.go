// Copyright © 2025 Ping Identity Corporation

package pingone_platform_testable_resources

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
)

func AlertChannel(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createAlertChannel,
		DeleteFunc:         deleteAlertChannel,
		Dependencies:       nil,
		ExportableResource: resources.AlertChannel(clientInfo),
	}
}

func createAlertChannel(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 0 {
		t.Errorf("Unexpected number of arguments provided to createAlertChannel(): %v", strArgs)

		return testutils_resource.ResourceInfo{}
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.AlertingApi.CreateAlertChannel(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID)
	clientStruct := management.AlertChannel{
		ChannelType: management.ENUMALERTCHANNELTYPE_EMAIL,
		AlertName:   utils.Pointer("Cert Expired Alert"),
		IncludeSeverities: []management.EnumAlertChannelSeverity{
			management.ENUMALERTCHANNELSEVERITY_ERROR,
		},
		IncludeAlertTypes: []management.EnumAlertChannelAlertType{
			management.ENUMALERTCHANNELALERTTYPE_CERTIFICATE_EXPIRED,
		},
		ExcludeAlertTypes: []management.EnumAlertChannelAlertType{},
		Addresses: []string{
			"example@example.com",
		},
	}

	request = request.AlertChannel(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateAlertChannel", resourceType)
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
			*resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   *resource.Id,
			testutils_resource.ENUM_NAME: *resource.AlertName,
		},
	}
}

func deleteAlertChannel(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteAlertChannel(): %v", ids)

		return
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.AlertingApi.DeleteAlertChannel(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteAlertChannel", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
