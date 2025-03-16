package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateSpAdapterResource{}
)

type PingFederateSpAdapterResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateSpAdapterResource
func SpAdapter(clientInfo *connector.ClientInfo) *PingFederateSpAdapterResource {
	return &PingFederateSpAdapterResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateSpAdapterResource) ResourceType() string {
	return "pingfederate_sp_adapter"
}

func (r *PingFederateSpAdapterResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	spAdapterData, err := r.getSpAdapterData()
	if err != nil {
		return nil, err
	}

	for spAdapterId, spAdapterName := range spAdapterData {
		commentData := map[string]string{
			"Sp Adapter ID":   spAdapterId,
			"Sp Adapter Name": spAdapterName,
			"Resource Type":   r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       spAdapterName,
			ResourceID:         spAdapterId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateSpAdapterResource) getSpAdapterData() (map[string]string, error) {
	spAdapterData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.SpAdaptersAPI.GetSpAdapters(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetSpAdapters", r.ResourceType())
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

	for _, spAdapter := range items {
		spAdapterId, spAdapterIdOk := spAdapter.GetIdOk()
		spAdapterName, spAdapterNameOk := spAdapter.GetNameOk()

		if spAdapterIdOk && spAdapterNameOk {
			spAdapterData[*spAdapterId] = *spAdapterName
		}
	}

	return spAdapterData, nil
}
