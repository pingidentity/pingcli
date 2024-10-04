package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneBrandingSettingsResource{}
)

type PingoneBrandingSettingsResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneBrandingSettingsResource
func BrandingSettings(clientInfo *connector.PingOneClientInfo) *PingoneBrandingSettingsResource {
	return &PingoneBrandingSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneBrandingSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	_, response, err := r.clientInfo.ApiClient.ManagementAPIClient.BrandingSettingsApi.ReadBrandingSettings(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	err = common.HandleClientResponse(response, err, "ReadBrandingSettings", r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

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
		ResourceName:       "branding_settings",
		ResourceID:         r.clientInfo.ExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingoneBrandingSettingsResource) ResourceType() string {
	return "pingone_branding_settings"
}
