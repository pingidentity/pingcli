package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
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
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	editorPolicyData, err := r.getEditorPolicyData()
	if err != nil {
		return nil, err
	}

	for editorPolicyId, editorPolicyName := range editorPolicyData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Editor Policy ID":      editorPolicyId,
			"Editor Policy Name":    editorPolicyName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       editorPolicyName,
			ResourceID:         fmt.Sprintf("%s", r.clientInfo.ExportEnvironmentID),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizePolicyManagementPolicyResource) getEditorPolicyData() (map[string]string, error) {
	editorPolicyData := make(map[string]string)

	iter := r.clientInfo.ApiClient.AuthorizeAPIClient.AuthorizeEditorPoliciesApi.ListRootPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	editorPolicys, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.AuthorizeEditorDataPoliciesReferenceablePolicyDTO](iter, "ListRootPolicies", "GetAuthorizationPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, editorPolicy := range editorPolicys {

		if me, ok := editorPolicy.GetManagedEntityOk(); ok {
			if restrictions, ok := me.GetRestrictionsOk(); ok {
				if readOnly, ok := restrictions.GetReadOnlyOk(); ok {
					if *readOnly {
						continue
					}
				}
			}
		}

		editorPolicyId, editorPolicyIdOk := editorPolicy.GetIdOk()
		editorPolicyName, editorPolicyNameOk := editorPolicy.GetNameOk()

		if editorPolicyIdOk && editorPolicyNameOk {
			editorPolicyData[*editorPolicyId] = *editorPolicyName
		}
	}

	return editorPolicyData, nil
}

func (r *PingoneAuthorizePolicyManagementPolicyResource) ResourceType() string {
	return "pingone_authorize_policy_management_root_policy"
}
