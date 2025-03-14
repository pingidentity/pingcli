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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateSpIdpConnectionResource
func SpIdpConnection(clientInfo *connector.ClientInfo) *PingFederateSpIdpConnectionResource {
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

	spIdpConnectionData, err := r.getSpIdpConnectionData()
	if err != nil {
		return nil, err
	}

	for spIdpConnectionId, spIdpConnectionName := range spIdpConnectionData {
		commentData := map[string]string{
			"Sp Idp Connection ID":   spIdpConnectionId,
			"Sp Idp Connection Name": spIdpConnectionName,
			"Resource Type":          r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       spIdpConnectionName,
			ResourceID:         spIdpConnectionId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateSpIdpConnectionResource) getSpIdpConnectionData() (map[string]string, error) {
	spIdpConnectionData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.SpIdpConnectionsAPI.GetConnections(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetConnections", r.ResourceType())
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

	for _, spIdpConnection := range items {
		spIdpConnectionId, spIdpConnectionIdOk := spIdpConnection.GetIdOk()
		spIdpConnectionName, spIdpConnectionNameOk := spIdpConnection.GetNameOk()

		if spIdpConnectionIdOk && spIdpConnectionNameOk {
			spIdpConnectionData[*spIdpConnectionId] = *spIdpConnectionName
		}
	}

	return spIdpConnectionData, nil
}
