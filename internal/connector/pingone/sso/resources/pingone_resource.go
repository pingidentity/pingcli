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
	_ connector.ExportableResource = &PingOneResourceResource{}
)

type PingOneResourceResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneResourceResource
func Resource(clientInfo *connector.ClientInfo) *PingOneResourceResource {
	return &PingOneResourceResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneResourceResource) ResourceType() string {
	return "pingone_resource"
}

func (r *PingOneResourceResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	resourceData, err := r.getResourceData()
	if err != nil {
		return nil, err
	}

	for resourceId, resourceName := range resourceData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
			"PingOne Resource ID":   resourceId,
			"PingOne Resource Name": resourceName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       resourceName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, resourceId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneResourceResource) getResourceData() (map[string]string, error) {
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

			if resourceIdOk && resourceNameOk {
				resourceData[*resourceId] = *resourceName
			}
		}
	}

	return resourceData, nil
}
