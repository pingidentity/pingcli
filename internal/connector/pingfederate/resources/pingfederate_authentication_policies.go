package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateAuthenticationPoliciesResource{}
)

type PingFederateAuthenticationPoliciesResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateAuthenticationPoliciesResource
func AuthenticationPolicies(clientInfo *connector.PingFederateClientInfo) *PingFederateAuthenticationPoliciesResource {
	return &PingFederateAuthenticationPoliciesResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateAuthenticationPoliciesResource) ResourceType() string {
	return "pingfederate_authentication_policies"
}

func (r *PingFederateAuthenticationPoliciesResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	authnPoliciesId := "authentication_policies_singleton_id"
	authnPoliciesName := "Authentication Policies"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       authnPoliciesName,
		ResourceID:         authnPoliciesId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
