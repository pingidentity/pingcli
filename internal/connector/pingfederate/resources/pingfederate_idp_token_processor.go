package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateIdpTokenProcessorResource{}
)

type PingFederateIdpTokenProcessorResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateIdpTokenProcessorResource
func IdpTokenProcessor(clientInfo *connector.PingFederateClientInfo) *PingFederateIdpTokenProcessorResource {
	return &PingFederateIdpTokenProcessorResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateIdpTokenProcessorResource) ResourceType() string {
	return "pingfederate_idp_token_processor"
}

func (r *PingFederateIdpTokenProcessorResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	tokenProcessorData, err := r.getTokenProcessorData()
	if err != nil {
		return nil, err
	}

	for tokenProcessorId, tokenProcessorName := range tokenProcessorData {
		commentData := map[string]string{
			"IDP Token Processor ID":   tokenProcessorId,
			"IDP Token Processor Name": tokenProcessorName,
			"Resource Type":            r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       tokenProcessorName,
			ResourceID:         tokenProcessorId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateIdpTokenProcessorResource) getTokenProcessorData() (map[string]string, error) {
	tokenProcessorData := make(map[string]string)

	tokenProcessors, response, err := r.clientInfo.ApiClient.IdpTokenProcessorsAPI.GetTokenProcessors(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetTokenProcessors", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if tokenProcessors == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	tokenProcessorsItems, tokenProcessorsItemsOk := tokenProcessors.GetItemsOk()
	if !tokenProcessorsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, tokenProcessor := range tokenProcessorsItems {
		tokenProcessorId, tokenProcessorIdOk := tokenProcessor.GetIdOk()
		tokenProcessorName, tokenProcessorNameOk := tokenProcessor.GetNameOk()

		if tokenProcessorIdOk && tokenProcessorNameOk {
			tokenProcessorData[*tokenProcessorId] = *tokenProcessorName
		}
	}

	return tokenProcessorData, nil
}
