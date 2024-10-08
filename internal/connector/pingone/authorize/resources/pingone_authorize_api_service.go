package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAuthorizeAPIServiceResource{}
)

type PingoneAuthorizeAPIServiceResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizeAPIServiceResource
func AuthorizeAPIService(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizeAPIServiceResource {
	return &PingoneAuthorizeAPIServiceResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeAPIServiceResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.AuthorizeAPIClient.APIServersApi.ReadAllAPIServers(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllAPIServers"

	embedded, err := common.GetAuthorizeEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, apiServer := range embedded.GetApiServers() {
		apiServerName, apiServerNameOk := apiServer.GetNameOk()
		apiServerId, apiServerIdOk := apiServer.GetIdOk()

		if apiServerNameOk && apiServerIdOk {
			commentData := map[string]string{
				"Resource Type":              r.ResourceType(),
				"Authorize API Service Name": *apiServerName,
				"Export Environment ID":      r.clientInfo.ExportEnvironmentID,
				"Authorize API Service ID":   *apiServerId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *apiServerName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *apiServerId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeAPIServiceResource) ResourceType() string {
	return "pingone_authorize_api_service"
}
