// Copyright © 2025 Ping Identity Corporation

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
	_ connector.ExportableResource = &PingOneBrandingThemeDefaultResource{}
)

type PingOneBrandingThemeDefaultResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneBrandingThemeDefaultResource
func BrandingThemeDefault(clientInfo *connector.ClientInfo) *PingOneBrandingThemeDefaultResource {
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
	if defaultBrandingThemeName == nil {
		return &importBlocks, nil
	}

	commentData := map[string]string{
		"Default Branding Theme Name": *defaultBrandingThemeName,
		"Export Environment ID":       r.clientInfo.PingOneExportEnvironmentID,
		"Resource Type":               r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_default_theme", *defaultBrandingThemeName),
		ResourceID:         r.clientInfo.PingOneExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}

func (r *PingOneBrandingThemeDefaultResource) getDefaultBrandingThemeName() (*string, error) {
	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.BrandingThemesApi.ReadBrandingThemes(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	brandingThemes, err := pingone.GetManagementAPIObjectsFromIterator[management.BrandingTheme](iter, "ReadBrandingThemes", "GetThemes", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, brandingTheme := range brandingThemes {
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

	return nil, fmt.Errorf("failed to export resource '%s'. No default branding theme found", r.ResourceType())
}
