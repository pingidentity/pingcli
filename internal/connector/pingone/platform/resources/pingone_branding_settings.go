package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneBrandingSettingsResource{}
)

type PingOneBrandingSettingsResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneBrandingSettingsResource
func BrandingSettings(clientInfo *connector.PingOneClientInfo) *PingOneBrandingSettingsResource {
	return &PingOneBrandingSettingsResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneBrandingSettingsResource) ResourceType() string {
	return "pingone_branding_settings"
}

func (r *PingOneBrandingSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportBrandingSettings()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneBrandingSettingsResource) exportBrandingSettings() error {
	_, response, err := r.clientInfo.ApiClient.ManagementAPIClient.BrandingSettingsApi.ReadBrandingSettings(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	err = common.HandleClientResponse(response, err, "ReadBrandingSettings", r.ResourceType())
	if err != nil {
		return err
	}

	if response.StatusCode == 204 {
		return common.DataNilError(r.ResourceType(), response)
	}

	r.addImportBlock()

	return nil
}

func (r *PingOneBrandingSettingsResource) addImportBlock() {
	commentData := map[string]string{
		"Resource Type":         r.ResourceType(),
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       r.ResourceType(),
		ResourceID:         r.clientInfo.ExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
