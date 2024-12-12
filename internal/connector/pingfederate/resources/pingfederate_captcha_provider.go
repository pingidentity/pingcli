package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateCaptchaProviderResource{}
)

type PingFederateCaptchaProviderResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateCaptchaProviderResource
func CaptchaProvider(clientInfo *connector.PingFederateClientInfo) *PingFederateCaptchaProviderResource {
	return &PingFederateCaptchaProviderResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateCaptchaProviderResource) ResourceType() string {
	return "pingfederate_captcha_provider"
}

func (r *PingFederateCaptchaProviderResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	captchaProviderData, err := r.getCaptchaProviderData()
	if err != nil {
		return nil, err
	}

	for captchaProviderId, captchaProviderName := range *captchaProviderData {
		commentData := map[string]string{
			"Captcha Provider Resource ID":   captchaProviderId,
			"Captcha Provider Resource Name": captchaProviderId,
			"Resource Type":                  r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       captchaProviderName,
			ResourceID:         captchaProviderId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateCaptchaProviderResource) getCaptchaProviderData() (*map[string]string, error) {
	captchaProviderData := make(map[string]string)

	captchaProviders, response, err := r.clientInfo.ApiClient.CaptchaProvidersAPI.GetCaptchaProviders(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetCaptchaProviders", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if captchaProviders == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	captchaProvidersItems, captchaProvidersItemsOk := captchaProviders.GetItemsOk()
	if !captchaProvidersItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, captchaProvider := range captchaProvidersItems {
		captchaProviderId, captchaProviderIdOk := captchaProvider.GetIdOk()
		captchaProviderName, captchaProviderNameOk := captchaProvider.GetNameOk()

		if captchaProviderIdOk && captchaProviderNameOk {
			captchaProviderData[*captchaProviderId] = *captchaProviderName
		}
	}

	return &captchaProviderData, nil
}
