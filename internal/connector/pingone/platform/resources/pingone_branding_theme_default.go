package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneBrandingThemeDefaultResource{}
)

type PingOneBrandingThemeDefaultResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneBrandingThemeDefaultResource
func BrandingThemeDefault(clientInfo *connector.PingOneClientInfo) *PingOneBrandingThemeDefaultResource {
	return &PingOneBrandingThemeDefaultResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneBrandingThemeDefaultResource) ResourceType() string {
	return "pingone_branding_theme_default"
}

func (r *PingOneBrandingThemeDefaultResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportBrandingThemeDefault()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneBrandingThemeDefaultResource) exportBrandingThemeDefault() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.BrandingThemesApi.ReadBrandingThemes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadBrandingThemes", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, brandingTheme := range embedded.GetThemes() {
			brandingThemeDefault, brandingThemeDefaultOk := brandingTheme.GetDefaultOk()

			if brandingThemeDefaultOk && *brandingThemeDefault {
				brandingThemeConfiguration, brandingThemeConfigurationOk := brandingTheme.GetConfigurationOk()

				if brandingThemeConfigurationOk {
					brandingThemeName, brandingThemeNameOk := brandingThemeConfiguration.GetNameOk()

					if brandingThemeNameOk {
						r.addImportBlock(*brandingThemeName)
					}
				}
			}
		}
	}

	return nil
}

func (r *PingOneBrandingThemeDefaultResource) addImportBlock(brandingThemeName string) {
	commentData := map[string]string{
		"Default Branding Theme Name": brandingThemeName,
		"Export Environment ID":       r.clientInfo.ExportEnvironmentID,
		"Resource Type":               r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       brandingThemeName,
		ResourceID:         r.clientInfo.ExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
