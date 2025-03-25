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

func Form(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createForm,
		DeleteFunc:         deleteForm,
		Dependencies:       nil,
		ExportableResource: resources.Form(clientInfo),
	}
}

func createForm(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 0 {
		t.Fatalf("Unexpected number of arguments provided to createForm(): %v", strArgs)
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.FormManagementApi.CreateForm(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID)
	clientStruct := management.Form{
		Name:              "TestFormName",
		Description:       utils.Pointer("TestFormDescription"),
		Category:          management.ENUMFORMCATEGORY_CUSTOM,
		Cols:              utils.Pointer(int32(4)),
		MarkOptional:      true,
		MarkRequired:      false,
		TranslationMethod: utils.Pointer(management.ENUMFORMTRANSLATIONMETHOD_TRANSLATE),
		FieldTypes: []management.EnumFormFieldType{
			management.ENUMFORMFIELDTYPE_ERROR_DISPLAY,
			management.ENUMFORMFIELDTYPE_TEXT,
			management.ENUMFORMFIELDTYPE_SUBMIT_BUTTON,
		},
		LanguageBundle: &map[string]string{
			"button.text": "Submit",
		},
	}

	request = request.Form(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateForm", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
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

func deleteForm(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Fatalf("Unexpected number of arguments provided to deleteForm(): %v", ids)
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.FormManagementApi.DeleteForm(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteForm", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
