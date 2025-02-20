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
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneBrandingSettingsResource
func BrandingSettings(clientInfo *connector.PingOneClientInfo) *PingOneBrandingSettingsResource {
	return &PingOneBrandingSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneBrandingSettingsResource) ResourceType() string {
	return "pingone_branding_settings"
}

func (r *PingOneBrandingSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	ok, err := checkBrandingSettingsData(r.clientInfo, r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return &importBlocks, nil
	}

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

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}

func checkBrandingSettingsData(clientInfo *connector.PingOneClientInfo, resourceType string) (bool, error) {
	_, response, err := clientInfo.ApiClient.ManagementAPIClient.BrandingSettingsApi.ReadBrandingSettings(clientInfo.Context, clientInfo.ExportEnvironmentID).Execute()
	return common.CheckSingletonResource(response, err, "ReadBrandingSettings", resourceType)
}
