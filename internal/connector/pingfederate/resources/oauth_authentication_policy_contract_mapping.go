package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOauthAuthenticationPolicyContractMappingResource{}
)

type PingFederateOauthAuthenticationPolicyContractMappingResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateOauthAuthenticationPolicyContractMappingResource
func OauthAuthenticationPolicyContractMapping(clientInfo *connector.PingFederateClientInfo) *PingFederateOauthAuthenticationPolicyContractMappingResource {
	return &PingFederateOauthAuthenticationPolicyContractMappingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOauthAuthenticationPolicyContractMappingResource) ResourceType() string {
	return "pingfederate_oauth_authentication_policy_contract_mapping"
}

func (r *PingFederateOauthAuthenticationPolicyContractMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthAuthenticationPolicyContractMappingData, err := r.getOauthAuthenticationPolicyContractMappingData()
	if err != nil {
		return nil, err
	}

	for _, oauthAuthenticationPolicyContractMappingId := range *oauthAuthenticationPolicyContractMappingData {
		commentData := map[string]string{
			"Oauth Authentication Policy Contract Mapping ID": oauthAuthenticationPolicyContractMappingId,
			"Resource Type": r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_mapping", oauthAuthenticationPolicyContractMappingId),
			ResourceID:         oauthAuthenticationPolicyContractMappingId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOauthAuthenticationPolicyContractMappingResource) getOauthAuthenticationPolicyContractMappingData() (*[]string, error) {
	oauthAuthenticationPolicyContractMappingData := []string{}

	apiObj, response, err := r.clientInfo.ApiClient.OauthAuthenticationPolicyContractMappingsAPI.GetApcMappings(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetApcMappings", r.ResourceType())
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

	for _, oauthAuthenticationPolicyContractMapping := range items {
		oauthAuthenticationPolicyContractMappingId, oauthAuthenticationPolicyContractMappingIdOk := oauthAuthenticationPolicyContractMapping.GetIdOk()

		if oauthAuthenticationPolicyContractMappingIdOk {
			oauthAuthenticationPolicyContractMappingData = append(oauthAuthenticationPolicyContractMappingData, *oauthAuthenticationPolicyContractMappingId)
		}
	}

	return &oauthAuthenticationPolicyContractMappingData, nil
}
