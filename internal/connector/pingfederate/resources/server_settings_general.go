// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateServerSettingsGeneralResource{}
)

type PingFederateServerSettingsGeneralResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateServerSettingsGeneralResource
func ServerSettingsGeneral(clientInfo *connector.ClientInfo) *PingFederateServerSettingsGeneralResource {
	return &PingFederateServerSettingsGeneralResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateServerSettingsGeneralResource) ResourceType() string {
	return "pingfederate_server_settings_general"
}

func (r *PingFederateServerSettingsGeneralResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	serverSettingsGeneralId := "server_settings_general_singleton_id"
	serverSettingsGeneralName := "Server Settings General"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       serverSettingsGeneralName,
		ResourceID:         serverSettingsGeneralId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
