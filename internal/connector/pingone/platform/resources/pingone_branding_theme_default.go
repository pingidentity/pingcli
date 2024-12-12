package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneBrandingThemeDefaultResource{}
)

type PingOneBrandingThemeDefaultResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneBrandingThemeDefaultResource
func BrandingThemeDefault(clientInfo *connector.PingOneClientInfo) *PingOneBrandingThemeDefaultResource {
	return &PingOneBrandingThemeDefaultResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneBrandingThemeDefaultResource) ResourceType() string {
	return "pingone_branding_theme_default"
}

func (r *PingOneBrandingThemeDefaultResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	defaultBrandingThemeName, err := r.getDefaultBrandingThemeName()
	if err != nil {
		return nil, err
	}

	commentData := map[string]string{
		"Default Branding Theme Name": *defaultBrandingThemeName,
		"Export Environment ID":       r.clientInfo.ExportEnvironmentID,
		"Resource Type":               r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_default_theme", *defaultBrandingThemeName),
		ResourceID:         r.clientInfo.ExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}

func (r *PingOneBrandingThemeDefaultResource) getDefaultBrandingThemeName() (*string, error) {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.BrandingThemesApi.ReadBrandingThemes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadBrandingThemes", r.ResourceType())
		if err != nil {
			return nil, err
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, brandingTheme := range embedded.GetThemes() {
			brandingThemeDefault, brandingThemeDefaultOk := brandingTheme.GetDefaultOk()

			if brandingThemeDefaultOk && *brandingThemeDefault {
				brandingThemeConfiguration, brandingThemeConfigurationOk := brandingTheme.GetConfigurationOk()

				if brandingThemeConfigurationOk {
					brandingThemeName, brandingThemeNameOk := brandingThemeConfiguration.GetNameOk()

					if brandingThemeNameOk {
						return brandingThemeName, nil
					}
				}
			}
		}
	}

	return nil, fmt.Errorf("failed to export resource '%s'. No default branding theme found.", r.ResourceType())
}
