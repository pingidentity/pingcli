package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOAuthAuthenticationPolicyContractMappingResource{}
)

type PingFederateOAuthAuthenticationPolicyContractMappingResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateOAuthAuthenticationPolicyContractMappingResource
func OAuthAuthenticationPolicyContractMapping(clientInfo *connector.PingFederateClientInfo) *PingFederateOAuthAuthenticationPolicyContractMappingResource {
	return &PingFederateOAuthAuthenticationPolicyContractMappingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOAuthAuthenticationPolicyContractMappingResource) ResourceType() string {
	return "pingfederate_oauth_authentication_policy_contract_mapping"
}

func (r *PingFederateOAuthAuthenticationPolicyContractMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	apcToPersistentGrantMappingData, err := r.getApcToPersistentGrantMappingData()
	if err != nil {
		return nil, err
	}

	for _, mappingId := range *apcToPersistentGrantMappingData {
		commentData := map[string]string{
			"Authentication Policy Contract Mapping ID": mappingId,
			"Resource Type": r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_mapping", mappingId),
			ResourceID:         mappingId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOAuthAuthenticationPolicyContractMappingResource) getApcToPersistentGrantMappingData() (*[]string, error) {
	apcToPersistentGrantMappingData := []string{}

	apcToPersistentGrantMappings, response, err := r.clientInfo.ApiClient.OauthAuthenticationPolicyContractMappingsAPI.GetApcMappings(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetApcMappings", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if apcToPersistentGrantMappings == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	apcToPersistentGrantMappingsItems, apcToPersistentGrantMappingsItemsOk := apcToPersistentGrantMappings.GetItemsOk()
	if !apcToPersistentGrantMappingsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, apcToPersistentGrantMapping := range apcToPersistentGrantMappingsItems {
		apcToPersistentGrantMappingId, apcToPersistentGrantMappingIdOk := apcToPersistentGrantMapping.GetIdOk()

		if apcToPersistentGrantMappingIdOk {
			apcToPersistentGrantMappingData = append(apcToPersistentGrantMappingData, *apcToPersistentGrantMappingId)
		}
	}

	return &apcToPersistentGrantMappingData, nil
}
