package resources

import (
	"fmt"

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

func (r *PingFederatePingOneConnectionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.PingOneConnectionsAPI.GetPingOneConnections(r.clientInfo.Context).Execute
	apiFunctionName := "GetPingOneConnections"

	pingoneConnections, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if pingoneConnections == nil {
		l.Error().Msgf("Returned %s() pingoneConnections is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	pingoneConnectionsItems, ok := pingoneConnections.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() pingoneConnections items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, pingoneConnection := range pingoneConnectionsItems {
		pingoneConnectionId, pingoneConnectionIdOk := pingoneConnection.GetIdOk()
		pingoneConnectionName, pingoneConnectionNameOk := pingoneConnection.GetNameOk()

		if pingoneConnectionIdOk && pingoneConnectionNameOk {
			commentData := map[string]string{
				"Resource Type":                    r.ResourceType(),
				"PingOne Connection Resource ID":   *pingoneConnectionId,
				"PingOne Connection Resource Name": *pingoneConnectionName,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *pingoneConnectionName,
				ResourceID:         *pingoneConnectionId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederatePingOneConnectionResource) ResourceType() string {
	return "pingfederate_pingone_connection"
}
