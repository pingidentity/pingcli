package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOAuthClientResource{}
)

type PingFederateOAuthCibaServerPolicyRequestPolicyResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateOAuthCibaServerPolicyRequestPolicyResource
func OAuthCibaServerPolicyRequestPolicy(clientInfo *connector.PingFederateClientInfo) *PingFederateOAuthCibaServerPolicyRequestPolicyResource {
	return &PingFederateOAuthCibaServerPolicyRequestPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOAuthCibaServerPolicyRequestPolicyResource) ResourceType() string {
	return "pingfederate_oauth_ciba_server_policy_request_policy"
}

func (r *PingFederateOAuthCibaServerPolicyRequestPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthClientData, err := r.getRequestPolicyData()
	if err != nil {
		return nil, err
	}

	for requestPolicyId, requestPolicyName := range oauthClientData {
		commentData := map[string]string{
			"OAuth CIBA Server Policy Request Policy ID":   requestPolicyId,
			"OAuth CIBA Server Policy Request Policy Name": requestPolicyName,
			"Resource Type": r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       requestPolicyName,
			ResourceID:         requestPolicyId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOAuthCibaServerPolicyRequestPolicyResource) getRequestPolicyData() (map[string]string, error) {
	requestPolicyData := make(map[string]string)

	requestPolicies, response, err := r.clientInfo.ApiClient.OauthCibaServerPolicyAPI.GetCibaServerPolicies(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetCibaServerPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if requestPolicies == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	requestPoliciesItems, requestPoliciesItemsOk := requestPolicies.GetItemsOk()
	if !requestPoliciesItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, requestPolicy := range requestPoliciesItems {
		requestPolicyId, requestPolicyIdOk := requestPolicy.GetIdOk()
		requestPolicyName, requestPolicyNameOk := requestPolicy.GetNameOk()

		if requestPolicyIdOk && requestPolicyNameOk {
			requestPolicyData[*requestPolicyId] = *requestPolicyName
		}
	}

	return requestPolicyData, nil
}
