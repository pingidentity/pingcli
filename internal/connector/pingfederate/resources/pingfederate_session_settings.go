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
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateSessionSettingsResource
func SessionSettings(clientInfo *connector.PingFederateClientInfo) *PingFederateSessionSettingsResource {
	return &PingFederateSessionSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateSessionSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	sessionSettingsId := "session_settings_singleton_id"
	sessionSettingsName := "Session Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       sessionSettingsName,
		ResourceID:         sessionSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateSessionSettingsResource) ResourceType() string {
	return "pingfederate_session_settings"
}
