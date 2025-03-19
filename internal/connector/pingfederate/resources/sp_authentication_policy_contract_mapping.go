// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateSpAuthenticationPolicyContractMappingResource{}
)

type PingFederateSpAuthenticationPolicyContractMappingResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateSpAuthenticationPolicyContractMappingResource
func SpAuthenticationPolicyContractMapping(clientInfo *connector.ClientInfo) *PingFederateSpAuthenticationPolicyContractMappingResource {
	return &PingFederateSpAuthenticationPolicyContractMappingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateSpAuthenticationPolicyContractMappingResource) ResourceType() string {
	return "pingfederate_sp_authentication_policy_contract_mapping"
}

func (r *PingFederateSpAuthenticationPolicyContractMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	spAuthenticationPolicyContractMappingData, err := r.getSpAuthenticationPolicyContractMappingData()
	if err != nil {
		return nil, err
	}

	for spAuthenticationPolicyContractMappingId, spAuthenticationPolicyContractMappingInfo := range spAuthenticationPolicyContractMappingData {
		spAuthenticationPolicyContractMappingSourceId := spAuthenticationPolicyContractMappingInfo[0]
		spAuthenticationPolicyContractMappingTargetId := spAuthenticationPolicyContractMappingInfo[1]

		commentData := map[string]string{
			"Sp Authentication Policy Contract Mapping ID": spAuthenticationPolicyContractMappingId,
			"Authentication Policy Contract ID":            spAuthenticationPolicyContractMappingSourceId,
			"Sp Adapter ID":                                spAuthenticationPolicyContractMappingTargetId,
			"Resource Type":                                r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_to_%s", spAuthenticationPolicyContractMappingSourceId, spAuthenticationPolicyContractMappingTargetId),
			ResourceID:         spAuthenticationPolicyContractMappingId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateSpAuthenticationPolicyContractMappingResource) getSpAuthenticationPolicyContractMappingData() (map[string][]string, error) {
	spAuthenticationPolicyContractMappingData := make(map[string][]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.SpAuthenticationPolicyContractMappingsAPI.GetApcToSpAdapterMappings(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetApcToSpAdapterMappings", r.ResourceType())
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

	for _, spAuthenticationPolicyContractMapping := range items {
		spAuthenticationPolicyContractMappingId, spAuthenticationPolicyContractMappingIdOk := spAuthenticationPolicyContractMapping.GetIdOk()
		spAuthenticationPolicyContractMappingSourceId, spAuthenticationPolicyContractMappingSourceIdOk := spAuthenticationPolicyContractMapping.GetSourceIdOk()
		spAuthenticationPolicyContractMappingTargetId, spAuthenticationPolicyContractMappingTargetIdOk := spAuthenticationPolicyContractMapping.GetTargetIdOk()

		if spAuthenticationPolicyContractMappingIdOk && spAuthenticationPolicyContractMappingSourceIdOk && spAuthenticationPolicyContractMappingTargetIdOk {
			spAuthenticationPolicyContractMappingData[*spAuthenticationPolicyContractMappingId] = []string{*spAuthenticationPolicyContractMappingSourceId, *spAuthenticationPolicyContractMappingTargetId}
		}
	}

	return spAuthenticationPolicyContractMappingData, nil
}
