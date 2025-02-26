package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
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
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	APIServerData, err := r.getAPIServerData()
	if err != nil {
		return nil, err
	}

	for apiServerId, apiServerName := range APIServerData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"API Server ID":         apiServerId,
			"API Server Name":       apiServerName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       apiServerName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, apiServerId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeAPIServiceResource) getAPIServerData() (map[string]string, error) {
	apiServerData := make(map[string]string)

	iter := r.clientInfo.ApiClient.AuthorizeAPIClient.APIServersApi.ReadAllAPIServers(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	apiServers, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.APIServer](iter, "ReadAllAPIServers", "GetApiServers", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, apiServer := range apiServers {
		apiServerId, apiServerIdOk := apiServer.GetIdOk()
		apiServerName, apiServerNameOk := apiServer.GetNameOk()

		if apiServerIdOk && apiServerNameOk {
			apiServerData[*apiServerId] = *apiServerName
		}
	}

	return apiServerData, nil
}

func (r *PingoneAuthorizeAPIServiceResource) ResourceType() string {
	return "pingone_authorize_api_service"
}
