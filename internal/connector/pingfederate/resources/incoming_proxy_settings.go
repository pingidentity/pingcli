// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateIncomingProxySettingsResource{}
)

type PingFederateIncomingProxySettingsResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateIncomingProxySettingsResource
func IncomingProxySettings(clientInfo *connector.ClientInfo) *PingFederateIncomingProxySettingsResource {
	return &PingFederateIncomingProxySettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateIncomingProxySettingsResource) ResourceType() string {
	return "pingfederate_incoming_proxy_settings"
}

func (r *PingFederateIncomingProxySettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	incomingProxySettingsId := "incoming_proxy_settings_singleton_id"
	incomingProxySettingsName := "Incoming Proxy Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       incomingProxySettingsName,
		ResourceID:         incomingProxySettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
