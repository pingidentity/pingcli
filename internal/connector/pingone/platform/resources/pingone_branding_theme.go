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
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneBrandingThemeResource
func BrandingTheme(clientInfo *connector.PingOneClientInfo) *PingOneBrandingThemeResource {
	return &PingOneBrandingThemeResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneBrandingThemeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.BrandingThemesApi.ReadBrandingThemes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadBrandingThemes"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, theme := range embedded.GetThemes() {
		themeId, themeIdOk := theme.GetIdOk()
		themeConfiguration, themeConfigurationOk := theme.GetConfigurationOk()
		var themeName *string
		var themeNameOk = false
		if themeConfigurationOk {
			themeName, themeNameOk = themeConfiguration.GetNameOk()
		}

		if themeIdOk && themeConfigurationOk && themeNameOk {
			commentData := map[string]string{
				"Resource Type":         r.ResourceType(),
				"Branding Theme Name":   *themeName,
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"Branding Theme ID":     *themeId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *themeName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *themeId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingOneBrandingThemeResource) ResourceType() string {
	return "pingone_branding_theme"
}
