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
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateSessionAuthenticationPoliciesGlobalResource
func SessionAuthenticationPoliciesGlobal(clientInfo *connector.PingFederateClientInfo) *PingFederateSessionAuthenticationPoliciesGlobalResource {
	return &PingFederateSessionAuthenticationPoliciesGlobalResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateSessionAuthenticationPoliciesGlobalResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	sessionAuthenticationPoliciesGlobalId := "pingfederate_session_authentication_policies_global_singleton_id"
	sessionAuthenticationPoliciesGlobalName := "Session Authentication Policies Global"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       sessionAuthenticationPoliciesGlobalName,
		ResourceID:         sessionAuthenticationPoliciesGlobalId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateSessionAuthenticationPoliciesGlobalResource) ResourceType() string {
	return "pingfederate_session_authentication_policies_global"
}
