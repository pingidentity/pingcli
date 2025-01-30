package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateCaptchaProvider(t *testing.T) *testutils_resource.TestableResource {
	t.Helper()

	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	return &testutils_resource.TestableResource{
		ClientInfo:         pingfederateClientInfo,
		ExportableResource: resources.AuthenticationApiApplication(pingfederateClientInfo),
		TestResource: testutils_resource.TestResource{
			Dependencies: nil,
			CreateFunc:   createCaptchaProvider,
			DeleteFunc:   deleteCaptchaProvider,
		},
	}
}

func Test_PingFederateCaptchaProvider(t *testing.T) {
	tr := TestableResource_PingFederateCaptchaProvider(t)

	creationInfo := tr.CreateResource(t, tr.TestResource)
	defer tr.DeleteResource(t, tr.TestResource)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: creationInfo[testutils_resource.ENUM_NAME],
			ResourceID:   creationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}

func createCaptchaProvider(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createCaptchaProvider(): %v", strArgs)
	}
	resourceType := strArgs[0]

	request := clientInfo.PingFederateApiClient.CaptchaProvidersAPI.CreateCaptchaProvider(clientInfo.Context)
	result := client.CaptchaProvider{
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

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteCaptchaProvider(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.CaptchaProvidersAPI.DeleteCaptchaProvider(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteCaptchaProvider", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
