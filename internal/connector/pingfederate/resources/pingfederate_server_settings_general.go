package resources

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateServerSettingsGeneralResource{}
)

type PingFederateServerSettingsGeneralResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateServerSettingsGeneralResource
func ServerSettingsGeneral(clientInfo *connector.PingFederateClientInfo) *PingFederateServerSettingsGeneralResource {
	return &PingFederateServerSettingsGeneralResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateServerSettingsGeneralResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	serverSettingsGeneralId := "pingfederate_server_settings_general_singleton_id"
	serverSettingsGeneralName := "Server Settings General"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       serverSettingsGeneralName,
		ResourceID:         serverSettingsGeneralId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateServerSettingsGeneralResource) ResourceType() string {
	return "pingfederate_server_settings_general"
}
