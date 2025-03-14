package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOauthAccessTokenManagerSettingsResource{}
)

type PingFederateOauthAccessTokenManagerSettingsResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateOauthAccessTokenManagerSettingsResource
func OauthAccessTokenManagerSettings(clientInfo *connector.ClientInfo) *PingFederateOauthAccessTokenManagerSettingsResource {
	return &PingFederateOauthAccessTokenManagerSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOauthAccessTokenManagerSettingsResource) ResourceType() string {
	return "pingfederate_oauth_access_token_manager_settings"
}

func (r *PingFederateOauthAccessTokenManagerSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthAccessTokenManagerSettingsId := "oauth_access_token_manager_settings_singleton_id" //#nosec G101 -- This is not hard-coded credentials
	oauthAccessTokenManagerSettingsName := "Oauth Access Token Manager Settings"            //#nosec G101 -- This is not hard-coded credentials

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       oauthAccessTokenManagerSettingsName,
		ResourceID:         oauthAccessTokenManagerSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
