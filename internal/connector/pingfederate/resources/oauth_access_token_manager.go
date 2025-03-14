package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOauthAccessTokenManagerResource{}
)

type PingFederateOauthAccessTokenManagerResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateOauthAccessTokenManagerResource
func OauthAccessTokenManager(clientInfo *connector.ClientInfo) *PingFederateOauthAccessTokenManagerResource {
	return &PingFederateOauthAccessTokenManagerResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOauthAccessTokenManagerResource) ResourceType() string {
	return "pingfederate_oauth_access_token_manager"
}

func (r *PingFederateOauthAccessTokenManagerResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthAccessTokenManagerData, err := r.getOauthAccessTokenManagerData()
	if err != nil {
		return nil, err
	}

	for oauthAccessTokenManagerId, oauthAccessTokenManagerName := range oauthAccessTokenManagerData {
		commentData := map[string]string{
			"Oauth Access Token Manager ID":   oauthAccessTokenManagerId,
			"Oauth Access Token Manager Name": oauthAccessTokenManagerName,
			"Resource Type":                   r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       oauthAccessTokenManagerName,
			ResourceID:         oauthAccessTokenManagerId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOauthAccessTokenManagerResource) getOauthAccessTokenManagerData() (map[string]string, error) {
	oauthAccessTokenManagerData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.OauthAccessTokenManagersAPI.GetTokenManagers(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetTokenManagers", r.ResourceType())
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

	for _, oauthAccessTokenManager := range items {
		oauthAccessTokenManagerId, oauthAccessTokenManagerIdOk := oauthAccessTokenManager.GetIdOk()
		oauthAccessTokenManagerName, oauthAccessTokenManagerNameOk := oauthAccessTokenManager.GetNameOk()

		if oauthAccessTokenManagerIdOk && oauthAccessTokenManagerNameOk {
			oauthAccessTokenManagerData[*oauthAccessTokenManagerId] = *oauthAccessTokenManagerName
		}
	}

	return oauthAccessTokenManagerData, nil
}
