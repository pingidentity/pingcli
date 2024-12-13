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

func (r *PingFederateOAuthAccessTokenMappingResource) ResourceType() string {
	return "pingfederate_oauth_access_token_mapping"
}

func (r *PingFederateOAuthAccessTokenMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	mappingData, err := r.getMappingData()
	if err != nil {
		return nil, err
	}

	for mappingId, mappingContextType := range *mappingData {
		commentData := map[string]string{
			"OAuth Access Token Mapping ID":           mappingId,
			"OAuth Access Token Mapping Context Type": mappingContextType,
			"Resource Type":                           r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s", mappingId, mappingContextType),
			ResourceID:         mappingId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOAuthAccessTokenMappingResource) getMappingData() (*map[string]string, error) {
	mappingData := make(map[string]string)

	mappings, response, err := r.clientInfo.ApiClient.OauthAccessTokenMappingsAPI.GetMappings(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetMappings", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, mapping := range mappings {
		mappingId, mappingIdOk := mapping.GetIdOk()
		mappingContext, mappingContextOk := mapping.GetContextOk()

		if mappingIdOk && mappingContextOk {
			mappingContextType, mappingContextTypeOk := mappingContext.GetTypeOk()

			if mappingContextTypeOk {
				mappingData[*mappingId] = *mappingContextType
			}
		}
	}

	return &mappingData, nil
}
