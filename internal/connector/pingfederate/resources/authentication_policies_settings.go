// Copyright © 2025 Ping Identity Corporation

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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateAuthenticationPoliciesSettingsResource
func AuthenticationPoliciesSettings(clientInfo *connector.ClientInfo) *PingFederateAuthenticationPoliciesSettingsResource {
	return &PingFederateAuthenticationPoliciesSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateAuthenticationPoliciesSettingsResource) ResourceType() string {
	return "pingfederate_authentication_policies_settings"
}

func (r *PingFederateAuthenticationPoliciesSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	authenticationPoliciesSettingsId := "authentication_policies_settings_singleton_id"
	authenticationPoliciesSettingsName := "Authentication Policies Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       authenticationPoliciesSettingsName,
		ResourceID:         authenticationPoliciesSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
