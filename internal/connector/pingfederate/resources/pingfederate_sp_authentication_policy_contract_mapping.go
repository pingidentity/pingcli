package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateSPAuthenticationPolicyContractMappingResource{}
)

type PingFederateSPAuthenticationPolicyContractMappingResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateSPAuthenticationPolicyContractMappingResource
func SPAuthenticationPolicyContractMapping(clientInfo *connector.PingFederateClientInfo) *PingFederateSPAuthenticationPolicyContractMappingResource {
	return &PingFederateSPAuthenticationPolicyContractMappingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateSPAuthenticationPolicyContractMappingResource) ResourceType() string {
	return "pingfederate_sp_authentication_policy_contract_mapping"
}

func (r *PingFederateSPAuthenticationPolicyContractMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	apcToSpAdapterMappingData, err := r.getApcToSpAdapterMappingData()
	if err != nil {
		return nil, err
	}

	for apcToSpAdapterMappingId, apcToSpAdapterMappingInfo := range apcToSpAdapterMappingData {
		apcToSpAdapterMappingSourceID := apcToSpAdapterMappingInfo[0]
		apcToSpAdapterMappingTargetID := apcToSpAdapterMappingInfo[1]

		commentData := map[string]string{
			"Resource Type": r.ResourceType(),
			"Source Authentication Policy Contract ID":     apcToSpAdapterMappingSourceID,
			"SP Authentication Policy Contract Mapping ID": apcToSpAdapterMappingId,
			"Target SP Adapter ID":                         apcToSpAdapterMappingTargetID,
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_to_%s", apcToSpAdapterMappingSourceID, apcToSpAdapterMappingTargetID),
			ResourceID:         apcToSpAdapterMappingId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateSPAuthenticationPolicyContractMappingResource) getApcToSpAdapterMappingData() (map[string][]string, error) {
	apcToSpAdapterMappingData := make(map[string][]string)

	apcToSpAdapterMappings, response, err := r.clientInfo.ApiClient.SpAuthenticationPolicyContractMappingsAPI.GetApcToSpAdapterMappings(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetApcToSpAdapterMappings", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if apcToSpAdapterMappings == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	apcToSpAdapterMappingsItems, apcToSpAdapterMappingsItemsOk := apcToSpAdapterMappings.GetItemsOk()
	if !apcToSpAdapterMappingsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, apcToSpAdapterMapping := range apcToSpAdapterMappingsItems {
		apcToSpAdapterMappingId, apcToSpAdapterMappingIdOk := apcToSpAdapterMapping.GetIdOk()
		apcToSpAdapterMappingSourceID, apcToSpAdapterMappingSourceIDOk := apcToSpAdapterMapping.GetSourceIdOk()
		apcToSpAdapterMappingTargetID, apcToSpAdapterMappingTargetIDOk := apcToSpAdapterMapping.GetTargetIdOk()

		if apcToSpAdapterMappingIdOk && apcToSpAdapterMappingSourceIDOk && apcToSpAdapterMappingTargetIDOk {
			apcToSpAdapterMappingData[*apcToSpAdapterMappingId] = []string{*apcToSpAdapterMappingSourceID, *apcToSpAdapterMappingTargetID}
		}
	}

	return apcToSpAdapterMappingData, nil
}
