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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateServerSettingsSystemKeysRotateResource
func ServerSettingsSystemKeysRotate(clientInfo *connector.ClientInfo) *PingFederateServerSettingsSystemKeysRotateResource {
	return &PingFederateServerSettingsSystemKeysRotateResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateServerSettingsSystemKeysRotateResource) ResourceType() string {
	return "pingfederate_server_settings_system_keys_rotate"
}

func (r *PingFederateServerSettingsSystemKeysRotateResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	serverSettingsSystemKeysRotateId := "server_settings_system_keys_rotate_singleton_id"
	serverSettingsSystemKeysRotateName := "Server Settings System Keys Rotate"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       serverSettingsSystemKeysRotateName,
		ResourceID:         serverSettingsSystemKeysRotateId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
