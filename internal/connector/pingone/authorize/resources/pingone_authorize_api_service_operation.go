package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAuthorizeAPIServiceOperationResource{}
)

type PingoneAuthorizeAPIServiceOperationResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizeAPIServiceOperationResource
func AuthorizeAPIServiceOperation(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizeAPIServiceOperationResource {
	return &PingoneAuthorizeAPIServiceOperationResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeAPIServiceOperationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteAPIServersFunc := r.clientInfo.ApiClient.AuthorizeAPIClient.APIServersApi.ReadAllAPIServers(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiAPIServersFunctionName := "ReadAllAPIServers"

	embedded, err := common.GetAuthorizeEmbedded(apiExecuteAPIServersFunc, apiAPIServersFunctionName, r.ResourceType())
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
			apiExecuteOperationsFunc := r.clientInfo.ApiClient.AuthorizeAPIClient.APIServerOperationsApi.ReadAllAPIServerOperations(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *apiServerId).Execute
			apiOperationsFunctionName := "ReadAllAPIServerOperations"

			operationsEmbedded, err := common.GetAuthorizeEmbedded(apiExecuteOperationsFunc, apiOperationsFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, apiServerOperation := range operationsEmbedded.GetOperations() {
				apiServerOperationId, apiServerOperationIdOk := apiServerOperation.GetIdOk()
				apiServerOperationName, apiServerOperationNameOk := apiServerOperation.GetNameOk()

				if apiServerOperationNameOk && apiServerOperationIdOk {

					commentData := map[string]string{
						"Resource Type":                        r.ResourceType(),
						"Authorize API Service Name":           *apiServerName,
						"Authorize API Service ID":             *apiServerId,
						"Export Environment ID":                r.clientInfo.ExportEnvironmentID,
						"Authorize API Service Operation Name": *apiServerOperationName,
						"Authorize API Service Operation ID":   *apiServerOperationId,
					}

					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType:       r.ResourceType(),
						ResourceName:       fmt.Sprintf("%s_%s", *apiServerName, *apiServerOperationName),
						ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *apiServerId, *apiServerOperationId),
						CommentInformation: common.GenerateCommentInformation(commentData),
					})

				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeAPIServiceOperationResource) ResourceType() string {
	return "pingone_authorize_api_service_operation"
}
