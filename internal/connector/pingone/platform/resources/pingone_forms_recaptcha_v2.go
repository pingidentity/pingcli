package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneFormRecaptchaV2Resource{}
)

type PingoneFormRecaptchaV2Resource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneFormRecaptchaV2Resource
func FormRecaptchaV2(clientInfo *connector.PingOneClientInfo) *PingoneFormRecaptchaV2Resource {
	return &PingoneFormRecaptchaV2Resource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneFormRecaptchaV2Resource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	// Fetch FormRecaptchaV2 Resource from API.
	// If response is 204 No Content, then return empty import blocks.
	_, response, err := r.clientInfo.ApiClient.ManagementAPIClient.RecaptchaConfigurationApi.ReadRecaptchaConfiguration(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	err = common.HandleClientResponse(response, err, "ReadRecaptchaConfiguration", r.ResourceType())
	if err != nil {
		return nil, err
	}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	if response.StatusCode == 204 {
		l.Debug().Msgf("No exportable %s resource found", r.ResourceType())
		return &importBlocks, nil
	}

	commentData := map[string]string{
		"Resource Type":         r.ResourceType(),
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       "recaptcha_configuration",
		ResourceID:         r.clientInfo.ExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingoneFormRecaptchaV2Resource) ResourceType() string {
	return "pingone_forms_recaptcha_v2"
}
