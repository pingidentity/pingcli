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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateIdpTokenProcessorResource
func IdpTokenProcessor(clientInfo *connector.ClientInfo) *PingFederateIdpTokenProcessorResource {
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

	idpTokenProcessorData, err := r.getIdpTokenProcessorData()
	if err != nil {
		return nil, err
	}

	for idpTokenProcessorId, idpTokenProcessorName := range *idpTokenProcessorData {
		commentData := map[string]string{
			"Idp Token Processor ID":   idpTokenProcessorId,
			"Idp Token Processor Name": idpTokenProcessorName,
			"Resource Type":            r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       idpTokenProcessorName,
			ResourceID:         idpTokenProcessorId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateIdpTokenProcessorResource) getIdpTokenProcessorData() (*map[string]string, error) {
	idpTokenProcessorData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.IdpTokenProcessorsAPI.GetTokenProcessors(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetTokenProcessors", r.ResourceType())
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

	for _, idpTokenProcessor := range items {
		idpTokenProcessorId, idpTokenProcessorIdOk := idpTokenProcessor.GetIdOk()
		idpTokenProcessorName, idpTokenProcessorNameOk := idpTokenProcessor.GetNameOk()

		if idpTokenProcessorIdOk && idpTokenProcessorNameOk {
			idpTokenProcessorData[*idpTokenProcessorId] = *idpTokenProcessorName
		}
	}

	return &idpTokenProcessorData, nil
}
