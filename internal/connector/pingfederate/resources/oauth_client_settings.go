package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOauthClientSettingsResource{}
)

type PingFederateOauthClientSettingsResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateOauthClientSettingsResource
func OauthClientSettings(clientInfo *connector.ClientInfo) *PingFederateOauthClientSettingsResource {
	return &PingFederateOauthClientSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOauthClientSettingsResource) ResourceType() string {
	return "pingfederate_oauth_client_settings"
}

func (r *PingFederateOauthClientSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthClientSettingsId := "oauth_client_settings_singleton_id"
	oauthClientSettingsName := "Oauth Client Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       oauthClientSettingsName,
		ResourceID:         oauthClientSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
