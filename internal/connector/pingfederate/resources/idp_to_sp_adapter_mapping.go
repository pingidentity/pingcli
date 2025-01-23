package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateIdpToSpAdapterMappingResource{}
)

type PingFederateIdpToSpAdapterMappingResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateIdpToSpAdapterMappingResource
func IdpToSpAdapterMapping(clientInfo *connector.PingFederateClientInfo) *PingFederateIdpToSpAdapterMappingResource {
	return &PingFederateIdpToSpAdapterMappingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateIdpToSpAdapterMappingResource) ResourceType() string {
	return "pingfederate_idp_to_sp_adapter_mapping"
}

func (r *PingFederateIdpToSpAdapterMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	idpToSpAdapterMappingData, err := r.getIdpToSpAdapterMappingData()
	if err != nil {
		return nil, err
	}

	for idpToSpAdapterMappingId, idpToSpAdapterMappingInfo := range *idpToSpAdapterMappingData {
		idpToSpAdapterMappingSourceId := idpToSpAdapterMappingInfo[0]
		idpToSpAdapterMappingTargetId := idpToSpAdapterMappingInfo[1]

		commentData := map[string]string{
			"Idp To Sp Adapter Mapping ID":        idpToSpAdapterMappingId,
			"Idp To Sp Adapter Mapping Source ID": idpToSpAdapterMappingSourceId,
			"Idp To Sp Adapter Mapping Target ID": idpToSpAdapterMappingTargetId,
			"Resource Type":                       r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_to_%s", idpToSpAdapterMappingSourceId, idpToSpAdapterMappingTargetId),
			ResourceID:         idpToSpAdapterMappingId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateIdpToSpAdapterMappingResource) getIdpToSpAdapterMappingData() (*map[string][]string, error) {
	idpToSpAdapterMappingData := make(map[string][]string)

	apiObj, response, err := r.clientInfo.ApiClient.IdpToSpAdapterMappingAPI.GetIdpToSpAdapterMappings(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetIdpToSpAdapterMappings", r.ResourceType())
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

	for _, idpToSpAdapterMapping := range items {
		idpToSpAdapterMappingId, idpToSpAdapterMappingIdOk := idpToSpAdapterMapping.GetIdOk()
		idpToSpAdapterMappingSourceId, idpToSpAdapterMappingSourceIdOk := idpToSpAdapterMapping.GetSourceIdOk()
		idpToSpAdapterMappingTargetId, idpToSpAdapterMappingTargetIdOk := idpToSpAdapterMapping.GetTargetIdOk()

		if idpToSpAdapterMappingIdOk && idpToSpAdapterMappingSourceIdOk && idpToSpAdapterMappingTargetIdOk {
			idpToSpAdapterMappingData[*idpToSpAdapterMappingId] = []string{*idpToSpAdapterMappingSourceId, *idpToSpAdapterMappingTargetId}
		}
	}

	return &idpToSpAdapterMappingData, nil
}
