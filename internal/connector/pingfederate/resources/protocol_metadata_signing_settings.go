// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateProtocolMetadataSigningSettingsResource{}
)

type PingFederateProtocolMetadataSigningSettingsResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateProtocolMetadataSigningSettingsResource
func ProtocolMetadataSigningSettings(clientInfo *connector.ClientInfo) *PingFederateProtocolMetadataSigningSettingsResource {
	return &PingFederateProtocolMetadataSigningSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateProtocolMetadataSigningSettingsResource) ResourceType() string {
	return "pingfederate_protocol_metadata_signing_settings"
}

func (r *PingFederateProtocolMetadataSigningSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	protocolMetadataSigningSettingsId := "protocol_metadata_signing_settings_singleton_id"
	protocolMetadataSigningSettingsName := "Protocol Metadata Signing Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       protocolMetadataSigningSettingsName,
		ResourceID:         protocolMetadataSigningSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
