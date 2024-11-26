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
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneMFASettingsResource
func MFASettings(clientInfo *connector.PingOneClientInfo) *PingOneMFASettingsResource {
	return &PingOneMFASettingsResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneMFASettingsResource) ResourceType() string {
	return "pingone_mfa_settings"
}

func (r *PingOneMFASettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportMFASettings()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneMFASettingsResource) exportMFASettings() error {
	_, response, err := r.clientInfo.ApiClient.MFAAPIClient.MFASettingsApi.ReadMFASettings(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	err = common.HandleClientResponse(response, err, "ReadMFASettings", r.ResourceType())
	if err != nil {
		return err
	}

	if response.StatusCode == 204 {
		return common.DataNilError(r.ResourceType(), response)
	}

	r.addImportBlock()

	return nil
}

func (r *PingOneMFASettingsResource) addImportBlock() {
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

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
