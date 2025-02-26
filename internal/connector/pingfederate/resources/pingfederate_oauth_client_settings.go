package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOAuthClientSettingsResource{}
)

type PingFederateOAuthClientSettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateOAuthClientSettingsResource
func OAuthClientSettings(clientInfo *connector.PingFederateClientInfo) *PingFederateOAuthClientSettingsResource {
	return &PingFederateOAuthClientSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOAuthClientSettingsResource) ResourceType() string {
	return "pingfederate_oauth_client_settings"
}

func (r *PingFederateOAuthClientSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oAuthClientSettingsId := "oauth_client_settings_singleton_id"
	oAuthClientSettingsName := "OAuth Client Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       oAuthClientSettingsName,
		ResourceID:         oAuthClientSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
