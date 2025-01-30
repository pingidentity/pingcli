package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneFormRecaptchaV2Resource{}
)

type PingOneFormRecaptchaV2Resource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneFormRecaptchaV2Resource
func FormRecaptchaV2(clientInfo *connector.PingOneClientInfo) *PingOneFormRecaptchaV2Resource {
	return &PingOneFormRecaptchaV2Resource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneFormRecaptchaV2Resource) ResourceType() string {
	return "pingone_forms_recaptcha_v2"
}

func (r *PingOneFormRecaptchaV2Resource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	ok, err := r.checkFormRecaptchaV2Data()
	if err != nil {
		return nil, err
	}

	if !ok {
		l.Debug().Msgf("No '%s' resources to export. Skipping...", r.ResourceType())
		return &importBlocks, nil
	}

	commentData := map[string]string{
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       r.ResourceType(),
		ResourceID:         r.clientInfo.ExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}

func (r *PingOneFormRecaptchaV2Resource) checkFormRecaptchaV2Data() (bool, error) {
	_, response, err := r.clientInfo.ApiClient.ManagementAPIClient.RecaptchaConfigurationApi.ReadRecaptchaConfiguration(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	err = common.HandleClientResponse(response, err, "ReadRecaptchaConfiguration", r.ResourceType())
	if err != nil {
		return false, err
	}

	if response.StatusCode == 204 {
		return false, nil
	}

	return true, nil
}
