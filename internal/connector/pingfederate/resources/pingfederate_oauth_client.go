package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOAuthClientResource{}
)

type PingFederateOAuthClientResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateOAuthClientResource
func OAuthClient(clientInfo *connector.PingFederateClientInfo) *PingFederateOAuthClientResource {
	return &PingFederateOAuthClientResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOAuthClientResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.OauthClientsAPI.GetOauthClients(r.clientInfo.Context).Execute
	apiFunctionName := "GetOauthClients"

	clients, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if clients == nil {
		l.Error().Msgf("Returned %s() clients is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	clientsItems, ok := clients.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() clients items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, client := range clientsItems {
		clientId, clientIdOk := client.GetClientIdOk()
		clientName, clientNameOk := client.GetNameOk()

		if clientIdOk && clientNameOk {
			commentData := map[string]string{
				"Resource Type":              r.ResourceType(),
				"OAuth Client Resource ID":   *clientId,
				"OAuth Client Resource Name": *clientName,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *clientName,
				ResourceID:         *clientId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederateOAuthClientResource) ResourceType() string {
	return "pingfederate_oauth_client"
}
