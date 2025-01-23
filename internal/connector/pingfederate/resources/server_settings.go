package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateServerSettingsResource{}
)

type PingFederateServerSettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateServerSettingsResource
func ServerSettings(clientInfo *connector.PingFederateClientInfo) *PingFederateServerSettingsResource {
	return &PingFederateServerSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateServerSettingsResource) ResourceType() string {
	return "pingfederate_server_settings"
}

func (r *PingFederateServerSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	serverSettingsId := "server_settings__singleton_id"
	serverSettingsName := "Server Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       serverSettingsName,
		ResourceID:         serverSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
