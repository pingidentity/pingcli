// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneResourceSecretResource{}
)

type PingOneResourceSecretResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneResourceSecretResource
func ResourceSecret(clientInfo *connector.ClientInfo) *PingOneResourceSecretResource {
	return &PingOneResourceSecretResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneResourceSecretResource) ResourceType() string {
	return "pingone_resource_secret"
}

func (r *PingOneResourceSecretResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	resourceData, err := r.getResourceData()
	if err != nil {
		return nil, err
	}

	for resourceId, resourceName := range resourceData {
		resourceSecretOk, err := r.getResourceSecret(resourceId)
		if err != nil {
			return nil, err
		}

		if !resourceSecretOk {
			continue
		}

		commentData := map[string]string{
			"Resource ID":           resourceId,
			"Resource Name":         resourceName,
			"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_secret", resourceName),
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, resourceId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneResourceSecretResource) getResourceData() (map[string]string, error) {
	resourceData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ResourcesApi.ReadAllResources(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	resourceInners, err := pingone.GetManagementAPIObjectsFromIterator[management.EntityArrayEmbeddedResourcesInner](iter, "ReadAllResources", "GetResources", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, resourceInner := range resourceInners {
		if resourceInner.Resource != nil {
			resourceId, resourceIdOk := resourceInner.Resource.GetIdOk()
			resourceName, resourceNameOk := resourceInner.Resource.GetNameOk()
			resourceType, resourceTypeOk := resourceInner.Resource.GetTypeOk()

			if resourceIdOk && resourceNameOk && resourceTypeOk && *resourceType == management.ENUMRESOURCETYPE_CUSTOM {
				resourceData[*resourceId] = *resourceName
			}
		}
	}

	return resourceData, nil
}

func (r *PingOneResourceSecretResource) getResourceSecret(resourceId string) (bool, error) {
	_, response, err := r.clientInfo.PingOneApiClient.ManagementAPIClient.ResourceClientSecretApi.ReadResourceSecret(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, resourceId).Execute()
	return common.HandleClientResponse(response, err, "ReadResourceSecret", r.ResourceType())
}
