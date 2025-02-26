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

func (r *PingFederateOpenIDConnectSettingsResource) ResourceType() string {
	return "pingfederate_openid_connect_settings"
}

func (r *PingFederateOpenIDConnectSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	openIDConnectSettingsId := "openid_connect_settings_singleton_id"
	openIDConnectSettingsName := "OpenID Connect Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       openIDConnectSettingsName,
		ResourceID:         openIDConnectSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
