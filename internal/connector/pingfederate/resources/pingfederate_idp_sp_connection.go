package resources

import (
	"fmt"

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

func (r *PingFederateIDPSPConnectionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.IdpSpConnectionsAPI.GetSpConnections(r.clientInfo.Context).Execute
	apiFunctionName := "GetSpConnections"

	spConnections, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if spConnections == nil {
		l.Error().Msgf("Returned %s() spConnections is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	spConnectionsItems, ok := spConnections.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() spConnections items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, spConnection := range spConnectionsItems {
		spConnectionId, spConnectionIdOk := spConnection.GetIdOk()
		spConnectionName, spConnectionNameOk := spConnection.GetNameOk()

		if spConnectionIdOk && spConnectionNameOk {
			commentData := map[string]string{
				"Resource Type":                   r.ResourceType(),
				"IDP SP Connection Resource ID":   *spConnectionId,
				"IDP SP Connection Resource Name": *spConnectionName,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *spConnectionName,
				ResourceID:         *spConnectionId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederateIDPSPConnectionResource) ResourceType() string {
	return "pingfederate_idp_sp_connection"
}
