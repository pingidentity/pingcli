package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOauthCibaServerPolicyRequestPolicyResource{}
)

type PingFederateOauthCibaServerPolicyRequestPolicyResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateOauthCibaServerPolicyRequestPolicyResource
func OauthCibaServerPolicyRequestPolicy(clientInfo *connector.ClientInfo) *PingFederateOauthCibaServerPolicyRequestPolicyResource {
	return &PingFederateOauthCibaServerPolicyRequestPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOauthCibaServerPolicyRequestPolicyResource) ResourceType() string {
	return "pingfederate_oauth_ciba_server_policy_request_policy"
}

func (r *PingFederateOauthCibaServerPolicyRequestPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthCibaServerPolicyRequestPolicyData, err := r.getOauthCibaServerPolicyRequestPolicyData()
	if err != nil {
		return nil, err
	}

	for oauthCibaServerPolicyRequestPolicyId, oauthCibaServerPolicyRequestPolicyName := range *oauthCibaServerPolicyRequestPolicyData {
		commentData := map[string]string{
			"Oauth Ciba Server Policy Request Policy ID":   oauthCibaServerPolicyRequestPolicyId,
			"Oauth Ciba Server Policy Request Policy Name": oauthCibaServerPolicyRequestPolicyName,
			"Resource Type": r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       oauthCibaServerPolicyRequestPolicyName,
			ResourceID:         oauthCibaServerPolicyRequestPolicyId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOauthCibaServerPolicyRequestPolicyResource) getOauthCibaServerPolicyRequestPolicyData() (*map[string]string, error) {
	oauthCibaServerPolicyRequestPolicyData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.OauthCibaServerPolicyAPI.GetCibaServerPolicies(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetCibaServerPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if apiObj == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	items, itemsOk := apiObj.GetItemsOk()
	if !itemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, oauthCibaServerPolicyRequestPolicy := range items {
		oauthCibaServerPolicyRequestPolicyId, oauthCibaServerPolicyRequestPolicyIdOk := oauthCibaServerPolicyRequestPolicy.GetIdOk()
		oauthCibaServerPolicyRequestPolicyName, oauthCibaServerPolicyRequestPolicyNameOk := oauthCibaServerPolicyRequestPolicy.GetNameOk()

		if oauthCibaServerPolicyRequestPolicyIdOk && oauthCibaServerPolicyRequestPolicyNameOk {
			oauthCibaServerPolicyRequestPolicyData[*oauthCibaServerPolicyRequestPolicyId] = *oauthCibaServerPolicyRequestPolicyName
		}
	}

	return &oauthCibaServerPolicyRequestPolicyData, nil
}
