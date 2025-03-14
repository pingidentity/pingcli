package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOauthClientResource{}
)

type PingFederateOauthClientResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateOauthClientResource
func OauthClient(clientInfo *connector.ClientInfo) *PingFederateOauthClientResource {
	return &PingFederateOauthClientResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOauthClientResource) ResourceType() string {
	return "pingfederate_oauth_client"
}

func (r *PingFederateOauthClientResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthClientData, err := r.getOauthClientData()
	if err != nil {
		return nil, err
	}

	for oauthClientId, oauthClientName := range oauthClientData {
		commentData := map[string]string{
			"Oauth Client ID":   oauthClientId,
			"Oauth Client Name": oauthClientName,
			"Resource Type":     r.ResourceType(),
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

func (r *PingFederateOauthClientResource) getOauthClientData() (map[string]string, error) {
	oauthClientData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.OauthClientsAPI.GetOauthClients(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetOauthClients", r.ResourceType())
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

	for _, oauthClient := range items {
		oauthClientId, oauthClientIdOk := oauthClient.GetClientIdOk()
		oauthClientName, oauthClientNameOk := oauthClient.GetNameOk()

		if oauthClientIdOk && oauthClientNameOk {
			oauthClientData[*oauthClientId] = *oauthClientName
		}
	}

	return oauthClientData, nil
}
