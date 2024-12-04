package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneBrandingThemeResource{}
)

type PingOneBrandingThemeResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneBrandingThemeResource
func BrandingTheme(clientInfo *connector.PingOneClientInfo) *PingOneBrandingThemeResource {
	return &PingOneBrandingThemeResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneBrandingThemeResource) ResourceType() string {
	return "pingone_branding_theme"
}

func (r *PingOneBrandingThemeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportBrandingThemes()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneBrandingThemeResource) exportBrandingThemes() error {
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
			brandingThemeId, brandingThemeIdOk := brandingTheme.GetIdOk()
			brandingThemeConfiguration, brandingThemeConfigurationOk := brandingTheme.GetConfigurationOk()

			if brandingThemeIdOk && brandingThemeConfigurationOk {
				brandingThemeName, brandingThemeNameOk := brandingThemeConfiguration.GetNameOk()

				if brandingThemeNameOk {
					r.addImportBlock(*brandingThemeId, *brandingThemeName)
				}
			}
		}
	}

	return nil
}

func (r *PingOneBrandingThemeResource) addImportBlock(brandingThemeId, brandingThemeName string) {
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

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
