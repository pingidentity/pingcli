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
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateTokenProcessorToTokenGeneratorMappingResource
func TokenProcessorToTokenGeneratorMapping(clientInfo *connector.PingFederateClientInfo) *PingFederateTokenProcessorToTokenGeneratorMappingResource {
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

	tokenToTokenMappingsData, err := r.getTokenToTokenMappingsData()
	if err != nil {
		return nil, err
	}

	for tokenToTokenMappingId, tokenToTokenMappingInfo := range tokenToTokenMappingsData {
		tokenToTokenMappingSourceId := tokenToTokenMappingInfo[0]
		tokenToTokenMappingTargetId := tokenToTokenMappingInfo[1]

		commentData := map[string]string{
			"Resource Type": r.ResourceType(),
			"Token Processor to Token Generator Mapping ID": tokenToTokenMappingId,
			"Token Processor ID":                            tokenToTokenMappingSourceId,
			"Token Generator ID":                            tokenToTokenMappingTargetId,
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_to_%s", tokenToTokenMappingSourceId, tokenToTokenMappingTargetId),
			ResourceID:         tokenToTokenMappingId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateTokenProcessorToTokenGeneratorMappingResource) getTokenToTokenMappingsData() (map[string][]string, error) {
	tokenToTokenMappingsData := make(map[string][]string)

	tokenToTokenMappings, response, err := r.clientInfo.ApiClient.TokenProcessorToTokenGeneratorMappingsAPI.GetTokenToTokenMappings(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetTokenToTokenMappings", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if tokenToTokenMappings == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	tokenToTokenMappingsItems, tokenToTokenMappingsItemsOk := tokenToTokenMappings.GetItemsOk()
	if !tokenToTokenMappingsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, tokenToTokenMapping := range tokenToTokenMappingsItems {
		tokenToTokenMappingId, tokenToTokenMappingIdOk := tokenToTokenMapping.GetIdOk()
		tokenToTokenMappingSourceId, tokenToTokenMappingSourceIdOk := tokenToTokenMapping.GetSourceIdOk()
		tokenToTokenMappingTargetId, tokenToTokenMappingTargetIdOk := tokenToTokenMapping.GetTargetIdOk()

		if tokenToTokenMappingIdOk && tokenToTokenMappingSourceIdOk && tokenToTokenMappingTargetIdOk {
			tokenToTokenMappingsData[*tokenToTokenMappingId] = []string{*tokenToTokenMappingSourceId, *tokenToTokenMappingTargetId}
		}
	}

	return tokenToTokenMappingsData, nil
}
