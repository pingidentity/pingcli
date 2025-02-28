package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOAuthClientRegistrationPolicyResource{}
)

type PingFederateOAuthClientRegistrationPolicyResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateOAuthClientRegistrationPolicyResource
func OAuthClientRegistrationPolicy(clientInfo *connector.PingFederateClientInfo) *PingFederateOAuthClientRegistrationPolicyResource {
	return &PingFederateOAuthClientRegistrationPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOAuthClientRegistrationPolicyResource) ResourceType() string {
	return "pingfederate_oauth_client_registration_policy"
}

func (r *PingFederateOAuthClientRegistrationPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	clientRegistrationPolicyData, err := r.getClientRegistrationPolicyData()
	if err != nil {
		return nil, err
	}

	for clientRegistrationPolicyId, clientRegistrationPolicyName := range clientRegistrationPolicyData {
		commentData := map[string]string{
			"OAuth Client Registration Policy ID":   clientRegistrationPolicyId,
			"OAuth Client Registration Policy Name": clientRegistrationPolicyName,
			"Resource Type":                         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       clientRegistrationPolicyName,
			ResourceID:         clientRegistrationPolicyId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOAuthClientRegistrationPolicyResource) getClientRegistrationPolicyData() (map[string]string, error) {
	clientRegistrationPolicyData := make(map[string]string)

	clientRegistrationPolicies, response, err := r.clientInfo.ApiClient.OauthClientRegistrationPoliciesAPI.GetDynamicClientRegistrationPolicies(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetDynamicClientRegistrationPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if clientRegistrationPolicies == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	clientRegistrationPoliciesItems, clientRegistrationPoliciesItemsOk := clientRegistrationPolicies.GetItemsOk()
	if !clientRegistrationPoliciesItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, clientRegistrationPolicy := range clientRegistrationPoliciesItems {
		clientRegistrationPolicyId, clientRegistrationPolicyIdOk := clientRegistrationPolicy.GetIdOk()
		clientRegistrationPolicyName, clientRegistrationPolicyNameOk := clientRegistrationPolicy.GetNameOk()

		if clientRegistrationPolicyIdOk && clientRegistrationPolicyNameOk {
			clientRegistrationPolicyData[*clientRegistrationPolicyId] = *clientRegistrationPolicyName
		}
	}

	return clientRegistrationPolicyData, nil
}
