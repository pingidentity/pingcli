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

func (r *PingFederateSPAuthenticationPolicyContractMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.SpAuthenticationPolicyContractMappingsAPI.GetApcToSpAdapterMappings(r.clientInfo.Context).Execute
	apiFunctionName := "GetApcToSpAdapterMappings"

	apcToSpAdapterMappings, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if apcToSpAdapterMappings == nil {
		l.Error().Msgf("Returned %s() apcToSpAdapterMappings is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	apcToSpAdapterMappingsItems, ok := apcToSpAdapterMappings.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() apcToSpAdapterMappings items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, apcToSpAdapterMapping := range apcToSpAdapterMappingsItems {
		apcToSpAdapterMappingId, apcToSpAdapterMappingIdOk := apcToSpAdapterMapping.GetIdOk()
		apcToSpAdapterMappingSourceID, apcToSpAdapterMappingSourceIDOk := apcToSpAdapterMapping.GetSourceIdOk()
		apcToSpAdapterMappingTargetID, apcToSpAdapterMappingTargetIDOk := apcToSpAdapterMapping.GetTargetIdOk()

		if apcToSpAdapterMappingIdOk && apcToSpAdapterMappingSourceIDOk && apcToSpAdapterMappingTargetIDOk {
			commentData := map[string]string{
				"Resource Type": r.ResourceType(),
				"SP Authentication Policy Contract Mapping Resource ID": *apcToSpAdapterMappingId,
				"Source Authentication Policy Contract Resource ID":     *apcToSpAdapterMappingSourceID,
				"Target SP Adapter Resource ID":                         *apcToSpAdapterMappingTargetID,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_to_%s", *apcToSpAdapterMappingSourceID, *apcToSpAdapterMappingTargetID),
				ResourceID:         *apcToSpAdapterMappingId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederateSPAuthenticationPolicyContractMappingResource) ResourceType() string {
	return "pingfederate_sp_authentication_policy_contract_mapping"
}
