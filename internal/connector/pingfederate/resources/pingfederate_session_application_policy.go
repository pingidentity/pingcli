package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateSessionApplicationPolicyResource{}
)

type PingFederateSessionApplicationPolicyResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateSessionApplicationPolicyResource
func SessionApplicationPolicy(clientInfo *connector.PingFederateClientInfo) *PingFederateSessionApplicationPolicyResource {
	return &PingFederateSessionApplicationPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateSessionApplicationPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	sessionApplicationPolicyId := "pingfederate_session_application_policy_singleton_id"
	sessionApplicationPolicyName := "Session Application Policy"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       sessionApplicationPolicyName,
		ResourceID:         sessionApplicationPolicyId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateSessionApplicationPolicyResource) ResourceType() string {
	return "pingfederate_session_application_policy"
}
