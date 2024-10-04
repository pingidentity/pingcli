package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateServerSettingsSystemKeysResource{}
)

type PingFederateServerSettingsSystemKeysResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateServerSettingsSystemKeysResource
func ServerSettingsSystemKeys(clientInfo *connector.PingFederateClientInfo) *PingFederateServerSettingsSystemKeysResource {
	return &PingFederateServerSettingsSystemKeysResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateServerSettingsSystemKeysResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	serverSettingsSystemKeysId := "server_settings_system_keys_singleton_id"
	serverSettingsSystemKeysName := "Server Settings System Keys"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       serverSettingsSystemKeysName,
		ResourceID:         serverSettingsSystemKeysId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateServerSettingsSystemKeysResource) ResourceType() string {
	return "pingfederate_server_settings_system_keys"
}
