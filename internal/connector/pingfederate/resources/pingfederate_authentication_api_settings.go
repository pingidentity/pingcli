package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateAuthenticationApiSettingsResource{}
)

type PingFederateAuthenticationApiSettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateAuthenticationApiSettingsResource
func AuthenticationApiSettings(clientInfo *connector.PingFederateClientInfo) *PingFederateAuthenticationApiSettingsResource {
	return &PingFederateAuthenticationApiSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateAuthenticationApiSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	authnApiSettingsId := "authentication_api_settings_singleton_id"
	authnApiSettingsName := "Authentication API Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       authnApiSettingsName,
		ResourceID:         authnApiSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateAuthenticationApiSettingsResource) ResourceType() string {
	return "pingfederate_authentication_api_settings"
}
