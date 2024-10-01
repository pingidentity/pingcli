package resources

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneMFASettingsResource{}
)

type PingoneMFASettingsResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneMFASettingsResource
func MFASettings(clientInfo *connector.PingOneClientInfo) *PingoneMFASettingsResource {
	return &PingoneMFASettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneMFASettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	_, response, err := r.clientInfo.ApiClient.MFAAPIClient.MFASettingsApi.ReadMFASettings(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	err = common.HandleClientResponse(response, err, "ReadMFASettings", r.ResourceType())
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
		ResourceName:       "mfa_settings",
		ResourceID:         r.clientInfo.ExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingoneMFASettingsResource) ResourceType() string {
	return "pingone_mfa_settings"
}
