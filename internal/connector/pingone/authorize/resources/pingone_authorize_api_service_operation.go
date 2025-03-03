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
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	apiServiceData, err := r.getAPIServiceData()
	if err != nil {
		return nil, err
	}

	for apiServiceId, apiServiceName := range apiServiceData {
		apiServiceOperationData, err := r.getAPIServiceOperationData(apiServiceId)
		if err != nil {
			return nil, err
		}

		for apiServiceOperationId, apiServiceOperationName := range apiServiceOperationData {
			commentData := map[string]string{
				"API Service ID":             apiServiceId,
				"API Service Name":           apiServiceName,
				"API Service Operation ID":   apiServiceOperationId,
				"API Service Operation Name": apiServiceOperationName,
				"Export Environment ID":      r.clientInfo.ExportEnvironmentID,
				"Resource Type":              r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", apiServiceName, apiServiceOperationName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, apiServiceId, apiServiceOperationId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeAPIServiceOperationResource) getAPIServiceData() (map[string]string, error) {
	apiServiceData := make(map[string]string)

	iter := r.clientInfo.ApiClient.AuthorizeAPIClient.APIServersApi.ReadAllAPIServers(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
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

func (r *PingoneAuthorizeAPIServiceOperationResource) getAPIServiceOperationData(apiServiceId string) (map[string]string, error) {
	apiServiceOperationData := make(map[string]string)

	iter := r.clientInfo.ApiClient.AuthorizeAPIClient.APIServerOperationsApi.ReadAllAPIServerOperations(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, apiServiceId).Execute()
	apiServiceOperations, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.APIServerOperation](iter, "ReadAllAPIServerOperations", "GetOperations", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, apiServiceOperation := range apiServiceOperations {
		apiServiceOperationId, apiServiceOperationIdOk := apiServiceOperation.GetIdOk()
		apiServiceOperationName, apiServiceOperationNameOk := apiServiceOperation.GetNameOk()

		if apiServiceOperationIdOk && apiServiceOperationNameOk {
			apiServiceOperationData[*apiServiceOperationId] = *apiServiceOperationName
		}
	}

	return apiServiceOperationData, nil
}

func (r *PingoneAuthorizeAPIServiceOperationResource) ResourceType() string {
	return "pingone_authorize_api_service_operation"
}
