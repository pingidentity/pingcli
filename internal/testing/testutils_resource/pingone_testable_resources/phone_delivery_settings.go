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
	"github.com/pingidentity/pingcli/internal/utils"
)

func PhoneDeliverySettings(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createPhoneDeliverySettings,
		DeleteFunc:         deletePhoneDeliverySettings,
		Dependencies:       nil,
		ExportableResource: resources.PhoneDeliverySettings(clientInfo),
	}
}

func createPhoneDeliverySettings(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 0 {
		t.Fatalf("Unexpected number of arguments provided to createPhoneDeliverySettings(): %v", strArgs)
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.PhoneDeliverySettingsApi.CreatePhoneDeliverySettings(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID)
	clientStruct := management.NotificationsSettingsPhoneDeliverySettings{
		NotificationsSettingsPhoneDeliverySettingsCustom: &management.NotificationsSettingsPhoneDeliverySettingsCustom{
			Name: "CustomProviderName1",
			Authentication: management.NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication{
				Method:   management.ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSCUSTOMAUTHMETHOD_BASIC,
				Username: utils.Pointer("testUsername"),
				Password: utils.Pointer("testPassword1"),
			},
			Requests: []management.NotificationsSettingsPhoneDeliverySettingsCustomRequest{
				{
					DeliveryMethod: management.ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSCUSTOMDELIVERYMETHOD_SMS,
					Url:            "https://example.com/message",
					Method:         management.ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSCUSTOMREQUESTMETHOD_POST,
					Body:           utils.Pointer("ARN&message=${message}&phoneNumber=${to}&sender=${from}"),
				},
				{
					DeliveryMethod: management.ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSCUSTOMDELIVERYMETHOD_VOICE,
					Url:            "https://example.com/voice",
					Method:         management.ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSCUSTOMREQUESTMETHOD_POST,
					Body:           utils.Pointer("ARN&message=${message}&phoneNumber=${to}&sender=${from}"),
				},
			},
			Numbers: []management.NotificationsSettingsPhoneDeliverySettingsCustomNumbers{
				{
					Type:      management.ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSCUSTOMNUMBERSTYPE_PHONE_NUMBER,
					Number:    "+1 222 333",
					Selected:  utils.Pointer(true),
					Available: utils.Pointer(true),
					Capabilities: []management.EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersCapability{
						management.ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSCUSTOMNUMBERSCAPABILITY_SMS,
						management.ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSCUSTOMNUMBERSCAPABILITY_VOICE,
					},
				},
				{
					Type:      management.ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSCUSTOMNUMBERSTYPE_TOLL_FREE,
					Number:    "+18544440099",
					Selected:  utils.Pointer(false),
					Available: utils.Pointer(true),
					Capabilities: []management.EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersCapability{
						management.ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSCUSTOMNUMBERSCAPABILITY_SMS,
						management.ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSCUSTOMNUMBERSCAPABILITY_VOICE,
					},
					SupportedCountries: []string{
						"US",
						"CA",
					},
				},
			},
		},
	}

	request = request.NotificationsSettingsPhoneDeliverySettings(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreatePhoneDeliverySettings", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		DepIds: []string{},
		SelfInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   *resource.NotificationsSettingsPhoneDeliverySettingsCustom.Id,
			testutils_resource.ENUM_NAME: resource.NotificationsSettingsPhoneDeliverySettingsCustom.Name,
		},
	}
}

func deletePhoneDeliverySettings(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Fatalf("Unexpected number of arguments provided to deletePhoneDeliverySettings(): %v", ids)
	}
	id := ids[0]

	request := clientInfo.PingOneApiClient.ManagementAPIClient.PhoneDeliverySettingsApi.DeletePhoneDeliverySettings(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeletePhoneDeliverySettings", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
