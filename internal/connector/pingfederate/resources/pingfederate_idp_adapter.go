package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateIDPAdapterResource{}
)

type PingFederateIDPAdapterResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateIDPAdapterResource
func IDPAdapter(clientInfo *connector.PingFederateClientInfo) *PingFederateIDPAdapterResource {
	return &PingFederateIDPAdapterResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateIDPAdapterResource) ResourceType() string {
	return "pingfederate_idp_adapter"
}

func (r *PingFederateIDPAdapterResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	idpAdapterData, err := r.getIDPAdapterData()
	if err != nil {
		return nil, err
	}

	for idpAdapterId, idpAdapterName := range idpAdapterData {
		commentData := map[string]string{
			"IDP Adapter ID":   idpAdapterId,
			"IDP Adapter Name": idpAdapterName,
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

func (r *PingFederateIDPAdapterResource) getIDPAdapterData() (map[string]string, error) {
	idpAdapterData := make(map[string]string)

	idpAdapters, response, err := r.clientInfo.ApiClient.IdpAdaptersAPI.GetIdpAdapters(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetIdpAdapters", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if idpAdapters == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	idpAdaptersItems, idpAdaptersItemsOk := idpAdapters.GetItemsOk()
	if !idpAdaptersItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, idpAdapter := range idpAdaptersItems {
		idpAdapterId, idpAdapterIdOk := idpAdapter.GetIdOk()
		idpAdapterName, idpAdapterNameOk := idpAdapter.GetNameOk()

		if idpAdapterIdOk && idpAdapterNameOk {
			idpAdapterData[*idpAdapterId] = *idpAdapterName
		}
	}

	return idpAdapterData, nil
}
