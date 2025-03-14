package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateIdpAdapterResource{}
)

type PingFederateIdpAdapterResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateIdpAdapterResource
func IdpAdapter(clientInfo *connector.ClientInfo) *PingFederateIdpAdapterResource {
	return &PingFederateIdpAdapterResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateIdpAdapterResource) ResourceType() string {
	return "pingfederate_idp_adapter"
}

func (r *PingFederateIdpAdapterResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	idpAdapterData, err := r.getIdpAdapterData()
	if err != nil {
		return nil, err
	}

	for idpAdapterId, idpAdapterName := range idpAdapterData {
		commentData := map[string]string{
			"Idp Adapter ID":   idpAdapterId,
			"Idp Adapter Name": idpAdapterName,
			"Resource Type":    r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       idpAdapterName,
			ResourceID:         idpAdapterId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateIdpAdapterResource) getIdpAdapterData() (map[string]string, error) {
	idpAdapterData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.IdpAdaptersAPI.GetIdpAdapters(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetIdpAdapters", r.ResourceType())
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

	for _, idpAdapter := range items {
		idpAdapterId, idpAdapterIdOk := idpAdapter.GetIdOk()
		idpAdapterName, idpAdapterNameOk := idpAdapter.GetNameOk()

		if idpAdapterIdOk && idpAdapterNameOk {
			idpAdapterData[*idpAdapterId] = *idpAdapterName
		}
	}

	return idpAdapterData, nil
}
