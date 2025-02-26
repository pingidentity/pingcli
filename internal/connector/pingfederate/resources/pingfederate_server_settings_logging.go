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

type PingFederateServerSettingsLoggingResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateServerSettingsLoggingResource
func ServerSettingsLogging(clientInfo *connector.PingFederateClientInfo) *PingFederateServerSettingsLoggingResource {
	return &PingFederateServerSettingsLoggingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateServerSettingsLoggingResource) ResourceType() string {
	return "pingfederate_server_settings_logging"
}

func (r *PingFederateServerSettingsLoggingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	serverSettingsLoggingId := "server_settings_logging_singleton_id"
	serverSettingsLoggingName := "Server Settings Logging"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       serverSettingsLoggingName,
		ResourceID:         serverSettingsLoggingId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
