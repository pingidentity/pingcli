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
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	ok, err := checkMFASettingsData(r.clientInfo, r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
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

func checkMFASettingsData(clientInfo *connector.PingOneClientInfo, resourceType string) (bool, error) {
	_, response, err := clientInfo.ApiClient.MFAAPIClient.MFASettingsApi.ReadMFASettings(clientInfo.Context, clientInfo.ExportEnvironmentID).Execute()
	return common.CheckSingletonResource(response, err, "ReadMFASettings", resourceType)
}
