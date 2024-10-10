package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAuthorizeAPIServiceDeploymentResource{}
)

type PingoneAuthorizeAPIServiceDeploymentResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizeAPIServiceDeploymentResource
func AuthorizeAPIServiceDeployment(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizeAPIServiceDeploymentResource {
	return &PingoneAuthorizeAPIServiceDeploymentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeAPIServiceDeploymentResource) ExportAll() (*[]connector.ImportBlock, error) {
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
		var (
			apiServerId     *string
			apiServerIdOk   bool
			apiServerName   *string
			apiServerNameOk bool
		)

		apiServerId, apiServerIdOk = apiServer.GetIdOk()
		apiServerName, apiServerNameOk = apiServer.GetNameOk()

		if apiServerIdOk && apiServerNameOk {

			_, response, err := r.clientInfo.ApiClient.AuthorizeAPIClient.APIServerDeploymentApi.ReadDeploymentStatus(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *apiServerId).Execute()
			err = common.HandleClientResponse(response, err, "ReadDeploymentStatus", r.ResourceType())
			if err != nil {
				return nil, err
			}

			importBlocks := []connector.ImportBlock{}

			l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

			if response.StatusCode == 204 {
				l.Debug().Msgf("No exportable %s resource found", r.ResourceType())
				return &importBlocks, nil
			}

			commentData := map[string]string{
				"Resource Type":              r.ResourceType(),
				"Authorize API Service Name": *apiServerName,
				"Authorize API Service ID":   *apiServerId,
				"Export Environment ID":      r.clientInfo.ExportEnvironmentID,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *apiServerName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *apiServerId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})

			return &importBlocks, nil
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeAPIServiceDeploymentResource) ResourceType() string {
	return "pingone_authorize_api_service_deployment"
}
