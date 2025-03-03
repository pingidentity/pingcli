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
	_ connector.ExportableResource = &PingOneApplicationResourceResource{}
)

type PingOneApplicationResourceResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneApplicationResourceResource
func ApplicationResource(clientInfo *connector.PingOneClientInfo) *PingOneApplicationResourceResource {
	return &PingOneApplicationResourceResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneApplicationResourceResource) ResourceType() string {
	return "pingone_application_resource"
}

type applicationResourceObj struct {
	applicationResourceName string
	resourceId              string
	resourceName            string
}

func (r *PingOneApplicationResourceResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationResourceData, err := r.getApplicationResourceData()
	if err != nil {
		return nil, err
	}

	for applicationResourceId, applicationResourceObj := range applicationResourceData {
		commentData := map[string]string{
			"PingOne Resource ID":       applicationResourceObj.resourceId,
			"PingOne Resource Name":     applicationResourceObj.resourceName,
			"Application Resource ID":   applicationResourceId,
			"Application Resource Name": applicationResourceObj.applicationResourceName,
			"Export Environment ID":     r.clientInfo.ExportEnvironmentID,
			"Resource Type":             r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s", applicationResourceObj.resourceName, applicationResourceObj.applicationResourceName),
			ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, applicationResourceObj.resourceId, applicationResourceId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneApplicationResourceResource) getApplicationResourceData() (map[string]applicationResourceObj, error) {
	applicationResourceData := make(map[string]applicationResourceObj)

	iter := r.clientInfo.ApiClient.AuthorizeAPIClient.ApplicationResourcesApi.ReadApplicationResources(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	applicationResources, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.ApplicationResource](iter, "ReadApplicationResources", "GetResources", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, applicationResource := range applicationResources {
		applicationResourceId, applicationResourceIdOk := applicationResource.GetIdOk()
		applicationResourceName, applicationResourceNameOk := applicationResource.GetNameOk()
		resourceId, resourceIdOk := applicationResource.Parent.GetIdOk()

		if applicationResourceIdOk && applicationResourceNameOk && resourceIdOk {

			resourceObj, httpResponse, err := r.clientInfo.ApiClient.ManagementAPIClient.ResourcesApi.ReadOneResource(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *resourceId).Execute()
			ok, err := common.HandleClientResponse(httpResponse, err, "ReadOneResource", r.ResourceType())
			if err != nil {
				return nil, err
			}
			// A warning was given when handling the client response. Return nil apiObjects to skip export of resource
			if !ok {
				return nil, nil
			}

			applicationResourceData[*applicationResourceId] = applicationResourceObj{
				applicationResourceName: *applicationResourceName,
				resourceId:              *resourceId,
				resourceName:            resourceObj.GetName(),
			}
		}
	}

	return applicationResourceData, nil
}
