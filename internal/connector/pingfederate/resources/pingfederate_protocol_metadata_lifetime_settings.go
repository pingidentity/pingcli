package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateProtocolMetadataLifetimeSettingsResource{}
)

type PingFederateProtocolMetadataLifetimeSettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateProtocolMetadataLifetimeSettingsResource
func ProtocolMetadataLifetimeSettings(clientInfo *connector.PingFederateClientInfo) *PingFederateProtocolMetadataLifetimeSettingsResource {
	return &PingFederateProtocolMetadataLifetimeSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateProtocolMetadataLifetimeSettingsResource) ResourceType() string {
	return "pingfederate_protocol_metadata_lifetime_settings"
}

func (r *PingFederateProtocolMetadataLifetimeSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	protocolMetadataLifetimeSettingsId := "protocol_metadata_lifetime_settings_singleton_id"
	protocolMetadataLifetimeSettingsName := "Protocol Metadata Lifetime Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       protocolMetadataLifetimeSettingsName,
		ResourceID:         protocolMetadataLifetimeSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}