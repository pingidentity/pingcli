package resources

import (
	"fmt"

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

func (r *PingFederateDataStoreResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.DataStoresAPI.GetDataStores(r.clientInfo.Context).Execute
	apiFunctionName := "GetDataStores"

	dataStores, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if dataStores == nil {
		l.Error().Msgf("Returned %s() dataStores is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	dataStoresItems, ok := dataStores.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() dataStores items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, dataStore := range dataStoresItems {
		dataStoreId, dataStoreIdOk := dataStore.GetIdOk()
		dataStoreType, dataStoreTypeOk := dataStore.GetTypeOk()

		if dataStoreIdOk && dataStoreTypeOk {
			commentData := map[string]string{
				"Resource Type":          r.ResourceType(),
				"Data Store Resource ID": *dataStoreId,
				"Data Store Type":        *dataStoreType,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", *dataStoreId, *dataStoreType),
				ResourceID:         *dataStoreId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederateDataStoreResource) ResourceType() string {
	return "pingfederate_data_store"
}
