// Copyright © 2025 Ping Identity Corporation

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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneApplicationResourceResource
func ApplicationResource(clientInfo *connector.ClientInfo) *PingOneApplicationResourceResource {
	return &PingOneApplicationResourceResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneApplicationResourceResource) ResourceType() string {
	return "pingone_application_resource"
}

func (r *PingOneApplicationResourceResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationResourceData, err := r.getApplicationResourceData()
	if err != nil {
		return nil, err
	}

	for applicationResourceId, applicationResourceInfo := range applicationResourceData {
		applicationResourceName := applicationResourceInfo[0]
		resourceId := applicationResourceInfo[1]

		resourceName, resourceNameOk, err := r.getResourceName(resourceId)
		if err != nil {
			return nil, err
		}

		if !resourceNameOk {
			continue
		}

		commentData := map[string]string{
			"PingOne Resource ID":       resourceId,
			"PingOne Resource Name":     *resourceName,
			"Application Resource ID":   applicationResourceId,
			"Application Resource Name": applicationResourceName,
			"Export Environment ID":     r.clientInfo.PingOneExportEnvironmentID,
			"Resource Type":             r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s", *resourceName, applicationResourceName),
			ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, resourceId, applicationResourceId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneApplicationResourceResource) getApplicationResourceData() (map[string][]string, error) {
	applicationResourceData := make(map[string][]string)

	iter := r.clientInfo.PingOneApiClient.AuthorizeAPIClient.ApplicationResourcesApi.ReadApplicationResources(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	applicationResources, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.ApplicationResource](iter, "ReadApplicationResources", "GetResources", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, applicationResource := range applicationResources {
		applicationResourceId, applicationResourceIdOk := applicationResource.GetIdOk()
		applicationResourceName, applicationResourceNameOk := applicationResource.GetNameOk()
		resourceId, resourceIdOk := applicationResource.Parent.GetIdOk()

		if applicationResourceIdOk && applicationResourceNameOk && resourceIdOk {
			applicationResourceData[*applicationResourceId] = []string{*applicationResourceName, *resourceId}
		}
	}

	return applicationResourceData, nil
}

func (r *PingOneApplicationResourceResource) getResourceName(resourceId string) (*string, bool, error) {
	apiObj, httpResponse, err := r.clientInfo.PingOneApiClient.ManagementAPIClient.ResourcesApi.ReadOneResource(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, resourceId).Execute()
	ok, err := common.HandleClientResponse(httpResponse, err, "ReadOneResource", r.ResourceType())
	if err != nil {
		return nil, false, err
	}
	if !ok {
		return nil, false, nil
	}

	if apiObj == nil {
		return nil, false, nil
	}

	resourceName, resourceNameOk := apiObj.GetNameOk()
	if !resourceNameOk {
		return nil, false, nil
	}

	return resourceName, true, nil
}
