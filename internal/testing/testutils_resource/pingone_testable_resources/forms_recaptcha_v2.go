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

func FormsRecaptchaV2(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createFormsRecaptchaV2,
		DeleteFunc:         deleteFormsRecaptchaV2,
		Dependencies:       nil,
		ExportableResource: resources.FormsRecaptchaV2(clientInfo),
	}
}

func createFormsRecaptchaV2(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 0 {
		t.Errorf("Unexpected number of arguments provided to createFormsRecaptchaV2(): %v", strArgs)
		return testutils_resource.ResourceInfo{}
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.RecaptchaConfigurationApi.UpdateRecaptchaConfiguration(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID)
	clientStruct := management.RecaptchaConfiguration{
		SiteKey:   "siteKey",
		SecretKey: "secretKey",
	}

	request = request.RecaptchaConfiguration(clientStruct)

	_, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "UpdateRecaptchaConfiguration", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
		return testutils_resource.ResourceInfo{}
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
		return testutils_resource.ResourceInfo{}
	}

	return testutils_resource.ResourceInfo{
		DeletionIds:  []string{},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{},
	}
}

func deleteFormsRecaptchaV2(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	request := clientInfo.PingOneApiClient.ManagementAPIClient.RecaptchaConfigurationApi.DeleteRecaptchaConfiguration(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteRecaptchaConfiguration", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
		return
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
		return
	}
}
