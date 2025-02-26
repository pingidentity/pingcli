package resources

// import (
// 	"fmt"

// 	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
// 	"github.com/pingidentity/pingcli/internal/connector"
// 	"github.com/pingidentity/pingcli/internal/connector/common"
// 	"github.com/pingidentity/pingcli/internal/connector/pingone"
// 	"github.com/pingidentity/pingcli/internal/logger"
// )

// // Verify that the resource satisfies the exportable resource interface
// var (
// 	_ connector.ExportableResource = &PingoneAuthorizeAPIServiceDeploymentResource{}
// )

// type PingoneAuthorizeAPIServiceDeploymentResource struct {
// 	clientInfo *connector.PingOneClientInfo
// }

// // Utility method for creating a PingoneAuthorizeAPIServiceDeploymentResource
// func AuthorizeAPIServiceDeployment(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizeAPIServiceDeploymentResource {
// 	return &PingoneAuthorizeAPIServiceDeploymentResource{
// 		clientInfo: clientInfo,
// 	}
// }

// func (r *PingoneAuthorizeAPIServiceDeploymentResource) ExportAll() (*[]connector.ImportBlock, error) {
// 	l := logger.Get()
// 	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

// 	importBlocks := []connector.ImportBlock{}

// 	apiServiceData, err := r.getAPIServiceData()
// 	if err != nil {
// 		return nil, err
// 	}

// 	for apiServiceId, apiServiceName := range apiServiceData {
// 		apiServiceDeploymentData, err := r.getAPIServiceDeploymentData(apiServiceId)
// 		if err != nil {
// 			return nil, err
// 		}

// 		for apiServiceDeploymentId, apiServiceDeploymentName := range apiServiceDeploymentData {
// 			commentData := map[string]string{
// 				"API Service ID":              apiServiceId,
// 				"API Service Name":            apiServiceName,
// 				"API Service Deployment ID":   apiServiceDeploymentId,
// 				"API Service Deployment Name": apiServiceDeploymentName,
// 				"Export Environment ID":       r.clientInfo.ExportEnvironmentID,
// 				"Resource Type":               r.ResourceType(),
// 			}

// 			importBlock := connector.ImportBlock{
// 				ResourceType:       r.ResourceType(),
// 				ResourceName:       fmt.Sprintf("%s_%s", apiServiceName, apiServiceDeploymentId),
// 				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, apiServiceId, apiServiceDeploymentId),
// 				CommentInformation: common.GenerateCommentInformation(commentData),
// 			}

// 			importBlocks = append(importBlocks, importBlock)
// 		}
// 	}

// 	return &importBlocks, nil
// }

// func (r *PingoneAuthorizeAPIServiceDeploymentResource) getAPIServiceData() (map[string]string, error) {
// 	apiServiceData := make(map[string]string)

// 	iter := r.clientInfo.ApiClient.AuthorizeAPIClient.APIServersApi.ReadAllAPIServers(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
// 	apiServices, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.APIServer](iter, "ReadAllAPIServers", "GetAPIServers", r.ResourceType())
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, apiService := range apiServices {
// 		apiServiceId, apiServiceIdOk := apiService.GetIdOk()
// 		apiServiceName, apiServiceNameOk := apiService.GetNameOk()

// 		if apiServiceIdOk && apiServiceNameOk {
// 			apiServiceData[*apiServiceId] = *apiServiceName
// 		}
// 	}

// 	return apiServiceData, nil
// }

// func (r *PingoneAuthorizeAPIServiceDeploymentResource) getAPIServiceDeploymentData(apiServiceId string) (map[string]string, error) {
// 	apiServiceDeploymentData := make(map[string]string)

// 	iter := r.clientInfo.ApiClient.AuthorizeAPIClient.APIServerDeploymentApi.ReadDeploymentStatus(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, apiServiceId).Execute()
// 	apiServiceDeployments, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.APIServerDeployment](iter, "ReadAPIServiceDeployments", "GetRolePermissions", r.ResourceType())
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, apiServiceDeployment := range apiServiceDeployments {
// 		apiServiceDeploymentId, apiServiceDeploymentIdOk := apiServiceDeployment.GetIdOk()
// 		apiServiceDeploymentName, apiServiceDeploymentNameOk := apiServiceDeployment.GetNameOk()

// 		if apiServiceDeploymentIdOk && apiServiceDeploymentNameOk {
// 			apiServiceDeploymentData[*apiServiceDeploymentId] = *apiServiceDeploymentName
// 		}
// 	}

// 	return apiServiceDeploymentData, nil
// }

// func (r *PingoneAuthorizeAPIServiceDeploymentResource) ResourceType() string {
// 	return "pingone_authorize_api_service_deployment"
// }
