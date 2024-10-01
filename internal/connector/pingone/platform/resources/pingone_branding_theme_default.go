package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneBrandingThemeDefaultResource{}
)

type PingoneBrandingThemeDefaultResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneBrandingThemeDefaultResource
func BrandingThemeDefault(clientInfo *connector.PingOneClientInfo) *PingoneBrandingThemeDefaultResource {
	return &PingoneBrandingThemeDefaultResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneBrandingThemeDefaultResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.BrandingThemesApi.ReadBrandingThemes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadBrandingThemes"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	foundDefault := false
	var defaultBrandingTheme management.BrandingTheme

	for _, brandingTheme := range embedded.GetThemes() {
		if brandingTheme.GetDefault() {
			foundDefault = true
			defaultBrandingTheme = brandingTheme
			break
		}
	}

	if !foundDefault {
		l.Debug().Msgf("No exportable %s resource found", r.ResourceType())
		return &[]connector.ImportBlock{}, nil
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	defaultBrandingThemeConfiguration, defaultBrandingThemeConfigurationOk := defaultBrandingTheme.GetConfigurationOk()
	var (
		defaultBrandingThemeName   *string
		defaultBrandingThemeNameOk = false
	)
	if defaultBrandingThemeConfigurationOk {
		defaultBrandingThemeName, defaultBrandingThemeNameOk = defaultBrandingThemeConfiguration.GetNameOk()
	}

	if defaultBrandingThemeConfigurationOk && defaultBrandingThemeNameOk {
		commentData := map[string]string{
			"Resource Type":         r.ResourceType(),
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		}

		importBlocks = append(importBlocks, connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_default_theme", *defaultBrandingThemeName),
			ResourceID:         r.clientInfo.ExportEnvironmentID,
			CommentInformation: common.GenerateCommentInformation(commentData),
		})
	}

	return &importBlocks, nil
}

func (r *PingoneBrandingThemeDefaultResource) ResourceType() string {
	return "pingone_branding_theme_default"
}
