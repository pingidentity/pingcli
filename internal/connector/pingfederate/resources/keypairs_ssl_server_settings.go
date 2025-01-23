package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateKeypairsSslServerSettingsResource{}
)

type PingFederateKeypairsSslServerSettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateKeypairsSslServerSettingsResource
func KeypairsSslServerSettings(clientInfo *connector.PingFederateClientInfo) *PingFederateKeypairsSslServerSettingsResource {
	return &PingFederateKeypairsSslServerSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateKeypairsSslServerSettingsResource) ResourceType() string {
	return "pingfederate_keypairs_ssl_server_settings"
}

func (r *PingFederateKeypairsSslServerSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	keypairsSslServerSettingsId := "keypairs_ssl_server_settings_singleton_id"
	keypairsSslServerSettingsName := "Keypairs Ssl Server Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       keypairsSslServerSettingsName,
		ResourceID:         keypairsSslServerSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
