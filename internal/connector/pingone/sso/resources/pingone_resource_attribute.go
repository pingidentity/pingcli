package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneResourceAttributeResource{}
)

type PingoneResourceAttributeResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneResourceAttributeResource
func ResourceAttribute(clientInfo *connector.PingOneClientInfo) *PingoneResourceAttributeResource {
	return &PingoneResourceAttributeResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneResourceAttributeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.ResourcesApi.ReadAllResources(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllResources"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, resourceInner := range embedded.GetResources() {
		resource := resourceInner.Resource
		resourceId, resourceIdOk := resource.GetIdOk()
		resourceName, resourceNameOk := resource.GetNameOk()
		resourceType, resourceTypeOk := resource.GetTypeOk()

		if resourceIdOk && resourceNameOk && resourceTypeOk {
			apiExecuteFunc = r.clientInfo.ApiClient.ManagementAPIClient.ResourceAttributesApi.ReadAllResourceAttributes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *resourceId).Execute
			apiFunctionName = "ReadAllResourceAttributes"

			embedded, err = common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, attributeInner := range embedded.GetAttributes() {
				if attributeInner.ResourceAttribute == nil {
					continue
				}
				resourceAttribute := attributeInner.ResourceAttribute

				resourceAttributeId, resourceAttributeIdOk := resourceAttribute.GetIdOk()
				resourceAttributeName, resourceAttributeNameOk := resourceAttribute.GetNameOk()
				resourceAttributeType, resourceAttributeTypeOk := resourceAttribute.GetTypeOk()

				if resourceAttributeTypeOk {
					switch {
					// Any CORE attribute is required and cannot be overridden
					case *resourceAttributeType == management.ENUMRESOURCEATTRIBUTETYPE_CORE:
						// Handle the special case where a CUSTOM resource can override the sub attribute
						if *resourceType != management.ENUMRESOURCETYPE_CUSTOM {
							continue
						}
						if *resourceAttributeName != "sub" {
							continue
						}
					}
				}

				if resourceAttributeIdOk && resourceAttributeNameOk && resourceAttributeTypeOk {
					commentData := map[string]string{
						"Resource Type":           r.ResourceType(),
						"Resource Name":           *resourceName,
						"Resource Attribute Name": *resourceAttributeName,
						"Export Environment ID":   r.clientInfo.ExportEnvironmentID,
						"Resource ID":             *resourceId,
						"Resource Attribute ID":   *resourceAttributeId,
					}

					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType:       r.ResourceType(),
						ResourceName:       fmt.Sprintf("%s_%s", *resourceName, *resourceAttributeName),
						ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *resourceId, *resourceAttributeId),
						CommentInformation: common.GenerateCommentInformation(commentData),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneResourceAttributeResource) ResourceType() string {
	return "pingone_resource_attribute"
}
