package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAuthorizePolicyManagementPolicyResource{}
)

type PingoneAuthorizePolicyManagementPolicyResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizePolicyManagementPolicyResource
func AuthorizePolicyManagementPolicy(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizePolicyManagementPolicyResource {
	return &PingoneAuthorizePolicyManagementPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizePolicyManagementPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.AuthorizeAPIClient.AuthorizeEditorPoliciesApi.ListRootPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ListRootPolicies"

	embedded, err := common.GetAuthorizeEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, authorizationPolicy := range embedded.GetAuthorizationPolicies() {
		authorizationPolicyName, authorizationPolicyNameOk := authorizationPolicy.GetNameOk()
		authorizationPolicyId, authorizationPolicyIdOk := authorizationPolicy.GetIdOk()

		if authorizationPolicyNameOk && authorizationPolicyIdOk {
			commentData := map[string]string{
				"Resource Type": r.ResourceType(),
				"Authorize Policy Management Authorization Policy Name": *authorizationPolicyName,
				"Export Environment ID":                                 r.clientInfo.ExportEnvironmentID,
				"Authorize Policy Management Authorization Policy ID":   *authorizationPolicyId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *authorizationPolicyName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *authorizationPolicyId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizePolicyManagementPolicyResource) ResourceType() string {
	return "pingone_authorize_policy_management_policy"
}
