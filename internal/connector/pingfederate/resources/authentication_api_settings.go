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

func (r *PingFederateAuthenticationApiSettingsResource) ResourceType() string {
	return "pingfederate_authentication_api_settings"
}

func (r *PingFederateAuthenticationApiSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	authenticationApiSettingsId := "authentication_api_settings_singleton_id"
	authenticationApiSettingsName := "Authentication Api Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       authenticationApiSettingsName,
		ResourceID:         authenticationApiSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
