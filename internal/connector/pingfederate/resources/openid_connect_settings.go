package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOpenidConnectSettingsResource{}
)

type PingFederateOpenidConnectSettingsResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateOpenidConnectSettingsResource
func OpenidConnectSettings(clientInfo *connector.ClientInfo) *PingFederateOpenidConnectSettingsResource {
	return &PingFederateOpenidConnectSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOpenidConnectSettingsResource) ResourceType() string {
	return "pingfederate_openid_connect_settings"
}

func (r *PingFederateOpenidConnectSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	openidConnectSettingsId := "openid_connect_settings_singleton_id"
	openidConnectSettingsName := "Openid Connect Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       openidConnectSettingsName,
		ResourceID:         openidConnectSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
