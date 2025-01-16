package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateDataStoreResource{}
)

type PingFederateDataStoreResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateDataStoreResource
func DataStore(clientInfo *connector.PingFederateClientInfo) *PingFederateDataStoreResource {
	return &PingFederateDataStoreResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateDataStoreResource) ResourceType() string {
	return "pingfederate_data_store"
}

func (r *PingFederateDataStoreResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	dataStoreData, err := r.getDataStoreData()
	if err != nil {
		return nil, err
	}

	for dataStoreId, dataStoreName := range *dataStoreData {
		commentData := map[string]string{
			"Data Store ID":   dataStoreId,
			"Data Store Name": dataStoreName,
			"Resource Type":   r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       dataStoreName,
			ResourceID:         dataStoreId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateDataStoreResource) getDataStoreData() (*map[string]string, error) {
	dataStoreData := make(map[string]string)

	apiObj, response, err := r.clientInfo.ApiClient.DataStoresAPI.GetDataStores(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetDataStores", r.ResourceType())
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

	for _, dataStore := range items {
		dataStoreId, dataStoreIdOk := dataStore.GetIdOk()
		dataStoreType, dataStoreTypeOk := dataStore.GetTypeOk()

		if dataStoreIdOk && dataStoreTypeOk {
			dataStoreData[*dataStoreId] = *dataStoreType
		}
	}

	return &dataStoreData, nil
}
