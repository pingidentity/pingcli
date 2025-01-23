package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOauthAccessTokenMappingResource{}
)

type PingFederateOauthAccessTokenMappingResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateOauthAccessTokenMappingResource
func OauthAccessTokenMapping(clientInfo *connector.PingFederateClientInfo) *PingFederateOauthAccessTokenMappingResource {
	return &PingFederateOauthAccessTokenMappingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOauthAccessTokenMappingResource) ResourceType() string {
	return "pingfederate_oauth_access_token_mapping"
}

func (r *PingFederateOauthAccessTokenMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthAccessTokenMappingData, err := r.getOauthAccessTokenMappingData()
	if err != nil {
		return nil, err
	}

	for oauthAccessTokenMappingId, oauthAccessTokenMappingContextType := range *oauthAccessTokenMappingData {
		commentData := map[string]string{
			"Oauth Access Token Mapping ID":           oauthAccessTokenMappingId,
			"Oauth Access Token Mapping Context Type": oauthAccessTokenMappingContextType,
			"Resource Type":                           r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s", oauthAccessTokenMappingContextType, oauthAccessTokenMappingId),
			ResourceID:         oauthAccessTokenMappingId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOauthAccessTokenMappingResource) getOauthAccessTokenMappingData() (*map[string]string, error) {
	oauthAccessTokenMappingData := make(map[string]string)

	apiObj, response, err := r.clientInfo.ApiClient.OauthAccessTokenMappingsAPI.GetMappings(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetMappings", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if apiObj == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, oauthAccessTokenMapping := range apiObj {
		oauthAccessTokenMappingId, oauthAccessTokenMappingIdOk := oauthAccessTokenMapping.GetIdOk()
		oauthAccessTokenMappingContext, oauthAccessTokenMappingContextOk := oauthAccessTokenMapping.GetContextOk()

		if oauthAccessTokenMappingIdOk && oauthAccessTokenMappingContextOk {
			oauthAccessTokenMappingContextType, oauthAccessTokenMappingContextTypeOk := oauthAccessTokenMappingContext.GetTypeOk()

			if oauthAccessTokenMappingContextTypeOk {
				oauthAccessTokenMappingData[*oauthAccessTokenMappingId] = *oauthAccessTokenMappingContextType
			}
		}
	}

	return &oauthAccessTokenMappingData, nil
}
