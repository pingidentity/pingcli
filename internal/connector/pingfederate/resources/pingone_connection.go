package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederatePingoneConnectionResource{}
)

type PingFederatePingoneConnectionResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederatePingoneConnectionResource
func PingoneConnection(clientInfo *connector.ClientInfo) *PingFederatePingoneConnectionResource {
	return &PingFederatePingoneConnectionResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederatePingoneConnectionResource) ResourceType() string {
	return "pingfederate_pingone_connection"
}

func (r *PingFederatePingoneConnectionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	pingoneConnectionData, err := r.getPingoneConnectionData()
	if err != nil {
		return nil, err
	}

	for pingoneConnectionId, pingoneConnectionName := range pingoneConnectionData {
		commentData := map[string]string{
			"Pingone Connection ID":   pingoneConnectionId,
			"Pingone Connection Name": pingoneConnectionName,
			"Resource Type":           r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       pingoneConnectionName,
			ResourceID:         pingoneConnectionId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederatePingoneConnectionResource) getPingoneConnectionData() (map[string]string, error) {
	pingoneConnectionData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.PingOneConnectionsAPI.GetPingOneConnections(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetPingOneConnections", r.ResourceType())
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

	for _, pingoneConnection := range items {
		pingoneConnectionId, pingoneConnectionIdOk := pingoneConnection.GetIdOk()
		pingoneConnectionName, pingoneConnectionNameOk := pingoneConnection.GetNameOk()

		if pingoneConnectionIdOk && pingoneConnectionNameOk {
			pingoneConnectionData[*pingoneConnectionId] = *pingoneConnectionName
		}
	}

	return pingoneConnectionData, nil
}
