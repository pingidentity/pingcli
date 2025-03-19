// Copyright © 2025 Ping Identity Corporation

package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func TestableResource_PingFederateCaptchaProvider(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createCaptchaProvider,
		DeleteFunc:         deleteCaptchaProvider,
		Dependencies:       nil,
		ExportableResource: resources.CaptchaProvider(clientInfo),
	}
}

func createCaptchaProvider(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createCaptchaProvider(): %v", strArgs)
	}
	resourceType := strArgs[0]

	request := clientInfo.PingFederateApiClient.CaptchaProvidersAPI.CreateCaptchaProvider(clientInfo.PingFederateContext)
	clientStruct := client.CaptchaProvider{
		Configuration: client.PluginConfiguration{
			Fields: []client.ConfigField{
				{
					Name:  "Site Key",
					Value: utils.Pointer("TestSiteKey"),
				},
				{
					Name:  "Secret Key",
					Value: utils.Pointer("TestSecretKey"),
				},
				{
					Name:  "Pass Score Threshold",
					Value: utils.Pointer("0.8"),
				},
			},
		},
		Id:   "TestCaptchaProviderId",
		Name: "TestCaptchaProviderName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "com.pingidentity.captcha.recaptchaV3.ReCaptchaV3Plugin",
		},
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateCaptchaProvider", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteCaptchaProvider(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.CaptchaProvidersAPI.DeleteCaptchaProvider(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteCaptchaProvider", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
