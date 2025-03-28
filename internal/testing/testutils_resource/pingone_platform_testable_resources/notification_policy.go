// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

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

func NotificationPolicy(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createNotificationPolicy,
		DeleteFunc:         deleteNotificationPolicy,
		Dependencies:       nil,
		ExportableResource: resources.NotificationPolicy(clientInfo),
	}
}

func createNotificationPolicy(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 0 {
		t.Errorf("Unexpected number of arguments provided to createNotificationPolicy(): %v", strArgs)

		return testutils_resource.ResourceInfo{}
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.NotificationsPoliciesApi.CreateNotificationsPolicy(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID)
	clientStruct := management.NotificationsPolicy{
		Name: "Notification policy with environment limit and country limitation",
		Quotas: []management.NotificationsPolicyQuotasInner{
			{
				Type: management.ENUMNOTIFICATIONSPOLICYQUOTAITEMTYPE_ENVIRONMENT,
				DeliveryMethods: []management.EnumNotificationsPolicyQuotaDeliveryMethods{
					management.ENUMNOTIFICATIONSPOLICYQUOTADELIVERYMETHODS_SMS,
					management.ENUMNOTIFICATIONSPOLICYQUOTADELIVERYMETHODS_VOICE,
				},
				Total: utils.Pointer(int32(1000)),
			},
		},
		CountryLimit: &management.NotificationsPolicyCountryLimit{
			Type: management.ENUMNOTIFICATIONSPOLICYCOUNTRYLIMITTYPE_ALLOWED,
			DeliveryMethods: []management.EnumNotificationsPolicyCountryLimitDeliveryMethod{
				management.ENUMNOTIFICATIONSPOLICYCOUNTRYLIMITDELIVERYMETHOD_SMS,
				management.ENUMNOTIFICATIONSPOLICYCOUNTRYLIMITDELIVERYMETHOD_VOICE,
			},
			Countries: []string{
				"US",
				"CA",
			},
		},
		Default: utils.Pointer(false),
	}

	request = request.NotificationsPolicy(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateNotificationsPolicy", resourceType)
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
			testutils_resource.ENUM_NAME: resource.Name,
		},
	}
}

func deleteNotificationPolicy(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteNotificationPolicy(): %v", ids)

		return
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.NotificationsPoliciesApi.DeleteNotificationsPolicy(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteNotificationsPolicy", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
