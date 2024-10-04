package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOAuthAccessTokenMappingResource{}
)

type PingFederateOAuthAccessTokenMappingResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateOAuthAccessTokenMappingResource
func OAuthAccessTokenMapping(clientInfo *connector.PingFederateClientInfo) *PingFederateOAuthAccessTokenMappingResource {
	return &PingFederateOAuthAccessTokenMappingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOAuthAccessTokenMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.OauthAccessTokenMappingsAPI.GetMappings(r.clientInfo.Context).Execute
	apiFunctionName := "GetMappings"

	accessTokenMappings, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if accessTokenMappings == nil {
		l.Error().Msgf("Returned %s() accessTokenMappings is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, accessTokenMapping := range accessTokenMappings {
		accessTokenMappingId, accessTokenMappingIdOk := accessTokenMapping.GetIdOk()
		accessTokenMappingContext, accessTokenMappingContextOk := accessTokenMapping.GetContextOk()
		var (
			accessTokenMappingContextType   *string
			accessTokenMappingContextTypeOk bool
		)
		if accessTokenMappingContextOk {
			accessTokenMappingContextType, accessTokenMappingContextTypeOk = accessTokenMappingContext.GetTypeOk()
		}

		if accessTokenMappingIdOk && accessTokenMappingContextOk && accessTokenMappingContextTypeOk {
			commentData := map[string]string{
				"Resource Type":                           r.ResourceType(),
				"OAuth Access Token Mapping Resource ID":  *accessTokenMappingId,
				"OAuth Access Token Mapping Context Type": *accessTokenMappingContextType,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", *accessTokenMappingId, *accessTokenMappingContextType),
				ResourceID:         *accessTokenMappingId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederateOAuthAccessTokenMappingResource) ResourceType() string {
	return "pingfederate_oauth_access_token_mapping"
}
