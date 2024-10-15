package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateServerSettingsSystemKeysRotateResource{}
)

type PingFederateServerSettingsSystemKeysRotateResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateServerSettingsSystemKeysResource
func ServerSettingsSystemKeysRotate(clientInfo *connector.PingFederateClientInfo) *PingFederateServerSettingsSystemKeysRotateResource {
	return &PingFederateServerSettingsSystemKeysRotateResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateServerSettingsSystemKeysRotateResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	serverSettingsSystemKeysRotateId := "server_settings_system_keys_rotate_singleton_id"
	serverSettingsSystemKeysRotateName := "Server Settings System Keys Rotate"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       serverSettingsSystemKeysRotateName,
		ResourceID:         serverSettingsSystemKeysRotateId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateServerSettingsSystemKeysRotateResource) ResourceType() string {
	return "pingfederate_server_settings_system_keys_rotate"
}
