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
	_ connector.ExportableResource = &PingOneResourceAttributeResource{}
)

type PingOneResourceAttributeResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneResourceAttributeResource
func ResourceAttribute(clientInfo *connector.PingOneClientInfo) *PingOneResourceAttributeResource {
	return &PingOneResourceAttributeResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneResourceAttributeResource) ResourceType() string {
	return "pingone_resource_attribute"
}

func (r *PingOneResourceAttributeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	resourceData, err := r.getResourceData()
	if err != nil {
		return nil, err
	}

	for resourceId, resourceNameAndType := range *resourceData {
		resourceName := resourceNameAndType[0]
		resourceType := resourceNameAndType[1]

		resourceAttributeData, err := r.getResourceAttributeData(resourceId, resourceType)
		if err != nil {
			return nil, err
		}

		for resourceAttributeId, resourceAttributeName := range *resourceAttributeData {
			commentData := map[string]string{
				"Export Environment ID":   r.clientInfo.ExportEnvironmentID,
				"Resource Attribute ID":   resourceAttributeId,
				"Resource Attribute Name": resourceAttributeName,
				"Resource ID":             resourceId,
				"Resource Name":           resourceName,
				"Resource Type":           r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", resourceName, resourceAttributeName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, resourceId, resourceAttributeId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneResourceAttributeResource) getResourceData() (*map[string][]string, error) {
	resourceData := make(map[string][]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.ResourcesApi.ReadAllResources(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllResources", r.ResourceType())
		if err != nil {
			return nil, err
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, resourceInner := range embedded.GetResources() {
			if resourceInner.Resource != nil {
				resourceId, resourceIdOk := resourceInner.Resource.GetIdOk()
				resourceName, resourceNameOk := resourceInner.Resource.GetNameOk()
				resourceType, resourceTypeOk := resourceInner.Resource.GetTypeOk()

				if resourceIdOk && resourceNameOk && resourceTypeOk {
					resourceData[*resourceId] = []string{*resourceName, string(*resourceType)}
				}
			}
		}
	}

	return &resourceData, nil
}

func (r *PingOneResourceAttributeResource) getResourceAttributeData(resourceId string, resourceType string) (*map[string]string, error) {
	resourceAttributeData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.ResourceAttributesApi.ReadAllResourceAttributes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, resourceId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllResourceAttributes", r.ResourceType())
		if err != nil {
			return nil, err
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, attributeInner := range embedded.GetAttributes() {
			if attributeInner.ResourceAttribute != nil {
				resourceAttributeId, resourceAttributeIdOk := attributeInner.ResourceAttribute.GetIdOk()
				resourceAttributeName, resourceAttributeNameOk := attributeInner.ResourceAttribute.GetNameOk()
				resourceAttributeType, resourceAttributeTypeOk := attributeInner.ResourceAttribute.GetTypeOk()

				if resourceAttributeIdOk && resourceAttributeNameOk && resourceAttributeTypeOk {
					// Any CORE attribute is required and cannot be overridden
					// Do not export CORE attributes
					// There is one exception where a CUSTOM resource can override the sub CORE attribute
					if *resourceAttributeType == management.ENUMRESOURCEATTRIBUTETYPE_CORE {
						if resourceType == string(management.ENUMRESOURCETYPE_CUSTOM) {
							// Skip export of all CORE attributes except for the sub attribute for CUSTOM resources
							if *resourceAttributeName != "sub" {
								continue
							}
						} else {
							// Skip export of all CORE attributes for non-CUSTOM resources
							continue
						}
					}

					resourceAttributeData[*resourceAttributeId] = *resourceAttributeName
				}
			}
		}
	}

	return &resourceAttributeData, nil
}
