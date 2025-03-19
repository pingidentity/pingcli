// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateSessionSettingsResource{}
)

type PingFederateSessionSettingsResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateSessionSettingsResource
func SessionSettings(clientInfo *connector.ClientInfo) *PingFederateSessionSettingsResource {
	return &PingFederateSessionSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateSessionSettingsResource) ResourceType() string {
	return "pingfederate_session_settings"
}

func (r *PingFederateSessionSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	sessionSettingsId := "session_settings_singleton_id"
	sessionSettingsName := "Session Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       sessionSettingsName,
		ResourceID:         sessionSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
