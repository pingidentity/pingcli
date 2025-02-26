package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateServerSettingsWsTrustStsSettingsResource{}
)

type PingFederateServerSettingsWsTrustStsSettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateServerSettingsWsTrustStsSettingsResource
func ServerSettingsWsTrustStsSettings(clientInfo *connector.PingFederateClientInfo) *PingFederateServerSettingsWsTrustStsSettingsResource {
	return &PingFederateServerSettingsWsTrustStsSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateServerSettingsWsTrustStsSettingsResource) ResourceType() string {
	return "pingfederate_server_settings_ws_trust_sts_settings"
}

func (r *PingFederateServerSettingsWsTrustStsSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	serverSettingsWsTrustStsSettingsId := "server_settings_ws_trust_sts_settings_singleton_id"
	serverSettingsWsTrustStsSettingsName := "Server Settings WS-Trust STS Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       serverSettingsWsTrustStsSettingsName,
		ResourceID:         serverSettingsWsTrustStsSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
