package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateSpIdpConnectionResource{}
)

type PingFederateSpIdpConnectionResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateSpIdpConnectionResource
func SpIdpConnection(clientInfo *connector.PingFederateClientInfo) *PingFederateSpIdpConnectionResource {
	return &PingFederateSpIdpConnectionResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateSpIdpConnectionResource) ResourceType() string {
	return "pingfederate_sp_idp_connection"
}

func (r *PingFederateSpIdpConnectionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	idpConnectionData, err := r.getIdpConnectionData()
	if err != nil {
		return nil, err
	}

	for idpConnectionId, idpConnectionName := range *idpConnectionData {
		commentData := map[string]string{
			"Resource Type":          r.ResourceType(),
			"SP IDP Connection ID":   idpConnectionId,
			"SP IDP Connection Name": idpConnectionName,
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       idpConnectionName,
			ResourceID:         idpConnectionId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateSpIdpConnectionResource) getIdpConnectionData() (*map[string]string, error) {
	idpConnectionData := make(map[string]string)

	idpConnections, response, err := r.clientInfo.ApiClient.SpIdpConnectionsAPI.GetConnections(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetConnections", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if idpConnections == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	idpConnectionsItems, idpConnectionsItemsOk := idpConnections.GetItemsOk()
	if !idpConnectionsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, idpConnection := range idpConnectionsItems {
		idpConnectionId, idpConnectionIdOk := idpConnection.GetIdOk()
		idpConnectionName, idpConnectionNameOk := idpConnection.GetNameOk()

		if idpConnectionIdOk && idpConnectionNameOk {
			idpConnectionData[*idpConnectionId] = *idpConnectionName
		}
	}

	return &idpConnectionData, nil
}
