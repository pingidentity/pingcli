package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateIdpSpConnectionResource{}
)

type PingFederateIdpSpConnectionResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateIdpSpConnectionResource
func IdpSpConnection(clientInfo *connector.PingFederateClientInfo) *PingFederateIdpSpConnectionResource {
	return &PingFederateIdpSpConnectionResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateIdpSpConnectionResource) ResourceType() string {
	return "pingfederate_idp_sp_connection"
}

func (r *PingFederateIdpSpConnectionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	idpSpConnectionData, err := r.getIdpSpConnectionData()
	if err != nil {
		return nil, err
	}

	for idpSpConnectionId, idpSpConnectionName := range *idpSpConnectionData {
		commentData := map[string]string{
			"Idp Sp Connection ID":   idpSpConnectionId,
			"Idp Sp Connection Name": idpSpConnectionName,
			"Resource Type":          r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       idpSpConnectionName,
			ResourceID:         idpSpConnectionId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateIdpSpConnectionResource) getIdpSpConnectionData() (*map[string]string, error) {
	idpSpConnectionData := make(map[string]string)

	apiObj, response, err := r.clientInfo.ApiClient.IdpSpConnectionsAPI.GetSpConnections(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetSpConnections", r.ResourceType())
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

	for _, idpSpConnection := range items {
		idpSpConnectionId, idpSpConnectionIdOk := idpSpConnection.GetIdOk()
		idpSpConnectionName, idpSpConnectionNameOk := idpSpConnection.GetNameOk()

		if idpSpConnectionIdOk && idpSpConnectionNameOk {
			idpSpConnectionData[*idpSpConnectionId] = *idpSpConnectionName
		}
	}

	return &idpSpConnectionData, nil
}
