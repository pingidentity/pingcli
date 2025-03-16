package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestSchemaAttributeExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.SchemaAttribute(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_preferredLanguage",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/364bc187-e88f-4853-87f3-64aa13d9a099", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_timezone",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/4d3de681-a822-4633-bc42-8c67f9052fd3", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_lastSignOn",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/6b6992f5-78f6-4a22-97a1-69ba30c591d0", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_title",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/47cdeaa0-5cf0-4964-83b5-b3fe125c092e", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_type",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/acce5383-16ff-4973-8ded-2b19fd9146ed", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_locale",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/b9ff90eb-188e-40b1-9725-92b55e40f1eb", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_enabled",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/022607db-04b7-4d37-a034-798342d32060", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_identityProvider",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/b77ef54a-54c1-4636-83d2-b410ed23aeee", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_lifecycle",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/e3541156-0fe1-4177-aa69-dce02420d8cc", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_createdAt",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/d3221cc9-fb62-42e9-a14c-971a7c7a1e74", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_verifyStatus",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/23a02ffd-9250-401e-8aa5-f8eb71b72c6c", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_nickname",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/87efff27-cdb5-4829-9976-a80ebb4f8ee5", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_mfaEnabled",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/a49c17e2-ce8f-45e5-8e71-d51c8c4d140a", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_id",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/e4bb0b09-3f8d-485e-94ca-20e312471633", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_email",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/0473af5a-1294-4462-8a19-8567e5dccd9c", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_updatedAt",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/8fe405d2-c620-4267-805d-371c2092eb59", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_memberOfGroupIDs",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/d85f4528-54a8-49c7-a643-c098ad28b860", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_address",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/bc10ef33-b7cf-4efd-afc2-44bbd8f572a9", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_externalId",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/985359b0-a6a7-49e3-9079-be770e49b37f", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_photo",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/d991bc82-002d-4872-b544-9f2562452269", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_memberOfGroupNames",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/6a2aa3a6-9926-4070-8827-3bf84f7033fb", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_population",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/fb572f33-8944-4a35-846c-e548dbdeb49f", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_primaryPhone",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/f48f844b-3ba2-45ad-ba3c-de473a12ca4d", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_accountId",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/3658a298-8ce8-446d-ada5-cebb24678506", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_emailVerified",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/c4034a55-f6ae-406e-b3ad-5da3c66d77a2", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_mobilePhone",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/b4a939bf-f60a-41c3-9aad-1482ddf31d32", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_name",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/c8cac0ca-31c6-43d4-a2e6-63b07c936a43", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_account",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/eeed302f-8ca8-4993-aeb0-5d8d08587d8d", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_username",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/77d3f22e-00ca-49d1-98a1-fc0ee48d2542", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_schema_attribute",
			ResourceName: "User_bypassMFA",
			ResourceID:   fmt.Sprintf("%s/ff3cb03d-4896-4d20-8612-f014c4048d01/355c50dc-0eb6-4c5b-ab36-2b3152e0534c", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
