package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederatePingOneConnectionResource{}
)

type PingFederatePingOneConnectionResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederatePingOneConnectionResource
func PingOneConnection(clientInfo *connector.PingFederateClientInfo) *PingFederatePingOneConnectionResource {
	return &PingFederatePingOneConnectionResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederatePingOneConnectionResource) ResourceType() string {
	return "pingfederate_pingone_connection"
}

func (r *PingFederatePingOneConnectionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	pingoneConnectionData, err := r.getPingOneConnectionData()
	if err != nil {
		return nil, err
	}

	for pingoneConnectionId, pingoneConnectionName := range *pingoneConnectionData {
		commentData := map[string]string{
			"PingOne Connection ID":   pingoneConnectionId,
			"PingOne Connection Name": pingoneConnectionName,
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

func (r *PingFederatePingOneConnectionResource) getPingOneConnectionData() (*map[string]string, error) {
	pingoneConnectionData := make(map[string]string)

	pingoneConnections, response, err := r.clientInfo.ApiClient.PingOneConnectionsAPI.GetPingOneConnections(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetPingOneConnections", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if pingoneConnections == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	pingoneConnectionsItems, pingoneConnectionsItemsOk := pingoneConnections.GetItemsOk()
	if !pingoneConnectionsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, pingoneConnection := range pingoneConnectionsItems {
		pingoneConnectionId, pingoneConnectionIdOk := pingoneConnection.GetIdOk()
		pingoneConnectionName, pingoneConnectionNameOk := pingoneConnection.GetNameOk()

		if pingoneConnectionIdOk && pingoneConnectionNameOk {
			pingoneConnectionData[*pingoneConnectionId] = *pingoneConnectionName
		}
	}

	return &pingoneConnectionData, nil
}
