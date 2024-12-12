package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOAuthAccessTokenManagerResource{}
)

type PingFederateOAuthAccessTokenManagerResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateOAuthAccessTokenManagerResource
func OAuthAccessTokenManager(clientInfo *connector.PingFederateClientInfo) *PingFederateOAuthAccessTokenManagerResource {
	return &PingFederateOAuthAccessTokenManagerResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOAuthAccessTokenManagerResource) ResourceType() string {
	return "pingfederate_oauth_access_token_manager"
}

func (r *PingFederateOAuthAccessTokenManagerResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	tokenManagerData, err := r.getTokenManagerData()
	if err != nil {
		return nil, err
	}

	for tokenManagerId, tokenManagerName := range *tokenManagerData {
		commentData := map[string]string{
			"OAuth Access Token Manager Resource ID":   tokenManagerId,
			"OAuth Access Token Manager Resource Name": tokenManagerName,
			"Resource Type": r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       tokenManagerName,
			ResourceID:         tokenManagerId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOAuthAccessTokenManagerResource) getTokenManagerData() (*map[string]string, error) {
	tokenManagerData := make(map[string]string)

	tokenManagers, response, err := r.clientInfo.ApiClient.OauthAccessTokenManagersAPI.GetTokenManagers(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetTokenManagers", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if tokenManagers == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	tokenManagersItems, ok := tokenManagers.GetItemsOk()
	if !ok {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, tokenManager := range tokenManagersItems {
		tokenManagerId, tokenManagerIdOk := tokenManager.GetIdOk()
		tokenManagerName, tokenManagerNameOk := tokenManager.GetNameOk()

		if tokenManagerIdOk && tokenManagerNameOk {
			tokenManagerData[*tokenManagerId] = *tokenManagerName
		}
	}

	return &tokenManagerData, nil
}
