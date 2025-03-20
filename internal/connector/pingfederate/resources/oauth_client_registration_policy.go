// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOauthClientRegistrationPolicyResource{}
)

type PingFederateOauthClientRegistrationPolicyResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateOauthClientRegistrationPolicyResource
func OauthClientRegistrationPolicy(clientInfo *connector.ClientInfo) *PingFederateOauthClientRegistrationPolicyResource {
	return &PingFederateOauthClientRegistrationPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOauthClientRegistrationPolicyResource) ResourceType() string {
	return "pingfederate_oauth_client_registration_policy"
}

func (r *PingFederateOauthClientRegistrationPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthClientRegistrationPolicyData, err := r.getOauthClientRegistrationPolicyData()
	if err != nil {
		return nil, err
	}

	for oauthClientRegistrationPolicyId, oauthClientRegistrationPolicyName := range oauthClientRegistrationPolicyData {
		commentData := map[string]string{
			"Oauth Client Registration Policy ID":   oauthClientRegistrationPolicyId,
			"Oauth Client Registration Policy Name": oauthClientRegistrationPolicyName,
			"Resource Type":                         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       oauthClientRegistrationPolicyName,
			ResourceID:         oauthClientRegistrationPolicyId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOauthClientRegistrationPolicyResource) getOauthClientRegistrationPolicyData() (map[string]string, error) {
	oauthClientRegistrationPolicyData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.OauthClientRegistrationPoliciesAPI.GetDynamicClientRegistrationPolicies(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetDynamicClientRegistrationPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if apiObj == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	items, itemsOk := apiObj.GetItemsOk()
	if !itemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, oauthClientRegistrationPolicy := range items {
		oauthClientRegistrationPolicyId, oauthClientRegistrationPolicyIdOk := oauthClientRegistrationPolicy.GetIdOk()
		oauthClientRegistrationPolicyName, oauthClientRegistrationPolicyNameOk := oauthClientRegistrationPolicy.GetNameOk()

		if oauthClientRegistrationPolicyIdOk && oauthClientRegistrationPolicyNameOk {
			oauthClientRegistrationPolicyData[*oauthClientRegistrationPolicyId] = *oauthClientRegistrationPolicyName
		}
	}

	return oauthClientRegistrationPolicyData, nil
}
