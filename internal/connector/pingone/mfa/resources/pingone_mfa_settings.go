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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneMFASettingsResource
func MFASettings(clientInfo *connector.ClientInfo) *PingOneMFASettingsResource {
	return &PingOneMFASettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneMFASettingsResource) ResourceType() string {
	return "pingone_mfa_settings"
}

func (r *PingOneMFASettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	ok, err := r.checkMFASettingsData()
	if err != nil {
		return nil, err
	}
	if !ok {
		return &importBlocks, nil
	}

	commentData := map[string]string{
		"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       r.ResourceType(),
		ResourceID:         r.clientInfo.PingOneExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}

func (r *PingOneMFASettingsResource) checkMFASettingsData() (bool, error) {
	_, response, err := r.clientInfo.PingOneApiClient.MFAAPIClient.MFASettingsApi.ReadMFASettings(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	return common.CheckSingletonResource(response, err, "ReadMFASettings", r.ResourceType())
}
