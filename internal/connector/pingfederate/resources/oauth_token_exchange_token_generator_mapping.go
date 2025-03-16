package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOauthTokenExchangeTokenGeneratorMappingResource{}
)

type PingFederateOauthTokenExchangeTokenGeneratorMappingResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateOauthTokenExchangeTokenGeneratorMappingResource
func OauthTokenExchangeTokenGeneratorMapping(clientInfo *connector.ClientInfo) *PingFederateOauthTokenExchangeTokenGeneratorMappingResource {
	return &PingFederateOauthTokenExchangeTokenGeneratorMappingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOauthTokenExchangeTokenGeneratorMappingResource) ResourceType() string {
	return "pingfederate_oauth_token_exchange_token_generator_mapping"
}

func (r *PingFederateOauthTokenExchangeTokenGeneratorMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthTokenExchangeTokenGeneratorMappingData, err := r.getOauthTokenExchangeTokenGeneratorMappingData()
	if err != nil {
		return nil, err
	}

	for oauthTokenExchangeTokenGeneratorMappingId, oauthTokenExchangeTokenGeneratorMappingInfo := range oauthTokenExchangeTokenGeneratorMappingData {
		oauthTokenExchangeTokenGeneratorMappingSourceId := oauthTokenExchangeTokenGeneratorMappingInfo[0]
		oauthTokenExchangeTokenGeneratorMappingTargetId := oauthTokenExchangeTokenGeneratorMappingInfo[1]

		commentData := map[string]string{
			"Oauth Token Exchange Token Generator Mapping ID": oauthTokenExchangeTokenGeneratorMappingId,
			"Processor Policy ID":                             oauthTokenExchangeTokenGeneratorMappingSourceId,
			"Token Generator ID":                              oauthTokenExchangeTokenGeneratorMappingTargetId,
			"Resource Type":                                   r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_to_%s", oauthTokenExchangeTokenGeneratorMappingSourceId, oauthTokenExchangeTokenGeneratorMappingTargetId),
			ResourceID:         oauthTokenExchangeTokenGeneratorMappingId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOauthTokenExchangeTokenGeneratorMappingResource) getOauthTokenExchangeTokenGeneratorMappingData() (map[string][]string, error) {
	oauthTokenExchangeTokenGeneratorMappingData := make(map[string][]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.OauthTokenExchangeTokenGeneratorMappingsAPI.GetTokenGeneratorMappings(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetTokenGeneratorMappings", r.ResourceType())
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

	for _, oauthTokenExchangeTokenGeneratorMapping := range items {
		oauthTokenExchangeTokenGeneratorMappingId, oauthTokenExchangeTokenGeneratorMappingIdOk := oauthTokenExchangeTokenGeneratorMapping.GetIdOk()
		oauthTokenExchangeTokenGeneratorMappingSourceId, oauthTokenExchangeTokenGeneratorMappingSourceIdOk := oauthTokenExchangeTokenGeneratorMapping.GetSourceIdOk()
		oauthTokenExchangeTokenGeneratorMappingTargetId, oauthTokenExchangeTokenGeneratorMappingTargetIdOk := oauthTokenExchangeTokenGeneratorMapping.GetTargetIdOk()

		if oauthTokenExchangeTokenGeneratorMappingIdOk && oauthTokenExchangeTokenGeneratorMappingSourceIdOk && oauthTokenExchangeTokenGeneratorMappingTargetIdOk {
			oauthTokenExchangeTokenGeneratorMappingData[*oauthTokenExchangeTokenGeneratorMappingId] = []string{*oauthTokenExchangeTokenGeneratorMappingSourceId, *oauthTokenExchangeTokenGeneratorMappingTargetId}
		}
	}

	return oauthTokenExchangeTokenGeneratorMappingData, nil
}
