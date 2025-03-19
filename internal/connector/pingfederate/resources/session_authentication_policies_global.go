// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateSessionAuthenticationPoliciesGlobalResource{}
)

type PingFederateSessionAuthenticationPoliciesGlobalResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateSessionAuthenticationPoliciesGlobalResource
func SessionAuthenticationPoliciesGlobal(clientInfo *connector.ClientInfo) *PingFederateSessionAuthenticationPoliciesGlobalResource {
	return &PingFederateSessionAuthenticationPoliciesGlobalResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateSessionAuthenticationPoliciesGlobalResource) ResourceType() string {
	return "pingfederate_session_authentication_policies_global"
}

func (r *PingFederateSessionAuthenticationPoliciesGlobalResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	sessionAuthenticationPoliciesGlobalId := "session_authentication_policies_global_singleton_id"
	sessionAuthenticationPoliciesGlobalName := "Session Authentication Policies Global"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       sessionAuthenticationPoliciesGlobalName,
		ResourceID:         sessionAuthenticationPoliciesGlobalId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
