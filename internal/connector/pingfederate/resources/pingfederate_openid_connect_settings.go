package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOpenIDConnectSettingsResource{}
)

type PingFederateOpenIDConnectSettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateOpenIDConnectSettingsResource
func OpenIDConnectSettings(clientInfo *connector.PingFederateClientInfo) *PingFederateOpenIDConnectSettingsResource {
	return &PingFederateOpenIDConnectSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOpenIDConnectSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	openIDConnectSettingsId := "openid_connect_settings_singleton_id"
	openIDConnectSettingsName := "OpenID Connect Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       openIDConnectSettingsName,
		ResourceID:         openIDConnectSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateOpenIDConnectSettingsResource) ResourceType() string {
	return "pingfederate_openid_connect_settings"
}
