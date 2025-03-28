// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneBrandingSettingsResource
func BrandingSettings(clientInfo *connector.ClientInfo) *PingOneBrandingSettingsResource {
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

	ok, err := r.checkBrandingSettingsData()
	if err != nil {
		return nil, err
	}
	if !ok {
		return &importBlocks, nil
	}

	commentData := map[string]string{
		"Resource Type":         r.ResourceType(),
		"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       r.ResourceType(),
		ResourceID:         r.clientInfo.PingOneExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}

func (r *PingOneBrandingSettingsResource) checkBrandingSettingsData() (bool, error) {
	_, response, err := r.clientInfo.PingOneApiClient.ManagementAPIClient.BrandingSettingsApi.ReadBrandingSettings(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()

	return common.CheckSingletonResource(response, err, "ReadBrandingSettings", r.ResourceType())
}
