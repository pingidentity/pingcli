package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateCaptchaProviderExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.CaptchaProvider(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_captcha_provider",
			ResourceName: "exampleCaptchaProviderV2",
			ResourceID:   "exampleCaptchaProviderV2",
		},
		{
			ResourceType: "pingfederate_captcha_provider",
			ResourceName: "exampleCaptchaProvider",
			ResourceID:   "exampleCaptchaProvider",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
