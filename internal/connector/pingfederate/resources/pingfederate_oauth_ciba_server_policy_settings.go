package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOAuthCIBAServerPolicySettingsResource{}
)

type PingFederateOAuthCIBAServerPolicySettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateOAuthCIBAServerPolicySettingsResource
func OAuthCIBAServerPolicySettings(clientInfo *connector.PingFederateClientInfo) *PingFederateOAuthCIBAServerPolicySettingsResource {
	return &PingFederateOAuthCIBAServerPolicySettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOAuthCIBAServerPolicySettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	oAuthCIBAServerPolicySettingsId := "oauth_ciba_server_policy_settings_singleton_id"
	oAuthCIBAServerPolicySettingsName := "OAuth CIBA Server Policy Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       oAuthCIBAServerPolicySettingsName,
		ResourceID:         oAuthCIBAServerPolicySettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateOAuthCIBAServerPolicySettingsResource) ResourceType() string {
	return "pingfederate_oauth_ciba_server_policy_settings"
}
