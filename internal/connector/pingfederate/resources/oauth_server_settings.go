package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOauthServerSettingsResource{}
)

type PingFederateOauthServerSettingsResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateOauthServerSettingsResource
func OauthServerSettings(clientInfo *connector.ClientInfo) *PingFederateOauthServerSettingsResource {
	return &PingFederateOauthServerSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOauthServerSettingsResource) ResourceType() string {
	return "pingfederate_oauth_server_settings"
}

func (r *PingFederateOauthServerSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthServerSettingsId := "oauth_server_settings_singleton_id"
	oauthServerSettingsName := "Oauth Server Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       oauthServerSettingsName,
		ResourceID:         oauthServerSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
