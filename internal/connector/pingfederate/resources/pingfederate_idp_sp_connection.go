package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateIDPSPConnectionResource{}
)

type PingFederateIDPSPConnectionResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateIDPSPConnectionResource
func IDPSPConnection(clientInfo *connector.PingFederateClientInfo) *PingFederateIDPSPConnectionResource {
	return &PingFederateIDPSPConnectionResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateIDPSPConnectionResource) ResourceType() string {
	return "pingfederate_idp_sp_connection"
}

func (r *PingFederateIDPSPConnectionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	spConnectionData, err := r.getSpConnectionData()
	if err != nil {
		return nil, err
	}

	for spConnectionId, spConnectionName := range *spConnectionData {
		commentData := map[string]string{
			"IDP SP Connection ID":   spConnectionId,
			"IDP SP Connection Name": spConnectionName,
			"Resource Type":          r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       spConnectionName,
			ResourceID:         spConnectionId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateIDPSPConnectionResource) getSpConnectionData() (*map[string]string, error) {
	spConnectionData := make(map[string]string)

	spConnections, response, err := r.clientInfo.ApiClient.IdpSpConnectionsAPI.GetSpConnections(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetSpConnections", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if spConnections == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	spConnectionsItems, spConnectionsItemsOk := spConnections.GetItemsOk()
	if !spConnectionsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, spConnection := range spConnectionsItems {
		spConnectionId, spConnectionIdOk := spConnection.GetIdOk()
		spConnectionName, spConnectionNameOk := spConnection.GetNameOk()

		if spConnectionIdOk && spConnectionNameOk {
			spConnectionData[*spConnectionId] = *spConnectionName
		}
	}

	return &spConnectionData, nil
}
