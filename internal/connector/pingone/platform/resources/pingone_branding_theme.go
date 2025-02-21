package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneBrandingThemeResource{}
)

type PingOneBrandingThemeResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneBrandingThemeResource
func BrandingTheme(clientInfo *connector.PingOneClientInfo) *PingOneBrandingThemeResource {
	return &PingOneBrandingThemeResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneBrandingThemeResource) ResourceType() string {
	return "pingone_branding_theme"
}

func (r *PingOneBrandingThemeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	brandingThemeData, err := r.getBrandingThemeData()
	if err != nil {
		return nil, err
	}

	for brandingThemeId, brandingThemeName := range brandingThemeData {
		commentData := map[string]string{
			"Branding Theme ID":     brandingThemeId,
			"Branding Theme Name":   brandingThemeName,
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       brandingThemeName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, brandingThemeId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneBrandingThemeResource) getBrandingThemeData() (map[string]string, error) {
	brandingThemeData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.BrandingThemesApi.ReadBrandingThemes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	brandingThemes, err := pingone.GetManagementAPIObjectsFromIterator[management.BrandingTheme](iter, "ReadBrandingThemes", "GetThemes", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, brandingTheme := range brandingThemes {
		brandingThemeId, brandingThemeIdOk := brandingTheme.GetIdOk()
		brandingThemeConfiguration, brandingThemeConfigurationOk := brandingTheme.GetConfigurationOk()

		if brandingThemeIdOk && brandingThemeConfigurationOk {
			brandingThemeName, brandingThemeNameOk := brandingThemeConfiguration.GetNameOk()

			if brandingThemeNameOk {
				brandingThemeData[*brandingThemeId] = *brandingThemeName
			}
		}
	}

	return brandingThemeData, nil
}
