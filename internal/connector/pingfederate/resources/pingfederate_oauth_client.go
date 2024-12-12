package resources

import (
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

func (r *PingFederateOAuthClientResource) ResourceType() string {
	return "pingfederate_oauth_client"
}

func (r *PingFederateOAuthClientResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthClientData, err := r.getOAuthClientData()
	if err != nil {
		return nil, err
	}

	for oauthClientId, oauthClientName := range *oauthClientData {
		commentData := map[string]string{
			"OAuth Client Resource ID":   oauthClientId,
			"OAuth Client Resource Name": oauthClientName,
			"Resource Type":              r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       oauthClientName,
			ResourceID:         oauthClientId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOAuthClientResource) getOAuthClientData() (*map[string]string, error) {
	oauthClientData := make(map[string]string)

	clients, response, err := r.clientInfo.ApiClient.OauthClientsAPI.GetOauthClients(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetOauthClients", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if clients == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	clientsItems, ok := clients.GetItemsOk()
	if !ok {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, client := range clientsItems {
		clientId, clientIdOk := client.GetClientIdOk()
		clientName, clientNameOk := client.GetNameOk()

		if clientIdOk && clientNameOk {
			oauthClientData[*clientId] = *clientName
		}
	}

	return &oauthClientData, nil
}
