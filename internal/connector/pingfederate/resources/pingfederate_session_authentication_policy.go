package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateSessionAuthenticationPolicyResource{}
)

type PingFederateSessionAuthenticationPolicyResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateSessionAuthenticationPolicyResource
func SessionAuthenticationPolicy(clientInfo *connector.PingFederateClientInfo) *PingFederateSessionAuthenticationPolicyResource {
	return &PingFederateSessionAuthenticationPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateSessionAuthenticationPolicyResource) ResourceType() string {
	return "pingfederate_session_authentication_policy"
}

func (r *PingFederateSessionAuthenticationPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	authenticationSessionPolicyData, err := r.getAuthenticationSessionPolicyData()
	if err != nil {
		return nil, err
	}

	for policyId, policyInfo := range *authenticationSessionPolicyData {
		authSourceType := policyInfo[0]
		authSourceRefId := policyInfo[1]

		commentData := map[string]string{
			"Resource Type":                      r.ResourceType(),
			"Session Authentication Policy ID":   policyId,
			"Session Authentication Source Type": authSourceType,
			"Session Authentication Source ID":   authSourceRefId,
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s_%s", policyId, authSourceType, authSourceRefId),
			ResourceID:         policyId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateSessionAuthenticationPolicyResource) getAuthenticationSessionPolicyData() (*map[string][]string, error) {
	authenticationSessionPolicyData := make(map[string][]string)

	authenticationSessionPolicies, response, err := r.clientInfo.ApiClient.SessionAPI.GetSourcePolicies(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetSourcePolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if authenticationSessionPolicies == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	authenticationSessionPoliciesItems, authenticationSessionPoliciesItemsOk := authenticationSessionPolicies.GetItemsOk()
	if !authenticationSessionPoliciesItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, authenticationSessionPolicy := range authenticationSessionPoliciesItems {
		policyId, policyIdOk := authenticationSessionPolicy.GetIdOk()
		authSource, authSourceOk := authenticationSessionPolicy.GetAuthenticationSourceOk()

		if policyIdOk && authSourceOk {
			authSourceType, authSourceTypeOk := authSource.GetTypeOk()
			authSourceRef, authSourceRefOk := authSource.GetSourceRefOk()

			if authSourceTypeOk && authSourceRefOk {
				authSourceRefId, authSourceRefIdOk := authSourceRef.GetIdOk()

				if authSourceRefIdOk {
					authenticationSessionPolicyData[*policyId] = []string{*authSourceType, *authSourceRefId}
				}
			}
		}
	}

	return &authenticationSessionPolicyData, nil
}
