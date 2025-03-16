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
	_ connector.ExportableResource = &PingoneAuthorizeAPIServiceDeploymentResource{}
)

type PingoneAuthorizeAPIServiceDeploymentResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingoneAuthorizeAPIServiceDeploymentResource
func AuthorizeAPIServiceDeployment(clientInfo *connector.ClientInfo) *PingoneAuthorizeAPIServiceDeploymentResource {
	return &PingoneAuthorizeAPIServiceDeploymentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeAPIServiceDeploymentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	apiServiceData, err := r.getAPIServiceData()
	if err != nil {
		return nil, err
	}

	for apiServiceId, apiServiceName := range apiServiceData {
		apiServiceDeployed, err := r.getAPIServiceDeployed(apiServiceId)
		if err != nil {
			return nil, err
		}

		if apiServiceDeployed {
			commentData := map[string]string{
				"API Service ID":        apiServiceId,
				"API Service Name":      apiServiceName,
				"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
				"Resource Type":         r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       apiServiceName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, apiServiceId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeAPIServiceDeploymentResource) getAPIServiceData() (map[string]string, error) {
	apiServiceData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.AuthorizeAPIClient.APIServersApi.ReadAllAPIServers(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiServices, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.APIServer](iter, "ReadAllAPIServers", "GetApiServers", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, apiService := range apiServices {
		apiServiceId, apiServiceIdOk := apiService.GetIdOk()
		apiServiceName, apiServiceNameOk := apiService.GetNameOk()

		if apiServiceIdOk && apiServiceNameOk {
			apiServiceData[*apiServiceId] = *apiServiceName
		}
	}

	return apiServiceData, nil
}

func (r *PingoneAuthorizeAPIServiceDeploymentResource) getAPIServiceDeployed(apiServiceId string) (bool, error) {

	apiServerDeployment, httpResponse, err := r.clientInfo.PingOneApiClient.AuthorizeAPIClient.APIServerDeploymentApi.ReadDeploymentStatus(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, apiServiceId).Execute()
	ok, err := common.HandleClientResponse(httpResponse, err, "ReadDeploymentStatus", r.ResourceType())
	if err != nil {
		return false, err
	}
	// A warning was given when handling the client response. Return nil apiObjects to skip export of resource
	if !ok {
		return false, nil
	}

	if status, ok := apiServerDeployment.GetStatusOk(); ok {
		if statusCode, ok := status.GetCodeOk(); ok && statusCode != nil && *statusCode != "DEPLOYMENT_UNINITIALIZED" {
			return true, nil
		}
	}

	return false, nil
}

func (r *PingoneAuthorizeAPIServiceDeploymentResource) ResourceType() string {
	return "pingone_authorize_api_service_deployment"
}
