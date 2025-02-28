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
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateSpAdapterResource
func SpAdapter(clientInfo *connector.PingFederateClientInfo) *PingFederateSpAdapterResource {
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

	spAdaptersData, err := r.getSpAdaptersData()
	if err != nil {
		return nil, err
	}

	for spAdapterId, spAdapterName := range spAdaptersData {
		commentData := map[string]string{
			"Resource Type":   r.ResourceType(),
			"SP Adapter ID":   spAdapterId,
			"SP Adapter Name": spAdapterName,
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

func (r *PingFederateSpAdapterResource) getSpAdaptersData() (map[string]string, error) {
	spAdaptersData := make(map[string]string)

	spAdapters, response, err := r.clientInfo.ApiClient.SpAdaptersAPI.GetSpAdapters(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetSpAdapters", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if spAdapters == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	spAdaptersItems, spAdaptersItemsOk := spAdapters.GetItemsOk()
	if !spAdaptersItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, spAdapter := range spAdaptersItems {
		spAdapterId, spAdapterIdOk := spAdapter.GetIdOk()
		spAdapterName, spAdapterNameOk := spAdapter.GetNameOk()

		if spAdapterIdOk && spAdapterNameOk {
			spAdaptersData[*spAdapterId] = *spAdapterName
		}
	}

	return spAdaptersData, nil
}
