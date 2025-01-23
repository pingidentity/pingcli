package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateCaptchaProvider_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.CaptchaProvider(PingFederateClientInfo)

	captchaProviderId, captchaProviderName := createCaptchaProvider(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteCaptchaProvider(t, PingFederateClientInfo, resource.ResourceType(), captchaProviderId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: captchaProviderName,
			ResourceID:   captchaProviderId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createCaptchaProvider(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.CaptchaProvidersAPI.CreateCaptchaProvider(clientInfo.Context)
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
	err = common.HandleClientResponse(response, err, "CreateCaptchaProvider", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteCaptchaProvider(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.CaptchaProvidersAPI.DeleteCaptchaProvider(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteCaptchaProvider", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
