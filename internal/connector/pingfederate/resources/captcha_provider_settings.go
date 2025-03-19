// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateCaptchaProviderSettingsResource{}
)

type PingFederateCaptchaProviderSettingsResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateCaptchaProviderSettingsResource
func CaptchaProviderSettings(clientInfo *connector.ClientInfo) *PingFederateCaptchaProviderSettingsResource {
	return &PingFederateCaptchaProviderSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateCaptchaProviderSettingsResource) ResourceType() string {
	return "pingfederate_captcha_provider_settings"
}

func (r *PingFederateCaptchaProviderSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	captchaProviderSettingsId := "captcha_provider_settings_singleton_id"
	captchaProviderSettingsName := "Captcha Provider Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       captchaProviderSettingsName,
		ResourceID:         captchaProviderSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
