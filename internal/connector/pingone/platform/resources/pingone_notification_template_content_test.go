// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestNotificationTemplateContentExport(t *testing.T) {
	// TODO remove this skip dependent upon STAGING-25369
	t.SkipNow()

	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.NotificationTemplateContent(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_email_en_625d98de_9f2d_4e1b_8417_d0ba139d36b2",
			ResourceID:   fmt.Sprintf("%s/device_pairing/625d98de-9f2d-4e1b-8417-d0ba139d36b2", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_sms_en_d4ca6154_bf87_4201_825b_6a1fecbd66ac",
			ResourceID:   fmt.Sprintf("%s/device_pairing/d4ca6154-bf87-4201-825b-6a1fecbd66ac", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_voice_en_d4ed6d8d_1b54_4903_970f_1c9896eed55d",
			ResourceID:   fmt.Sprintf("%s/device_pairing/d4ed6d8d-1b54-4903-970f-1c9896eed55d", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_email_en_2acfe36d_065c_465e_be21_cb95e46cee45",
			ResourceID:   fmt.Sprintf("%s/device_pairing/2acfe36d-065c-465e-be21-cb95e46cee45", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_voice_fr_f50e80bc_e84d_7124_0db5_4fd4cf72d7c9",
			ResourceID:   fmt.Sprintf("%s/device_pairing/f50e80bc-e84d-7124-0db5-4fd4cf72d7c9", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_email_fr_a0a13d00_a249_7ad1_3f7e_b6ba77a55955",
			ResourceID:   fmt.Sprintf("%s/device_pairing/a0a13d00-a249-7ad1-3f7e-b6ba77a55955", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_sms_fr_c21bda2c_64b4_7025_2c83_d04d0f72077f",
			ResourceID:   fmt.Sprintf("%s/device_pairing/c21bda2c-64b4-7025-2c83-d04d0f72077f", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_email_en_14651a6a_945b_725b_321f_e13cbe0fd9c6",
			ResourceID:   fmt.Sprintf("%s/device_pairing/14651a6a-945b-725b-321f-e13cbe0fd9c6", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_sms_en_f67b076d_bb78_4cbd_b945_f721be9c88f6",
			ResourceID:   fmt.Sprintf("%s/device_pairing/f67b076d-bb78-4cbd-b945-f721be9c88f6", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_voice_en_56b27d33_0110_7670_2c16_f0fca48b6340",
			ResourceID:   fmt.Sprintf("%s/device_pairing/56b27d33-0110-7670-2c16-f0fca48b6340", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_voice_en_d02693ae_8809_4a7f_a7f9_da9f272c8096",
			ResourceID:   fmt.Sprintf("%s/device_pairing/d02693ae-8809-4a7f-a7f9-da9f272c8096", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_sms_en_2512b56d_e14d_7cbd_3667_e1663d44fa41",
			ResourceID:   fmt.Sprintf("%s/device_pairing/2512b56d-e14d-7cbd-3667-e1663d44fa41", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "email_verification_admin_email_en_b130f9a6_a422_72c0_3afa_105d5f8fbb88",
			ResourceID:   fmt.Sprintf("%s/email_verification_admin/b130f9a6-a422-72c0-3afa-105d5f8fbb88", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "email_verification_user_email_en_5eda6f7b_59c6_7c22_3348_9821179c2b37",
			ResourceID:   fmt.Sprintf("%s/email_verification_user/5eda6f7b-59c6-7c22-3348-9821179c2b37", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_sms_fr_63501c32_723c_7d4c_1f93_4e3c8c0cb292",
			ResourceID:   fmt.Sprintf("%s/general/63501c32-723c-7d4c-1f93-4e3c8c0cb292", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_email_en_92adace7_5056_7d40_1c7e_adc71e57cc3f",
			ResourceID:   fmt.Sprintf("%s/general/92adace7-5056-7d40-1c7e-adc71e57cc3f", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_sms_en_1dd4c1a3_802b_70c0_3d10_5524eb9defc7",
			ResourceID:   fmt.Sprintf("%s/general/1dd4c1a3-802b-70c0-3d10-5524eb9defc7", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_voice_en_72618f82_18ed_7b7c_19ac_0b5899d92f0c",
			ResourceID:   fmt.Sprintf("%s/general/72618f82-18ed-7b7c-19ac-0b5899d92f0c", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_email_fr_28524b17_b60b_7fa2_131d_59816ac19864",
			ResourceID:   fmt.Sprintf("%s/general/28524b17-b60b-7fa2-131d-59816ac19864", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_voice_fr_831e9b77_5a05_7ed1_0fa6_c8cb637b5904",
			ResourceID:   fmt.Sprintf("%s/general/831e9b77-5a05-7ed1-0fa6-c8cb637b5904", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_email_fr_a7c11013_ca70_7071_3955_d647568f95d2",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/a7c11013-ca70-7071-3955-d647568f95d2", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_email_en_995558d3_39a9_72bf_32a6_e3c1e395aa1f",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/995558d3-39a9-72bf-32a6-e3c1e395aa1f", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_sms_fr_a5dacd1c_c395_74ab_216f_a17037b22cf6",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/a5dacd1c-c395-74ab-216f-a17037b22cf6", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_sms_en_daef917c_3695_7347_1ed0_bb03d80198c2",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/daef917c-3695-7347-1ed0-bb03d80198c2", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "recovery_code_template_email_fr_c558cd3c_eb16_7158_38c3_d87fb2e320f0",
			ResourceID:   fmt.Sprintf("%s/recovery_code_template/c558cd3c-eb16-7158-38c3-d87fb2e320f0", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "recovery_code_template_email_en_dc6755cd_123b_71f6_2fe0_5b74d3789001",
			ResourceID:   fmt.Sprintf("%s/recovery_code_template/dc6755cd-123b-71f6-2fe0-5b74d3789001", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_push_en_c6b2f1e9_fcde_4b64_b473_f5370219da76",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/c6b2f1e9-fcde-4b64-b473-f5370219da76", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_voice_fr_c6de3d50_d766_7533_25cf_bd28f72e2f86",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/c6de3d50-d766-7533-25cf-bd28f72e2f86", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_sms_en_5d2b94bc_d264_7f79_048f_a0f062f66d98",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/5d2b94bc-d264-7f79-048f-a0f062f66d98", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_push_en_b368bc5e_0815_7d16_178c_2631a620e00c",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/b368bc5e-0815-7d16-178c-2631a620e00c", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_email_fr_2e103d6d_af8e_70fa_282f_91821ed778fd",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/2e103d6d-af8e-70fa-282f-91821ed778fd", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_push_fr_d3d66f1b_b748_7afc_2d4b_a1daffd50a77",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/d3d66f1b-b748-7afc-2d4b-a1daffd50a77", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_voice_en_8a888c96_4cb9_7941_0c07_0a4e99a54a04",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/8a888c96-4cb9-7941-0c07-0a4e99a54a04", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_sms_en_41054e31_dacd_4591_a8c8_f44cbec6313f",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/41054e31-dacd-4591-a8c8-f44cbec6313f", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_sms_fr_807cd1a1_f3f8_7440_10f5_5f9cf944abb3",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/807cd1a1-f3f8-7440-10f5-5f9cf944abb3", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_email_en_e8539132_48c6_7061_1309_f33e99599a3e",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/e8539132-48c6-7061-1309-f33e99599a3e", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_email_en_d1235c66_48c6_46ae_ae6d_599513ab26d7",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/d1235c66-48c6-46ae-ae6d-599513ab26d7", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_push_fr_453937a6_e95b_78b9_0d13_be9e17cdda89",
			ResourceID:   fmt.Sprintf("%s/transaction/453937a6-e95b-78b9-0d13-be9e17cdda89", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_sms_fr_10458132_7361_7d6b_3e42_04128ae31625",
			ResourceID:   fmt.Sprintf("%s/transaction/10458132-7361-7d6b-3e42-04128ae31625", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_voice_en_b2509f63_c86c_7f76_0fea_52472af67df2",
			ResourceID:   fmt.Sprintf("%s/transaction/b2509f63-c86c-7f76-0fea-52472af67df2", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_email_fr_1953c10e_53a9_7c44_13eb_560823acacf6",
			ResourceID:   fmt.Sprintf("%s/transaction/1953c10e-53a9-7c44-13eb-560823acacf6", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_push_en_926802db_4abb_7369_249a_ff3c63c6a7d1",
			ResourceID:   fmt.Sprintf("%s/transaction/926802db-4abb-7369-249a-ff3c63c6a7d1", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_email_en_996806a7_0c4e_744c_117b_6312e08225d3",
			ResourceID:   fmt.Sprintf("%s/transaction/996806a7-0c4e-744c-117b-6312e08225d3", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_sms_en_d9751d12_2f37_7d6e_3cc4_33368770f6da",
			ResourceID:   fmt.Sprintf("%s/transaction/d9751d12-2f37-7d6e-3cc4-33368770f6da", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_voice_fr_0db4c9f4_c1d6_7adf_1870_22f70f5e95a1",
			ResourceID:   fmt.Sprintf("%s/transaction/0db4c9f4-c1d6-7adf-1870-22f70f5e95a1", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_email_fr_03bdf108_c71d_74fb_28e8_143f22b98125",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/03bdf108-c71d-74fb-28e8-143f22b98125", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_email_en_93688f61_e554_736d_227d_ac8ee610c254",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/93688f61-e554-736d-227d-ac8ee610c254", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
