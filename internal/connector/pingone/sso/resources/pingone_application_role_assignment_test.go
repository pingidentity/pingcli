package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

func TestApplicationRoleAssignmentExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.ApplicationRoleAssignment(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "OAuth Worker App_PingFederate Crypto Administrator_1",
			ResourceID:   fmt.Sprintf("%s/9d6c443b-6329-4d3c-949e-880eda3b9599/d4aa4aec-c521-4538-ab76-8776355d2b22", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "OAuth Worker App_PingFederate User Administrator_2",
			ResourceID:   fmt.Sprintf("%s/9d6c443b-6329-4d3c-949e-880eda3b9599/9f431f95-8df7-43cb-8419-e2b3898ca8c4", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "OAuth Worker App_PingFederate Administrator_3",
			ResourceID:   fmt.Sprintf("%s/9d6c443b-6329-4d3c-949e-880eda3b9599/28607a1f-b0b3-4c43-8807-4bf8a93c8d07", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "OAuth Worker App_PingFederate Expression Administrator_4",
			ResourceID:   fmt.Sprintf("%s/9d6c443b-6329-4d3c-949e-880eda3b9599/cbd5b6a0-1748-4ca6-b252-e02fd843897e", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "PingOne DaVinci Connection_Identity Data Admin_1",
			ResourceID:   fmt.Sprintf("%s/7b621870-7124-4426-b432-6c675642afcb/4331fc1a-434c-4cee-ba2a-ceb57974550c", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "PingOne DaVinci Connection_DaVinci Admin_2",
			ResourceID:   fmt.Sprintf("%s/7b621870-7124-4426-b432-6c675642afcb/ebcdd4c7-0014-4eb5-9aa9-15af45795c15", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "PingOne DaVinci Connection_Environment Admin_3",
			ResourceID:   fmt.Sprintf("%s/7b621870-7124-4426-b432-6c675642afcb/9e1d7f96-c4a9-49d3-bb2d-d2b1fef197dd", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_Identity Data Admin_1",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/9225c10f-b902-4107-8aba-b15b219d6c0e", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_Client Application Developer_2",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/0081f0ab-d02c-4718-b10c-35fd48b82f47", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_Identity Data Read Only_3",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/a0f34409-4d1b-4b22-911a-7b4a61ac68b1", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_Identity Data Admin_4",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/970667f1-26d5-4021-809f-e5d17fe44a7d", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_Client Application Developer_5",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/785b582f-eaf2-4a0b-ac8e-b7c7f9665762", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_Identity Data Read Only_6",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/91562725-239b-4854-8cef-c4efe35ea77f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_Identity Data Admin_7",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/ed54c262-38ab-4874-a206-2d13e34f21fd", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_Client Application Developer_8",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/3f112aa9-b712-4388-821d-8f37a429b071", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_Identity Data Read Only_9",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/1395d969-6527-45f4-b356-4ef36a5d6349", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_PingFederate Crypto Administrator_10",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/c01ef5c4-74c4-4074-8929-b0836aa9a783", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_DaVinci Admin_11",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/9bdbe295-e199-4952-8717-3405112eccad", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_Organization Admin_12",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/b57756a8-d9c6-4fbc-95d4-9d2aabf801e0", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_Environment Admin_13",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/3e77cca6-8820-4eb6-bcfd-761cf4e74ad1", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_PingFederate User Administrator_14",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/6600fad1-82c4-412f-aa2c-22e8668d8c3a", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_Configuration Read Only_15",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/6f01ea75-5e04-45a5-8614-186b58f9eb4e", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_PingFederate Auditor_16",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/cf1edf79-fd13-4d72-a049-7bdc4377ee0c", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_PingFederate Administrator_17",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/530824c1-675f-4282-8a61-6567fc3afee6", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_DaVinci Admin Read Only_18",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/e82d85ed-8687-4724-87ad-7f138cdbe673", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_role_assignment",
			ResourceName: "Worker App_PingFederate Expression Administrator_19",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/c090f7c9-4419-447b-8316-baf3e70030bc", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
