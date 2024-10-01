package resources

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateIncomingProxySettingsResource{}
)

type PingFederateIncomingProxySettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateIncomingProxySettingsResource
func IncomingProxySettings(clientInfo *connector.PingFederateClientInfo) *PingFederateIncomingProxySettingsResource {
	return &PingFederateIncomingProxySettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateIncomingProxySettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	incomingProxySettingsId := "incoming_proxy_settings_singleton_id"
	incomingProxySettingsName := "Incoming Proxy Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       incomingProxySettingsName,
		ResourceID:         incomingProxySettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateIncomingProxySettingsResource) ResourceType() string {
	return "pingfederate_incoming_proxy_settings"
}
