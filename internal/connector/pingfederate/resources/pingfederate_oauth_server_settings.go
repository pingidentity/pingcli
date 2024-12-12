package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOAuthServerSettingsResource{}
)

type PingFederateOAuthServerSettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateOAuthServerSettingsResource
func OAuthServerSettings(clientInfo *connector.PingFederateClientInfo) *PingFederateOAuthServerSettingsResource {
	return &PingFederateOAuthServerSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOAuthServerSettingsResource) ResourceType() string {
	return "pingfederate_oauth_server_settings"
}

func (r *PingFederateOAuthServerSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oAuthServerSettingsId := "oauth_server_settings_singleton_id"
	oAuthServerSettingsName := "OAuth Server Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       oAuthServerSettingsName,
		ResourceID:         oAuthServerSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
