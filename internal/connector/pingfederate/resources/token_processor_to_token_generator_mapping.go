package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateTokenProcessorToTokenGeneratorMappingResource{}
)

type PingFederateTokenProcessorToTokenGeneratorMappingResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateTokenProcessorToTokenGeneratorMappingResource
func TokenProcessorToTokenGeneratorMapping(clientInfo *connector.ClientInfo) *PingFederateTokenProcessorToTokenGeneratorMappingResource {
	return &PingFederateTokenProcessorToTokenGeneratorMappingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateTokenProcessorToTokenGeneratorMappingResource) ResourceType() string {
	return "pingfederate_token_processor_to_token_generator_mapping"
}

func (r *PingFederateTokenProcessorToTokenGeneratorMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	tokenProcessorToTokenGeneratorMappingData, err := r.getTokenProcessorToTokenGeneratorMappingData()
	if err != nil {
		return nil, err
	}

	for tokenProcessorToTokenGeneratorMappingId, tokenProcessorToTokenGeneratorMappingInfo := range tokenProcessorToTokenGeneratorMappingData {
		tokenProcessorToTokenGeneratorMappingSourceId := tokenProcessorToTokenGeneratorMappingInfo[0]
		tokenProcessorToTokenGeneratorMappingTargetId := tokenProcessorToTokenGeneratorMappingInfo[1]

		commentData := map[string]string{
			"Token Processor To Token Generator Mapping ID": tokenProcessorToTokenGeneratorMappingId,
			"Token Processor ID":                            tokenProcessorToTokenGeneratorMappingSourceId,
			"Token Generator ID":                            tokenProcessorToTokenGeneratorMappingTargetId,
			"Resource Type":                                 r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_to_%s", tokenProcessorToTokenGeneratorMappingSourceId, tokenProcessorToTokenGeneratorMappingTargetId),
			ResourceID:         tokenProcessorToTokenGeneratorMappingId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateTokenProcessorToTokenGeneratorMappingResource) getTokenProcessorToTokenGeneratorMappingData() (map[string][]string, error) {
	tokenProcessorToTokenGeneratorMappingData := make(map[string][]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.TokenProcessorToTokenGeneratorMappingsAPI.GetTokenToTokenMappings(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetTokenToTokenMappings", r.ResourceType())
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

	for _, tokenProcessorToTokenGeneratorMapping := range items {
		tokenProcessorToTokenGeneratorMappingId, tokenProcessorToTokenGeneratorMappingIdOk := tokenProcessorToTokenGeneratorMapping.GetIdOk()
		tokenProcessorToTokenGeneratorMappingSourceId, tokenProcessorToTokenGeneratorMappingSourceIdOk := tokenProcessorToTokenGeneratorMapping.GetSourceIdOk()
		tokenProcessorToTokenGeneratorMappingTargetId, tokenProcessorToTokenGeneratorMappingTargetIdOk := tokenProcessorToTokenGeneratorMapping.GetTargetIdOk()

		if tokenProcessorToTokenGeneratorMappingIdOk && tokenProcessorToTokenGeneratorMappingSourceIdOk && tokenProcessorToTokenGeneratorMappingTargetIdOk {
			tokenProcessorToTokenGeneratorMappingData[*tokenProcessorToTokenGeneratorMappingId] = []string{*tokenProcessorToTokenGeneratorMappingSourceId, *tokenProcessorToTokenGeneratorMappingTargetId}
		}
	}

	return tokenProcessorToTokenGeneratorMappingData, nil
}
