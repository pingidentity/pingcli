package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateAuthenticationPoliciesSettingsResource{}
)

type PingFederateAuthenticationPoliciesSettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateAuthenticationPoliciesSettingsResource
func AuthenticationPoliciesSettings(clientInfo *connector.PingFederateClientInfo) *PingFederateAuthenticationPoliciesSettingsResource {
	return &PingFederateAuthenticationPoliciesSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateAuthenticationPoliciesSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	authnPoliciesSettingsId := "authentication_policies_settings_singleton_id"
	authnPoliciesSettingsName := "Authentication Policies Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       authnPoliciesSettingsName,
		ResourceID:         authnPoliciesSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateAuthenticationPoliciesSettingsResource) ResourceType() string {
	return "pingfederate_authentication_policies_settings"
}
