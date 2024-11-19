package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneMFASettingsResource{}
)

type PingOneMFASettingsResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneMFASettingsResource
func MFASettings(clientInfo *connector.PingOneClientInfo) *PingOneMFASettingsResource {
	return &PingOneMFASettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneMFASettingsResource) ResourceType() string {
	return "pingone_mfa_settings"
}

func (r *PingOneMFASettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	_, response, err := r.clientInfo.ApiClient.MFAAPIClient.MFASettingsApi.ReadMFASettings(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	err = common.HandleClientResponse(response, err, "ReadMFASettings", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 204 {
		l.Debug().Msgf("No exportable %s resource found", r.ResourceType())
		return &importBlocks, nil
	}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	commentData := map[string]string{
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Resource Type":         r.ResourceType(),
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       r.ResourceType(),
		ResourceID:         r.clientInfo.ExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}
