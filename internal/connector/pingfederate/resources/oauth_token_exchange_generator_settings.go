package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOauthTokenExchangeGeneratorSettingsResource{}
)

type PingFederateOauthTokenExchangeGeneratorSettingsResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateOauthTokenExchangeGeneratorSettingsResource
func OauthTokenExchangeGeneratorSettings(clientInfo *connector.ClientInfo) *PingFederateOauthTokenExchangeGeneratorSettingsResource {
	return &PingFederateOauthTokenExchangeGeneratorSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOauthTokenExchangeGeneratorSettingsResource) ResourceType() string {
	return "pingfederate_oauth_token_exchange_generator_settings"
}

func (r *PingFederateOauthTokenExchangeGeneratorSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthTokenExchangeGeneratorSettingsId := "oauth_token_exchange_generator_settings_singleton_id" //#nosec G101 -- This is not hard-coded credentials
	oauthTokenExchangeGeneratorSettingsName := "Oauth Token Exchange Generator Settings"            //#nosec G101 -- This is not hard-coded credentials

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       oauthTokenExchangeGeneratorSettingsName,
		ResourceID:         oauthTokenExchangeGeneratorSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
