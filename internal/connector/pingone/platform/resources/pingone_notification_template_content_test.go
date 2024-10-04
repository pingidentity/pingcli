package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestNotificationTemplateContentExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.NotificationTemplateContent(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Voice_en_Test Duplication on Device Pairing",
			ResourceID:   fmt.Sprintf("%s/device_pairing/d4ed6d8d-1b54-4903-970f-1c9896eed55d", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Email_en_Test Duplication on Device Pairing",
			ResourceID:   fmt.Sprintf("%s/device_pairing/2acfe36d-065c-465e-be21-cb95e46cee45", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_SMS_en_Test Duplication on Device Pairing",
			ResourceID:   fmt.Sprintf("%s/device_pairing/f67b076d-bb78-4cbd-b945-f721be9c88f6", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Voice_en",
			ResourceID:   fmt.Sprintf("%s/device_pairing/d02693ae-8809-4a7f-a7f9-da9f272c8096", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Email_en",
			ResourceID:   fmt.Sprintf("%s/device_pairing/625d98de-9f2d-4e1b-8417-d0ba139d36b2", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_SMS_en",
			ResourceID:   fmt.Sprintf("%s/device_pairing/d4ca6154-bf87-4201-825b-6a1fecbd66ac", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_SMS_fr",
			ResourceID:   fmt.Sprintf("%s/device_pairing/c21bda2c-64b4-7025-2c83-d04d0f72077f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "email_verification_admin_Email_en",
			ResourceID:   fmt.Sprintf("%s/email_verification_admin/b130f9a6-a422-72c0-3afa-105d5f8fbb88", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "email_verification_user_Email_en",
			ResourceID:   fmt.Sprintf("%s/email_verification_user/5eda6f7b-59c6-7c22-3348-9821179c2b37", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_Voice_fr",
			ResourceID:   fmt.Sprintf("%s/general/831e9b77-5a05-7ed1-0fa6-c8cb637b5904", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_SMS_fr",
			ResourceID:   fmt.Sprintf("%s/general/63501c32-723c-7d4c-1f93-4e3c8c0cb292", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_SMS_en",
			ResourceID:   fmt.Sprintf("%s/general/1dd4c1a3-802b-70c0-3d10-5524eb9defc7", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_SMS_fr",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/a5dacd1c-c395-74ab-216f-a17037b22cf6", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_Email_en",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/995558d3-39a9-72bf-32a6-e3c1e395aa1f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Voice_en",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/b5664ec8-a329-4ced-92ce-bbab388e7329", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Email_en",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/d1235c66-48c6-46ae-ae6d-599513ab26d7", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_SMS_en",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/41054e31-dacd-4591-a8c8-f44cbec6313f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Push_en",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/c6b2f1e9-fcde-4b64-b473-f5370219da76", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_SMS_fr",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/807cd1a1-f3f8-7440-10f5-5f9cf944abb3", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Push_fr",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/d3d66f1b-b748-7afc-2d4b-a1daffd50a77", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_SMS_fr",
			ResourceID:   fmt.Sprintf("%s/transaction/10458132-7361-7d6b-3e42-04128ae31625", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_Email_fr",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/03bdf108-c71d-74fb-28e8-143f22b98125", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_Email_en",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/93688f61-e554-736d-227d-ac8ee610c254", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
