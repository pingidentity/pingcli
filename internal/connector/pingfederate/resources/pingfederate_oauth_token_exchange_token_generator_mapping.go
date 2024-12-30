package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOAuthTokenExchangeTokenGeneratorMappingResource{}
)

type PingFederateOAuthTokenExchangeTokenGeneratorMappingResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateOAuthTokenExchangeTokenGeneratorMappingResource
func OAuthTokenExchangeTokenGeneratorMapping(clientInfo *connector.PingFederateClientInfo) *PingFederateOAuthTokenExchangeTokenGeneratorMappingResource {
	return &PingFederateOAuthTokenExchangeTokenGeneratorMappingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOAuthTokenExchangeTokenGeneratorMappingResource) ResourceType() string {
	return "pingfederate_oauth_token_exchange_token_generator_mapping"
}

func (r *PingFederateOAuthTokenExchangeTokenGeneratorMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	processorPolicyToGeneratorMappingData, err := r.getProcessorPolicyToGeneratorMappingData()
	if err != nil {
		return nil, err
	}

	for mappingId, mappingInfo := range *processorPolicyToGeneratorMappingData {
		sourceId := mappingInfo[0]
		targetId := mappingInfo[1]

		commentData := map[string]string{
			"OAuth Token Exchange Token Generator Mapping ID": mappingId,
			"Processor Policy ID":                             sourceId,
			"Resource Type":                                   r.ResourceType(),
			"Token Generator ID":                              targetId,
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_to_%s", sourceId, targetId),
			ResourceID:         mappingId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOAuthTokenExchangeTokenGeneratorMappingResource) getProcessorPolicyToGeneratorMappingData() (*map[string][]string, error) {
	processorPolicyToGeneratorMappingData := make(map[string][]string)

	processorPolicyToGeneratorMappings, response, err := r.clientInfo.ApiClient.OauthTokenExchangeTokenGeneratorMappingsAPI.GetTokenGeneratorMappings(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetTokenGeneratorMappings", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if processorPolicyToGeneratorMappings == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	processorPolicyToGeneratorMappingsItems, processorPolicyToGeneratorMappingsItemsOk := processorPolicyToGeneratorMappings.GetItemsOk()
	if !processorPolicyToGeneratorMappingsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, mapping := range processorPolicyToGeneratorMappingsItems {
		mappingId, mappingIdOk := mapping.GetIdOk()
		mappingSourceId, mappingSourceIdOk := mapping.GetSourceIdOk()
		mappingTargetId, mappingTargetIdOk := mapping.GetTargetIdOk()

		if mappingIdOk && mappingSourceIdOk && mappingTargetIdOk {
			processorPolicyToGeneratorMappingData[*mappingId] = []string{*mappingSourceId, *mappingTargetId}
		}
	}

	return &processorPolicyToGeneratorMappingData, nil
}
