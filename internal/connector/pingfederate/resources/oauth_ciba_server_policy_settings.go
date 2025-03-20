// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOauthCibaServerPolicySettingsResource{}
)

type PingFederateOauthCibaServerPolicySettingsResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateOauthCibaServerPolicySettingsResource
func OauthCibaServerPolicySettings(clientInfo *connector.ClientInfo) *PingFederateOauthCibaServerPolicySettingsResource {
	return &PingFederateOauthCibaServerPolicySettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOauthCibaServerPolicySettingsResource) ResourceType() string {
	return "pingfederate_oauth_ciba_server_policy_settings"
}

func (r *PingFederateOauthCibaServerPolicySettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthCibaServerPolicySettingsId := "oauth_ciba_server_policy_settings_singleton_id"
	oauthCibaServerPolicySettingsName := "Oauth Ciba Server Policy Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       oauthCibaServerPolicySettingsName,
		ResourceID:         oauthCibaServerPolicySettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
