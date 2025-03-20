// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestResourceAttributeExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.ResourceAttribute(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "authorize-api-service_sub",
			ResourceID:   fmt.Sprintf("%s/3c6001a0-6110-4934-9d34-fa8c4a2894c2/3f79ad4f-588a-4652-bad8-d64d405bef8a", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "test_sub",
			ResourceID:   fmt.Sprintf("%s/4b9ef858-62ce-4bd0-9186-997b8527529d/c82b24b9-7ea3-4de4-8840-50b6c3cb1387", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "testing_sub",
			ResourceID:   fmt.Sprintf("%s/52afd89f-f3c0-4c78-b896-432c0a07329b/a7cf0daf-0e30-4ae5-bf88-7c5dc629d7cf", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_locale",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/0122f755-ac5d-4bc1-a755-0f56b6f582ec", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_preferred_username",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/09c1e110-7b3b-4f2d-a1ab-7d3054df8aa6", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_email",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/0b230ac5-25a4-4012-a393-e2529d91d4df", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_email_verified",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/1754e8e0-97b5-4477-b76e-a97d6d4fcb8d", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_phone_number",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/1b6c9496-a281-4379-a8ef-dc5b60cb1bf4", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_nickname",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/1c822a3c-52fb-4cb2-b8c7-99800d32221a", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_address.postal_code",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/1fdeb8af-3f4d-4979-a4eb-2344694f9ec2", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_name",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/28ce7067-77cb-460c-ad61-c7300b6b2ceb", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_updated_at",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/310677e7-2ece-4740-a17a-ec5cd9412b5c", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_family_name",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/37c492db-521d-4ef5-9fb2-dec64bb1de1e", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_profile",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/3c3f3bfe-d096-4f0e-9f9e-1ce9633cac5d", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_address.formatted",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/3dc6971b-d7e8-4019-8502-d43dc0ced872", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_address.region",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/4dab2957-276b-4132-886f-fd217d21c01d", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_phone_number_verified",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/791968e4-08bb-4aa8-bfc3-a28287fe0070", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_website",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/8447184f-1d5c-43cd-951b-a15c924b5bae", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_given_name",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/8bf9debc-5f13-45e4-81ba-cae3bc1c0d77", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_birthdate",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/91234ec7-61e8-4c5b-83f3-a08388e1a5f7", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_address.locality",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/b1f8538f-2b55-43ff-9b78-238fcad14b9d", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_zoneinfo",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/b298df9c-75c8-4b5a-b1a9-97b71bce415f", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_gender",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/c24c29ad-14d9-407e-a7c9-acb22a4792ce", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_address.street_address",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/d1c9d1eb-f988-4983-93e1-97ed6f0d835f", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_address.country",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/d2964153-b987-4688-a5dc-09b7a1d52667", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_exampleAttribute",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/dc35abb7-79bc-4449-8fcc-265fbb39345f", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_picture",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/dc9cd3f2-2076-44d2-b760-100a2beb49db", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "openid_middle_name",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/fd6180af-b339-47bb-a9e3-6e02b69fb7ad", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_resource_attribute",
			ResourceName: "Undeployed Test API Service_sub",
			ResourceID:   fmt.Sprintf("%s/a35fe5ea-084c-4245-80f1-85f9eaf4f063/be2a2418-127a-4d60-9c04-88a5a568e25c", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
