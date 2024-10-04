package resources

import (
	"fmt"

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

func (r *PingFederateOAuthAccessTokenManagerResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.OauthAccessTokenManagersAPI.GetTokenManagers(r.clientInfo.Context).Execute
	apiFunctionName := "GetTokenManagers"

	accessTokenManagers, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if accessTokenManagers == nil {
		l.Error().Msgf("Returned %s() accessTokenManagers is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	accessTokenManagersItems, ok := accessTokenManagers.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() accessTokenManagers items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, accessTokenManager := range accessTokenManagersItems {
		accessTokenManagerId, accessTokenManagerIdOk := accessTokenManager.GetIdOk()
		accessTokenManagerName, accessTokenManagerNameOk := accessTokenManager.GetNameOk()

		if accessTokenManagerIdOk && accessTokenManagerNameOk {
			commentData := map[string]string{
				"Resource Type":                            r.ResourceType(),
				"OAuth Access Token Manager Resource ID":   *accessTokenManagerId,
				"OAuth Access Token Manager Resource Name": *accessTokenManagerName,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *accessTokenManagerName,
				ResourceID:         *accessTokenManagerId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederateOAuthAccessTokenManagerResource) ResourceType() string {
	return "pingfederate_oauth_access_token_manager"
}
